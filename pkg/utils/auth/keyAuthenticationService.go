package auth

import (
	"context"
)

type KeyAuthenticationService interface {
	VerifyKey(context context.Context, key string) (bool, error)
}

func NewKeyAuthenticationService() KeyAuthenticationService {
	return &keyAuthenticationService{}
}

type keyAuthenticationService struct {
}

func (s *keyAuthenticationService) VerifyKey(context context.Context, key string) (bool, error) {
	return true, nil
}
