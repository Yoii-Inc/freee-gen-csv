package token

import (
	"golang.org/x/oauth2"
)

type MyFunc func(*oauth2.Token) error

type MyTokenSource struct {
	Src oauth2.TokenSource
	F   MyFunc
}

func (s *MyTokenSource) Token() (*oauth2.Token, error) {
	t, err := s.Src.Token()
	if err != nil {
		return nil, err
	}
	if err = s.F(t); err != nil {
		return t, err
	}
	return t, nil
}
