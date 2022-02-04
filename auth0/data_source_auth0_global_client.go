package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func newDataGlobalClient() *schema.Resource {
	return &schema.Resource{
		Read:   readDataGlobalClient,
		Schema: newClientSchema(),
	}
}

func readDataGlobalClient(d *schema.ResourceData, m interface{}) error {
	if err := readGlobalClientId(d, m); err != nil {
		return err
	}
	return readClient(d, m)
}
