package digitalocean


type digitalocean struct{
	OAuth IAuth
}

type IDigitalOcean interface{

}

func New(clientKey, clientSecret, callbackURL string) *digitalocean {
	return &digitalocean{OAuth:OAuth(clientKey,clientSecret, callbackURL, "")}
}