package auth0

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccClientDataSource(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccDataSourceClientConfigCreate, rand),
				Check: testAccDataSourceAuth0Client("auth0_client.my_client", "data.auth0_client.foo",
					[]string{"id", "name", "client_id", "client_secret", "description"}),
			},
			{
				Config:      random.Template(testAccDataSourceClientConfigCreateDuplicates, rand),
				ExpectError: regexp.MustCompile("Multiple Auth0 Clients with name Acceptance Test - Duplicate Name Check"),
			},
		},
	})
}

const testAccDataSourceClientConfigCreate = `

resource "auth0_client" "my_client" {
  name = "Acceptance Test - Zero Value Check - {{.random}}"
  description = "Terraform acceptance tests"
}

data "auth0_client" "foo" {
	name = "${auth0_client.my_client.name}"
  }
`
const testAccDataSourceClientConfigCreateDuplicates = `

resource "auth0_client" "my_client_1" {
  name = "Acceptance Test - Duplicate Name Check"
  description = "Terraform acceptance tests"
}

resource "auth0_client" "my_client_2" {
	name = "Acceptance Test - Duplicate Name Check"
	description = "Terraform acceptance tests"
}
data "auth0_client" "foo" {
	name = "${auth0_client.my_client_1.name}"
	depends_on = [
		auth0_client.my_client_1,
		auth0_client.my_client_2,
  ]
  }
`

func testAccDataSourceAuth0Client(resourceName, dataSourceName string, testAttributes []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := s.RootModule().Resources[resourceName]
		clientResource := client.Primary.Attributes

		search := s.RootModule().Resources[dataSourceName]
		searchResource := search.Primary.Attributes
		if searchResource["id"] == "" {
			return fmt.Errorf("Expected to get a Client ID from Auth0")
		}

		for _, attribute := range testAttributes {
			if searchResource[attribute] != clientResource[attribute] {
				return fmt.Errorf("Expected the Client's %s to be: %s, but got: %s", attribute, clientResource[attribute], searchResource[attribute])
			}
		}
		return nil
	}
}
