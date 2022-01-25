provider "auth0" {}

resource "auth0_log_stream" "example_http" {
  name = "HTTP log stream"
  type = "http"
  sink {
    http_endpoint       = "https://example.com/logs"
    http_content_type   = "application/json"
    http_content_format = "JSONOBJECT"
    http_authorization  = "AKIAXXXXXXXXXXXXXXXX"
  }
}

resource "auth0_log_stream" "example_aws" {
  name   = "AWS Eventbridge"
  type   = "eventbridge"
  status = "active"
  sink {
    aws_account_id = "my_account_id"
    aws_region     = "us-east-2"
  }
}
