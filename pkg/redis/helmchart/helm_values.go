package helmchart

import (
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/containerresources"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/helm/convertmaps"
	"github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/secret"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func getHelmChartValuesMap(i *input) pulumi.Map {

	// HelmVal https://github.com/bitnami/charts/blob/main/bitnami/redis/values.yaml
	return pulumi.Map{
		"fullnameOverride": pulumi.String(i.resourceName),
		"architecture":     pulumi.String("standalone"),
		"master": pulumi.Map{
			"podLabels": convertmaps.ConvertGoMapToPulumiMap(i.labels),
			"resources": containerresources.ConvertToPulumiMap(i.containerSpec.Resources),
			"persistence": pulumi.Map{
				"enabled": pulumi.Bool(i.containerSpec.IsPersistenceEnabled),
				"size":    pulumi.String(i.containerSpec.DiskSize),
			},
		},
		"replica": pulumi.Map{
			"podLabels":    convertmaps.ConvertGoMapToPulumiMap(i.labels),
			"replicaCount": pulumi.Int(i.containerSpec.Replicas),
			"resources":    containerresources.ConvertToPulumiMap(i.containerSpec.Resources),
			"persistence": pulumi.Map{
				"enabled": pulumi.Bool(i.containerSpec.IsPersistenceEnabled),
				"size":    pulumi.String(i.containerSpec.DiskSize),
			},
		},
		"auth": pulumi.Map{
			"existingSecret":            pulumi.String(i.resourceName),
			"existingSecretPasswordKey": pulumi.String(secret.RedisPasswordKey),
		},
	}
}
