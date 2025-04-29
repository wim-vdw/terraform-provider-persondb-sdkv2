package myprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePerson() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePersonCreate,
		ReadContext:   resourcePersonRead,
		UpdateContext: resourcePersonUpdate,
		DeleteContext: resourcePersonDelete,

		Schema: map[string]*schema.Schema{
			"name_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePersonCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := m.(*Client).AddPerson(d.Get("name_id").(string), d.Get("name").(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create person",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId("/person/" + d.Get("name_id").(string))
	return diags
}

func resourcePersonRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePersonUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourcePersonDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
