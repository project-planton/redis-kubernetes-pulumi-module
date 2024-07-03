package helmchart

import (
	rediskubernetesmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/rediskubernetes/model"
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	resourceId    string
	resourceName  string
	namespaceName string
	namespace     *kubernetescorev1.Namespace
	containerSpec *rediskubernetesmodel.RedisKubernetesSpecContainerSpec
	labels        map[string]string
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(rediscontextstate.Key).(rediscontextstate.ContextState)

	return &input{
		resourceId:    ctxState.Spec.ResourceId,
		resourceName:  ctxState.Spec.ResourceName,
		namespaceName: ctxState.Spec.NamespaceName,
		namespace:     ctxState.Status.AddedResources.Namespace,
		labels:        ctxState.Spec.Labels,
		containerSpec: ctxState.Spec.ContainerSpec,
	}
}
