package helmchart

import (
	"github.com/pkg/errors"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) error {
	if err := addHelmChart(ctx); err != nil {
		return errors.Wrap(err, "failed to add helm chart")
	}
	return nil
}

func addHelmChart(ctx *pulumi.Context) error {
	i := extractInput(ctx)

	// Deploying a Redis Helm chart from the Helm repository.
	_, err := helmv3.NewChart(ctx, i.resourceId, helmv3.ChartArgs{
		Chart:     pulumi.String("redis"),
		Version:   pulumi.String("17.10.1"), // Use the Helm chart version you want to install
		Namespace: pulumi.String(i.namespaceName),
		Values:    getHelmChartValuesMap(i),
		//if you need to add the repository, you can specify `repo url`:
		FetchArgs: helmv3.FetchArgs{
			Repo: pulumi.String("https://charts.bitnami.com/bitnami"), // The URL for the Helm chart repository
		},
	}, pulumi.Timeouts(&pulumi.CustomTimeouts{Create: "3m", Update: "3m", Delete: "3m"}), pulumi.Parent(i.namespace))
	if err != nil {
		return err
	}
	return nil
}
