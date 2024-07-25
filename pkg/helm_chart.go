package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/model"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func (s *ResourceStack) helmChart(ctx *pulumi.Context,
	createdNamespace *kubernetescorev1.Namespace) error {

	redisKubernetes := s.Input.ApiResource

	helmValues := getHelmValues(redisKubernetes)

	// Deploying a Locust Helm chart from the Helm repository.
	_, err := helmv3.NewChart(ctx,
		redisKubernetes.Metadata.Id,
		helmv3.ChartArgs{
			Chart:     pulumi.String("redis"),
			Version:   pulumi.String("5.1.5"), // Use the Helm chart version you want to install
			Namespace: createdNamespace.Metadata.Name().Elem(),
			Values:    helmValues,
			//if you need to add the repository, you can specify `repo url`:
			FetchArgs: helmv3.FetchArgs{
				Repo: pulumi.String("https://charts.redis.io"),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create helm chart")
	}

	return nil
}

func getHelmValues(redisKubernetes *model.RedisKubernetes) pulumi.Map {
	// https://github.com/redisci/helm-charts/blob/main/charts/redis/values.yaml
	var baseValues = pulumi.Map{}
	return baseValues
}
