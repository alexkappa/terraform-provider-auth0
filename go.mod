module github.com/alexkappa/terraform-provider-auth0

go 1.13

replace gopkg.in/auth0.v5 => github.com/sortlist/auth0-1 v0.0.0-20201125100601-156820b66f1b

require (
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/terraform-plugin-sdk v1.16.0
	gopkg.in/auth0.v5 v5.2.2
)
