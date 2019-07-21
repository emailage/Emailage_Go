package auth

import "strings"

const (
	HMACSHA256 = iota
	HMACSHA384
	HMACSHA512
)

const (
	c = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.~"
)

type oauth1 struct {
	rnd *Erandom
	chs []rune
}

// New creates a new value of type pointer oauth1
func New() (*oauth1, error) {

	ra, err := NewRandom()
	if err != nil {
		return nil, err
	}
	ch := []rune(c)

	oa := &oauth1{
		rnd: ra,
		chs: ch,
	}

	return oa, nil
}

func (oa *oauth1) GetRandomString(length int) {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune(oa.chs[oa.rnd.Random.Intn(25)])
	}

}
