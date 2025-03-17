// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/mailgun/mailgun-go/v4"

	"github.com/dimoschi/terraform-provider-mailgun/internal/provider/datasource_domains"
	"github.com/dimoschi/terraform-provider-mailgun/internal/provider/provider_mailgun"
	"github.com/dimoschi/terraform-provider-mailgun/internal/provider/resource_domain"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &mailgunProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &mailgunProvider{
			version: version,
		}
	}
}

// mailgunProvider is the provider implementation.
type mailgunProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *mailgunProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "mailgun"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *mailgunProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = provider_mailgun.MailgunProviderSchema(ctx)
}

// Configure prepares a Mailgun API client for data sources and resources.
func (p *mailgunProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config provider_mailgun.MailgunModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ApiKey.IsNull() {
		resp.Diagnostics.AddError(
			"Missing API Key Configuration",
			"While configuring the provider, the API key was not found in "+
				"the Mailgun provider configuration. "+
				"Please set the api_key value in the provider configuration.",
		)
		return
	}

	// Default values
	endpoint := "https://api.mailgun.net"
	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	// Create Mailgun client
	mg := mailgun.NewMailgun("", config.ApiKey.ValueString())

	// Set region if provided
	if !config.Region.IsNull() {
		region := config.Region.ValueString()
		if region == "EU" {
			mg.SetAPIBase("https://api.eu.mailgun.net/v3")
		}
	}

	// If endpoint is provided, override the API base
	if !config.Endpoint.IsNull() {
		mg.SetAPIBase(endpoint)
	}

	// Make the Mailgun client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = mg
	resp.ResourceData = mg
}

// DataSources defines the data sources implemented in the provider.
func (p *mailgunProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Add data sources here
		func() datasource.DataSource {
			return &datasource_domains.DomainsDataSource{}
		},
	}
}

// Resources defines the resources implemented in the provider.
func (p *mailgunProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// Add resources here
		func() resource.Resource {
			return &resource_domain.DomainResource{}
		},
	}
}
