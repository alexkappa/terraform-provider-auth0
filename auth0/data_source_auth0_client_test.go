package auth0

import (
	"fmt"
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccDataClientConfigByName = `
%v
data auth0_client my_client {
  name = "Acceptance Test - {{.random}}"
}
`

const testAccDataClientConfigById = `
%v
data auth0_client my_client {
  client_id = auth0_client.my_client.client_id
}
`

func clientDataSourceFields() []string {
	fields := clientGlobalDataSourceFields()
	fields = append(
		fields,
		"client_metadata.%",
		"client_metadata.foo",
	)
	return fields
}

func TestAccDataClientByName(t *testing.T) {
	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccClientConfig, rand), // must initialize resource before reading with data source
			},
			{
				Config: random.Template(fmt.Sprintf(testAccDataClientConfigByName, testAccClientConfig), rand),
				Check:  checkDataSourceStateMatchesResourceState("data.auth0_client.my_client", "auth0_client.my_client", clientDataSourceFields()),
			},
		},
	})
}

func TestAccDataClientById(t *testing.T) {
	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccClientConfig, rand),
			},
			{
				Config: random.Template(fmt.Sprintf(testAccDataClientConfigById, testAccClientConfig), rand),
				Check:  checkDataSourceStateMatchesResourceState("data.auth0_client.my_client", "auth0_client.my_client", clientDataSourceFields()),
			},
		},
	})
}
