module github.com/alexkappa/terraform-provider-auth0

go 1.16

require (
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/terraform-plugin-sdk v1.16.1
	github.com/kisielk/errcheck v1.6.0 // indirect
	golang.org/x/mod v0.5.0 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/grpc v1.31.1 // indirect
	gopkg.in/auth0.v5 v5.19.2
)

replace gopkg.in/auth0.v5 v5.19.2 => github.com/dev-usa/auth0 v1.3.1-0.20210824211100-82f62182b30e
