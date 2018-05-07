package management

import "os"

var m *Management

var (
	Auth0Domain       = os.Getenv("AUTH0_DOMAIN")
	Auth0ClientID     = os.Getenv("AUTH0_CLIENT_ID")
	Auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
)

func init() {
	var err error
	m, err = New(Auth0Domain, Auth0ClientID, Auth0ClientSecret)
	if err != nil {
		panic(err)
	}
}
