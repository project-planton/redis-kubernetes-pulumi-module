package loadbalancer

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/cloudaccount/enums/kubernetesprovider"
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	kubeProvider kubernetesprovider.KubernetesProvider
}

func extractInput(ctx *pulumi.Context) *input {
	var contextState = ctx.Value(rediscontextstate.Key).(rediscontextstate.ContextState)

	return &input{
		kubeProvider: contextState.Spec.EnvironmentInfo.KubernetesProvider,
	}
}
