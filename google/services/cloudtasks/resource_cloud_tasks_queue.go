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

package cloudtasks

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
)

func suppressOmittedMaxDuration(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "0s" {
		log.Printf("[INFO] max retry is 0s and api omitted field, suppressing diff")
		return true
	}
	return tpgresource.DurationDiffSuppress(k, old, new, d)
}

func ResourceCloudTasksQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudTasksQueueCreate,
		Read:   resourceCloudTasksQueueRead,
		Update: resourceCloudTasksQueueUpdate,
		Delete: resourceCloudTasksQueueDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudTasksQueueImport,
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
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the queue`,
			},
			"app_engine_routing_override": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Overrides for task-level appEngineRouting. These settings apply only
to App Engine tasks in this queue`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `App instance.

By default, the task is sent to an instance which is available when the task is attempted.`,
						},
						"service": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `App service.

By default, the task is sent to the service which is the default service when the task is attempted.`,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `App version.

By default, the task is sent to the version which is the default version when the task is attempted.`,
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The host that the task is sent to.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The queue name.`,
			},
			"rate_limits": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Description: `Rate limits for task dispatches.

The queue's actual dispatch rate is the result of:

* Number of tasks in the queue
* User-specified throttling: rateLimits, retryConfig, and the queue's state.
* System throttling due to 429 (Too Many Requests) or 503 (Service
  Unavailable) responses from the worker, high error rates, or to
  smooth sudden large traffic spikes.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_concurrent_dispatches": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `The maximum number of concurrent tasks that Cloud Tasks allows to
be dispatched for this queue. After this threshold has been
reached, Cloud Tasks stops dispatching tasks until the number of
concurrent requests decreases.`,
						},
						"max_dispatches_per_second": {
							Type:     schema.TypeFloat,
							Computed: true,
							Optional: true,
							Description: `The maximum rate at which tasks are dispatched from this queue.

If unspecified when the queue is created, Cloud Tasks will pick the default.`,
						},
						"max_burst_size": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: `The max burst size.

Max burst size limits how fast tasks in queue are processed when many tasks are
in the queue and the rate is high. This field allows the queue to have a high
rate so processing starts shortly after a task is enqueued, but still limits
resource usage when many tasks are enqueued in a short period of time.`,
						},
					},
				},
			},
			"retry_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Settings that determine the retry behavior.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `Number of attempts per task.

Cloud Tasks will attempt the task maxAttempts times (that is, if
the first attempt fails, then there will be maxAttempts - 1
retries). Must be >= -1.

If unspecified when the queue is created, Cloud Tasks will pick
the default.

-1 indicates unlimited attempts.`,
						},
						"max_backoff": {
							Type:             schema.TypeString,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: tpgresource.DurationDiffSuppress,
							Description: `A task will be scheduled for retry between minBackoff and
maxBackoff duration after it fails, if the queue's RetryConfig
specifies that the task should be retried.`,
						},
						"max_doublings": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `The time between retries will double maxDoublings times.

A task's retry interval starts at minBackoff, then doubles maxDoublings times,
then increases linearly, and finally retries retries at intervals of maxBackoff
up to maxAttempts times.`,
						},
						"max_retry_duration": {
							Type:             schema.TypeString,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: suppressOmittedMaxDuration,
							Description: `If positive, maxRetryDuration specifies the time limit for
retrying a failed task, measured from when the task was first
attempted. Once maxRetryDuration time has passed and the task has
been attempted maxAttempts times, no further attempts will be
made and the task will be deleted.

If zero, then the task age is unlimited.`,
						},
						"min_backoff": {
							Type:             schema.TypeString,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: tpgresource.DurationDiffSuppress,
							Description: `A task will be scheduled for retry between minBackoff and
maxBackoff duration after it fails, if the queue's RetryConfig
specifies that the task should be retried.`,
						},
					},
				},
			},
			"stackdriver_logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration options for writing logs to Stackdriver Logging.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sampling_ratio": {
							Type:     schema.TypeFloat,
							Required: true,
							Description: `Specifies the fraction of operations to write to Stackdriver Logging.
This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
default and means that no operations are logged.`,
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

func resourceCloudTasksQueueCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandCloudTasksQueueName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	appEngineRoutingOverrideProp, err := expandCloudTasksQueueAppEngineRoutingOverride(d.Get("app_engine_routing_override"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_routing_override"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineRoutingOverrideProp)) && (ok || !reflect.DeepEqual(v, appEngineRoutingOverrideProp)) {
		obj["appEngineRoutingOverride"] = appEngineRoutingOverrideProp
	}
	rateLimitsProp, err := expandCloudTasksQueueRateLimits(d.Get("rate_limits"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rate_limits"); !tpgresource.IsEmptyValue(reflect.ValueOf(rateLimitsProp)) && (ok || !reflect.DeepEqual(v, rateLimitsProp)) {
		obj["rateLimits"] = rateLimitsProp
	}
	retryConfigProp, err := expandCloudTasksQueueRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retry_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(retryConfigProp)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}
	stackdriverLoggingConfigProp, err := expandCloudTasksQueueStackdriverLoggingConfig(d.Get("stackdriver_logging_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("stackdriver_logging_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(stackdriverLoggingConfigProp)) && (ok || !reflect.DeepEqual(v, stackdriverLoggingConfigProp)) {
		obj["stackdriverLoggingConfig"] = stackdriverLoggingConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Queue: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Queue: %s", err)
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
		return fmt.Errorf("Error creating Queue: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Queue %q: %#v", d.Id(), res)

	return resourceCloudTasksQueueRead(d, meta)
}

func resourceCloudTasksQueueRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Queue: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CloudTasksQueue %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}

	if err := d.Set("name", flattenCloudTasksQueueName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}
	if err := d.Set("app_engine_routing_override", flattenCloudTasksQueueAppEngineRoutingOverride(res["appEngineRoutingOverride"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}
	if err := d.Set("rate_limits", flattenCloudTasksQueueRateLimits(res["rateLimits"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}
	if err := d.Set("retry_config", flattenCloudTasksQueueRetryConfig(res["retryConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}
	if err := d.Set("stackdriver_logging_config", flattenCloudTasksQueueStackdriverLoggingConfig(res["stackdriverLoggingConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Queue: %s", err)
	}

	return nil
}

func resourceCloudTasksQueueUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Queue: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	appEngineRoutingOverrideProp, err := expandCloudTasksQueueAppEngineRoutingOverride(d.Get("app_engine_routing_override"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_routing_override"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, appEngineRoutingOverrideProp)) {
		obj["appEngineRoutingOverride"] = appEngineRoutingOverrideProp
	}
	rateLimitsProp, err := expandCloudTasksQueueRateLimits(d.Get("rate_limits"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rate_limits"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, rateLimitsProp)) {
		obj["rateLimits"] = rateLimitsProp
	}
	retryConfigProp, err := expandCloudTasksQueueRetryConfig(d.Get("retry_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retry_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retryConfigProp)) {
		obj["retryConfig"] = retryConfigProp
	}
	stackdriverLoggingConfigProp, err := expandCloudTasksQueueStackdriverLoggingConfig(d.Get("stackdriver_logging_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("stackdriver_logging_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, stackdriverLoggingConfigProp)) {
		obj["stackdriverLoggingConfig"] = stackdriverLoggingConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Queue %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("app_engine_routing_override") {
		updateMask = append(updateMask, "appEngineRoutingOverride")
	}

	if d.HasChange("rate_limits") {
		updateMask = append(updateMask, "rateLimits")
	}

	if d.HasChange("retry_config") {
		updateMask = append(updateMask, "retryConfig")
	}

	if d.HasChange("stackdriver_logging_config") {
		updateMask = append(updateMask, "stackdriverLoggingConfig")
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
			return fmt.Errorf("Error updating Queue %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Queue %q: %#v", d.Id(), res)
		}

	}

	return resourceCloudTasksQueueRead(d, meta)
}

func resourceCloudTasksQueueDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Queue: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Queue %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Queue")
	}

	log.Printf("[DEBUG] Finished deleting Queue %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudTasksQueueImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/queues/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/queues/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudTasksQueueName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

// service, version, and instance are input-only. host is output-only.
func flattenCloudTasksQueueAppEngineRoutingOverride(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] = original["host"]
	if override, ok := d.GetOk("app_engine_routing_override"); ok && len(override.([]interface{})) > 0 {
		transformed["service"] = d.Get("app_engine_routing_override.0.service")
		transformed["version"] = d.Get("app_engine_routing_override.0.version")
		transformed["instance"] = d.Get("app_engine_routing_override.0.instance")
	}
	return []interface{}{transformed}
}

func flattenCloudTasksQueueRateLimits(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["max_dispatches_per_second"] =
		flattenCloudTasksQueueRateLimitsMaxDispatchesPerSecond(original["maxDispatchesPerSecond"], d, config)
	transformed["max_concurrent_dispatches"] =
		flattenCloudTasksQueueRateLimitsMaxConcurrentDispatches(original["maxConcurrentDispatches"], d, config)
	transformed["max_burst_size"] =
		flattenCloudTasksQueueRateLimitsMaxBurstSize(original["maxBurstSize"], d, config)
	return []interface{}{transformed}
}
func flattenCloudTasksQueueRateLimitsMaxDispatchesPerSecond(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudTasksQueueRateLimitsMaxConcurrentDispatches(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenCloudTasksQueueRateLimitsMaxBurstSize(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenCloudTasksQueueRetryConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["max_attempts"] =
		flattenCloudTasksQueueRetryConfigMaxAttempts(original["maxAttempts"], d, config)
	transformed["max_retry_duration"] =
		flattenCloudTasksQueueRetryConfigMaxRetryDuration(original["maxRetryDuration"], d, config)
	transformed["min_backoff"] =
		flattenCloudTasksQueueRetryConfigMinBackoff(original["minBackoff"], d, config)
	transformed["max_backoff"] =
		flattenCloudTasksQueueRetryConfigMaxBackoff(original["maxBackoff"], d, config)
	transformed["max_doublings"] =
		flattenCloudTasksQueueRetryConfigMaxDoublings(original["maxDoublings"], d, config)
	return []interface{}{transformed}
}
func flattenCloudTasksQueueRetryConfigMaxAttempts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenCloudTasksQueueRetryConfigMaxRetryDuration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudTasksQueueRetryConfigMinBackoff(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudTasksQueueRetryConfigMaxBackoff(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudTasksQueueRetryConfigMaxDoublings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenCloudTasksQueueStackdriverLoggingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["sampling_ratio"] =
		flattenCloudTasksQueueStackdriverLoggingConfigSamplingRatio(original["samplingRatio"], d, config)
	return []interface{}{transformed}
}
func flattenCloudTasksQueueStackdriverLoggingConfigSamplingRatio(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCloudTasksQueueName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/queues/{{name}}")
}

func expandCloudTasksQueueAppEngineRoutingOverride(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandCloudTasksQueueAppEngineRoutingOverrideService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandCloudTasksQueueAppEngineRoutingOverrideVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedInstance, err := expandCloudTasksQueueAppEngineRoutingOverrideInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedHost, err := expandCloudTasksQueueAppEngineRoutingOverrideHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	return transformed, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueAppEngineRoutingOverrideHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxDispatchesPerSecond, err := expandCloudTasksQueueRateLimitsMaxDispatchesPerSecond(original["max_dispatches_per_second"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDispatchesPerSecond); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDispatchesPerSecond"] = transformedMaxDispatchesPerSecond
	}

	transformedMaxConcurrentDispatches, err := expandCloudTasksQueueRateLimitsMaxConcurrentDispatches(original["max_concurrent_dispatches"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxConcurrentDispatches); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxConcurrentDispatches"] = transformedMaxConcurrentDispatches
	}

	transformedMaxBurstSize, err := expandCloudTasksQueueRateLimitsMaxBurstSize(original["max_burst_size"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBurstSize); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxBurstSize"] = transformedMaxBurstSize
	}

	return transformed, nil
}

func expandCloudTasksQueueRateLimitsMaxDispatchesPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimitsMaxConcurrentDispatches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRateLimitsMaxBurstSize(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxAttempts, err := expandCloudTasksQueueRetryConfigMaxAttempts(original["max_attempts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxAttempts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxAttempts"] = transformedMaxAttempts
	}

	transformedMaxRetryDuration, err := expandCloudTasksQueueRetryConfigMaxRetryDuration(original["max_retry_duration"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetryDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxRetryDuration"] = transformedMaxRetryDuration
	}

	transformedMinBackoff, err := expandCloudTasksQueueRetryConfigMinBackoff(original["min_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinBackoff); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minBackoff"] = transformedMinBackoff
	}

	transformedMaxBackoff, err := expandCloudTasksQueueRetryConfigMaxBackoff(original["max_backoff"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxBackoff); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxBackoff"] = transformedMaxBackoff
	}

	transformedMaxDoublings, err := expandCloudTasksQueueRetryConfigMaxDoublings(original["max_doublings"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDoublings); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDoublings"] = transformedMaxDoublings
	}

	return transformed, nil
}

func expandCloudTasksQueueRetryConfigMaxAttempts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxRetryDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMinBackoff(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxBackoff(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueRetryConfigMaxDoublings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudTasksQueueStackdriverLoggingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSamplingRatio, err := expandCloudTasksQueueStackdriverLoggingConfigSamplingRatio(original["sampling_ratio"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSamplingRatio); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["samplingRatio"] = transformedSamplingRatio
	}

	return transformed, nil
}

func expandCloudTasksQueueStackdriverLoggingConfigSamplingRatio(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
