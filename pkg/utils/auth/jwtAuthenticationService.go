package auth

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	key "onqlave.io/onqlave.core/pkg/utils/auth/key"
)

type JwtAuthenticationService interface {
	CreateSignupLink(context context.Context, url string, emailAddress string) (string, error)
	VerifyToken(context context.Context, token string) (*key.Token, error)
	KeyGenService() key.KeyGenService
}

func NewJwtAuthenticationService(client *auth.Client, keyGenService key.KeyGenService) JwtAuthenticationService {
	return &jwtAuthenticationService{client: client, keyGenService: keyGenService}
}

type jwtAuthenticationService struct {
	client        *auth.Client
	keyGenService key.KeyGenService
}

func (s *jwtAuthenticationService) KeyGenService() key.KeyGenService {
	return s.keyGenService
}

func (s *jwtAuthenticationService) CreateSignupLink(context context.Context, url string, emailAddress string) (string, error) {
	settings := &auth.ActionCodeSettings{
		URL:               url,
		HandleCodeInApp:   true,
		DynamicLinkDomain: "coolapp.page.link",
	}
	link, err := s.client.EmailSignInLink(context, emailAddress, settings)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (s *jwtAuthenticationService) VerifyToken(context context.Context, token string) (*key.Token, error) {
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "missing key in request header")
	}
	//"https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
	rawToken := strings.Replace(token, "Bearer ", "", -1)
	if s.keyGenService.IsCustomToken(rawToken) {
		//custom token
		t, err := s.keyGenService.Decrypt(rawToken)
		if err != nil {
			return nil, &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid onqlave token",
				Internal: err,
			}
		}
		return &t, nil
	} else {
		jwt, err := s.client.VerifyIDToken(context, rawToken)
		if err != nil {
			return nil, &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid jwt token",
				Internal: err,
			}
		}
		return &key.Token{AuthTime: jwt.AuthTime,
			Issuer:   jwt.Issuer,
			IssuedAt: jwt.IssuedAt,
			Audience: jwt.Audience,
			Expires:  jwt.Expires,
			Subject:  jwt.Subject,
			UID:      jwt.UID,
			Claims:   jwt.Claims,
		}, nil
	}
}
