package auth

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

/*
	Test Objects
*/

type mockFunc func()

/*
	TESTS
*/
func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *Oauth1
		wantErr bool
	}{
		{
			name: "empty constructor",
			want: &Oauth1{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert := assert.New(t)
			assert.NotNil(got.chs)
			assert.NotNil(got.rnd)
		})
	}
}

func Test_Oauth1_GetRandomString(t *testing.T) {
	type fields struct {
		rnd *Erandom
		chs []rune
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mocks   []mockFunc
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				length: 1,
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			random, _ := NewRandom()
			for _, mockSetup := range tt.mocks {
				mockSetup()
			}
			oauth := &Oauth1{
				rnd: random,
				chs: []rune("bbbbbbbbbbbbbbbbbbbbbbbbb"),
			}
			got := oauth.GetRandomString(tt.args.length)
			if reflect.TypeOf(got).String() != "string" {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			assert.Equal(t, got, "b")
		})
	}
}

func Test_Oauth1_HmacEncrypt(t *testing.T) {
	type fields struct {
		rnd *Erandom
		chs []rune
	}
	type args struct {
		value     string
		key       string
		algorithm HMACSHA
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oa := &Oauth1{
				rnd: tt.fields.rnd,
				chs: tt.fields.chs,
			}
			got, err := oa.HmacEncrypt(tt.args.value, tt.args.key, tt.args.algorithm)
			if (err != nil) != tt.wantErr {
				t.Errorf("HmacEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HmacEncrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
