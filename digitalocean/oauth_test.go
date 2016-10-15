package digitalocean

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func Test_Authenticate(t *testing.T) {
	Convey("Authenticate", t, func() {
		oauth := OAuth("digitalocean_key", "secret", "http://test.com/auth", "read", "write")
		url, err :=oauth.GetAccessUrl()
		So(err, ShouldBeNil)
		So(url, ShouldContainSubstring, "cloud.digitalocean.com/v1/oauth/authorize")
		So(url, ShouldContainSubstring, fmt.Sprintf("client_id=%s", "digitalocean_key"))
		So(url, ShouldContainSubstring, "scope=read")
	})
}