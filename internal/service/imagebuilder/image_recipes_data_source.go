// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	awstypes "github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/namevaluesfilters"
	namevaluesfiltersv2 "github.com/hashicorp/terraform-provider-aws/internal/namevaluesfilters/v2"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_imagebuilder_image_recipes", name="Image Recipes")
func DataSourceImageRecipes() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceImageRecipesRead,
		Schema: map[string]*schema.Schema{
			names.AttrARNs: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrFilter: namevaluesfilters.Schema(),
			names.AttrNames: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"owner": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: enum.Validate[awstypes.Ownership](),
			},
		},
	}
}

func dataSourceImageRecipesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).ImageBuilderClient(ctx)

	input := &imagebuilder.ListImageRecipesInput{}

	if v, ok := d.GetOk("owner"); ok {
		input.Owner = awstypes.Ownership(v.(string))
	}

	if v, ok := d.GetOk("filter"); ok {
		input.Filters = namevaluesfiltersv2.New(v.(*schema.Set)).ImageBuilderFilters()
	}

	out, err := conn.ListImageRecipes(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading Image Builder Image Recipes: %s", err)
	}

	var arns, nms []string

	for _, r := range out.ImageRecipeSummaryList {
		arns = append(arns, aws.ToString(r.Arn))
		nms = append(nms, aws.ToString(r.Name))
	}

	d.SetId(meta.(*conns.AWSClient).Region)
	d.Set(names.AttrARNs, arns)
	d.Set(names.AttrNames, nms)

	return diags
}
