terraform {
  required_providers {
    auth0 = {
      source  = "alexkappa/auth0"
      version = "0.24.3"
    }
  }
}

provider "auth0" {}

resource "auth0_action" "do" {

  name = "Test Action ${timestamp()}"

  supported_triggers {
    id      = "post-login"
    version = "v2"
  }

  runtime = "node16"
  code    = <<-EOT
	exports.onContinuePostLogin = async (event, api) => { 
		console.log(event) 
	};"
	EOT

  dependencies {
    name    = "lodash"
    version = "latest"
  }
  dependencies {
    name    = "request"
    version = "latest"
  }

  secrets {
    name  = "FOO"
    value = "Foo"
  }
  secrets {
    name  = "BAR"
    value = "Bar"
  }

  deploy = true
}
