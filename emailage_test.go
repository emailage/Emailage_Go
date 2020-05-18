package emailage

import (
	"github.com/emailage/Emailage_Go/auth"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

/*
	Test Objects
*/
type mockFuncAuth func(auth *authorizer)

type authorizer struct {
	mock.Mock
}

func (s *authorizer) GetSignature(fullURL string, hmacsha auth.HMACSHA, token string) (string, error) {
	args := s.Called(fullURL, hmacsha, token)
	return args.String(0), args.Error(1)
}

func (s *authorizer) RandomString(length int) string {
	args := s.Called(length)
	return args.String(0)
}

type mockFunc func()

/*
	TESTS
*/
func TestClientOpts_validate(t *testing.T) {
	type fields struct {
		Token      string
		AccountSID string
		Endpoint   string
		HTTP       *http.Client
		Format     ResponseFormat
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid opts",
			fields: fields{
				Token:      "token",
				AccountSID: "accountsid",
				Endpoint:   "someendpoint",
				Format:     JSON,
			},
			wantErr: false,
		},
		{
			name: "missing token",
			fields: fields{
				AccountSID: "accountsid",
				Endpoint:   "someendpoint",
				Format:     JSON,
			},
			wantErr: true,
		},
		{
			name: "missing accountSID",
			fields: fields{
				Token:    "token",
				Endpoint: "someendpoint",
				Format:   JSON,
			},
			wantErr: true,
		},
		{
			name: "missing endpoint",
			fields: fields{
				Token:      "token",
				AccountSID: "accountsid",
				Format:     JSON,
			},
			wantErr: true,
		},
		{
			name: "missing format",
			fields: fields{
				Token:      "token",
				AccountSID: "accountsid",
				Endpoint:   "someendpoint",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientOpts{
				Token:      tt.fields.Token,
				AccountSID: tt.fields.AccountSID,
				Endpoint:   tt.fields.Endpoint,
				Format:     tt.fields.Format,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("ClientOpts.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	co := &ClientOpts{
		Token:      "token",
		AccountSID: "accountsid",
		Endpoint:   "someendpoint",
		Format:     "json",
	}
	type args struct {
		co *ClientOpts
	}
	tests := []struct {
		name    string
		args    args
		want    *Emailage
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				co: co,
			},
			want: &Emailage{
				oc:   &auth.Oauth1{},
				opts: co,
			},
			wantErr: false,
		},
		{
			name: "nil client opts",
			args: args{
				co: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.co)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.TypeOf(got).String() != "*emailage.Emailage" && !tt.wantErr {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailage_IPAddressOnlyScore(t *testing.T) {
	type fields struct {
		opts *ClientOpts
		oc   *auth.Oauth1
	}
	type args struct {
		ip     string
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Emailage{
				opts: tt.fields.opts,
				oc:   tt.fields.oc,
			}
			got, err := e.IPAddressOnlyScore(tt.args.ip, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Emailage.IPAddressOnlyScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Emailage.IPAddressOnlyScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailage_EmailAndIPScore(t *testing.T) {
	type fields struct {
		opts *ClientOpts
		oc   *auth.Oauth1
	}
	type args struct {
		email  string
		ip     string
		params map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Emailage{
				opts: tt.fields.opts,
				oc:   tt.fields.oc,
			}
			got, err := e.EmailAndIPScore(tt.args.email, tt.args.ip, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Emailage.EmailAndIPScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Emailage.EmailAndIPScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeBOM(t *testing.T) {
	type args struct {
		d io.ReadCloser
	}
	tests := []struct {
		name    string
		args    args
		want    io.Reader
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeBOM(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeBOM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeBOM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorCodeLookup(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "invalid input",
			args: args{
				code: "400",
			},
			want: "Invalid input parameter. Error message should indicate which one",
		},
		{
			name: "not found",
			args: args{
				code: "9999",
			},
			want: "unknown error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorCodeLookup(tt.args.code); got != tt.want {
				t.Errorf("ErrorCodeLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailage_EmailOnlyScore(t *testing.T) {
	type fields struct {
		opts *ClientOpts
		oc   authorizer
	}
	type args struct {
		email  string
		params map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		authMocks   []mockFuncAuth
		clientMocks []mockFunc
		want        *Response
		wantErr     bool
	}{
		{
			name: "Simple Email Query",
			fields: fields{
				opts: &ClientOpts{
					Token:       "token",
					AccountSID:  "sid",
					Endpoint:    "https://api.exists.somewhere.com/",
					Format:      JSON,
					Algorithm:   auth.HMACSHA512,
					HTTPTimeout: 0,
				},
			},
			args: args{
				email:  "nigerian.prince@legit.ru",
				params: nil,
			},
			authMocks: []mockFuncAuth{
				func(auth *authorizer) {
					auth.On("RandomString", mock.Anything).Return("asdfghjkl;")
					auth.On("GetSignature", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("hmacsignature", nil)
				},
			},
			clientMocks: []mockFunc{
				func() {
					httpmock.RegisterResponder("POST",
						`=~^https://api.exists.somewhere.com/.*`,
						httpmock.NewStringResponder(200, `{"query":{"email":"nigerian.prince%40legit.ru","queryType":"EmailAgeVerification","count":1,"created":"2019-08-01T00:47:48Z","lang":"en-US","responseCount":1,"results":[{"userdefinedrecordid":"","email":"nigerian.prince@legit.ru","eName":"","emailAge":"","email_creation_days":"","domainAge":"2010-02-08T21:00:00Z","domain_creation_days":"3460","firstVerificationDate":"2018-05-30T17:28:08Z","first_seen_days":"427","lastVerificationDate":"2019-07-31T06:51:02Z","status":"ValidDomain","country":"RU","fraudRisk":"500 Moderate","EAScore":"500","EAReason":"Limited History for Email","EAStatusID":"4","EAReasonID":"8","EAAdviceID":"4","EAAdvice":"Moderate Fraud Risk","EARiskBandID":"3","EARiskBand":"Fraud Score 301 to 600","source_industry":"","fraud_type":"","lastflaggedon":"","dob":"","gender":"","location":"","smfriends":"","totalhits":"14","uniquehits":"1","imageurl":"","emailExists":"Not Sure","domainExists":"Yes","company":"","title":"","domainname":"legit.ru","domaincompany":"","domaincountryname":"Russian Federation","domaincategory":"","domaincorporate":"","domainrisklevel":"Moderate","domainrelevantinfo":"Valid  Domain from Russian Federation","domainrisklevelID":"3","domainrelevantinfoID":"508","domainriskcountry":"No","smlinks":[],"phone_status":"","shipforward":"","overallDigitalIdentityScore":"","emailToIpConfidence":"","emailToPhoneConfidence":"","emailToBillAddressConfidence":"","emailToShipAddressConfidence":"","emailToFullNameConfidence":"","emailToLastNameConfidence":"","ipToPhoneConfidence":"","ipToBillAddressConfidence":"","ipToShipAddressConfidence":"","ipToFullNameConfidence":"","ipToLastNameConfidence":"","phoneToBillAddressConfidence":"","phoneToShipAddressConfidence":"","phoneToFullNameConfidence":"","phoneToLastNameConfidence":"","billAddressToFullNameConfidence":"","billAddressToLastNameConfidence":"","shipAddressToBillAddressConfidence":"","shipAddressToFullNameConfidence":"","shipAddressToLastNameConfidence":""}]},"responseStatus":{"status":"success","errorCode":"0","description":""}}`))
				},
			},
			want: &Response{
				Query: &Query{
					Email:         "nigerian.prince@legit.ru",
					Count:         1,
					Created:       "2019-08-01T00:47:48Z",
					Lang:          "en-US",
					ResponseCount: 1,
					Results: []Result{{
						Country:      "RU",
						EAAdvice:     "Moderate Fraud Risk",
						EAAdviceID:   "4",
						EAReason:     "Limited History for Email",
						EAReasonID:   "8",
						EARiskBand:   "Fraud Score 301 to 600",
						EARiskBandID: "3",
						EAScore:      "500",
						EAStatusID:   "4",
						Email:        "nigerian.prince@legit.ru",
					}},
				},
				ResponseStatus: &ResponseStatus{
					Status:      "success",
					ErrorCode:   "0",
					Description: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Query With Standardized Addresses",
			fields: fields{
				opts: &ClientOpts{
					Token:       "token",
					AccountSID:  "sid",
					Endpoint:    "https://api.exists.somewhere.com/",
					Format:      JSON,
					Algorithm:   auth.HMACSHA512,
					HTTPTimeout: 0,
				},
			},
			args: args{
				email: "nigerian.prince@legit.ru",
				params: map[string]string{
					"billaddress":     "123 Any Bill Pl",
					"billcity":        "Chandler",
					"billregion":      "AZ",
					"billpostal":      "85225",
					"billcountry":     "US",
					"shipaddress":     "123 Any Bill Pl",
					"shipbillcity":    "Chandler",
					"shipbillregion":  "AZ",
					"shipbillpostal":  "85225",
					"shipbillcountry": "US",
				},
			},
			authMocks: []mockFuncAuth{
				func(auth *authorizer) {
					auth.On("RandomString", mock.Anything).Return("asdfghjkl;")
					auth.On("GetSignature", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("hmacsignature", nil)
				},
			},
			clientMocks: []mockFunc{
				func() {
					httpmock.RegisterResponder("POST",
						`=~^https://api.exists.somewhere.com/.*`,
						httpmock.NewStringResponder(200, `{"query":{"email":"nigerian.prince%40legit.ru","queryType":"EmailAgeVerification","count":1,"created":"2019-08-01T00:47:48Z","lang":"en-US","responseCount":1,"results":[{"userdefinedrecordid":"","email":"nigerian.prince@legit.ru","eName":"","emailAge":"","email_creation_days":"","domainAge":"2010-02-08T21:00:00Z","domain_creation_days":"3460","firstVerificationDate":"2018-05-30T17:28:08Z","first_seen_days":"427","lastVerificationDate":"2019-07-31T06:51:02Z","status":"ValidDomain","country":"RU","fraudRisk":"500 Moderate","EAScore":"500","EAReason":"Limited History for Email","EAStatusID":"4","EAReasonID":"8","EAAdviceID":"4","EAAdvice":"Moderate Fraud Risk","EARiskBandID":"3","EARiskBand":"Fraud Score 301 to 600","source_industry":"","fraud_type":"","lastflaggedon":"","dob":"","gender":"","location":"","smfriends":"","totalhits":"14","uniquehits":"1","imageurl":"","emailExists":"Not Sure","domainExists":"Yes","company":"","title":"","domainname":"legit.ru","domaincompany":"","domaincountryname":"Russian Federation","domaincategory":"","domaincorporate":"","domainrisklevel":"Moderate","domainrelevantinfo":"Valid  Domain from Russian Federation","domainrisklevelID":"3","domainrelevantinfoID":"508","domainriskcountry":"No","smlinks":[],"phone_status":"Invalid","shipforward":"","billriskcountry":"No","standardizedbillingaddress":"123 Any Bill Pl,Chandler,AZ,85225","standardizedshippingaddress":"123 Any Ship Pl,Chandler,AZ,85225","overallDigitalIdentityScore":"","emailToIpConfidence":"","emailToPhoneConfidence":"","emailToBillAddressConfidence":"","emailToShipAddressConfidence":"","emailToFullNameConfidence":"","emailToLastNameConfidence":"","ipToPhoneConfidence":"","ipToBillAddressConfidence":"","ipToShipAddressConfidence":"","ipToFullNameConfidence":"","ipToLastNameConfidence":"","phoneToBillAddressConfidence":"","phoneToShipAddressConfidence":"","phoneToFullNameConfidence":"","phoneToLastNameConfidence":"","billAddressToFullNameConfidence":"","billAddressToLastNameConfidence":"","shipAddressToBillAddressConfidence":"","shipAddressToFullNameConfidence":"","shipAddressToLastNameConfidence":""}]},"responseStatus":{"status":"success","errorCode":"0","description":""}}`))
				},
			},
			want: &Response{
				Query: &Query{
					Email:         "nigerian.prince@legit.ru",
					Count:         1,
					Created:       "2019-08-01T00:47:48Z",
					Lang:          "en-US",
					ResponseCount: 1,
					Results: []Result{{
						Country:                 "RU",
						EAAdvice:                "Moderate Fraud Risk",
						EAAdviceID:              "4",
						EAReason:                "Limited History for Email",
						EAReasonID:              "8",
						EARiskBand:              "Fraud Score 301 to 600",
						EARiskBandID:            "3",
						EAScore:                 "500",
						EAStatusID:              "4",
						Email:                   "nigerian.prince@legit.ru",
						StandardizedBillAddress: "123 Any Bill Pl,Chandler,AZ,85225",
						StandardizedShipAddress: "123 Any Ship Pl,Chandler,AZ,85225",
					}},
				},
				ResponseStatus: &ResponseStatus{
					Status:      "success",
					ErrorCode:   "0",
					Description: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			oa := new(authorizer)
			for _, mockSetup := range tt.authMocks {
				mockSetup(oa)
			}
			for _, mockSetup := range tt.clientMocks {
				mockSetup()
			}
			e := &Emailage{
				opts:       tt.fields.opts,
				oc:         oa,
				HTTPClient: http.Client{},
			}
			got, err := e.EmailOnlyScore(tt.args.email, tt.args.params)
			assert.Equal(t, 1, httpmock.GetTotalCallCount())
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailOnlyScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ResponseStatus, tt.want.ResponseStatus) {
				t.Errorf("EmailOnlyScore() got = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.want.Query.Results[0].EAAdviceID, got.Query.Results[0].EAAdviceID)
			assert.Equal(t, tt.want.Query.Results[0].EAReasonID, got.Query.Results[0].EAReasonID)
			assert.Equal(t, tt.want.Query.Results[0].EARiskBandID, got.Query.Results[0].EARiskBandID)
			assert.Equal(t, tt.want.Query.Results[0].EAScore, got.Query.Results[0].EAScore)
			assert.Equal(t, tt.want.Query.Results[0].EAStatusID, got.Query.Results[0].EAStatusID)
		})
	}
}

/*
 * BENCHMARK
 */
func BenchmarkEmailage_EmailOnlyScore(b *testing.B) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET",
		`=~^https://api.exists.somewhere.com/.*`,
		httpmock.NewStringResponder(200, `{"query":{"email":"nigerian.prince@legit.ru","queryType":"EmailAgeVerification","count":1,"created":"2019-08-01T00:47:48Z","lang":"en-US","responseCount":1,"results":[{"userdefinedrecordid":"","email":"nigerian.prince@legit.ru","eName":"","emailAge":"","email_creation_days":"","domainAge":"2010-02-08T21:00:00Z","domain_creation_days":"3460","firstVerificationDate":"2018-05-30T17:28:08Z","first_seen_days":"427","lastVerificationDate":"2019-07-31T06:51:02Z","status":"ValidDomain","country":"RU","fraudRisk":"500 Moderate","EAScore":"500","EAReason":"Limited History for Email","EAStatusID":"4","EAReasonID":"8","EAAdviceID":"4","EAAdvice":"Moderate Fraud Risk","EARiskBandID":"3","EARiskBand":"Fraud Score 301 to 600","source_industry":"","fraud_type":"","lastflaggedon":"","dob":"","gender":"","location":"","smfriends":"","totalhits":"14","uniquehits":"1","imageurl":"","emailExists":"Not Sure","domainExists":"Yes","company":"","title":"","domainname":"legit.ru","domaincompany":"","domaincountryname":"Russian Federation","domaincategory":"","domaincorporate":"","domainrisklevel":"Moderate","domainrelevantinfo":"Valid  Domain from Russian Federation","domainrisklevelID":"3","domainrelevantinfoID":"508","domainriskcountry":"No","smlinks":[],"phone_status":"","shipforward":"","overallDigitalIdentityScore":"","emailToIpConfidence":"","emailToPhoneConfidence":"","emailToBillAddressConfidence":"","emailToShipAddressConfidence":"","emailToFullNameConfidence":"","emailToLastNameConfidence":"","ipToPhoneConfidence":"","ipToBillAddressConfidence":"","ipToShipAddressConfidence":"","ipToFullNameConfidence":"","ipToLastNameConfidence":"","phoneToBillAddressConfidence":"","phoneToShipAddressConfidence":"","phoneToFullNameConfidence":"","phoneToLastNameConfidence":"","billAddressToFullNameConfidence":"","billAddressToLastNameConfidence":"","shipAddressToBillAddressConfidence":"","shipAddressToFullNameConfidence":"","shipAddressToLastNameConfidence":""}]},"responseStatus":{"status":"success","errorCode":"0","description":""}}`))

	for n := 0; n < b.N; n++ {
		opts := &ClientOpts{
			Format:      "json",
			Token:       "ASDASDFASDFFASDFASDFASFD",
			AccountSID:  "ASDFFASDFASDFASASDFASFASDFASF",
			Endpoint:    "https://api.exists.somewhere.com/",
			HTTPTimeout: 3 * time.Second,
			Algorithm:   auth.HMACSHA512,
		}
		client, _ := New(opts)
		client.EmailOnlyScore("nigerian.prince@legit.ru", nil)
	}
}
