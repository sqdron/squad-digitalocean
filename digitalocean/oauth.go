package digitalocean

import (
	"golang.org/x/oauth2"
	"github.com/sqdron/squad-oauth/oauth"
	"errors"
	"time"
	"github.com/sqdron/squad/util"
)

const (
	authURL string = "https://cloud.digitalocean.com/v1/oauth/authorize"
	tokenURL string = "https://cloud.digitalocean.com/v1/oauth/token"
)

type Session struct {
	Code         string
	State        string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type IAuth interface {
	GetAccessUrl() (string, error)
	Authorize(s *Session) (*Session, error)
}

type digitalOcean struct {
	config *oauth2.Config
}

func (p *digitalOcean) Name() string {
	return "digitalocean"
}

func (p *digitalOcean) GetAccessUrl() (string, error) {
	return p.config.AuthCodeURL(util.GenerateString(7)), nil
}

func (p *digitalOcean) Authorize(s *Session) (*Session, error) {
	token, err := p.config.Exchange(oauth2.NoContext, s.Code)
	if err != nil {
		return nil, err
	}

	if !token.Valid() {
		return nil, errors.New("Invalid token received from provider")
	}

	s.AccessToken = token.AccessToken
	s.RefreshToken = token.RefreshToken
	s.ExpiresAt = token.Expiry
	return s, err
}

func (p *digitalOcean) GetAccount(s *oauth.Session) (*oauth.Account, error) {
	return nil, errors.New("Not implemented")
}

func (p *digitalOcean) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{RefreshToken: refreshToken}
	ts := p.config.TokenSource(oauth2.NoContext, token)
	newToken, err := ts.Token()
	if err != nil {
		return nil, err
	}
	return newToken, err
}

func (p *digitalOcean) RefreshTokenAvailable() bool {
	return true
}

func OAuth(clientKey, clientSecret, callbackURL string, scopes ...string) IAuth {
	config := &oauth2.Config{
		ClientID:     clientKey,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL},
		Scopes: scopes}

	p := &digitalOcean{config}
	return p
}
