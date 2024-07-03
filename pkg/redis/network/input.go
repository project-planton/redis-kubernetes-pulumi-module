package network

import (
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	isIngressEnabled   bool
	endpointDomainName string
}

func extractInput(ctx *pulumi.Context) *input {
	var contextState = ctx.Value(rediscontextstate.Key).(rediscontextstate.ContextState)

	return &input{
		isIngressEnabled:   contextState.Spec.IsIngressEnabled,
		endpointDomainName: contextState.Spec.EndpointDomainName,
	}
}
