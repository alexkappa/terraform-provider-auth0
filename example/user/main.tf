provider "auth0" {}

resource "auth0_user" "user" {
  connection_name = "Username-Password-Authentication"
  user_id = "12345"
  email = "test@test.com"
  email_verified = true
  password = "testtest$12$12"
}
