package auth0

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-auth0/auth0/internal/random"
	"gopkg.in/auth0.v4/management"
)

var salesforceResourceTypes = []string{"auth0_connection_salesforce", "auth0_connection_salesforce_sandbox", "auth0_connection_salesforce_community"}

func init() {
	for _, r := range salesforceResourceTypes {
		resource.AddTestSweepers(r, &resource.Sweeper{
			Name: r,
			F: func(_ string) error {
				api, err := Auth0()
				if err != nil {
					return err
				}
				var page int
				for {
					l, err := api.Connection.List(
						management.WithFields("id", "name"),
						management.Page(page))
					if err != nil {
						return err
					}
					for _, connection := range l.Connections {
						if strings.Contains(connection.GetName(), "Test") {
							log.Printf("[DEBUG] Deleting connection %v\n", connection.GetName())
							if e := api.Connection.Delete(connection.GetID()); e != nil {
								multierror.Append(err, e)
							}
						}
					}
					if err != nil {
						return err
					}
					if !l.HasNext() {
						break
					}
					page++
				}
				return nil
			},
		})
	}
}

func TestAccConnectionSalesforce(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionSalesforceConfig("auth0_connection_salesforce", "my_connection"), rand+"1"),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_salesforce.my_connection", "name", "Acceptance-Test-Salesforce-Connection-{{.random}}", rand+"1"),
					resource.TestCheckResourceAttr("auth0_connection_salesforce.my_connection", "strategy", "salesforce"),
					resource.TestCheckResourceAttr("auth0_connection_salesforce.my_connection", "options.0.community_base_url", "https://salesforce.example.com"),
				),
			},
			{
				Config: random.Template(testAccConnectionSalesforceConfig("auth0_connection_salesforce_sandbox", "my_connection_sandbox"), rand+"2"),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_salesforce_sandbox.my_connection_sandbox", "name", "Acceptance-Test-Salesforce-Connection-{{.random}}", rand+"2"),
					resource.TestCheckResourceAttr("auth0_connection_salesforce_sandbox.my_connection_sandbox", "strategy", "salesforce-sandbox"),
					resource.TestCheckResourceAttr("auth0_connection_salesforce_sandbox.my_connection_sandbox", "options.0.community_base_url", "https://salesforce.example.com"),
				),
			},
			{
				Config: random.Template(testAccConnectionSalesforceConfig("auth0_connection_salesforce_community", "my_connection_community"), rand+"3"),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_salesforce_community.my_connection_community", "name", "Acceptance-Test-Salesforce-Connection-{{.random}}", rand+"3"),
					resource.TestCheckResourceAttr("auth0_connection_salesforce_community.my_connection_community", "strategy", "salesforce-community"),
					resource.TestCheckResourceAttr("auth0_connection_salesforce_community.my_connection_community", "options.0.community_base_url", "https://salesforce.example.com"),
				),
			},
		},
	})
}

func testAccConnectionSalesforceConfig(resource string, name string) string {
	return fmt.Sprintf(`
		resource %q %q {
			name = "Acceptance-Test-Salesforce-Connection-{{.random}}"
			is_domain_connection = false

			options {
				client_id = false
				client_secret = "sms-connection"
				community_base_url = "https://salesforce.example.com"
			}
		}`, resource, name)
}

func TestAccConnectionSalesforceWithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("salesforce"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_salesforce.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_salesforce.my_connection", "enabled_clients.#", "4"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
