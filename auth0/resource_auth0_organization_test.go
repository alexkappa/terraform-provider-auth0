package auth0

import (
	"log"
	"strings"
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"gopkg.in/auth0.v5/management"
)

func init() {
	resource.AddTestSweepers("auth0_organization", &resource.Sweeper{
		Name: "auth0_organization",
		F: func(_ string) error {
			api, err := Auth0()
			if err != nil {
				return err
			}
			var page int
			for {
				l, err := api.Organization.List(management.Page(page))
				if err != nil {
					return err
				}
				for _, organization := range l.Organizations {
					log.Printf("[DEBUG] ➝ %s", organization.GetName())
					if strings.Contains(organization.GetName(), "test") {
						if e := api.Organization.Delete(organization.GetID()); e != nil {
							multierror.Append(err, e)
						}
						log.Printf("[DEBUG] ✗ %s", organization.GetName())
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

func TestAccOrganization(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccOrganizationCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_organization.acme", "name", "test-{{.random}}", rand),
					random.TestCheckResourceAttr("auth0_organization.acme", "display_name", "Acme Inc. {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_organization.acme", "connections.#", "1"),
				),
			},
			{
				Config: random.Template(testAccOrganizationUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_organization.acme", "name", "test-{{.random}}", rand),
					random.TestCheckResourceAttr("auth0_organization.acme", "display_name", "Acme Inc. {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_organization.acme", "branding.#", "1"),
					resource.TestCheckResourceAttr("auth0_organization.acme", "branding.0.logo_url", "https://acme.com/logo.svg"),
					resource.TestCheckResourceAttr("auth0_organization.acme", "branding.0.colors.%", "2"),
					resource.TestCheckResourceAttr("auth0_organization.acme", "branding.0.colors.primary", "#e3e2f0"),
					resource.TestCheckResourceAttr("auth0_organization.acme", "branding.0.colors.page_background", "#e3e2ff"),
					resource.TestCheckResourceAttr("auth0_organization.acme", "connections.#", "2"),
				),
			},
			{
				Config: random.Template(testAccOrganizationUpdateAgain, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_organization.acme", "name", "test-{{.random}}", rand),
					random.TestCheckResourceAttr("auth0_organization.acme", "display_name", "Acme Inc. {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_organization.acme", "connections.#", "1"),
				),
			},
		},
	})
}

const testAccOrganizationAux = `

resource auth0_connection acme {
	name = "Acceptance-Test-Connection-Acme-{{.random}}"
	strategy = "auth0"
}

resource auth0_connection acmeinc {
	name = "Acceptance-Test-Connection-Acme-Inc-{{.random}}"
	strategy = "auth0"
}
`

const testAccOrganizationCreate = testAccOrganizationAux + `

resource auth0_organization acme {
	name = "test-{{.random}}"
	display_name = "Acme Inc. {{.random}}"
	
	connections {
		connection_id = auth0_connection.acme.id
	}
}
`

const testAccOrganizationUpdate = testAccOrganizationAux + `

resource auth0_organization acme {
	name = "test-{{.random}}"
	display_name = "Acme Inc. {{.random}}"
	branding {
		logo_url = "https://acme.com/logo.svg"
		colors = {
			primary = "#e3e2f0"
			page_background = "#e3e2ff"
		}
	}
	connections {
		connection_id = auth0_connection.acme.id
	}
	connections {
		connection_id = auth0_connection.acmeinc.id
		assign_membership_on_login = true
	}
}
`

const testAccOrganizationUpdateAgain = testAccOrganizationAux + `

resource auth0_organization acme {
	name = "test-{{.random}}"
	display_name = "Acme Inc. {{.random}}"
	branding {
		logo_url = "https://acme.com/logo.svg"
		colors = {
			primary = "#e3e2f0"
			page_background = "#e3e2ff"
		}
	}
	connections {
		connection_id = auth0_connection.acmeinc.id
		assign_membership_on_login = false
	}
}
`
