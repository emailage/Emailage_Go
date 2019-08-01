package emailage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/emailage/Emailage_Go/auth"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type ResponseFormat string

const (
	JSON ResponseFormat = "json"
	XML  ResponseFormat = "xml"
)

// ClientOpts contains fields used by the client
type ClientOpts struct {
	Token       string
	AccountSID  string
	Endpoint    string
	Format      ResponseFormat
	Algorithm   auth.HMACSHA
	HTTPTimeout time.Duration
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
	case c.Format != "json" && c.Format != "xml":
		return errors.New("unsupported or missing response format. Only JSON or XML is supported")
	}
	return nil
}

// Emailage
type Emailage struct {
	opts       *ClientOpts
	oc         auth.Authorizer
	HttpClient http.Client
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
		e.HttpClient.Timeout = co.HTTPTimeout
	}
	return e, nil
}

// EmailOnlyScore provides a risk score for the provided email address.
func (e *Emailage) EmailOnlyScore(email string, params map[string]string) (*Response, error) {
	return e.base(email, params)
}

// IPAddressOnlyScore provides a risk score for the provided IP address.
func (e *Emailage) IPAddressOnlyScore(ip string, params map[string]string) (*Response, error) {
	return e.base(ip, params)
}

// EmailAndIPScore provides a risk score for the provided email/IP address
// combination. IP4 and IP6 addresses are supported.
func (e *Emailage) EmailAndIPScore(email, ip string, params map[string]string) (*Response, error) {
	return e.base(email+"+"+ip, params)
}

// base is an intermediate method call that all exposed methods call which then proxy's
// that call to the API
func (e *Emailage) base(input string, params map[string]string) (*Response, error) {
	if params != nil {
		params["format"] = string(e.opts.Format)
		params["query"] = input
	} else {
		params = map[string]string{
			"format": string(e.opts.Format),
			"query":  input,
		}
	}
	var r Response
	if err := e.call(params, &r); err != nil {
		return nil, err
	}
	if (r.ResponseStatus != nil) && (r.ResponseStatus.Status == "failed") {
		return nil, errors.New(ErrorCodeLookup(r.ResponseStatus.ErrorCode))
	}
	return &r, nil
}

// removeBOM removes the first 3 bytes of the API response.  Those
// bytes take the form of a Byte Order Mark and will not allow
// the received JSON to be marshalled correctly otherwise
func removeBOM(d io.ReadCloser) (io.Reader, error) {
	buf := bufio.NewReader(d)
	r, _, err := buf.ReadRune()
	if err != nil {
		return nil, err
	}
	if r != '\uFEFF' {
		buf.UnreadRune() // Not a BOM so put the rune back
	}
	return buf, nil
}

// call setups up the request to the Classic API and executes it
func (e *Emailage) call(params map[string]string, fres interface{}) error {

	// populate authentication parameters
	ts := time.Now().Unix()
	params["format"] = "json"
	params["oauth_consumer_key"] = e.opts.AccountSID
	params["oauth_nonce"] = e.oc.GetRandomString(10)
	params["oauth_signature_method"] = string(e.opts.Algorithm)
	params["oauth_timestamp"] = strconv.FormatInt(ts, 10)
	params["oauth_version"] = "1.0"

	for k, v := range params {
		params[k] = url.QueryEscape(v)
	}

	// sort parameters in alphabetical order
	i := 0
	m := make([]string, len(params))
	for k := range params {
		m[i] = k
		i++
	}
	sort.Strings(m)

	// calculate signature
	var q bytes.Buffer
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
	qs := url.QueryEscape(q.String())

	s, err := e.oc.GetSignature("GET&"+url.QueryEscape(e.opts.Endpoint)+"&"+qs, auth.GET, e.opts.Algorithm, e.opts.Token)
	if err != nil {
		return err
	}

	q.WriteString("&oauth_signature=")
	q.WriteString(s)
	qs = e.opts.Endpoint + "?" + q.String()

	res, err := e.HttpClient.Get(qs)
	if err != nil {
		return err
	}
	defer res.Body.Close()

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
