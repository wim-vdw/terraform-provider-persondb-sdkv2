package myprovider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"database_filename": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CUSTOM_DATABASE_FILENAME", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"myprovider_person": resourcePerson(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	c := &Client{
		CustomDatabase: d.Get("database_filename").(string),
	}
	_, err := os.ReadFile(c.CustomDatabase)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Database client",
			Detail:   err.Error(),
		})
		return nil, diags
	}
	return c, diags
}
