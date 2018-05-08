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
