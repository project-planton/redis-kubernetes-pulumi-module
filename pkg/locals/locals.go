package locals

import (
	"fmt"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes/model"
	"github.com/plantoncloud/redis-kubernetes-pulumi-module/pkg/outputs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	IngressExternalHostname string
	IngressInternalHostname string
	KubePortForwardCommand  string
	KubeServiceFqdn         string
	KubeServiceName         string
	Namespace               string
	RedisKubernetes         *model.RedisKubernetes
	RedisPodSelectorLabels  map[string]string
)

// Initializer will be invoked by the stack-job-runner sdk before the pulumi operations are executed.
func Initializer(ctx *pulumi.Context, stackInput *model.RedisKubernetesStackInput) {
	//assign value for the local variable to make it available across the module.
	RedisKubernetes = stackInput.ApiResource

	redisKubernetes := stackInput.ApiResource

	//decide on the namespace
	Namespace = redisKubernetes.Metadata.Id

	RedisPodSelectorLabels = map[string]string{
		"app.kubernetes.io/component": "master",
		"app.kubernetes.io/instance":  redisKubernetes.Metadata.Id,
		"app.kubernetes.io/name":      "redis",
	}

	KubeServiceName = fmt.Sprintf("%s-master", redisKubernetes.Metadata.Name)

	//export kubernetes service name
	ctx.Export(outputs.Service, pulumi.String(KubeServiceName))

	KubeServiceFqdn = fmt.Sprintf("%s.%s.svc.cluster.local.", KubeServiceName, Namespace)

	//export kubernetes endpoint
	ctx.Export(outputs.KubeEndpoint, pulumi.String(KubeServiceFqdn))

	KubePortForwardCommand = fmt.Sprintf("kubectl port-forward -n %s service/%s 8080:8080",
		Namespace, KubeServiceName)

	//export kube-port-forward command
	ctx.Export(outputs.PortForwardCommand, pulumi.String(KubePortForwardCommand))

	if redisKubernetes.Spec.Ingress == nil ||
		!redisKubernetes.Spec.Ingress.IsEnabled ||
		redisKubernetes.Spec.Ingress.EndpointDomainName == "" {
		return
	}

	IngressExternalHostname = fmt.Sprintf("%s.%s", redisKubernetes.Metadata.Id,
		redisKubernetes.Spec.Ingress.EndpointDomainName)

	IngressInternalHostname = fmt.Sprintf("%s-internal.%s", redisKubernetes.Metadata.Id,
		redisKubernetes.Spec.Ingress.EndpointDomainName)

	//export ingress hostnames
	ctx.Export(outputs.IngressExternalHostname, pulumi.String(IngressExternalHostname))
	ctx.Export(outputs.IngressInternalHostname, pulumi.String(IngressInternalHostname))
}
