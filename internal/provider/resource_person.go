package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	persondbclient "github.com/wim-vdw/terraform-provider-persondb/internal/client"
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 30),
			},
			"first_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
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
	client := m.(*persondbclient.Client)
	personID := d.Get("person_id").(string)
	lastName := d.Get("last_name").(string)
	firstName := d.Get("first_name").(string)
	exists, err := client.CheckPersonExists(personID)
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
	err = client.CreatePerson(personID, lastName, firstName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create person",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId("/person/" + personID)

	// Best practice: Update state after modification
	// https://developer.hashicorp.com/terraform/plugin/sdkv2/best-practices/detecting-drift#update-state-after-modification
	resourcePersonRead(ctx, d, m)

	return diags
}

func resourcePersonRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func resourcePersonRead *****")
	client := m.(*persondbclient.Client)
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 3 || parts[1] != "person" {
		return diag.Errorf("invalid ID format: expected '/person/<person_id>', got: %s", d.Id())
	}
	personID := parts[2]
	lastName, firstName, err := client.ReadPerson(personID)
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
	client := m.(*persondbclient.Client)
	personID := d.Get("person_id").(string)
	lastName := d.Get("last_name").(string)
	firstName := d.Get("first_name").(string)
	err := client.UpdatePerson(personID, lastName, firstName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to update person",
			Detail:   err.Error(),
		})
		return diags
	}

	// Best practice: Update state after modification
	// https://developer.hashicorp.com/terraform/plugin/sdkv2/best-practices/detecting-drift#update-state-after-modification
	resourcePersonRead(ctx, d, m)

	return diags
}

func resourcePersonDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	tflog.Info(ctx, "***** func resourcePersonDelete *****")
	client := m.(*persondbclient.Client)
	personID := d.Get("person_id").(string)
	err := client.DeletePerson(personID)
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
