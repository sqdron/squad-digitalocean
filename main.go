package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad-digitalocean/digitalocean"
	"github.com/sqdron/squad/configurator"
)

type AuthOptions struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func main() {
	options := &AuthOptions{}
	configurator.New().ReadFromFile("./env/digitalocean.json", &options)
	do:= digitalocean.New(options.ClientID, options.ClientSecret, options.RedirectURL)
	client := squad.Client()
	client.Api.Route("oauth.digitalocean.url").Action(do.OAuth.GetAccessUrl)
	client.Api.Route("oauth.digitalocean.authorize").Action(do.OAuth.Authorize)
	client.Activate()
}
