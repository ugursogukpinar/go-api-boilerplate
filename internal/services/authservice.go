package services

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/config"
)

type AuthService struct {
	driver *auth.Service
}

type customLogger struct {
	l *slog.Logger
}

func (cl *customLogger) Logf(format string, args ...interface{}) {
	cl.l.Debug(format, args...)
}

func NewAuthService(cfg *config.Config, appLogger *slog.Logger) *AuthService {
	options := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) { // secret key for JWT
			return cfg.App.JWTSecret, nil
		}),
		TokenDuration:   time.Minute * 5, // token expires in 5 minutes
		CookieDuration:  time.Hour * 24,  // cookie expires in 1 day and will enforce re-login
		Issuer:          cfg.App.Name,
		URL:             "",
		AvatarStore:     avatar.NewNoOp(),
		AvatarRoutePath: "/",
		Logger:          &customLogger{appLogger},
		JWTHeaderKey:    "Authorization",
	}

	driver := auth.NewService(options)

	return &AuthService{
		driver: driver,
	}
}

func (s *AuthService) GetAuthRoutes() http.Handler {
	s.driver.AddDirectProvider("local", provider.CredCheckerFunc(s.LocalAuthenticator))
	authRoutes, _ := s.driver.Handlers()
	return authRoutes
}

func (s *AuthService) LocalAuthenticator(username, password string) (bool, error) {
	if username == "us" {
		return true, nil
	}

	return false, errors.New("invalid username or password")
}
