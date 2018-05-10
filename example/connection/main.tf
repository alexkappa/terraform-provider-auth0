provider "auth0" {}

resource "auth0_connection" "my_connection" {
  name = "Example-Connection"
  strategy = "auth0"
  options = {
    password_policy = "excellent"
  }
}
