package emailage

import (
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/mrjones/oauth"
)

func TestClientOpts_validate(t *testing.T) {
	type fields struct {
		Token      string
		AccountSID string
		Endpoint   string
		HTTP       *http.Client
		Format     string
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
				Format:     "json",
			},
			wantErr: false,
		},
		{
			name: "missing token",
			fields: fields{
				AccountSID: "accountsid",
				Endpoint:   "someendpoint",
				Format:     "json",
			},
			wantErr: true,
		},
		{
			name: "missing accountSID",
			fields: fields{
				Token:    "token",
				Endpoint: "someendpoint",
				Format:   "json",
			},
			wantErr: true,
		},
		{
			name: "missing endpoint",
			fields: fields{
				Token:      "token",
				AccountSID: "accountsid",
				Format:     "json",
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
				HTTP:       tt.fields.HTTP,
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
				oc:   oauth.NewConsumer("token", "accountsid", oauth.ServiceProvider{}),
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailage_EmailOnlyScore(t *testing.T) {
	type fields struct {
		opts *ClientOpts
		oc   *oauth.Consumer
	}
	type args struct {
		email  string
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
			got, err := e.EmailOnlyScore(tt.args.email, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Emailage.EmailOnlyScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Emailage.EmailOnlyScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailage_IPAddressOnlyScore(t *testing.T) {
	type fields struct {
		opts *ClientOpts
		oc   *oauth.Consumer
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
		oc   *oauth.Consumer
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
		code int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "invalid input",
			args: args{
				code: 400,
			},
			want: "Invalid input parameter. Error message should indicate which one",
		},
		{
			name: "not found",
			args: args{
				code: 9999,
			},
			want: "",
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
