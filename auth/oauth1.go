package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"net/url"
	"strings"
)

type HMACSHA string

const (
	HMACSHA256 HMACSHA = "HMAC-SHA256"
	HMACSHA384 HMACSHA = "HMAC-SHA384"
	HMACSHA512 HMACSHA = "HMAC-SHA512"
)

type RequestMethod uint8

const (
	GET RequestMethod = iota
	POST
)

const (
	c = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.~"
)

type Authorizer interface {
	GetSignature(string, RequestMethod, HMACSHA, string) (string, error)
	RandomString(length int) string
}

type Oauth1 struct {
	rnd *Erandom
	chs []rune
}

// New creates a new value of type pointer oauth1
func New() (*Oauth1, error) {

	ra, err := NewRandom()
	if err != nil {
		return nil, err
	}
	ch := []rune(c)

	oa := &Oauth1{
		rnd: ra,
		chs: ch,
	}

	return oa, nil
}

func (oa *Oauth1) RandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune(oa.chs[oa.rnd.Random.Intn(25)])
	}
	return sb.String()
}

// code adapted from this: https://golangcode.com/generate-sha256-hmac/
func (oa *Oauth1) HmacEncrypt(value string, key string, algorithm HMACSHA) ([]byte, error) {
	var a func() hash.Hash
	switch algorithm {
	case HMACSHA256:
		a = sha256.New
		break
	case HMACSHA384:
		a = sha512.New384
		break
	case HMACSHA512:
		a = sha512.New
		break
	default:
		return nil, fmt.Errorf("Unknown algorithm: %s", algorithm)
	}

	var h = hmac.New(a, []byte(key))
	h.Write([]byte(value))

	return h.Sum(nil), nil
}

func (oa *Oauth1) ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (oa *Oauth1) GetSignature(fullUrl string, method RequestMethod, hmacsha HMACSHA, token string) (string, error) {

	hs, err := oa.HmacEncrypt(fullUrl, token+"&", hmacsha)
	if err != nil {
		return "", err
	}

	var b64 = oa.ToBase64(hs)
	var r = url.QueryEscape(b64)
	return r, nil
}
