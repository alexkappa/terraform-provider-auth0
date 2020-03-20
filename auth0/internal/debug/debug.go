package debug

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func DumpAttr(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		log.Printf("[DEBUG] Attrs: \n")
		for key, value := range rs.Primary.Attributes {
			log.Printf("[DEBUG]\t %s: %q\n", key, value)
		}
		return nil
	}
}
