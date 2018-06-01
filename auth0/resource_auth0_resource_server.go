package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yieldr/go-auth0/management"
)

func newResourceServer() *schema.Resource {
	return &schema.Resource{

		Create: createResourceServer,
		Read:   readResourceServer,
		Update: updateResourceServer,
		Delete: deleteResourceServer,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"signing_alg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signing_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_offline_access": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token_lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"skip_consent_for_verifiable_first_party_clients": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"verification_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"options": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func createResourceServer(d *schema.ResourceData, m interface{}) error {
	s := buildResourceServer(d)
	api := m.(*management.Management)
	if err := api.ResourceServer.Create(s); err != nil {
		return err
	}
	d.SetId(s.ID)
	return readResourceServer(d, m)
}

func readResourceServer(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	s, err := api.ResourceServer.Read(d.Id())
	if err != nil {
		return err
	}
	d.SetId(s.ID)
	d.Set("name", s.Name)
	d.Set("identifier", s.Identifier)
	d.Set("scopes", func() (m []map[string]interface{}) {
		for _, scope := range s.Scopes {
			m = append(m, map[string]interface{}{
				"value":       scope.Value,
				"description": scope.Description,
			})
		}
		return m
	}())
	d.Set("signing_alg", s.SigningAlgorithm)
	d.Set("signing_secret", s.SigningSecret)
	d.Set("allow_offline_access", s.AllowOfflineAccess)
	d.Set("token_lifetime", s.TokenLifetime)
	d.Set("skip_consent_for_verifiable_first_party_clients", s.SkipConsentForVerifiableFirstPartyClients)
	d.Set("verification_location", s.VerificationLocation)
	d.Set("options", s.Options)
	return nil
}

func updateResourceServer(d *schema.ResourceData, m interface{}) error {
	s := buildResourceServer(d)
	s.Identifier = ""
	api := m.(*management.Management)
	err := api.ResourceServer.Update(d.Id(), s)
	if err != nil {
		return err
	}
	return readResourceServer(d, m)
}

func deleteResourceServer(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.ResourceServer.Delete(d.Id())
}

func buildResourceServer(d *schema.ResourceData) *management.ResourceServer {

	s := &management.ResourceServer{
		Name:                                      d.Get("name").(string),
		Identifier:                                d.Get("identifier").(string),
		SigningAlgorithm:                          d.Get("signing_alg").(string),
		SigningSecret:                             d.Get("signing_secret").(string),
		AllowOfflineAccess:                        d.Get("allow_offline_access").(bool),
		TokenLifetime:                             d.Get("token_lifetime").(int),
		SkipConsentForVerifiableFirstPartyClients: d.Get("skip_consent_for_verifiable_first_party_clients").(bool),
		VerificationLocation:                      d.Get("verification_location").(string),
		Options:                                   d.Get("options").(map[string]interface{}),
	}

	if v, ok := d.GetOk("scopes"); ok {

		for _, vI := range v.([]interface{}) {

			scopes := vI.(map[string]interface{})

			s.Scopes = append(s.Scopes, &management.ResourceServerScope{
				Value:       scopes["value"].(string),
				Description: scopes["description"].(string),
			})
		}
	}

	return s
}
