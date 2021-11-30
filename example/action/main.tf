terraform {
  required_providers {
    auth0 = {
      source = "alexkappa/auth0"
      version = "0.24.1"
    }
  }
}

provider "auth0" {}

resource "auth0_action" "do" {
	name = "Test Action ${timestamp()}"
	supported_triggers {
		id = "post-login"
		version = "v2"
	}
	code = <<-EOT
	exports.onContinuePostLogin = async (event, api) => { 
		console.log(event) 
	};"
	EOT
	deploy = true
}