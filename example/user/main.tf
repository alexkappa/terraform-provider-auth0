provider "auth0" {}

resource "auth0_user" "user" {
  connection_name = "Username-Password-Authentication"
  user_id = "12345"
  username = "unique_username"
  name = "Firstname Lastname"
  nickname = "some.nickname"
  email = "test@test.com"
  email_verified = true
  password = "passpass$12$12"
  picture = "https://www.example.com/a-valid-picture-url.jpg"
  roles = [ auth0_role.admin.id ]
}

resource "auth0_role" "admin" {
	name = "admin"
	description = "Administrator"
}