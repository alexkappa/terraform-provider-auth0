package auth0

import (
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
)

func newLogStream() *schema.Resource {
	return &schema.Resource{

		Create: createLogStream,
		Read:   readLogStream,
		Update: updateLogStream,
		Delete: deleteLogStream,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"eventbridge", "eventgrid", "http", "datadog", "splunk"}, true),
				ForceNew:    true,
				Description: "Type of the LogStream, which indicates the Sink provider",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"active", "paused", "suspended"}, false),
				Description: "Status of the LogStream",
			},
			"sink": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aws_account_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"aws_region": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"azure_subscription_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"azure_resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"azure_region": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"http_content_format": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"JSONLINES", "JSONARRAY"}, false),
						},
						"http_content_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_endpoint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_authorization": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"http_custom_headers": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Description: "IDs of the clients for which the connection is enabled",
						},

						"datadog_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"datadog_api_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"splunk_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"splunk_token": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"splunk_port": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"splunk_secure": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func createLogStream(d *schema.ResourceData, m interface{}) error {
	c := expandLogStream(d)
	api := m.(*management.Management)
	if err := api.LogStream.Create(c); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return readLogStream(d, m)
}

func readLogStream(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.LogStream.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(auth0.StringValue(c.ID))
	d.Set("name", c.Name)
	d.Set("status", c.Status)
	d.Set("type", c.Type)
	d.Set("sink", flattenLogStreamSink(d, c.Sink))
	return nil
}

func updateLogStream(d *schema.ResourceData, m interface{}) error {
	c := expandLogStream(d)
	api := m.(*management.Management)
	err := api.LogStream.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readLogStream(d, m)
}

func deleteLogStream(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.LogStream.Delete(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
	}
	return err
}

func flattenLogStreamSink(d ResourceData, sink interface{}) []interface{} {

	var m interface{}

	switch o := sink.(type) {
	case *management.EventBridgeSink:
		m = flattenLogStreamEventBridgeSink(o)
	case *management.EventGridSink:
		m = flattenLogStreamEventGridSink(o)
	case *management.HTTPSink:
		m = flattenLogStreamHTTPSink(o)
	case *management.DatadogSink:
		m = flattenLogStreamDatadogSink(o)
	case *management.SplunkSink:
		m = flattenLogStreamSplunkSink(o)
	}
	return []interface{}{m}
}

func flattenLogStreamEventBridgeSink(o *management.EventBridgeSink) interface{} {
	return map[string]interface{}{
		"aws_account_id":           o.GetAWSAccountID(),
		"aws_region":               o.GetAWSRegion(),
		"aws_partner_event_source": o.GetAWSPartnerEventSource(),
	}
}

func flattenLogStreamEventGridSink(o *management.EventGridSink) interface{} {
	return map[string]interface{}{
		"azure_subscription_id": o.GetAzureSubscriptionID(),
		"azure_resource_group":  o.GetAzureResourceGroup(),
		"azure_region":          o.GetAzureRegion(),
		"azure_partner_topic":   o.GetAzurePartnerTopic(),
	}
}

func flattenLogStreamHTTPSink(o *management.HTTPSink) interface{} {
	return map[string]interface{}{
		"http_endpoint":       o.GetHTTPEndpoint(),
		"http_contentFormat":  o.GetHTTPContentFormat(),
		"http_contentType":    o.GetHTTPContentType(),
		"http_authorization":  o.GetHTTPAuthorization(),
		"http_custom_headers": o.HTTPCustomHeaders,
	}
}

func flattenLogStreamDatadogSink(o *management.DatadogSink) interface{} {
	return map[string]interface{}{
		"datadog_region":  o.GetDatadogRegion(),
		"datadog_api_key": o.GetDatadogAPIKey(),
	}
}

func flattenLogStreamSplunkSink(o *management.SplunkSink) interface{} {
	return map[string]interface{}{
		"splunk_domain": o.GetSplunkDomain(),
		"splunk_token":  o.GetSplunkToken(),
		"splunk_port":   o.GetSplunkPort(),
		"splunk_secure": o.GetSplunkSecure(),
	}
}
func expandLogStream(d ResourceData) *management.LogStream {

	c := &management.LogStream{
		Name:   String(d, "name", IsNewResource()),
		Type:   String(d, "type", IsNewResource()),
		Status: String(d, "status"),
	}

	s := d.Get("type").(string)

	List(d, "sink").Elem(func(d ResourceData) {
		switch s {
		case management.LogStreamSinkEventBridge:
			c.Sink = expandLogStreamEventBridgeSink(d)
		case management.LogStreamSinkEventGrid:
			c.Sink = expandLogStreamEventGridSink(d)
		case management.LogStreamSinkHTTP:
			c.Sink = expandLogStreamHTTPSink(d)
		case management.LogStreamSinkDatadog:
			c.Sink = expandLogStreamDatadogSink(d)
		case management.LogStreamSinkSplunk:
			c.Sink = expandLogStreamSplunkSink(d)
		default:
			log.Printf("[WARN]: Raise an issue with the auth0 provider in order to support it:")
			log.Printf("[WARN]: 	https://github.com/alexkappa/terraform-provider-auth0/issues/new")
		}
	})

	return c
}

func expandLogStreamEventBridgeSink(d ResourceData) *management.EventBridgeSink {
	o := &management.EventBridgeSink{
		AWSAccountID:          String(d, "aws_account_id"),
		AWSRegion:             String(d, "aws_region"),
		AWSPartnerEventSource: String(d, "aws_partner-event_source"),
	}
	return o
}

func expandLogStreamEventGridSink(d ResourceData) *management.EventGridSink {
	o := &management.EventGridSink{
		AzureSubscriptionID: String(d, "azure_subscription_id"),
		AzureResourceGroup:  String(d, "azure_resource_group"),
		AzureRegion:         String(d, "azure_region"),
		AzurePartnerTopic:   String(d, "azure_partner_topic"),
	}
	return o
}

func expandLogStreamHTTPSink(d ResourceData) *management.HTTPSink {
	o := &management.HTTPSink{
		HTTPContentFormat: String(d, "http_content_format"),
		HTTPContentType:   String(d, "http_content_type"),
		HTTPEndpoint:      String(d, "http_endpoint"),
		HTTPAuthorization: String(d, "http_authorization"),
		HTTPCustomHeaders: Set(d, "http_custom_headers").List(),
	}
	return o
}
func expandLogStreamDatadogSink(d ResourceData) *management.DatadogSink {
	o := &management.DatadogSink{
		DatadogRegion: String(d, "datadog_region"),
		DatadogAPIKey: String(d, "datadog_api_key"),
	}
	return o
}
func expandLogStreamSplunkSink(d ResourceData) *management.SplunkSink {
	o := &management.SplunkSink{
		SplunkDomain: String(d, "splunk_domain"),
		SplunkToken:  String(d, "splunk_token"),
		SplunkPort:   String(d, "splunk_port"),
		SplunkSecure: Bool(d, "splunk_secure"),
	}
	return o
}
