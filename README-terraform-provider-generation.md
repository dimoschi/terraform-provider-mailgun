# Generating the Mailgun Terraform Provider

This document provides instructions on how to generate a Terraform provider for Mailgun using the OpenAPI Provider Spec Generator.

## Prerequisites

- Go 1.18 or later
- Terraform 1.0.0 or later
- OpenAPI Provider Spec Generator (`terraform-plugin-codegen-openapi`)

## Setup

1. Install the required code generation tools:

```bash
# Install the OpenAPI Provider Spec Generator
go install github.com/hashicorp/terraform-plugin-codegen-openapi/cmd/tfplugingen-openapi@latest

# Install the Terraform Plugin Framework code generator
go install github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest
```

2. Ensure you have the OpenAPI specification file (`openapi-final.yaml`) and the generator configuration file (`mailgun-openapi-generator-config.yaml`) in your project directory.

## Generate the Provider

The process involves two steps:

### 1. Generate the Provider Code Specification

First, generate a JSON specification file that describes the provider code to be generated:

```bash
tfplugingen-openapi generate \
  --config mailgun-openapi-generator-config.yaml \
  --output provider_code_spec.json openapi-final.yaml
```

This command processes your OpenAPI spec and configuration to create a provider code specification JSON file.

### 2. Generate the Provider Code

Next, use the Terraform Plugin Framework code generator to create the actual provider code from the specification:

```bash
tfplugingen-framework \
  generate all \
  --input provider_code_spec.json \
  --output internal/provider
```

This will generate the Go code for your Terraform provider in the current directory, creating the necessary files and directory structure.

## Configuration Details

The `mailgun-openapi-generator-config.yaml` file contains the following configurations:

### Provider Configuration

- **Name**: mailgun
- **Description**: Terraform provider for Mailgun API
- **Authentication**: API key-based authentication

### Resources

1. **mailgun_domain**: Manages a Mailgun domain
   - Create, read, update, and delete operations
   - Attributes for domain configuration including:
     - Name, spam action, wildcard settings
     - DKIM configuration
     - IP assignments
     - DNS records

### Data Sources

1. **mailgun_domains**: Get a list of all domains
   - Basic information about each domain including name, ID, state, and configuration

## Usage Example

After generating the provider, you can use it in your Terraform configurations like this:

```hcl
terraform {
  required_providers {
    mailgun = {
      source = "hashicorp/mailgun"
      version = "~> 1.0"
    }
  }
}

provider "mailgun" {
  api_key = var.mailgun_api_key
  region  = "US"  # or "EU"
}

resource "mailgun_domain" "example" {
  name                = "example.com"
  wildcard            = true
  spam_action         = "disabled"
  force_dkim_authority = true
  dkim_key_size       = 2048
  web_scheme          = "https"
}

# Using the data source to list all domains
data "mailgun_domains" "all" {}

output "all_domains" {
  value = data.mailgun_domains.all.domains
}
```

## Customizing the Generated Code

After generating the provider code, you may need to make some manual adjustments:

1. Review the generated code for any inconsistencies or missing functionality
2. Add custom validation logic if needed
3. Implement any complex operations that couldn't be automatically generated
4. Add comprehensive tests for the provider

## Building and Installing the Provider

After generating and customizing the code, build and install the provider:

```bash
go build -o terraform-provider-mailgun
```

For local development, you can install the provider in your local Terraform plugin directory:

```bash
mkdir -p ~/.terraform.d/plugins/hashicorp/mailgun/1.0.0/$(go env GOOS)_$(go env GOARCH)
cp terraform-provider-mailgun ~/.terraform.d/plugins/hashicorp/mailgun/1.0.0/$(go env GOOS)_$(go env GOARCH)/
```

## Notes

- The generated provider is based on the OpenAPI specification and may require additional customization for optimal functionality.
- Some complex operations or edge cases might not be fully covered by the automatic generation and may need manual implementation.
- Always test the provider thoroughly before using it in production environments.
