package digitalocean

import (
	"golang.org/x/oauth2"
	"github.com/sqdron/squad-oauth/oauth"
	"errors"
	"github.com/sqdron/squad/util"
	"fmt"
)

const (
	authURL string = "https://cloud.digitalocean.com/v1/oauth/authorize"
	tokenURL string = "https://cloud.digitalocean.com/v1/oauth/token"
)





type digitalOcean struct {
	config *oauth2.Config
}

func (p *digitalOcean) Name() string {
	return "digitalocean"
}

func (p *digitalOcean) GetAccessUrl() (string, error) {
	return p.config.AuthCodeURL(util.GenerateString(7)), nil
}

func (p *digitalOcean) Authorize(s *oauth.Session) (*oauth.Session, error) {
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

func (p *digitalOcean) Refresh(s *oauth.Session) (*oauth.Session, error) {
	token := &oauth2.Token{RefreshToken: s.RefreshToken}
	ts := p.config.TokenSource(oauth2.NoContext, token)
	newToken, err := ts.Token()
	if err != nil {
		return nil, err
	}
	s.AccessToken = newToken.AccessToken
	s.RefreshToken = newToken.RefreshToken
	s.ExpiresAt = newToken.Expiry
	return s, nil
}

func (p *digitalOcean) RefreshTokenAvailable() bool {
	return true
}

func OAuth(clientKey, clientSecret, callbackURL string, scopes ...string) oauth.IAuth {
	config := &oauth2.Config{
		ClientID:     clientKey,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL},
		Scopes: scopes}
	fmt.Println(config)
	p := &digitalOcean{config}
	return p
}
