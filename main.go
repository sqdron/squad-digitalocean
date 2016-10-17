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
	cloud:= digitalocean.New(options.ClientID, options.ClientSecret, options.RedirectURL)

	client := squad.Client()
	client.Api.Route("oauth.digitalocean.url").Action(cloud.OAuth.GetAccessUrl)
	client.Api.Route("oauth.digitalocean.authorize").Action(cloud.OAuth.Authorize)
	client.Api.Route("oauth.digitalocean.refresh").Action(cloud.OAuth.Refresh)

	client.Api.Route("cloud.digitalocean.unit.create").Action(cloud.Unit.Create)
	client.Api.Route("cloud.digitalocean.unit.list").Action(cloud.Unit.List)
	client.Api.Route("cloud.digitalocean.image.list").Action(cloud.Image.List)
	client.Api.Route("cloud.digitalocean.plan.list").Action(cloud.Plan.List)

	client.Api.Route("cloud.digitalocean.key.list").Action(cloud.Key.List)
	client.Api.Route("cloud.digitalocean.key.create").Action(cloud.Key.Create)
	client.Activate()
}
