package emailage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/mrjones/oauth"
)

// ClientOpts contains fields used by the client
type ClientOpts struct {
	Token      string
	AccountSID string
	Endpoint   string
	HTTP       *http.Client
	Format     string
}

// validate validates that the required fields are present
func (c *ClientOpts) validate() error {
	format := strings.ToLower(c.Format)
	switch {
	case c.Token == "":
		return errors.New("missing token")
	case c.AccountSID == "":
		return errors.New("missing AccountSID")
	case c.Endpoint == "":
		return errors.New("missing endpoint")
	case format != "json" && format != "xml":
		return errors.New("unsupported or missing response format. Only JSON or XML is supported")
	}
	return nil
}

// Emailage
type Emailage struct {
	opts *ClientOpts
	oc   *oauth.Consumer
}

// New creates a new value of type pointer Emailage
func New(co *ClientOpts) (*Emailage, error) {
	if co == nil {
		return nil, errors.New("client opts nil")
	}
	if err := co.validate(); err != nil {
		return nil, err
	}
	e := &Emailage{
		opts: co,
	}
	sp := oauth.ServiceProvider{}
	e.oc = oauth.NewConsumer(co.Token, co.AccountSID, sp)
	if co.HTTP != nil {
		e.oc.HttpClient = co.HTTP
	}
	return e, nil
}

// EmailOnlyScore provides a risk score for the provided email address.
func (e *Emailage) EmailOnlyScore(email string) (*Response, error) {
	if !IsEmail(email) {
		return nil, errors.New("invalid email address")
	}
	params := map[string]string{
		"format": e.opts.Format,
		"query":  email,
	}
	var r Response
	if err := e.call("email_only_score", params, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// IPAddressOnlyScore provides a risk score for the provided IP address.
func (e *Emailage) IPAddressOnlyScore(ip string) (*Response, error) {
	if !IsIP(ip) {
		return nil, errors.New("invalid ip address")
	}
	params := map[string]string{
		"format": e.opts.Format,
		"query":  ip,
	}
	var r Response
	if err := e.call("ip_only_score", params, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// EmailAndIPScore provides a risk score for the provided email/IP address
// combination. IP4 and IP6 addresses are supported.
func (e *Emailage) EmailAndIPScore(email, ip string) (*Response, error) {
	return nil, nil
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
func (e *Emailage) call(caller string, params map[string]string, fres interface{}) error {
	res, err := e.oc.Get(e.opts.Endpoint, params, &oauth.AccessToken{})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	buf, err := removeBOM(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		defer res.Body.Close()
		buf := bytes.NewBuffer(make([]byte, 0, res.ContentLength))
		if _, err := buf.ReadFrom(res.Body); err != nil {
			return err
		}
		return newServiceError(buf.String(), caller, res.StatusCode)
	}
	if fres == nil {
		return nil
	}
	return json.NewDecoder(buf).Decode(fres)
}
