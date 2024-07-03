package redis

import (
	"github.com/pkg/errors"
	environmentblueprinthostnames "github.com/plantoncloud/environment-pulumi-blueprint/pkg/gcpgke/endpointdomains/hostnames"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	"github.com/plantoncloud/pulumi-blueprint-golang-commons/pkg/kubernetes/pulumikubernetesprovider"
	rediscontextstate "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/contextstate"
	redisnetutilshostname "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/network/ingress/netutils/hostname"
	redisnetutilsservice "github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/network/ingress/netutils/service"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func loadConfig(ctx *pulumi.Context, resourceStack *ResourceStack) (*rediscontextstate.ContextState, error) {

	kubernetesProvider, err := pulumikubernetesprovider.GetWithStackCredentials(ctx, resourceStack.Input.CredentialsInput)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup kubernetes provider")
	}

	var resourceId = resourceStack.Input.ResourceInput.Metadata.Id
	var resourceName = resourceStack.Input.ResourceInput.Metadata.Name
	var environmentInfo = resourceStack.Input.ResourceInput.Spec.EnvironmentInfo
	var isIngressEnabled = false

	if resourceStack.Input.ResourceInput.Spec.Ingress != nil {
		isIngressEnabled = resourceStack.Input.ResourceInput.Spec.Ingress.IsEnabled
	}

	var endpointDomainName = ""
	var envDomainName = ""
	var ingressType = kubernetesworkloadingresstype.KubernetesWorkloadIngressType_unspecified
	var internalHostname = ""
	var externalHostname = ""
	var certSecretName = ""

	if isIngressEnabled {
		endpointDomainName = resourceStack.Input.ResourceInput.Spec.Ingress.EndpointDomainName
		envDomainName = environmentblueprinthostnames.GetExternalEnvHostname(environmentInfo.EnvironmentName, endpointDomainName)
		ingressType = resourceStack.Input.ResourceInput.Spec.Ingress.IngressType

		internalHostname = redisnetutilshostname.GetInternalHostname(resourceId, environmentInfo.EnvironmentName, endpointDomainName)
		externalHostname = redisnetutilshostname.GetExternalHostname(resourceId, environmentInfo.EnvironmentName, endpointDomainName)
	}

	return &rediscontextstate.ContextState{
		Spec: &rediscontextstate.Spec{
			KubeProvider:       kubernetesProvider,
			ResourceId:         resourceId,
			ResourceName:       resourceName,
			Labels:             resourceStack.KubernetesLabels,
			WorkspaceDir:       resourceStack.WorkspaceDir,
			NamespaceName:      resourceId,
			EnvironmentInfo:    resourceStack.Input.ResourceInput.Spec.EnvironmentInfo,
			ContainerSpec:      resourceStack.Input.ResourceInput.Spec.Container,
			IsIngressEnabled:   isIngressEnabled,
			IngressType:        ingressType,
			EndpointDomainName: endpointDomainName,
			EnvDomainName:      envDomainName,
			InternalHostname:   internalHostname,
			ExternalHostname:   externalHostname,
			KubeServiceName:    redisnetutilsservice.GetKubeServiceName(resourceName),
			KubeLocalEndpoint:  redisnetutilsservice.GetKubeServiceNameFqdn(resourceName, resourceId),
			CertSecretName:     certSecretName,
		},
		Status: &rediscontextstate.Status{},
	}, nil
}
