module github.com/alexkappa/terraform-provider-auth0

go 1.15

require (
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/terraform-plugin-sdk v1.16.1
	gopkg.in/auth0.v5 v5.15.0
)

replace gopkg.in/auth0.v5 => github.com/Abacus-Insights/auth0 v1.3.1-0.20210512201735-a335ec727e5e
