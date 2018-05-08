package emailage

import (
	"testing"
)

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

// same checks to see if the two given service
// errors are the same
func same(se1, se2 *serviceError) bool {
	if se1.StatusCode != se2.StatusCode {
		return false
	}
	if se1.CallCode != se2.CallCode {
		return false
	}
	if se1.Message != se2.Message {
		return false
	}
	return true
}

func Test_newServiceError(t *testing.T) {
	type args struct {
		msg    string
		code   string
		status int
	}
	tests := []struct {
		name string
		args args
		want *serviceError
	}{
		{
			name: "successful service error",
			args: args{
				msg:    "some error message",
				code:   "some caller",
				status: 503,
			},
			want: &serviceError{
				Message:    "some error message",
				CallCode:   "some caller",
				StatusCode: 503,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := newServiceError(tt.args.msg, tt.args.code, tt.args.status); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("newServiceError() = %v, want %v", got, tt.want)
			// }
			if got := newServiceError(tt.args.msg, tt.args.code, tt.args.status); !same(got, tt.want) {
				t.Errorf("newServiceError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceError_Error(t *testing.T) {
	type fields struct {
		Message    string
		CallCode   string
		StatusCode int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "error message check",
			fields: fields{
				Message:    "error message",
				CallCode:   "caller",
				StatusCode: 503,
			},
			want: "error message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &serviceError{
				Message:    tt.fields.Message,
				CallCode:   tt.fields.CallCode,
				StatusCode: tt.fields.StatusCode,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("serviceError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
