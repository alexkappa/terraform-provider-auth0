package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("auth0_log_stream", &resource.Sweeper{
		Name: "auth0_log_stream",
		/*		F: func(_ string) error {
				api, err := Auth0()
				if err != nil {
					return err
				}
				l, err := api.LogStream.List()
				if err != nil {
					return err
				}
				for _, logstream := range l {
					if strings.Contains(logstream.GetName(), "Test") {
						log.Printf("[DEBUG] Deleting logstream %v\n", logstream.GetName())
						if e := api.LogStream.Delete(logstream.GetID()); e != nil {
							multierror.Append(err, e)
						}
					}
				}
				if err != nil {
					return err
				}
				return nil
			},*/
	})
}

func TestAccLogStreamHttp(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				resource "auth0_log_stream" "my_log_stream" {
					name = "Acceptance-Test-LogStream-http"
					type = "http"
					sink {
						http_endpoint = "https://example.com/webhook/logs"
						http_content_type = "application/json"
						http_content_format = "JSONLINES"
						http_authorization = "AKIAXXXXXXXXXXXXXXXX"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "name", "Acceptance-Test-LogStream-http"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "type", "http"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.http_endpoint", "https://example.com/webhook/logs"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.http_content_type", "application/json"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.http_content_format", "JSONLINES"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.http_authorization", "AKIAXXXXXXXXXXXXXXXX"),
				),
			},
		},
	})
}
func TestAccLogStreamAWS(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				resource "auth0_log_stream" "my_log_stream" {
					name = "Acceptance-Test-LogStream-aws"
					type = "eventbridge"
					sink {
						aws_account_id = "999999999999"
						aws_region = "us-west-2"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "name", "Acceptance-Test-LogStream-aws"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "type", "eventbridge"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.aws_account_id", "999999999999"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.aws_region", "us-west-2"),
				),
			},
		},
	})
}
func TestAccLogStreamAzure(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				resource "auth0_log_stream" "my_log_stream" {
					name = "Acceptance-Test-LogStream-azure"
					type = "eventgrid"
					sink {
						azure_subscription_id = "b69a6835-57c7-4d53-b0d5-1c6ae580b6d5"
						azure_region = "northeurope"
						azure_resource_group = "azure-logs-rg"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "name", "Acceptance-Test-LogStream-azure"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "type", "eventgrid"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.azure_subscription_id", "b69a6835-57c7-4d53-b0d5-1c6ae580b6d5"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.azure_region", "northeurope"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.azure_resource_group", "azure-logs-rg"),
				),
			},
		},
	})
}
func TestAccLogStreamDatadog(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				resource "auth0_log_stream" "my_log_stream" {
					name = "Acceptance-Test-LogStream-datadog"
					type = "datadog"
					sink {
						datadog_region = "us"
						datadog_api_key = "121233123455"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "name", "Acceptance-Test-LogStream-datadog"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "type", "datadog"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.datadog_region", "us"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.datadog_api_key", "121233123455"),
				),
			},
		},
	})
}
func TestAccLogStreamSplunk(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
				resource "auth0_log_stream" "my_log_stream" {
					name = "eAcceptance-Test-LogStream-splunk"
					type = "splunk"
					sink {
						splunk_domain = "demo.splunk.com"
						splunk_token = "12a34ab5-c6d7-8901-23ef-456b7c89d0c1"
						splunk_port = "8080"
						splunk_secure = "true"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "name", "Acceptance-Test-LogStream-splunk"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "type", "splunk"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.splunk_domain", "demo.splunk.com"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.splunk_token", "12a34ab5-c6d7-8901-23ef-456b7c89d0c1"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.splunk_port", "8080"),
					resource.TestCheckResourceAttr("auth0_log_stream.my_log_stream", "sink.0.splunk_secure", "true"),
				),
			},
		},
	})
}
