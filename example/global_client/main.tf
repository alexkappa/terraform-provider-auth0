provider "auth0" {}

resource "auth0_global_client" "global" {
    // Auth0 Universal Login - Custom Login Page
    custom_login_page_on = true
    custom_login_page = <<PAGE
<html>
    <head><title>My Custom Login Page</title></head>
    <body>
        I should probably have a login form here
    </body>
</html>
PAGE
    callbacks = [ "http://somehostname.com/a/callback" ]
}

// Generally should never be used as it is non-expiring access token to every part of your auth0 tenant
output "auth0_global_client_id" {
    value = auth0_global_client.global.client_id
}

output "auth0_global_client_secret" {
    value = auth0_global_client.global.client_secret
    sensitive = true
}
