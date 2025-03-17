// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package provider_mailgun

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func MailgunProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Required:            true,
				Description:         "The API key for Mailgun",
				MarkdownDescription: "The API key for Mailgun",
			},
			"endpoint": schema.StringAttribute{
				Optional:            true,
				Description:         "The Mailgun API endpoint (default https://api.mailgun.net)",
				MarkdownDescription: "The Mailgun API endpoint (default https://api.mailgun.net)",
			},
			"region": schema.StringAttribute{
				Optional:            true,
				Description:         "The Mailgun region (US or EU)",
				MarkdownDescription: "The Mailgun region (US or EU)",
			},
		},
	}
}

type MailgunModel struct {
	ApiKey   types.String `tfsdk:"api_key"`
	Endpoint types.String `tfsdk:"endpoint"`
	Region   types.String `tfsdk:"region"`
}
