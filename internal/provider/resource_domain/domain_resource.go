// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &DomainResource{}
)

// DomainResource is the resource implementation.
type DomainResource struct {
}

// Metadata returns the resource type name.
func (r *DomainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

// Schema defines the schema for the resource.
func (r *DomainResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = DomainResourceSchema(ctx)
}

// Create creates a new resource.
func (r *DomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Implementation will be added here
	// This would typically include making API calls to Mailgun to create a domain
}

// Read refreshes the Terraform state with the latest data.
func (r *DomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Implementation will be added here
	// This would typically include making API calls to Mailgun to retrieve domain data
}

// Update updates an existing resource.
func (r *DomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Implementation will be added here
	// This would typically include making API calls to Mailgun to update a domain
}

// Delete deletes an existing resource.
func (r *DomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Implementation will be added here
	// This would typically include making API calls to Mailgun to delete a domain
}
