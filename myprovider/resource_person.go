package myprovider

import (
	"context"
	"fmt"

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
	// Load DB first (if needed to populate internal state)
	err = m.(*Client).LoadDB()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error loading database",
			Detail:   err.Error(),
		})
		return diags
	}
	err = m.(*Client).SaveDB()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to save person",
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
	err := m.(*Client).LoadDB()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error loading database",
			Detail:   err.Error(),
		})
		return diags
	}
	person, err := m.(*Client).GetPerson(d.Get("name_id").(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error reading person from the database",
			Detail:   err.Error(),
		})
		return diags
	}
	d.Set("last_name", person.LastName)
	d.Set("first_name", person.FirstName)
	return diags
}

func resourcePersonUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "*** function resourcePersonUpdate called ***")

	client := m.(*Client)

	// Load DB first (if needed to populate internal state)
	err := client.LoadDB()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error loading database",
			Detail:   err.Error(),
		})
		return diags
	}

	nameID := d.Get("name_id").(string)

	// Get the existing person (optional: could skip if overwriting)
	person, err := client.GetPerson(nameID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("could not find person for update: %w", err))
	}

	// Update fields from Terraform config
	if d.HasChange("first_name") {
		person.FirstName = d.Get("first_name").(string)
	}
	// You don't need to handle "last_name" since it's ForceNew

	// Save updated person back to the in-memory DB
	client.Persons[nameID] = person

	// Now save to file
	err = client.SaveDB()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to save person",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourcePersonDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "*** function resourcePersonDelete called ***")
	return diags
}
