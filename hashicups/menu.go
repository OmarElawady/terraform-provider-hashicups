package hashicups

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMenu() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceMenusRead,
		UpdateContext: dataSourceMenusUpdate,
		CreateContext: dataSourceMenusCreate,
		DeleteContext: dataSourceMenusDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
		},
	}
}

func dataSourceMenusCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	log.Printf("creating menus")
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func dataSourceMenusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	log.Printf("reading menus")

	// always run

	return diags
}

func dataSourceMenusUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Printf("Updating menus object")
	return diags
}

func dataSourceMenusDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Printf("Deleting menus object")
	return diags
}
