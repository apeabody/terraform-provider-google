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

package compute

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeDiskResourcePolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeDiskResourcePolicyAttachmentCreate,
		Read:   resourceComputeDiskResourcePolicyAttachmentRead,
		Delete: resourceComputeDiskResourcePolicyAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeDiskResourcePolicyAttachmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
			tpgresource.DefaultProviderZone,
		),

		Schema: map[string]*schema.Schema{
			"disk": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the disk in which the resource policies are attached to.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource policy to be attached to the disk for scheduling snapshot
creation. Do not specify the self link.`,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to the zone where the disk resides.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeDiskResourcePolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNestedComputeDiskResourcePolicyAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceComputeDiskResourcePolicyAttachmentEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/disks/{{disk}}/addResourcePolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DiskResourcePolicyAttachment: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DiskResourcePolicyAttachment: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating DiskResourcePolicyAttachment: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{disk}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating DiskResourcePolicyAttachment", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create DiskResourcePolicyAttachment: %s", err)
	}

	log.Printf("[DEBUG] Finished creating DiskResourcePolicyAttachment %q: %#v", d.Id(), res)

	return resourceComputeDiskResourcePolicyAttachmentRead(d, meta)
}

func resourceComputeDiskResourcePolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/disks/{{disk}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DiskResourcePolicyAttachment: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeDiskResourcePolicyAttachment %q", d.Id()))
	}

	res, err = flattenNestedComputeDiskResourcePolicyAttachment(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeDiskResourcePolicyAttachment because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceComputeDiskResourcePolicyAttachmentDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeDiskResourcePolicyAttachment because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DiskResourcePolicyAttachment: %s", err)
	}

	zone, err := tpgresource.GetZone(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("zone", zone); err != nil {
		return fmt.Errorf("Error reading DiskResourcePolicyAttachment: %s", err)
	}

	if err := d.Set("name", flattenNestedComputeDiskResourcePolicyAttachmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DiskResourcePolicyAttachment: %s", err)
	}

	return nil
}

func resourceComputeDiskResourcePolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DiskResourcePolicyAttachment: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/disks/{{disk}}/removeResourcePolicies")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	obj = make(map[string]interface{})

	zone, err := tpgresource.GetZone(d, config)
	if err != nil {
		return err
	}
	if zone == "" {
		return fmt.Errorf("zone must be non-empty - set in resource or at provider-level")
	}

	// resourcePolicies are referred to by region but affixed to zonal disks.
	// We construct the regional name from the zone:
	//
	//	projects/{project}/regions/{region}/resourcePolicies/{resourceId}
	region := tpgresource.GetRegionFromZone(zone)
	if region == "" {
		return fmt.Errorf("invalid zone %q, unable to infer region from zone", zone)
	}

	name, err := expandNestedComputeDiskResourcePolicyAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(name)) && (ok || !reflect.DeepEqual(v, name)) {
		obj["resourcePolicies"] = []interface{}{fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", project, region, name)}
	}

	log.Printf("[DEBUG] Deleting DiskResourcePolicyAttachment %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "DiskResourcePolicyAttachment")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting DiskResourcePolicyAttachment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting DiskResourcePolicyAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeDiskResourcePolicyAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/disks/(?P<disk>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<disk>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<zone>[^/]+)/(?P<disk>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<disk>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{disk}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeDiskResourcePolicyAttachmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedComputeDiskResourcePolicyAttachmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeDiskResourcePolicyAttachmentEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	zone, err := tpgresource.GetZone(d, config)
	if err != nil {
		return nil, err
	}
	if zone == "" {
		return nil, fmt.Errorf("zone must be non-empty - set in resource or at provider-level")
	}

	// resourcePolicies are referred to by region but affixed to zonal disks.
	// We construct the regional name from the zone:
	//
	//	projects/{project}/regions/{region}/resourcePolicies/{resourceId}
	region := tpgresource.GetRegionFromZone(zone)
	if region == "" {
		return nil, fmt.Errorf("invalid zone %q, unable to infer region from zone", zone)
	}

	obj["resourcePolicies"] = []interface{}{fmt.Sprintf("projects/%s/regions/%s/resourcePolicies/%s", project, region, obj["name"])}
	delete(obj, "name")
	return obj, nil
}

func flattenNestedComputeDiskResourcePolicyAttachment(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["resourcePolicies"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value resourcePolicies. Actual value: %v", v)
	}

	_, item, err := resourceComputeDiskResourcePolicyAttachmentFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeDiskResourcePolicyAttachmentFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName, err := expandNestedComputeDiskResourcePolicyAttachmentName(d.Get("name"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedName := flattenNestedComputeDiskResourcePolicyAttachmentName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		// List response only contains the ID - construct a response object.
		item := map[string]interface{}{
			"name": itemRaw,
		}

		// Decode list item before comparing.
		item, err := resourceComputeDiskResourcePolicyAttachmentDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemName := flattenNestedComputeDiskResourcePolicyAttachmentName(item["name"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemName)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceComputeDiskResourcePolicyAttachmentDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	res["name"] = tpgresource.GetResourceNameFromSelfLink(res["name"].(string))
	return res, nil
}
