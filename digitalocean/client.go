package digitalocean

import (
	"golang.org/x/oauth2"
	"github.com/digitalocean/godo"
	"fmt"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	fmt.Println(token)
	return token, nil
}

func NewClient(token string) *godo.Client {
	oauthClient := oauth2.NewClient(oauth2.NoContext, &TokenSource{
		AccessToken: token,
	})
	return godo.NewClient(oauthClient)
}
