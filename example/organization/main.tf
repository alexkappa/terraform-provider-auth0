terraform {
  required_providers {
    auth0 = {
      source  = "alexkappa/auth0"
      version = "0.22.0"
    }
  }
}

provider "auth0" {}

resource "auth0_organization" "organization" {
  name         = "alex-inc"
  display_name = "Alex Inc."
  branding {
    logo_url = "https://alexkappa.com/assets/icons/icon.png"
    colors = {
      primary         = "#f2f2f2"
      page_background = "#e1e1e1"
    }
  }
  connections {
    connection_id = "con_X7iCWk8xB076gRi2"
  }
}

output "auth0_organization_id" {
  value = auth0_organization.organization.id
}
