## 0.17.0

BUG FIXES:

* resource/auth0_client: Fix handling of `mobile` attributes ([#305](https://github.com/alexkappa/terraform-provider-auth0/pull/305)).

## 0.16.1

BUG FIXES:

* resource/auth0_connection: Fix `validation` field to correctly define a `min` and `max` `username` lengths ([#258](https://github.com/alexkappa/terraform-provider-auth0/pull/258)).

## 0.16.0

FEATURES:

* resource/auth0_log_stream: Support for LogStreams ([#270](https://github.com/alexkappa/terraform-provider-auth0/pull/270)).

NOTES:

* Upgrade to `gopkg.in/auth0.v5` (`v5.2.2`)

## 0.15.2

ENHANCEMENTS:

* resource/auth0_connection: Support for additional fields for `samlp` connection ([#268](https://github.com/alexkappa/terraform-provider-auth0/pull/268)).

## 0.15.1

ENHANCEMENTS:

* resource/auth0_hook: Fix documentation rendering issue.

## 0.15.0 (September 24, 2020)

ENHANCEMENTS:

* resource/auth0_connection: Support for `oauth2` connection options ([#267](https://github.com/alexkappa/terraform-provider-auth0/pull/267)).

## 0.14.0 (August 24, 2020)

ENHANCEMENTS:

* resource/auth0_client: Validate `description` length to be 140 characters ([#260](https://github.com/alexkappa/terraform-provider-auth0/pull/260))
* resource/auth0_tenant: Validate `session_lifetime` to be > 1 ([#229](https://github.com/alexkappa/terraform-provider-auth0/pull/229))

BUG FIXES:

* resource/auth0_connection: Fix `debug` field in SAML connections and change iterating from a `Set` to a `List` ([#261](https://github.com/alexkappa/terraform-provider-auth0/pull/261))

NOTES:

Fixed typo in the documentation ([#263](https://github.com/alexkappa/terraform-provider-auth0/pull/263))

## 0.13.0 (August 17, 2020)

FEATURES:

* resource/auth0_client: support for setting `refresh_token` ([#255](https://github.com/alexkappa/terraform-provider-auth0/pull/255))

## 0.12.2 (July 07, 2020)

BUG FIXES:

* resource/auth0_client: `samlp` addon is now correctly translated to camel case ([#226](https://github.com/alexkappa/terraform-provider-auth0/issues/226))

## 0.12.1 (July 03, 2020)

BUG FIXES:

* resource/auth0_rule_config: forces new resource when `key` has change ([#246](https://github.com/alexkappa/terraform-provider-auth0/pull/246)).

## 0.12.0 (July 01, 2020)

FEATURES:

* resource/auth0_connection: support for the `saml` connection strategy ([#244](https://github.com/alexkappa/terraform-provider-auth0/pull/244)).

## 0.11.0 (June 04, 2020)

BUG FIXES:

* resource/auth0_connection: inconsistent state after applying changes to `options` ([#237](https://github.com/alexkappa/terraform-provider-auth0/pull/237)).
* resource/auth0_client_grant: force a new resource if `audience` or `client_id` has changed ([#239](https://github.com/alexkappa/terraform-provider-auth0/pull/239), [#186](https://github.com/alexkappa/terraform-provider-auth0/pull/186)).

## 0.10.3 (June 02, 2020)

BUG FIXES:

* resource/auth0_hook: allow creating hooks with the `send-phone-message` trigger ([#240](https://github.com/alexkappa/terraform-provider-auth0/pull/240)).

## 0.10.2 (May 19, 2020)

BUG FIXES:

* resource/auth0_user: fix issue causing an `Error: unexpected end of JSON input`.

## 0.10.1 (May 13, 2020)

BUG FIXES:

* resource/auth0_connection: migrate state for `strategy_version` causing an `Error: a number is required` error.

## 0.10.0 (May 11, 2020)

FEATURES:

* resource/auth0_connection: support for the `apple` connection strategy ([#216](https://github.com/alexkappa/terraform-provider-auth0/pull/216)).
* resource/auth0_connection: support for the `facebook` connection strategy ([#221](https://github.com/alexkappa/terraform-provider-auth0/pull/221)).
* resource/auth0_connection: support for the `linkedin` connection strategy ([#222](https://github.com/alexkappa/terraform-provider-auth0/pull/222)).
* resource/auth0_connection: support for the `oidc` connection strategy ([#215](https://github.com/alexkappa/terraform-provider-auth0/pull/215))

## 0.9.3 (April 24, 2020)

BUG FIXES: 

* resource/auth0_hook: avoid sending `trigger_id` during updates ([#210](https://github.com/alexkappa/terraform-provider-auth0/pull/210)).

## 0.9.2 (April 20, 2020)

BUG FIXES: 

* resource/auth0_connection: `configuration` properties are now write-only ([#208](https://github.com/alexkappa/terraform-provider-auth0/pull/208)).

## 0.9.1 (April 16, 2020)

BUG FIXES:

* resource/auth0_client, resource/auth0_global_client: fix `null` scope issue ([#204](https://github.com/alexkappa/terraform-provider-auth0/pull/204))
* resource/auth0_connection: various bug fixes for auth0 type connections.
* resource/auth0_role: paginating role permissions for large amounts of permissions defined per role.

## 0.9.0 (April 14, 2020)

BUG FIXES:

* resource/auth0_resource_server: fixed rename scope bug ([#197](https://github.com/alexkappa/terraform-provider-auth0/issues/197))
* resource/auth0_tenant: fix "too few properties defined" bug by sending certain fields in every update ([#185](https://github.com/alexkappa/terraform-provider-auth0/issues/185))

NOTES:

* User Agent is now more accurate and follows the package version of `go-auth0/auth0`.
* Updates (PATCH) will include most fields in requests by default even if no changes were observed. [#194](https://github.com/alexkappa/terraform-provider-auth0/pull/194)

## 0.8.2 (April 08, 2020)

BUG FIXES:

* resource/auth0_connection: with `email` strategy `totp` settings were not handled correctly ([#191](https://github.com/alexkappa/terraform-provider-auth0/pull/191)).

## 0.8.1 (March 27, 2020)

FEATURES:

* resource/auth0_connection: support for the `github` connection strategy ([#184](https://github.com/alexkappa/terraform-provider-auth0/pull/184)).

## 0.8.0 (March 24, 2020)

FEATURES:

* **New Resource:** auth0_prompt ([#8](https://github.com/terraform-providers/terraform-provider-auth0/pull/8))
* resource/auth0_tenant: add `use_scope_descriptions_for_consent` flag ([#180](https://github.com/alexkappa/terraform-provider-auth0/pull/180)).

BUG FIXES:
* resource/auth0_tenant: fix crash when the `change_password` field was not defined ([#181](https://github.com/alexkappa/terraform-provider-auth0/issues/181)).

## 0.7.0 (March 23, 2020)

FEATURES:

* resource/auth0_connection: support for the passwordless `email` connection strategy.

ENHANCEMENTS:

* resource/auth0_connection: now using the more powerful connection options from `gopkg.in/auth0.v4`.

BUG FIXES:
* resource/auth0_tenant, resource/auth0_connection: issues setting boolean attributes within nested blocks should now be resolved ([#163](https://github.com/alexkappa/terraform-provider-auth0/issues/163), [#160](https://github.com/alexkappa/terraform-provider-auth0/issues/160))

NOTES:

* Upgrade to `gopkg.in/auth0.v4` (`v4.0.0`)

## 0.6.0 (March 03, 2020)

FEATURES:

* **New Resource:** auth0_hook ([#171](https://github.com/alexkappa/terraform-provider-auth0/pull/171))
* **New Resource:** auth0_global_client ([#172](https://github.com/alexkappa/terraform-provider-auth0/pull/172))

ENHANCEMENTS:

* resource/auth0_user: `name`, `family_name`, `given_name`, `blocked` and `picture` are added ([#166](https://github.com/alexkappa/terraform-provider-auth0/pull/166))
* resource/auth0_client: add `initiate_login_uri` ([#2](https://github.com/terraform-providers/terraform-provider-auth0/pull/2))
* resource/auth0_tenant: add `default_redirection_uri` ([#2](https://github.com/terraform-providers/terraform-provider-auth0/pull/2))
* resource/auth0_connection: `strategy` is now required and the `apple`, `oidc` and `line` strategies are added ([#6](https://github.com/terraform-providers/terraform-provider-auth0/pull/6))

BUG FIXES:

* resource/auth0_user: unassigning a role won't fail if the role has already been deleted.

## v0.5.1 (January 29, 2020)

Initial release under releases.hashicorp.com

BUG FIXES:

* resource/auth0_email: fix `api_key` issue when reading back the resource from Auth0 ([#161](https://github.com/alexkappa/terraform-provider-auth0/issues/161))

## v0.5.0 (January 20, 2020)

ENHANCEMENTS:

* resource/auth0_email: add `domain` field to allow configuring of mailgun provider ([#164](https://github.com/alexkappa/terraform-provider-auth0/pull/164))

NOTES:

* Upgrade to `gopkg.in/auth0.v3` (`v3.0.3`)


## v0.4.3 (January 16, 2020)

BUG FIXES:

* resource/auth0_client_grant: fix empty scope issue ([#162](https://github.com/alexkappa/terraform-provider-auth0/pull/162))

## v0.4.2 (December 30, 2019)

ENHANCEMENTS:

* resource/*: update and destroy operations now do not fail if the resource has been deleted manually ([#155](https://github.com/alexkappa/terraform-provider-auth0/pull/155)).

## v0.4.1 (December 18, 2019)

ENHANCEMENTS:

* resource/auth0_client: support rotating `client_secret` by changing the value of `client_secret_rotation_trigger` ([#153](https://github.com/alexkappa/terraform-provider-auth0/pull/153)).

## v0.4.0 (December 13, 2019)

ENHANCEMENTS:

* resource/auth0_connection: Introduce `password_complexity_options` ([#132](https://github.com/alexkappa/terraform-provider-auth0/pull/132)).
* resource/auth0_resource_server: `signing_secret` is now also a computed field. If set it's validated to be at least 16 characters ([#146](https://github.com/alexkappa/terraform-provider-auth0/pull/146)).
* resource/auth0_resource_server: `identifier` update forces new resource ([#147](https://github.com/alexkappa/terraform-provider-auth0/pull/147)).
* resource/auth0_role (**breaking change**): `user_ids` is removed. In its place the following is introduced ([#149](https://github.com/alexkappa/terraform-provider-auth0/pull/149)).
* resource/auth0_user: `roles` is added ([#149](https://github.com/alexkappa/terraform-provider-auth0/pull/149)).

BUG FIXES:

* resource/auth0_connection: Fix `password_dictionary` [#128](https://github.com/alexkappa/terraform-provider-auth0/pull/128)
* resource/auth0_client: Fix `is_first_party` setting if set to zero value ([#148](https://github.com/alexkappa/terraform-provider-auth0/pull/148)). 

## v0.3.1 (December 10, 2019)

ENHANCEMENTS:

* resource/auth0_tenant: Support `flags` and `universal_login` settings [#133](https://github.com/alexkappa/terraform-provider-auth0/pull/133)

## v0.3.0 (December 10, 2019)

BUG FIXES:

* resource/auth0_email_template: Fix 404 issue when modifying templates ([#144](https://github.com/alexkappa/terraform-provider-auth0/pull/144)).

NOTES:

* Upgrade to `gopkg.in/auth0.v2`

## v0.2.2 (December 10, 2019)

ENHANCEMENTS:

* Switch to using Github Actions in favor of Wercker.

## v0.2.1 (October 21, 2019)

ENHANCEMENTS:

* resource/auth0_connection: Improved support for `enabled_clients` by using a set instead of a list ([#105](https://github.com/alexkappa/terraform-provider-auth0/pull/105)).
* resource/auth0_connection: Add `community_base_url` to `salesforce-community` type connections ([#130](https://github.com/alexkappa/terraform-provider-auth0/pull/130)).
* resource/auth0_client: Validate `app_type` ([#112](https://github.com/alexkappa/terraform-provider-auth0/pull/)).
* resource/auth0_user: Make `password` a sensitive field ([#117](https://github.com/alexkappa/terraform-provider-auth0/pull/117)).

BUG FIXES

* resource/auth0_connection: Fix incorrect schema for `password_no_personal_info` ([#107](https://github.com/alexkappa/terraform-provider-auth0/pull/107)).
* resource/auth0_user: Fix bugs in `user_metadata`, `app_metadata` and `password` ([#131](https://github.com/alexkappa/terraform-provider-auth0/pull/131)).

NOTES:

* Improve documentation on supported features ([#113](https://github.com/alexkappa/terraform-provider-auth0/pull/113)).
* Update examples after upgrade to 0.2.x ([#114](https://github.com/alexkappa/terraform-provider-auth0/pull/114)).
* Update terraform and auth0 dependencies to latest release ([#120](https://github.com/alexkappa/terraform-provider-auth0/pull/120)).
* Add tenant example ([#125](https://github.com/alexkappa/terraform-provider-auth0/pull/125)).

FEATURES:

* resource/auth0_user: Support for `nickname` attribute ([#108](https://github.com/alexkappa/terraform-provider-auth0/pull/108))

## v0.2.0 (June 27, 2019)

ENHANCEMENTS:

* resource/auth0_user: Add support for user attribute `nickname`

BUG FIXES:

* resource/auth0_connection: Fix incorrect schema of `password_no_personal_info`

NOTES:

* Provider is compatible with Terraform v0.12.3

## v0.1.20 (May 17, 2019)

FEATURES:

* resource/auth0_connection: Add twillio for guardian MFA

## v0.1.19 (May 14, 2019)

ENHANCEMENTS

* resource/auth0_connection: Add `adfs_server` field.

## v0.1.18 (March 26, 2019)

NOTES:

* Extract version from changelog for release notes.

## v0.1.17 (March 26, 2019)

NOTES:

* Update Travis to build on tag push.

## v0.1.16 (March 25, 2019)

NOTES:

* Added contributing, code of conduct, issue templates to the project.

## v0.1.15 (March 25, 2019)

FEATURES:

* **New Resource:** auth0_tenant ([#79](https://github.com/alexkappa/terraform-provider-auth0/pull/79))

ENHANCEMENTS:

* resource/auth0_connection: `enabled_clients` will suppress diff if the list order is different.
* resource/auth0_connection: set `client_secret` to sensitive ([#91](https://github.com/alexkappa/terraform-provider-auth0/pull/91)).
* resource/auth0_resource_server: introduce `token_lifetime_for_web` ([#84](https://github.com/alexkappa/terraform-provider-auth0/pull/84)).

NOTES:

* Re-imported Auth0 SDK from `gopkg.in/auth0.v1`.
