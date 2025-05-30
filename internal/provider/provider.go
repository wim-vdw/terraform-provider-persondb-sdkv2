package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	persondbclient "github.com/wim-vdw/terraform-provider-persondb/internal/client"
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
			"persondb_person": resourcePerson(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"persondb_person": dataSourcePerson(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func providerConfigure *****")
	databaseFilename := d.Get("database_filename").(string)
	client, err := persondbclient.NewClient(databaseFilename)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to initialize database",
			Detail:   err.Error(),
		})
		return nil, diags
	}
	return client, diags
}
