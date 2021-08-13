package hashicups

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{"hashicups_coffees": dataSourceCoffees(), "hashicups_menu": dataSourceMenu()},
	}
}
