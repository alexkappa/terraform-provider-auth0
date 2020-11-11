package auth0

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
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
				ForceNew: true,
			},
			"scopes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Required: true,
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
				Computed: true,
				ValidateFunc: func(i interface{}, k string) (s []string, es []error) {
					v, ok := i.(string)
					if !ok {
						es = append(es, fmt.Errorf("expected type of %s to be string", k))
						return
					}
					min := 16
					if len(v) < min {
						es = append(es, fmt.Errorf("expected length of %s to be at least %d, %q is %d", k, min, v, len(v)))
					}
					return
				},
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
			"token_lifetime_for_web": {
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
			"enforce_policies": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token_dialect": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"access_token",
					"access_token_authz",
				}, true),
			},
		},
	}
}

func createResourceServer(d *schema.ResourceData, m interface{}) error {
	s := expandResourceServer(d)
	api := m.(*management.Management)
	if err := api.ResourceServer.Create(s); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(s.ID))
	return readResourceServer(d, m)
}

func readResourceServer(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	s, err := api.ResourceServer.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(auth0.StringValue(s.ID))
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
	d.Set("token_lifetime_for_web", s.TokenLifetimeForWeb)
	d.Set("skip_consent_for_verifiable_first_party_clients", s.SkipConsentForVerifiableFirstPartyClients)
	d.Set("verification_location", s.VerificationLocation)
	d.Set("options", s.Options)
	d.Set("enforce_policies", s.EnforcePolicies)
	d.Set("token_dialect", s.TokenDialect)
	return nil
}

func updateResourceServer(d *schema.ResourceData, m interface{}) error {
	s := expandResourceServer(d)
	s.Identifier = nil
	api := m.(*management.Management)
	err := api.ResourceServer.Update(d.Id(), s)
	if err != nil {
		return err
	}
	return readResourceServer(d, m)
}

func deleteResourceServer(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.ResourceServer.Delete(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
	}
	return err
}

func expandResourceServer(d *schema.ResourceData) *management.ResourceServer {

	s := &management.ResourceServer{
		Name:                 String(d, "name"),
		Identifier:           String(d, "identifier"),
		SigningAlgorithm:     String(d, "signing_alg"),
		SigningSecret:        String(d, "signing_secret", IsNewResource(), HasChange()),
		AllowOfflineAccess:   Bool(d, "allow_offline_access"),
		TokenLifetime:        Int(d, "token_lifetime"),
		TokenLifetimeForWeb:  Int(d, "token_lifetime_for_web"),
		VerificationLocation: String(d, "verification_location"),
		Options:              Map(d, "options"),
		EnforcePolicies:      Bool(d, "enforce_policies"),
		TokenDialect:         String(d, "token_dialect", IsNewResource(), HasChange()),

		SkipConsentForVerifiableFirstPartyClients: Bool(d, "skip_consent_for_verifiable_first_party_clients"),
	}

	Set(d, "scopes").Elem(func(d ResourceData) {
		s.Scopes = append(s.Scopes, &management.ResourceServerScope{
			Value:       String(d, "value"),
			Description: String(d, "description"),
		})
	})

	return s
}
