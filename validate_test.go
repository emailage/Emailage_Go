package emailage

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid email",
			args: args{
				email: "nigerian.prince@legit.ru",
			},
			want: true,
		},
		{
			name: "invalid email",
			args: args{
				email: "nigerian.princelegit.ru",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.email); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid IP",
			args: args{
				ip: "192.168.0.1",
			},
			want: true,
		},
		{
			name: "invalid IP",
			args: args{
				ip: "192.168.0.1999",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIP(tt.args.ip); got != tt.want {
				t.Errorf("IsIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv4(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid I4P",
			args: args{
				ip: "192.168.0.1",
			},
			want: true,
		},
		{
			name: "invalid IP4",
			args: args{
				ip: "192.168.0.1999",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPv4(tt.args.ip); got != tt.want {
				t.Errorf("IsIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv6(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid IP6",
			args: args{
				ip: "2600:100a:b019:8b06:b47d:a3c1:12a8:4f8c",
			},
			want: true,
		},
		{
			name: "invalid IP6",
			args: args{
				ip: "2600:100a:b019:8b06:b47d:a3c1:12a8:4f8cx",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPv6(tt.args.ip); got != tt.want {
				t.Errorf("IsIPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}
