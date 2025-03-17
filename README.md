# Terraform Provider for Mailgun

This Terraform provider allows you to manage [Mailgun](https://www.mailgun.com/) resources through Terraform. It provides the ability to create, read, update, and delete Mailgun domains, as well as retrieve information about existing domains.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.22 (for development)
- A Mailgun account with an API key

## Installation

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    mailgun = {
      source = "registry.terraform.io/dimoschi/mailgun"
    }
  }
}
```

## Provider Configuration

The provider needs to be configured with your Mailgun API key. You can also optionally specify the region (US or EU) and a custom API endpoint if needed.

```hcl
provider "mailgun" {
  api_key  = var.mailgun_api_key  # Required
  region   = "US"                 # Optional, can be "US" or "EU"
  endpoint = null                 # Optional, defaults to "https://api.mailgun.net"
}
```

### Configuration Parameters

| Parameter | Description | Required | Default |
|-----------|-------------|----------|---------|
| `api_key` | Your Mailgun API key | Yes | - |
| `region` | The Mailgun region (US or EU) | No | US |
| `endpoint` | The Mailgun API endpoint | No | https://api.mailgun.net |

## Resources and Data Sources

### Resources

#### `mailgun_domain`

Manages a Mailgun domain.

```hcl
resource "mailgun_domain" "example" {
  name                = "example.com"
  wildcard            = true
  spam_action         = "disabled"
  force_dkim_authority = true
  dkim_key_size       = "2048"
  web_scheme          = "https"
}
```

See the [domain resource documentation](docs/resources/domain.md) for all available parameters.

### Data Sources

#### `mailgun_domains`

Retrieves a list of all domains in your Mailgun account.

```hcl
data "mailgun_domains" "all" {}

output "all_domains" {
  value = data.mailgun_domains.all.items
}
```

See the [domains data source documentation](docs/data-sources/domains.md) for all available parameters.

## Usage Example

```hcl
terraform {
  required_providers {
    mailgun = {
      source = "registry.terraform.io/dimoschi/mailgun"
    }
  }
}

variable "mailgun_api_key" {
  description = "Mailgun API key"
  type        = string
  sensitive   = true
}

provider "mailgun" {
  api_key = var.mailgun_api_key
  region  = "US"
}

# Create a new domain
resource "mailgun_domain" "example" {
  name        = "example.com"
  wildcard    = true
  spam_action = "tag"
  web_scheme  = "https"
}

# List all domains
data "mailgun_domains" "all" {}

output "all_domains" {
  value = data.mailgun_domains.all.items
}
```

## Development

### Building the Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

### Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

### Testing

To run the provider tests:

```shell
make testacc
```

*Note:* Acceptance tests create real resources, and often cost money to run.

### Generating the Provider

This provider was generated using the OpenAPI Provider Spec Generator. For more information on how to regenerate the provider, see [README-terraform-provider-generation.md](README-terraform-provider-generation.md).

## License

This provider is licensed under the [Mozilla Public License v2.0](LICENSE).
