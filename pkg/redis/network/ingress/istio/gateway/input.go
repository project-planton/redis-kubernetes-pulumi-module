package gateway

import (
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	workspaceDir     string
	kubeProvider     *pulumikubernetes.Provider
	resourceName     string
	resourceId       string
	labels           map[string]string
	externalHostname string
	envDomainName    string
	namespaceName    string
}

func extractInput(ctx *pulumi.Context) *input {
	var contextState = ctx.Value(rediscontextstate.Key).(rediscontextstate.ContextState)

	return &input{
		workspaceDir:     contextState.Spec.WorkspaceDir,
		kubeProvider:     contextState.Spec.KubeProvider,
		resourceName:     contextState.Spec.ResourceName,
		resourceId:       contextState.Spec.ResourceId,
		labels:           contextState.Spec.Labels,
		externalHostname: contextState.Spec.ExternalHostname,
		envDomainName:    contextState.Spec.EnvDomainName,
		namespaceName:    contextState.Spec.NamespaceName,
	}
}
