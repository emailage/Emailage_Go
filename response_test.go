package emailage

import "testing"

func TestFraudCodeLookup(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "card not present",
			args: args{
				code: 1,
			},
			want: "Card Not Present Fraud",
		},
		{
			name: "not found",
			args: args{
				code: 1000,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FraudCodeLookup(tt.args.code); got != tt.want {
				t.Errorf("FraudCodeLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneOwnerMatchLookup(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "full match",
			args: args{
				s: "Y",
			},
			want: "Full Match",
		},
		{
			name: "not found",
			args: args{
				s: "A",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhoneOwnerMatchLookup(tt.args.s); got != tt.want {
				t.Errorf("PhoneOwnerMatchLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRiskLevelLookup(t *testing.T) {
	type args struct {
		level int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "very high",
			args: args{
				level: 1,
			},
			want: "Very High",
		},
		{
			name: "not found",
			args: args{
				level: 1000,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RiskLevelLookup(tt.args.level); got != tt.want {
				t.Errorf("RiskLevelLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRiskReasonLookup(t *testing.T) {
	type args struct {
		reason int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "moderate risk",
			args: args{
				reason: 301,
			},
			want: "Moderate Risk",
		},
		{
			name: "not found",
			args: args{
				reason: 9999,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RiskReasonLookup(tt.args.reason); got != tt.want {
				t.Errorf("RiskReasonLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
