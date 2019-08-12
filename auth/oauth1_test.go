package auth

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
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
			got := oauth.RandomString(tt.args.length)
			if reflect.TypeOf(got).String() != "string" {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			assert.Equal(t, got, "b")
		})
	}
}

func Test_Oauth1_HmacEncrypt(t *testing.T) {
	type fields struct {
		seed int64
		chs  []rune
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
		{
			name: "Simple Encrypt HMACSHA256",
			fields: fields{
				chs:  []rune(c),
				seed: 42,
			},
			args: args{
				algorithm: HMACSHA256,
				key:       "key",
				value:     "value",
			},
			want:    []byte{144, 251, 252, 241, 94, 116, 163, 107, 137, 219, 219, 42, 114, 29, 154, 236, 255, 223, 221, 220, 92, 131, 226, 127, 117, 146, 89, 79, 113, 147, 36, 129},
			wantErr: false,
		},
		{
			name: "Simple Encrypt HMACSHA384",
			fields: fields{
				chs:  []rune(c),
				seed: 42,
			},
			args: args{
				algorithm: HMACSHA384,
				key:       "key",
				value:     "value",
			},
			want:    []byte{169, 197, 31, 177, 11, 18, 216, 255, 110, 36, 71, 211, 206, 19, 56, 131, 231, 39, 213, 173, 9, 180, 239, 242, 48, 52, 98, 235, 171, 238, 245, 28, 31, 208, 14, 142, 177, 144, 36, 99, 2, 147, 62, 208, 98, 175, 0, 20},
			wantErr: false,
		},
		{
			name: "Simple Encrypt HMACSHA512",
			fields: fields{
				chs:  []rune(c),
				seed: 42,
			},
			args: args{
				algorithm: HMACSHA512,
				key:       "key",
				value:     "value",
			},
			want:    []byte{134, 149, 29, 199, 101, 190, 249, 95, 148, 116, 102, 156, 209, 141, 247, 112, 93, 153, 174, 71, 234, 62, 118, 162, 202, 76, 34, 247, 22, 86, 244, 46, 166, 110, 58, 205, 200, 152, 201, 63, 71, 80, 9, 250, 89, 157, 11, 184, 59, 213, 54, 95, 54, 169, 203, 146, 197, 112, 112, 143, 141, 229, 250, 232},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(tt.fields.seed)
			random := &Erandom{}
			oa := &Oauth1{
				rnd: random,
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

func TestOauth1_GetSignature(t *testing.T) {
	type fields struct {
		seed int64
		chs  []rune
	}
	type args struct {
		fullUrl string
		method  RequestMethod
		hmacsha HMACSHA
		token   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Simple Signature",
			fields: fields{
				seed: 42,
				chs:  []rune(c),
			},
			args: args{
				fullUrl: "http://test.com/test/",
				method:  GET,
				hmacsha: HMACSHA512,
				token:   "token",
			},
			want:    "bhcgBcS9bVIUVTN1k5IvKT%2FW2q%2FlniGrdV73HIi6ALLQAG2WNPd5%2B503PuhFCVK%2BHswtXSB1TtMMFEQVCVLPlg%3D%3D",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		rand.Seed(tt.fields.seed)
		random := &Erandom{}
		t.Run(tt.name, func(t *testing.T) {
			oa := &Oauth1{
				rnd: random,
				chs: tt.fields.chs,
			}
			got, err := oa.GetSignature(tt.args.fullUrl, tt.args.method, tt.args.hmacsha, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSignature() got = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
 BENCHMARKS
*/
func BenchmarkOauth1_GetRandomString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		oauth, _ := New()
		oauth.RandomString(10)
	}
}
func BenchmarkOauth1_HmacEncrypt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		oauth, _ := New()
		oauth.HmacEncrypt(RandStringBytesMaskImprSrcSB(1000), "ASDFASASDFASDFASDADFS", HMACSHA512)
	}
}
func BenchmarkOauth1_GetSignature(b *testing.B) {
	for n := 0; n < b.N; n++ {
		oauth, _ := New()
		oauth.GetSignature("http://test.com/test/", GET, HMACSHA512, "ADSFASDFASDFFASDFASDFFAF")
	}
}

// Utilities
// borrowed from this stack overflow article: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
