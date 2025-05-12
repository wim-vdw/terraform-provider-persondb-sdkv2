package provider

import (
	"context"
	"strings"

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
			"person_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePersonCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func resourcePersonCreate *****")
	client := m.(*Client)
	personID := d.Get("person_id").(string)
	lastName := d.Get("last_name").(string)
	firstName := d.Get("first_name").(string)
	exists, err := client.checkPersonExists(personID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create person",
			Detail:   err.Error(),
		})
		return diags
	}
	if exists {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "person already exists",
			Detail:   "Person with person_id '" + personID + "' already exists. Use 'terraform import' to manage it in Terraform.",
		})
		return diags
	}
	err = client.createPerson(personID, lastName, firstName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create person",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId("/person/" + personID)
	return diags
}

func resourcePersonRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func resourcePersonRead *****")
	client := m.(*Client)
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 3 || parts[1] != "person" {
		return diag.Errorf("invalid ID format: expected '/person/<person_id>', got: %s", d.Id())
	}
	personID := parts[2]
	lastName, firstName, err := client.readPerson(personID)
	if err != nil {
		// Person could not be found, so we set the ID to empty
		d.SetId("")
	}
	d.Set("person_id", personID)
	d.Set("last_name", lastName)
	d.Set("first_name", firstName)
	return diags
}

func resourcePersonUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func resourcePersonUpdate *****")
	client := m.(*Client)
	personID := d.Get("person_id").(string)
	lastName := d.Get("last_name").(string)
	firstName := d.Get("first_name").(string)
	err := client.updatePerson(personID, lastName, firstName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to update person",
			Detail:   err.Error(),
		})
		return diags
	}
	return diags
}

func resourcePersonDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func resourcePersonDelete *****")
	client := m.(*Client)
	personID := d.Get("person_id").(string)
	err := client.deletePerson(personID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete person",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId("")
	return diags
}
