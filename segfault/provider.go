package segfault

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"segfault": dataSegfault(),
		},
	}
}

func dataSegfault() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSegfaultRead,
		Schema: map[string]*schema.Schema{},
	}
}

func dataSegfaultRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var p *int
	_ = *p
	return nil
}
