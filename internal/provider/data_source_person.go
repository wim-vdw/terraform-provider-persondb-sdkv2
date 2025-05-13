package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePerson() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePersonRead,
		Schema: map[string]*schema.Schema{
			"person_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePersonRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func dataSourcePersonRead *****")
	client := m.(*Client)
	personID := d.Get("person_id").(string)
	lastName, firstName, err := client.readPerson(personID)
	if err != nil {
		return diag.Errorf("unable to read person with person_id '%s' from the database", personID)
	}
	d.SetId("/person/" + personID)
	d.Set("last_name", lastName)
	d.Set("first_name", firstName)
	return diags
}
