provider "auth0" {}

resource "auth0_user" "user" {
  connection_name = "Username-Password-Authentication"
  user_id         = "auth0|1234567890"
  email           = "test@test.com"
  password        = "passpass$12$12"
  nickname        = "testnick"
}

resource "auth0_role" "my_role" {
  name        = "My Role - (Managed by Terraform)"
  description = "Role Description..."
}

resource "auth0_user_role" "my_user_role" {
  user_id = "${auth0_user.user.id}"
  roles   = ["${auth0_role.my_role.id}"]

  depends_on = ["auth0_user.user"]
}
