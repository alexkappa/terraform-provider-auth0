package auth0

import (
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5/management"
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
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							ForceNew:     true,
							RequiredWith: []string{"sink.0.aws_region"},
						},
						"aws_region": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							RequiredWith: []string{"sink.0.aws_account_id"},
						},
						"aws_partner_event_source": {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Optional:    true,
							Description: "Name of the Partner Event Source to be used with AWS, if the type is 'eventbridge'",
						},
						"azure_subscription_id": {
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							ForceNew:     true,
							RequiredWith: []string{"sink.0.azure_resource_group", "sink.0.azure_region"},
						},
						"azure_resource_group": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							RequiredWith: []string{"sink.0.azure_subscription_id", "sink.0.azure_region"},
						},
						"azure_region": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							RequiredWith: []string{"sink.0.azure_subscription_id", "sink.0.azure_resource_group"},
						},
						"azure_partner_topic": {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Optional:    true,
							Description: "Name of the Partner Topic to be used with Azure, if the type is 'eventgrid'",
						},
						"http_content_format": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: []string{"sink.0.http_endpoint", "sink.0.http_authorization", "sink.0.http_content_type"},
							Description:  "HTTP Content Format can be JSONLINES or JSONARRAY",
							ValidateFunc: validation.StringInSlice([]string{
								"JSONLINES", "JSONARRAY"}, false),
						},
						"http_content_type": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "HTTP Content Type",
							RequiredWith: []string{"sink.0.http_endpoint", "sink.0.http_authorization", "sink.0.http_content_format"},
						},
						"http_endpoint": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "HTTP endpoint",
							RequiredWith: []string{"sink.0.http_content_format", "sink.0.http_authorization", "sink.0.http_content_type"},
						},
						"http_authorization": {
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							RequiredWith: []string{"sink.0.http_content_format", "sink.0.http_endpoint", "sink.0.http_content_type"},
						},
						"http_custom_headers": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Default:     nil,
							Description: "custom HTTP headers",
						},

						"datadog_region": {
							Type:         schema.TypeString,
							Sensitive:    true,
							Optional:     true,
							RequiredWith: []string{"sink.0.datadog_api_key"},
						},
						"datadog_api_key": {
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							RequiredWith: []string{"sink.0.datadog_region"},
						},
						"splunk_domain": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: []string{"sink.0.splunk_token", "sink.0.splunk_port", "sink.0.splunk_secure"},
						},
						"splunk_token": {
							Type:         schema.TypeString,
							Optional:     true,
							Sensitive:    true,
							RequiredWith: []string{"sink.0.splunk_domain", "sink.0.splunk_port", "sink.0.splunk_secure"},
						},
						"splunk_port": {
							Type:         schema.TypeString,
							Optional:     true,
							RequiredWith: []string{"sink.0.splunk_domain", "sink.0.splunk_token", "sink.0.splunk_secure"},
						},
						"splunk_secure": {
							Type:         schema.TypeBool,
							Optional:     true,
							Default:      nil,
							RequiredWith: []string{"sink.0.splunk_domain", "sink.0.splunk_port", "sink.0.splunk_token"},
						},
					},
				},
			},
		},
	}
}

func createLogStream(d *schema.ResourceData, m interface{}) error {
	ls := expandLogStream(d)

	api := m.(*management.Management)
	if err := api.LogStream.Create(ls); err != nil {
		return err
	}
	d.SetId(ls.GetID())

	// The Management API only allows updating a log stream's status. Therefore
	// if the status field was present in the configuration, we perform an
	// additional operation to modify it.
	s := String(d, "status")
	if s != nil && s != ls.Status {
		err := api.LogStream.Update(ls.GetID(), &management.LogStream{Status: s})
		if err != nil {
			return err
		}
	}
	return readLogStream(d, m)
}

func readLogStream(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	ls, err := api.LogStream.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(ls.GetID())
	d.Set("name", ls.Name)
	d.Set("status", ls.Status)
	d.Set("type", ls.Type)
	d.Set("sink", flattenLogStreamSink(d, ls.Sink))
	return nil
}

func updateLogStream(d *schema.ResourceData, m interface{}) error {
	ls := expandLogStream(d)

	api := m.(*management.Management)
	err := api.LogStream.Update(d.Id(), ls)
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
	case *management.LogStreamSinkAmazonEventBridge:
		m = flattenLogStreamSinkAmazonEventBridge(o)
	case *management.LogStreamSinkAzureEventGrid:
		m = flattenLogStreamSinkAzureEventGrid(o)
	case *management.LogStreamSinkHTTP:
		m = flattenLogStreamSinkHTTP(o)
	case *management.LogStreamSinkDatadog:
		m = flattenLogStreamSinkDatadog(o)
	case *management.LogStreamSinkSplunk:
		m = flattenLogStreamSinkSplunk(o)
	}
	return []interface{}{m}
}

func flattenLogStreamSinkAmazonEventBridge(o *management.LogStreamSinkAmazonEventBridge) interface{} {
	return map[string]interface{}{
		"aws_account_id":           o.GetAccountID(),
		"aws_region":               o.GetRegion(),
		"aws_partner_event_source": o.GetPartnerEventSource(),
	}
}

func flattenLogStreamSinkAzureEventGrid(o *management.LogStreamSinkAzureEventGrid) interface{} {
	return map[string]interface{}{
		"azure_subscription_id": o.GetSubscriptionID(),
		"azure_resource_group":  o.GetResourceGroup(),
		"azure_region":          o.GetRegion(),
		"azure_partner_topic":   o.GetPartnerTopic(),
	}
}

func flattenLogStreamSinkHTTP(o *management.LogStreamSinkHTTP) interface{} {
	return map[string]interface{}{
		"http_endpoint":       o.GetEndpoint(),
		"http_contentFormat":  o.GetContentFormat(),
		"http_contentType":    o.GetContentType(),
		"http_authorization":  o.GetAuthorization(),
		"http_custom_headers": o.CustomHeaders,
	}
}

func flattenLogStreamSinkDatadog(o *management.LogStreamSinkDatadog) interface{} {
	return map[string]interface{}{
		"datadog_region":  o.GetRegion(),
		"datadog_api_key": o.GetAPIKey(),
	}
}

func flattenLogStreamSinkSplunk(o *management.LogStreamSinkSplunk) interface{} {
	return map[string]interface{}{
		"splunk_domain": o.GetDomain(),
		"splunk_token":  o.GetToken(),
		"splunk_port":   o.GetPort(),
		"splunk_secure": o.GetSecure(),
	}
}
func expandLogStream(d ResourceData) *management.LogStream {

	ls := &management.LogStream{
		Name:   String(d, "name", HasChange()),
		Type:   String(d, "type", HasChange()),
		Status: String(d, "status", Not(IsNewResource())),
	}

	s := d.Get("type").(string)

	List(d, "sink").Elem(func(d ResourceData) {
		switch s {
		case management.LogStreamTypeAmazonEventBridge:
			// LogStreamTypeAmazonEventBridge cannot be updated
			if d.IsNewResource() {
				ls.Sink = expandLogStreamSinkAmazonEventBridge(d)
			}
		case management.LogStreamTypeAzureEventGrid:
			// LogStreamTypeAzureEventGrid cannot be updated
			if d.IsNewResource() {
				ls.Sink = expandLogStreamSinkAzureEventGrid(d)
			}
		case management.LogStreamTypeHTTP:
			ls.Sink = expandLogStreamSinkHTTP(d)
		case management.LogStreamTypeDatadog:
			ls.Sink = expandLogStreamSinkDatadog(d)
		case management.LogStreamTypeSplunk:
			ls.Sink = expandLogStreamSinkSplunk(d)
		default:
			log.Printf("[WARN]: Unsupported log stream sink %s", s)
			log.Printf("[WARN]: Raise an issue with the auth0 provider in order to support it:")
			log.Printf("[WARN]: 	https://github.com/alexkappa/terraform-provider-auth0/issues/new")
		}
	})

	return ls
}

func expandLogStreamSinkAmazonEventBridge(d ResourceData) *management.LogStreamSinkAmazonEventBridge {
	o := &management.LogStreamSinkAmazonEventBridge{
		AccountID:          String(d, "aws_account_id"),
		Region:             String(d, "aws_region"),
		PartnerEventSource: String(d, "aws_partner_event_source"),
	}
	return o
}

func expandLogStreamSinkAzureEventGrid(d ResourceData) *management.LogStreamSinkAzureEventGrid {
	o := &management.LogStreamSinkAzureEventGrid{
		SubscriptionID: String(d, "azure_subscription_id"),
		ResourceGroup:  String(d, "azure_resource_group"),
		Region:         String(d, "azure_region"),
		PartnerTopic:   String(d, "azure_partner_topic"),
	}
	return o
}

func expandLogStreamSinkHTTP(d ResourceData) *management.LogStreamSinkHTTP {
	o := &management.LogStreamSinkHTTP{
		ContentFormat: String(d, "http_content_format"),
		ContentType:   String(d, "http_content_type"),
		Endpoint:      String(d, "http_endpoint"),
		Authorization: String(d, "http_authorization"),
		CustomHeaders: Set(d, "http_custom_headers").List(),
	}
	return o
}
func expandLogStreamSinkDatadog(d ResourceData) *management.LogStreamSinkDatadog {
	o := &management.LogStreamSinkDatadog{
		Region: String(d, "datadog_region"),
		APIKey: String(d, "datadog_api_key"),
	}
	return o
}
func expandLogStreamSinkSplunk(d ResourceData) *management.LogStreamSinkSplunk {
	o := &management.LogStreamSinkSplunk{
		Domain: String(d, "splunk_domain"),
		Token:  String(d, "splunk_token"),
		Port:   String(d, "splunk_port"),
		Secure: Bool(d, "splunk_secure"),
	}
	return o
}
