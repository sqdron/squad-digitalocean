package digitalocean

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func Test_Authenticate(t *testing.T) {
	Convey("Authenticate", t, func() {
		provider := DigitalOcean("digitalocean_key", "secret", "http://test.com/auth", "read", "write")
		session, err :=provider.OpenSession("test_state")
		So(err, ShouldBeNil)
		So(session.AuthURL, ShouldContainSubstring, "cloud.digitalocean.com/v1/oauth/authorize")
		So(session.AuthURL, ShouldContainSubstring, fmt.Sprintf("client_id=%s", "digitalocean_key"))
		So(session.AuthURL, ShouldContainSubstring, "state=test_state")
		So(session.AuthURL, ShouldContainSubstring, "scope=read")
	})
}