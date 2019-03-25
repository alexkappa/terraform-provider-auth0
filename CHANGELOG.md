## v0.1.16 (March 25, 2019)

NOTES:

* Added contributing, code of conduct, issue templates to the project.

## v0.1.15 (March 25, 2019)

FEATURES:

* **New Resource:** auth0_tenant ([#79](https://github.com/yieldr/terraform-provider-auth0/pull/79))

ENHANCEMENTS:

* resource/auth0_connection: `enabled_clients` will suppress diff if the list order is different.
* resource/auth0_connection: set `client_secret` to sensitive ([#91](https://github.com/yieldr/terraform-provider-auth0/pull/91)).
* resource/auth0_resource_server: introduce `token_lifetime_for_web` ([#84](https://github.com/yieldr/terraform-provider-auth0/pull/84)).

NOTES:

* Re-imported Auth0 SDK from `gopkg.in/auth0.v1`.
