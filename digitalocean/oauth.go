package digitalocean

import (
	"golang.org/x/oauth2"
	"github.com/sqdron/squad-oauth/oauth"
	"errors"
	"fmt"
)

const (
	authURL string = "https://cloud.digitalocean.com/v1/oauth/authorize"
	tokenURL string = "https://cloud.digitalocean.com/v1/oauth/token"
	//endpointProfile string = "https://api.digitalocean.com/v2/account"
)

type digitalOcean struct {
	config *oauth2.Config
}

func (p *digitalOcean) Name() string {
	return "digitalocean"
}

func (p *digitalOcean) OpenSession(state string) (*oauth.Session, error) {
	session := &oauth.Session{}
	session.AuthURL = p.config.AuthCodeURL(state)
	session.ID = state
	return session, nil
}

func (p *digitalOcean) Authorize(session *oauth.Session, code string) (string, error) {
	fmt.Println(session)
	fmt.Println(p.config)
	fmt.Println(session.ID)
	token, err := p.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return "", err
	}

	if !token.Valid() {
		return "", errors.New("Invalid token received from provider")
	}

	session.AccessToken = token.AccessToken
	session.RefreshToken = token.RefreshToken
	session.ExpiresAt = token.Expiry
	return token.AccessToken, err
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

func DigitalOcean(clientKey, clientSecret, callbackURL string, scopes ...string) oauth.IProvider {
	config := &oauth2.Config{
		ClientID:     clientKey,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL},
		Scopes: scopes}

	p := &digitalOcean{config}
	return oauth.NewProvider(p)
}
