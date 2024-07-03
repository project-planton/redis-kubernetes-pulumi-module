package istio

import (
	"github.com/pkg/errors"
	rediskubernetesnetworkgateway "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/network/ingress/istio/gateway"
	rediskubernetesnetworkvirtualservice "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/network/ingress/istio/virtualservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context) error {
	if err := rediskubernetesnetworkgateway.Resources(ctx); err != nil {
		return errors.Wrap(err, "failed to add gateway resources")
	}
	if err := rediskubernetesnetworkvirtualservice.Resources(ctx); err != nil {
		return errors.Wrap(err, "failed to add virtual resources")
	}
	return nil
}
