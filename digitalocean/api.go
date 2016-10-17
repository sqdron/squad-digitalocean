package digitalocean

import (
	"github.com/sqdron/squad-oauth/oauth"
	"github.com/sqdron/squad-cloud/cloud"
)

type digitalocean struct {
	OAuth oauth.IAuth
	Unit  cloud.ICloudUnit
	Image cloud.IImage
	Plan  cloud.ICloudPlan
	Key   cloud.ICloudKey
}

type IDigitalOcean interface {

}

func New(clientKey, clientSecret, callbackURL string) *digitalocean {
	return &digitalocean{OAuth:OAuth(clientKey, clientSecret, callbackURL, "read", "write"),
		Unit:UnitService(),
		Image:ImageService(),
		Plan:PlanService(),
		Key:KeysService()}
}
