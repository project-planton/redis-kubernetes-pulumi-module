package ingress

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	redisistio "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/network/ingress/istio"
	redisloadbalancer "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/network/ingress/loadbalancer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) (newCtx *pulumi.Context, err error) {
	i := extractInput(ctx)
	switch i.ingressType {
	case kubernetesworkloadingresstype.KubernetesWorkloadIngressType_load_balancer:
		ctx, err = redisloadbalancer.Resources(ctx)
		if err != nil {
			return ctx, errors.Wrap(err, "failed to add load balancer resources")
		}
	case kubernetesworkloadingresstype.KubernetesWorkloadIngressType_ingress_controller:
		if err = redisistio.Resources(ctx); err != nil {
			return ctx, errors.Wrap(err, "failed to add redisistio resources")
		}
	}
	return ctx, nil
}
