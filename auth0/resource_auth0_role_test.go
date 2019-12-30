package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccRole(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRole_create,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_role.the_one", "name", "The One - Role - Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.the_one", "description", "The One - Role - Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.the_one", "permissions.#", "1"),
				),
			},
			{
				Config: testAccRole_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_role.the_one", "description", "The One who will bring peace - Role - Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.the_one", "permissions.#", "2"),
				),
			},
		},
	})
}

const testAccRole_create = `
provider auth0 {}

resource auth0_resource_server matrix {
	name = "The One - Resource Server - Acceptance Test"
	identifier = "https://matrix.com/"
	scopes {
		value = "stop:bullets"
		description = "Stop bullets"
	}
	scopes {
		value = "bring:peace"
		description = "Bring peace"
	}
  }

resource auth0_role the_one {
  name = "The One - Role - Acceptance Test"
  description = "The One - Role - Acceptance Test"
  permissions {
	name = "stop:bullets"
	resource_server_identifier = auth0_resource_server.matrix.identifier
  }
}
`

const testAccRole_update = `
provider auth0 {}

resource auth0_resource_server matrix {
	name = "The One - Resource Server - Acceptance Test"
	identifier = "https://matrix.com/"
	scopes {
		value = "stop:bullets"
		description = "Create bars"
	}
	scopes {
		value = "bring:peace"
		description = "Bring peace"
	}
  }

resource auth0_role the_one {
  name = "The One - Role - Acceptance Test"
  description = "The One who will bring peace - Role - Acceptance Test"
  permissions {
	name = "stop:bullets"
	resource_server_identifier = auth0_resource_server.matrix.identifier
  }
  permissions {
	name = "bring:peace"
	resource_server_identifier = auth0_resource_server.matrix.identifier
  }
}
`
