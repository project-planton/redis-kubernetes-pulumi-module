package virtualservice

import (
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	resourceId       string
	resourceName     string
	namespaceName    string
	workspaceDir     string
	namespace        *kubernetescorev1.Namespace
	externalHostname string
	internalHostname string
	kubeEndpoint     string
	envDomainName    string
}

func extractInput(ctx *pulumi.Context) *input {
	var contextState = ctx.Value(rediscontextstate.Key).(rediscontextstate.ContextState)

	return &input{
		resourceId:       contextState.Spec.ResourceId,
		resourceName:     contextState.Spec.ResourceName,
		workspaceDir:     contextState.Spec.WorkspaceDir,
		namespaceName:    contextState.Spec.NamespaceName,
		namespace:        contextState.Status.AddedResources.Namespace,
		externalHostname: contextState.Spec.ExternalHostname,
		internalHostname: contextState.Spec.InternalHostname,
		kubeEndpoint:     contextState.Spec.KubeLocalEndpoint,
		envDomainName:    contextState.Spec.EnvDomainName,
	}
}
