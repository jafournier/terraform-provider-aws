// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	imagebuilder_sdkv2 "github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceLifecyclePolicy,
			Name:    "Lifecycle Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceComponent,
			TypeName: "aws_imagebuilder_component",
		},
		{
			Factory:  DataSourceComponents,
			TypeName: "aws_imagebuilder_components",
			Name:     "Components",
		},
		{
			Factory:  DataSourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
		},
		{
			Factory:  DataSourceContainerRecipes,
			TypeName: "aws_imagebuilder_container_recipes",
			Name:     "Container Recipes",
		},
		{
			Factory:  DataSourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
		},
		{
			Factory:  DataSourceDistributionConfigurations,
			TypeName: "aws_imagebuilder_distribution_configurations",
		},
		{
			Factory:  DataSourceImage,
			TypeName: "aws_imagebuilder_image",
		},
		{
			Factory:  DataSourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
		},
		{
			Factory:  DataSourceImagePipelines,
			TypeName: "aws_imagebuilder_image_pipelines",
		},
		{
			Factory:  DataSourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
		},
		{
			Factory:  DataSourceImageRecipes,
			TypeName: "aws_imagebuilder_image_recipes",
			Name:     "Image Recipes",
		},
		{
			Factory:  DataSourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
		},
		{
			Factory:  DataSourceInfrastructureConfigurations,
			TypeName: "aws_imagebuilder_infrastructure_configurations",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceComponent,
			TypeName: "aws_imagebuilder_component",
			Name:     "Component",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
			Name:     "Container Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
			Name:     "Distribution Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceImage,
			TypeName: "aws_imagebuilder_image",
			Name:     "Image",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
			Name:     "Image Pipeline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
			Name:     "Image Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
			Name:     "Infrastructure Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceWorkflow,
			TypeName: "aws_imagebuilder_workflow",
			Name:     "Workflow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ImageBuilder
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*imagebuilder_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return imagebuilder_sdkv2.NewFromConfig(cfg,
		imagebuilder_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
