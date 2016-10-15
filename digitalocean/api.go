package digitalocean

import "github.com/sqdron/squad-oauth/oauth"

type digitalocean struct{
	OAuth oauth.IAuth
}

type IDigitalOcean interface{

}

func New(clientKey, clientSecret, callbackURL string) *digitalocean {
	return &digitalocean{OAuth:OAuth(clientKey,clientSecret, callbackURL, "")}
}