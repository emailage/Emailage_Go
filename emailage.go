package emailage

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/emailage/Emailage_Go/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// AuthenticationScheme describes the authentication scheme that will be used - OAUTH1 or OAUTH2.
type AuthenticationScheme string

const (
	// OAUTH1 is an older, stable authentication scheme
	OAUTH1 AuthenticationScheme = "oauth1"

	// OAUTH2 is a newer authentication scheme
	OAUTH2 AuthenticationScheme = "oauth2"
)

// ClientOpts contains fields used by the client
type ClientOpts struct {
	Token         string
	AccountSID    string
	Endpoint      string
	TokenEndpoint string
	AuthType      AuthenticationScheme
	Algorithm     auth.HMACSHA
	HTTPTimeout   time.Duration
}

// validate validates that the required fields are present
func (c *ClientOpts) validate() error {
	switch {
	case c.Token == "":
		return errors.New("missing token")
	case c.AccountSID == "":
		return errors.New("missing AccountSID")
	case c.Endpoint == "":
		return errors.New("missing endpoint")
	case c.TokenEndpoint == "" && c.AuthType == OAUTH2:
		return errors.New("a token endpoint is required with oauth2")
	}
	return nil
}

// Emailage Configuration object for the emailage type
type Emailage struct {
	opts       *ClientOpts
	oc         auth.Authorizer
	HTTPClient http.Client
	ctx        context.Context
	cccfg      *clientcredentials.Config
	oa2Token   *oauth2.Token
}

// New creates a new value of type pointer Emailage
func New(co *ClientOpts) (*Emailage, error) {
	if co == nil {
		return nil, errors.New("client opts nil")
	}
	if err := co.validate(); err != nil {
		return nil, err
	}
	a, err := auth.New()
	if err != nil {
		return nil, err
	}

	e := &Emailage{
		opts: co,
		oc:   a,
	}

	if co.HTTPTimeout > 0 {
		e.HTTPClient.Timeout = co.HTTPTimeout
	}
	e.ctx = context.Background()

	e.cccfg = &clientcredentials.Config{
		ClientID:     e.opts.AccountSID,
		ClientSecret: e.opts.Token,
		Scopes:       []string{},
		AuthStyle:    oauth2.AuthStyleInParams,
		TokenURL:     e.opts.TokenEndpoint,
	}

	return e, nil
}

// EmailOnlyScore provides a risk score for the provided email address.
func (e *Emailage) EmailOnlyScore(email string, params map[string]string) (*Response, error) {
	return e.base(email, params)
}

// IPAddressOnlyScore provides a risk score for the provided IP address.
func (e *Emailage) IPAddressOnlyScore(ip string, params map[string]string, scheme AuthenticationScheme) (*Response, error) {
	return e.base(ip, params)
}

// EmailAndIPScore provides a risk score for the provided email/IP address
// combination. IP4 and IP6 addresses are supported.
func (e *Emailage) EmailAndIPScore(email, ip string, params map[string]string, scheme AuthenticationScheme) (*Response, error) {
	return e.base(email+"+"+ip, params)
}

// base is an intermediate method call that all exposed methods call which then proxy's
// that call to the API
func (e *Emailage) base(input string, params map[string]string) (*Response, error) {
	var method auth.RequestMethod
	if e.opts.AuthType == OAUTH2 {
		method = auth.OAUTH2
	} else {
		method = auth.POST
	}

	if params != nil {
		params["format"] = "json"
		params["query"] = input
	} else {
		params = map[string]string{
			"format": "json",
			"query":  input,
		}
	}
	var r Response
	if err := e.call(params, &r, method); err != nil {
		return nil, err
	}
	if (r.ResponseStatus != nil) && (r.ResponseStatus.Status == "failed") {
		return nil, errors.New(ErrorCodeLookup(r.ResponseStatus.ErrorCode))
	}
	var err error
	r.Query.Email, err = url.QueryUnescape(r.Query.Email)
	if err != nil {
		return nil, errors.New("Could not escape returned email")
	}
	for _, result := range r.Query.Results {
		result.Email, err = url.QueryUnescape(result.Email)
		if err != nil {
			return nil, errors.New("Could not escape returned email")
		}
	}

	return &r, nil
}

// removeBOM removes the first 3 bytes of the API response.  Those
// bytes take the form of a Byte Order Mark and will not allow
// the received JSON to be marshalled correctly otherwise
func removeBOM(d io.ReadCloser) (io.Reader, error) {
	var buf = bufio.NewReader(d)
	r, _, err := buf.ReadRune()
	if err != nil {
		return nil, err
	}
	if r != '\uFEFF' {
		buf.UnreadRune() // Not a BOM so put the rune back
	}
	return buf, nil
}

func (e *Emailage) call(params map[string]string, fres interface{}, method auth.RequestMethod) error {
	switch method {
	case auth.GET:
		return e.getRequest(params, fres)
	case auth.POST:
		return e.postRequest(params, fres)
	case auth.OAUTH2:
		return e.oauth2Request(params, fres)
	}
	return errors.New("request method not supported")
}

func (e *Emailage) oauth2Request(params map[string]string, fres interface{}) error {
	query := mapToURLValues(params)

	var err error
	if e.oa2Token == nil || !e.oa2Token.Valid() {
		e.oa2Token, err = e.cccfg.Token(e.ctx)
		if err != nil {
			return err
		}
	}

	client := e.cccfg.Client(e.ctx)
	res, err := client.PostForm(e.opts.Endpoint+"?format=json", query)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return handleResponse(err, res, fres)
}

func mapToURLValues(params map[string]string) url.Values {
	if params == nil {
		return nil
	}
	result := url.Values{}
	for k, v := range params {
		result.Add(k, v)
	}
	return result
}

func (e *Emailage) postRequest(params map[string]string, fres interface{}) error {

	// populate authentication parameters
	ap := e.populateAuthParameters()
	var qsa strings.Builder
	qsa.WriteString("format=json&")
	qsa.WriteString("oauth_consumer_key=" + ap["oauth_consumer_key"] + "&")
	qsa.WriteString("oauth_nonce=" + ap["oauth_nonce"] + "&")
	qsa.WriteString("oauth_signature_method=" + ap["oauth_signature_method"] + "&")
	qsa.WriteString("oauth_timestamp=" + ap["oauth_timestamp"] + "&")
	qsa.WriteString("oauth_version=1.0")

	// calculate full Auth url
	var u strings.Builder
	u.WriteString("POST&")
	u.WriteString(url.QueryEscape(e.opts.Endpoint) + "&")
	u.WriteString(url.QueryEscape(qsa.String()))

	s, err := e.oc.GetSignature(u.String(), e.opts.Algorithm, e.opts.Token)
	if err != nil {
		return err
	}

	var qs string
	var res *http.Response
	qa := fmt.Sprintf("?format=%s&oauth_version=%s&oauth_consumer_key=%s&oauth_timestamp=%s&oauth_nonce=%s&oauth_signature_method=%s&oauth_signature=%s",
		ap["format"], ap["oauth_version"], ap["oauth_consumer_key"], ap["oauth_timestamp"], ap["oauth_nonce"], ap["oauth_signature_method"], s)
	qs = e.opts.Endpoint + qa
	req, err := http.NewRequest("POST", qs, bytes.NewReader([]byte(appendQSParameters(params))))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Language", "en-US")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Accept-Charset", "utf-8")
	res, err = e.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return handleResponse(err, res, fres)
}

// call setups up the request to the Classic API and executes it
func (e *Emailage) getRequest(params map[string]string, fres interface{}) error {

	// populate authentication parameters
	ap := e.populateAuthParameters()
	q := prepareQSForSignatureCalc(ap, params)

	// calculate full Auth url
	var u strings.Builder
	fmt.Fprintf(&u, "GET&%s&%s", url.QueryEscape(e.opts.Endpoint), url.QueryEscape(q))

	s, err := e.oc.GetSignature(u.String(), e.opts.Algorithm, e.opts.Token)
	if err != nil {
		return err
	}

	var qs string
	var res *http.Response
	var qa strings.Builder
	qa.WriteString(q)
	fmt.Fprintf(&qa, "&oauth_signature=%s", s)
	qs = e.opts.Endpoint + "?" + qa.String()
	res, err = e.HTTPClient.Get(qs)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return handleResponse(err, res, fres)
}

func handleResponse(err error, res *http.Response, fres interface{}) error {
	buf, err := removeBOM(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		buf := bytes.NewBuffer(make([]byte, 0, res.ContentLength))
		if _, err := buf.ReadFrom(res.Body); err != nil {
			return err
		}
		return errors.New(buf.String())
	}
	if fres == nil {
		return nil
	}
	return json.NewDecoder(buf).Decode(fres)
}

func prepareQSForSignatureCalc(ap map[string]string, params map[string]string) string {
	qp := make(map[string]string, len(ap)+len(params))
	for k, v := range params {
		qp[k] = url.QueryEscape(v)
	}
	for k, v := range ap {
		qp[k] = url.QueryEscape(v)
	}

	// sort parameters in alphabetical order
	var i int
	var m = make([]string, len(qp))
	for k := range qp {
		m[i] = k
		i++
	}
	sort.Strings(m)

	// encode Auth parameters for calculation
	var q strings.Builder
	for _, v := range m {
		if v != "" {
			if q.Len() > 1 {
				q.WriteRune('&')
			}
			q.WriteString(v)
			q.WriteRune('=')
			q.WriteString(params[v])
		}
	}

	return q.String()
}

func (e *Emailage) populateAuthParameters() map[string]string {
	ap := make(map[string]string)
	ts := time.Now().Unix()
	ap["format"] = "json"
	ap["oauth_consumer_key"] = e.opts.AccountSID
	ap["oauth_nonce"] = e.oc.RandomString(10)
	ap["oauth_signature_method"] = string(e.opts.Algorithm)
	ap["oauth_timestamp"] = strconv.FormatInt(ts, 10)
	ap["oauth_version"] = "1.0"
	return ap
}
func appendQSParameters(params map[string]string) string {
	var r strings.Builder
	for k, v := range params {
		if v != "" {
			if r.Len() > 1 {
				r.WriteRune('&')
			}
			r.WriteString(k)
			r.WriteRune('=')
			r.WriteString(v)
		}
	}
	return r.String()
}

// ErrorCodeLookup provies the ability to look up an error code
// returned from the API
func ErrorCodeLookup(code string) string {
	switch code {
	case "400":
		return "Invalid input parameter. Error message should indicate which one"
	case "401":
		return "Invalid or expired token. This can happen if an access token was either revoked or has expired. This can be fixed by re-authenticating the user"
	case "403":
		return "Invalid oAuth request (wrong consumer key, bad nonce, expired timestamp...)"
	case "404":
		return "File or folder was not found at the specified path"
	case "405":
		return "Request method not expected (generally should be GET or POST)"
	case "503":
		return "Calls are exceeding the defined throttle limit."
	case "3001":
		return "Authentication Error: The signature doesn't match or the user/consumer key file wasn't found"
	case "3002":
		return "Authentication Error: Your account status is inactive or disabled. Please contact us at support@Emailage.com to activate your account"
	case "3003":
		return "Authentication Error: Your account is currently expired. The free trial access period has ended. Please contact support@Emailage.com to upgrade your account"
	case "3004":
		return "Authentication Error: You do not have enough query credits available. Please contact support@Emailage.com to upgrade your account"
	case "3005":
		return "Authentication Error: A general error has occurred which prevented the proper authorization by our API in response to your request. Please contact us at apisupport@Emailage.com"
	case "3006":
		return "Invalid Parameter: Parameter not provided or empty"
	case "3007":
		return "Invalid Parameter: Malformed or wrong parameter provided"
	case "3008":
		return "Authentication Error: Invalid timestamp provided"
	case "3009":
		return "Authentication Error: Invalid nonce provided"
	case "3010":
		return "Invalid Parameter: Invalid partnerid provided"
	case "3011":
		return "Invalid Parameter: Invalid CustomerIdentifierID provided"
	case "3012":
		return "Invalid Parameter: The user_email is not valid or is inactive"
	case "3013":
		return "Invalid Parameter: The user_email and departmentid do not match"
	case "3014":
		return "Invalid Parameter: Consumer Key not found"
	case "3015":
		return "Invalid Parameter: No account found for your consumer Key"
	case "3016":
		return "Authentication Error: Invalid Signature Method"
	default:
		return "unknown error"
	}
}
