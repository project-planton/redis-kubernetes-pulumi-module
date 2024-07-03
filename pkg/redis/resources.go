package redis

import (
	"github.com/pkg/errors"
	code2cloudv1deployrdcstackk8smodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/rediskubernetes/stack/model"
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	redishelmchart "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/helmchart"
	redisnamespace "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/namespace"
	rediskubernetesnetwork "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/network"
	"github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/outputs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	WorkspaceDir     string
	Input            *code2cloudv1deployrdcstackk8smodel.RedisKubernetesStackInput
	KubernetesLabels map[string]string
}

func (resourceStack *ResourceStack) Resources(ctx *pulumi.Context) error {
	//load context config
	contextState, err := loadConfig(ctx, resourceStack)
	if err != nil {
		return errors.Wrap(err, "failed to initiate context config")
	}
	ctx = ctx.WithValue(rediscontextstate.Key, *contextState)

	// Create the namespace resource
	ctx, err = redisnamespace.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create namespace resource")
	}

	// Deploying a Redis Helm chart from the Helm repository.
	err = redishelmchart.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to add redis kubernetes helm chart resources")
	}

	ctx, err = rediskubernetesnetwork.Resources(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to add redis kubernetes ingress resources")
	}

	err = outputs.Export(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to export redis kubernetes outputs")
	}
	return nil
}
