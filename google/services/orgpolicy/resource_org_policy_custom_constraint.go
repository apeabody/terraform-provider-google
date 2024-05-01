// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package orgpolicy

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceOrgPolicyCustomConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrgPolicyCustomConstraintCreate,
		Read:   resourceOrgPolicyCustomConstraintRead,
		Update: resourceOrgPolicyCustomConstraintUpdate,
		Delete: resourceOrgPolicyCustomConstraintDelete,

		Importer: &schema.ResourceImporter{
			State: resourceOrgPolicyCustomConstraintImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"action_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ALLOW", "DENY"}),
				Description:  `The action to take if the condition is met. Possible values: ["ALLOW", "DENY"]`,
			},
			"condition": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `A CEL condition that refers to a supported service resource, for example 'resource.management.autoUpgrade == false'. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).`,
			},
			"method_types": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `A list of RESTful methods for which to enforce the constraint. Can be 'CREATE', 'UPDATE', or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).`,
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Immutable. The name of the custom constraint. This is unique within the organization.`,
			},
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.`,
			},
			"resource_types": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, 'container.googleapis.com/NodePool'.`,
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-friendly description of the constraint to display as an error message when the policy is violated.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-friendly name for the constraint.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The timestamp representing when the constraint was last updated.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceOrgPolicyCustomConstraintCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandOrgPolicyCustomConstraintName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandOrgPolicyCustomConstraintDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandOrgPolicyCustomConstraintDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	conditionProp, err := expandOrgPolicyCustomConstraintCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}
	actionTypeProp, err := expandOrgPolicyCustomConstraintActionType(d.Get("action_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionTypeProp)) && (ok || !reflect.DeepEqual(v, actionTypeProp)) {
		obj["actionType"] = actionTypeProp
	}
	methodTypesProp, err := expandOrgPolicyCustomConstraintMethodTypes(d.Get("method_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("method_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(methodTypesProp)) && (ok || !reflect.DeepEqual(v, methodTypesProp)) {
		obj["methodTypes"] = methodTypesProp
	}
	resourceTypesProp, err := expandOrgPolicyCustomConstraintResourceTypes(d.Get("resource_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTypesProp)) && (ok || !reflect.DeepEqual(v, resourceTypesProp)) {
		obj["resourceTypes"] = resourceTypesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{OrgPolicyBasePath}}{{parent}}/customConstraints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new CustomConstraint: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating CustomConstraint: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/customConstraints/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating CustomConstraint %q: %#v", d.Id(), res)

	return resourceOrgPolicyCustomConstraintRead(d, meta)
}

func resourceOrgPolicyCustomConstraintRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{OrgPolicyBasePath}}{{parent}}/customConstraints/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("OrgPolicyCustomConstraint %q", d.Id()))
	}

	if err := d.Set("name", flattenOrgPolicyCustomConstraintName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}
	if err := d.Set("display_name", flattenOrgPolicyCustomConstraintDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}
	if err := d.Set("description", flattenOrgPolicyCustomConstraintDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}
	if err := d.Set("condition", flattenOrgPolicyCustomConstraintCondition(res["condition"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}
	if err := d.Set("action_type", flattenOrgPolicyCustomConstraintActionType(res["actionType"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}
	if err := d.Set("method_types", flattenOrgPolicyCustomConstraintMethodTypes(res["methodTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}
	if err := d.Set("resource_types", flattenOrgPolicyCustomConstraintResourceTypes(res["resourceTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}
	if err := d.Set("update_time", flattenOrgPolicyCustomConstraintUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CustomConstraint: %s", err)
	}

	return nil
}

func resourceOrgPolicyCustomConstraintUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandOrgPolicyCustomConstraintDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandOrgPolicyCustomConstraintDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	conditionProp, err := expandOrgPolicyCustomConstraintCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}
	actionTypeProp, err := expandOrgPolicyCustomConstraintActionType(d.Get("action_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("action_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, actionTypeProp)) {
		obj["actionType"] = actionTypeProp
	}
	methodTypesProp, err := expandOrgPolicyCustomConstraintMethodTypes(d.Get("method_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("method_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, methodTypesProp)) {
		obj["methodTypes"] = methodTypesProp
	}

	obj, err = resourceOrgPolicyCustomConstraintUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{OrgPolicyBasePath}}{{parent}}/customConstraints/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating CustomConstraint %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating CustomConstraint %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating CustomConstraint %q: %#v", d.Id(), res)
	}

	return resourceOrgPolicyCustomConstraintRead(d, meta)
}

func resourceOrgPolicyCustomConstraintDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{OrgPolicyBasePath}}{{parent}}/customConstraints/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting CustomConstraint %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "CustomConstraint")
	}

	log.Printf("[DEBUG] Finished deleting CustomConstraint %q: %#v", d.Id(), res)
	return nil
}

func resourceOrgPolicyCustomConstraintImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/customConstraints/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/customConstraints/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenOrgPolicyCustomConstraintName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenOrgPolicyCustomConstraintDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOrgPolicyCustomConstraintDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOrgPolicyCustomConstraintCondition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOrgPolicyCustomConstraintActionType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOrgPolicyCustomConstraintMethodTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOrgPolicyCustomConstraintResourceTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenOrgPolicyCustomConstraintUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandOrgPolicyCustomConstraintName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "{{parent}}/customConstraints/{{name}}")
}

func expandOrgPolicyCustomConstraintDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintActionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintMethodTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintResourceTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceOrgPolicyCustomConstraintUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// need to send resource_types in all PATCH requests
	resourceTypesProp := d.Get("resource_types")
	if v, ok := d.GetOkExists("resource_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, resourceTypesProp)) {
		obj["resourceTypes"] = resourceTypesProp
	}

	return obj, nil
}
