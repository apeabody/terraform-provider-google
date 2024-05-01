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

package appengine

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func sslSettingsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	// If certificate id is empty, and ssl management type is `MANUAL`, then
	// ssl settings will not be configured, and ssl_settings block is not returned

	if k == "ssl_settings.#" &&
		old == "0" && new == "1" &&
		d.Get("ssl_settings.0.certificate_id") == "" &&
		d.Get("ssl_settings.0.ssl_management_type") == "MANUAL" {
		return true
	}

	return false
}

func ResourceAppEngineDomainMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineDomainMappingCreate,
		Read:   resourceAppEngineDomainMappingRead,
		Update: resourceAppEngineDomainMappingUpdate,
		Delete: resourceAppEngineDomainMappingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineDomainMappingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Relative name of the domain serving the application. Example: example.com.`,
			},
			"override_strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"STRICT", "OVERRIDE", ""}),
				Description: `Whether the domain creation should override any existing mappings for this domain.
By default, overrides are rejected. Default value: "STRICT" Possible values: ["STRICT", "OVERRIDE"]`,
				Default: "STRICT",
			},
			"ssl_settings": {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: sslSettingsDiffSuppress,
				Description:      `SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.`,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_management_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"AUTOMATIC", "MANUAL"}),
							Description: `SSL management type for this domain. If 'AUTOMATIC', a managed certificate is automatically provisioned.
If 'MANUAL', 'certificateId' must be manually specified in order to configure SSL for this domain. Possible values: ["AUTOMATIC", "MANUAL"]`,
						},
						"certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will
remove SSL support.
By default, a managed certificate is automatically created for every domain mapping. To omit SSL support
or to configure SSL manually, specify 'SslManagementType.MANUAL' on a 'CREATE' or 'UPDATE' request. You must be
authorized to administer the 'AuthorizedCertificate' resource to manually map it to a DomainMapping resource.
Example: 12345.`,
						},
						"pending_managed_certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `ID of the managed 'AuthorizedCertificate' resource currently being provisioned, if applicable. Until the new
managed certificate has been successfully provisioned, the previous SSL state will be preserved. Once the
provisioning process completes, the 'certificateId' field will reflect the new managed certificate and this
field will be left empty. To remove SSL support while there is still a pending managed certificate, clear the
'certificateId' field with an update request.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.`,
			},
			"resource_records": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `The resource records required to configure this domain mapping. These records must be added to the domain's DNS
configuration in order to serve the application via this domain mapping.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Relative name of the object affected by this record. Only applicable for CNAME records. Example: 'www'.`,
						},
						"rrdata": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Data for this record. Values vary by record type, as defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1).`,
						},
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"A", "AAAA", "CNAME", ""}),
							Description:  `Resource record type. Example: 'AAAA'. Possible values: ["A", "AAAA", "CNAME"]`,
						},
					},
				},
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

func resourceAppEngineDomainMappingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	sslSettingsProp, err := expandAppEngineDomainMappingSslSettings(d.Get("ssl_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(sslSettingsProp)) && (ok || !reflect.DeepEqual(v, sslSettingsProp)) {
		obj["sslSettings"] = sslSettingsProp
	}
	idProp, err := expandAppEngineDomainMappingDomainName(d.Get("domain_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("domain_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DomainMapping: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DomainMapping: %s", err)
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
		return fmt.Errorf("Error creating DomainMapping: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = AppEngineOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating DomainMapping", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create DomainMapping: %s", err)
	}

	if err := d.Set("name", flattenAppEngineDomainMappingName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DomainMapping %q: %#v", d.Id(), res)

	return resourceAppEngineDomainMappingRead(d, meta)
}

func resourceAppEngineDomainMappingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DomainMapping: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AppEngineDomainMapping %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	if err := d.Set("name", flattenAppEngineDomainMappingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("ssl_settings", flattenAppEngineDomainMappingSslSettings(res["sslSettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("resource_records", flattenAppEngineDomainMappingResourceRecords(res["resourceRecords"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}
	if err := d.Set("domain_name", flattenAppEngineDomainMappingDomainName(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading DomainMapping: %s", err)
	}

	return nil
}

func resourceAppEngineDomainMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DomainMapping: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	sslSettingsProp, err := expandAppEngineDomainMappingSslSettings(d.Get("ssl_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslSettingsProp)) {
		obj["sslSettings"] = sslSettingsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DomainMapping %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("ssl_settings") {
		updateMask = append(updateMask, "ssl_settings.certificate_id",
			"ssl_settings.ssl_management_type")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
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
			return fmt.Errorf("Error updating DomainMapping %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating DomainMapping %q: %#v", d.Id(), res)
		}

		err = AppEngineOperationWaitTime(
			config, res, project, "Updating DomainMapping", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceAppEngineDomainMappingRead(d, meta)
}

func resourceAppEngineDomainMappingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for DomainMapping: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting DomainMapping %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "DomainMapping")
	}

	err = AppEngineOperationWaitTime(
		config, res, project, "Deleting DomainMapping", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting DomainMapping %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineDomainMappingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^apps/(?P<project>[^/]+)/domainMappings/(?P<domain_name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<domain_name>[^/]+)$",
		"^(?P<domain_name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}/domainMappings/{{domain_name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineDomainMappingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineDomainMappingSslSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["certificate_id"] =
		flattenAppEngineDomainMappingSslSettingsCertificateId(original["certificateId"], d, config)
	transformed["ssl_management_type"] =
		flattenAppEngineDomainMappingSslSettingsSslManagementType(original["sslManagementType"], d, config)
	transformed["pending_managed_certificate_id"] =
		flattenAppEngineDomainMappingSslSettingsPendingManagedCertificateId(original["pendingManagedCertificateId"], d, config)
	return []interface{}{transformed}
}
func flattenAppEngineDomainMappingSslSettingsCertificateId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineDomainMappingSslSettingsSslManagementType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineDomainMappingSslSettingsPendingManagedCertificateId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineDomainMappingResourceRecords(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":   flattenAppEngineDomainMappingResourceRecordsName(original["name"], d, config),
			"rrdata": flattenAppEngineDomainMappingResourceRecordsRrdata(original["rrdata"], d, config),
			"type":   flattenAppEngineDomainMappingResourceRecordsType(original["type"], d, config),
		})
	}
	return transformed
}
func flattenAppEngineDomainMappingResourceRecordsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineDomainMappingResourceRecordsRrdata(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineDomainMappingResourceRecordsType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineDomainMappingDomainName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAppEngineDomainMappingSslSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificateId, err := expandAppEngineDomainMappingSslSettingsCertificateId(original["certificate_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificateId"] = transformedCertificateId
	}

	transformedSslManagementType, err := expandAppEngineDomainMappingSslSettingsSslManagementType(original["ssl_management_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSslManagementType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sslManagementType"] = transformedSslManagementType
	}

	transformedPendingManagedCertificateId, err := expandAppEngineDomainMappingSslSettingsPendingManagedCertificateId(original["pending_managed_certificate_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPendingManagedCertificateId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pendingManagedCertificateId"] = transformedPendingManagedCertificateId
	}

	return transformed, nil
}

func expandAppEngineDomainMappingSslSettingsCertificateId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingSslSettingsSslManagementType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingSslSettingsPendingManagedCertificateId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineDomainMappingDomainName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
