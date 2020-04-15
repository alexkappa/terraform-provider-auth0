package debug

import (
	"fmt"
	"log"
	"sort"

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
		attributes := rs.Primary.Attributes
		keys := keys(attributes)
		sort.Strings(keys)
		for _, key := range keys {
			log.Printf("[DEBUG]\t %s: %q\n", key, attributes[key])
		}
		return nil
	}
}

func keys(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
