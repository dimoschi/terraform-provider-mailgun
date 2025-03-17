# Mailgun Provider Verification Example

This example demonstrates how to use the Mailgun Terraform provider to retrieve domain information.

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) 0.13+
- A Mailgun account with an API key

## Usage

1. Copy the example variables file and update it with your Mailgun API key:

```bash
cp terraform.tfvars.example terraform.tfvars
```

Then edit `terraform.tfvars` with your actual API key:

```hcl
api_key = "your-mailgun-api-key"
```

2. Initialize Terraform:

```bash
terraform init
```

3. Apply the configuration:

```bash
terraform apply
```

4. View the domains:

```bash
terraform output domains
```

## Configuration Options

The following variables can be configured in the `terraform.tfvars` file:

- `api_key` (required): Your Mailgun API key
- `region` (optional): The Mailgun region to use (US or EU)
- `endpoint` (optional): A custom API endpoint URL
