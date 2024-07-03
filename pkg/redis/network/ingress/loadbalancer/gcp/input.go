package gcp

import (
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	resourceId         string
	resourceName       string
	namespace          *kubernetescorev1.Namespace
	externalHostname   string
	internalHostname   string
	endpointDomainName string
	serviceName        string
}

func extractInput(ctx *pulumi.Context) *input {
	var contextState = ctx.Value(rediscontextstate.Key).(rediscontextstate.ContextState)

	return &input{
		resourceId:         contextState.Spec.ResourceId,
		resourceName:       contextState.Spec.ResourceName,
		namespace:          contextState.Status.AddedResources.Namespace,
		externalHostname:   contextState.Spec.ExternalHostname,
		internalHostname:   contextState.Spec.InternalHostname,
		endpointDomainName: contextState.Spec.EndpointDomainName,
		serviceName:        contextState.Spec.KubeServiceName,
	}
}
