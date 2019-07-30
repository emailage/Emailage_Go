package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

/*
	Test Objects
*/
type MockedRandom struct {
	mock.Mock
}

func (m *MockedRandom) NewRandom() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

type mockFunc func()

/*
	TESTS
*/
func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *oauth1
		wantErr bool
	}{
		{
			name: "empty constructor",
			want: &oauth1{},
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

func Test_oauth1_GetRandomString(t *testing.T) {
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
			want:    "b",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			random := new(MockedRandom)
			for _, mockSetup := range tt.mocks {
				mockSetup()
			}
			oauth := &oauth1{
				rnd: random,
				chs: []rune{'b'},
			}
		})
	}
}
