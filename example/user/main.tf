provider "auth0" {}

resource "auth0_user" "user" {
  connection_name = "Username-Password-Authentication"
  user_id = "12345"
  username = "test"
  nickname = "testnick"
  email = "test@test.com"
  email_verified = true
  password = "passpass$12$12"
  roles = [ auth0_role.admin.id ]
}

resource "auth0_role" "admin" {
	name = "admin"
	description = "Administrator"
}