package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/rediskubernetes"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/kubernetes"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

const (
	Namespace               = "namespace"
	Service                 = "service"
	KubePortForwardCommand  = "port-forward-command"
	KubeEndpoint            = "kube-endpoint"
	IngressExternalHostname = "ingress-external-hostname"
	IngressInternalHostname = "ingress-internal-hostname"
	RedisUsername           = "redis-username"
	RedisPasswordSecretName = "redis-password-secret-name"
	RedisPasswordSecretKey  = "redis-password-secret-key"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *rediskubernetes.RedisKubernetesStackInput) *rediskubernetes.RedisKubernetesStackOutputs {
	return &rediskubernetes.RedisKubernetesStackOutputs{
		Namespace:          autoapistackoutput.GetVal(pulumiOutputs, Namespace),
		Service:            autoapistackoutput.GetVal(pulumiOutputs, Service),
		PortForwardCommand: autoapistackoutput.GetVal(pulumiOutputs, KubePortForwardCommand),
		KubeEndpoint:       autoapistackoutput.GetVal(pulumiOutputs, KubeEndpoint),
		ExternalHostname:   autoapistackoutput.GetVal(pulumiOutputs, IngressExternalHostname),
		InternalHostname:   autoapistackoutput.GetVal(pulumiOutputs, IngressInternalHostname),
		Username:           autoapistackoutput.GetVal(pulumiOutputs, RedisUsername),
		PasswordSecret: &kubernetes.KubernernetesSecretKey{
			Name: autoapistackoutput.GetVal(pulumiOutputs, RedisPasswordSecretName),
			Key:  autoapistackoutput.GetVal(pulumiOutputs, RedisPasswordSecretKey),
		},
	}
}
