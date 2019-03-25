# Auth0 Go SDK

[![GoDoc](https://godoc.org/gopkg.in/auth0.v1?status.svg)](https://godoc.org/gopkg.in/auth0.v1)
[![wercker status](https://app.wercker.com/status/f2c3f70b3219eada66488b8c527f19f9/s/master "wercker status")](https://app.wercker.com/project/byKey/f2c3f70b3219eada66488b8c527f19f9)
[![Maintainability](https://api.codeclimate.com/v1/badges/3610191501844db862e8/maintainability)](https://codeclimate.com/github/yieldr/go-auth0/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/3610191501844db862e8/test_coverage)](https://codeclimate.com/github/yieldr/go-auth0/test_coverage)

## Documentation

You can find this library documentation in this [page](https://godoc.org/gopkg.in/auth0.v1).

For more information about [auth0](http://auth0.com/) check their [documentation page](http://docs.auth0.com/)

## Management API Client

The Auth0 Management API is meant to be used by back-end servers or trusted parties performing administrative tasks. Generally speaking, anything that can be done through the Auth0 dashboard (and more) can also be done through this API.

Initialize your client class with an API v2 token and a domain.

```go
import "gopkg.in/auth0.v1/management"

m, err := management.New("<auth0-domain>", "<auth0-client-id>", "<auth0-client-secret>")
if err != nil {
	// handle err
}
```

With an authenticated management client we can now interact with the Auth0 Management API.

```go
c := &Client{
	Name: "Client Name",
	Description: "Long description of client",
}

err = m.Client.Create(c)
if err != nil {
	// handle err
}
```

Following is a list of supported Auth0 resources.

- [x] [Clients (Applications)](https://auth0.com/docs/api/management/v2#!/Clients/get_clients)
- [x] [Client Grants](https://auth0.com/docs/api/management/v2#!/Client_Grants/get_client_grants)
- [x] [Connections](https://auth0.com/docs/api/management/v2#!/Connections/get_connections)
- [x] [Custom Domains](https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains)
- [ ] [Device Credentials](https://auth0.com/docs/api/management/v2#!/Device_Credentials/get_device_credentials)
- [x] [Grants](https://auth0.com/docs/api/management/v2#!/Grants/get_grants)
- [x] [Logs](https://auth0.com/docs/api/management/v2#!/Logs/get_logs)
- [x] [Resource Servers (APIs)](https://auth0.com/docs/api/management/v2#!/Resource_Servers/get_resource_servers)
- [x] [Rules](https://auth0.com/docs/api/management/v2#!/Rules/get_rules)
- [x] [Rules Configs](https://auth0.com/docs/api/management/v2#!/Rules_Configs/get_rules_configs)
- [ ] [User Blocks](https://auth0.com/docs/api/management/v2#!/User_Blocks/get_user_blocks)
- [x] [Users](https://auth0.com/docs/api/management/v2#!/Users/get_users)
- [ ] [Users By Email](https://auth0.com/docs/api/management/v2#!/Users_By_Email/get_users_by_email)
- [ ] [Blacklists](https://auth0.com/docs/api/management/v2#!/Blacklists/get_tokens)
- [x] [Email Templates](https://auth0.com/docs/api/management/v2#!/Email_Templates/get_email_templates_by_templateName)
- [x] [Emails](https://auth0.com/docs/api/management/v2#!/Emails/get_provider)
- [ ] [Guardian](https://auth0.com/docs/api/management/v2#!/Guardian/get_factors)
- [ ] [Jobs](https://auth0.com/docs/api/management/v2#!/Jobs/get_jobs_by_id)
	- [ ] [Get](https://auth0.com/docs/api/management/v2#!/Jobs/get_jobs_by_id)
	- [ ] [Get Errors](https://auth0.com/docs/api/management/v2#!/Jobs/get_errors)
	- [ ] [Get Results](https://auth0.com/docs/api/management/v2#!/Jobs/get_results)
	- [ ] [Export Users](https://auth0.com/docs/api/management/v2#!/Jobs/post_users_exports)
	- [ ] [Import Users](https://auth0.com/docs/api/management/v2#!/Jobs/post_users_imports)
	- [x] [Verification Email](https://auth0.com/docs/api/management/v2#!/Jobs/post_verification_email)
- [x] [Stats](https://auth0.com/docs/api/management/v2#!/Stats/get_active_users)
- [x] [Tenants](https://auth0.com/docs/api/management/v2#!/Tenants/get_settings)
- [x] [Tickets](https://auth0.com/docs/api/management/v2#!/Tickets/post_email_verification)
