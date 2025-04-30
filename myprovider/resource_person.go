package myprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePersonCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "*** function resourcePersonCreate called ***")
	err := m.(*Client).AddPerson(d.Get("name_id").(string), d.Get("last_name").(string), d.Get("first_name").(string))
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
	tflog.Info(ctx, "*** function resourcePersonRead called ***")
	return diags
}

func resourcePersonUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "*** function resourcePersonUpdate called ***")
	return diags
}

func resourcePersonDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "*** function resourcePersonDelete called ***")
	return diags
}
