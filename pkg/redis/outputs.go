package redis

import (
	"context"
	"github.com/plantoncloud/redis-kubernetes-pulumi-blueprint/pkg/redis/outputs"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	"github.com/pkg/errors"
	code2cloudv1deployrdcmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/rediskubernetes/model"
	code2cloudv1deployrdcstackk8smodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/rediskubernetes/stack/model"
	"github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/org"
	"github.com/plantoncloud/pulumi-stack-runner-go-sdk/pkg/stack/output/backend"
)

func Outputs(ctx context.Context, input *code2cloudv1deployrdcstackk8smodel.RedisKubernetesStackInput) (*code2cloudv1deployrdcmodel.RedisKubernetesStatusStackOutputs, error) {
	pulumiOrgName, err := org.GetOrgName()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pulumi org name")
	}
	stackOutput, err := backend.StackOutput(pulumiOrgName, input.StackJob)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stack output")
	}
	return OutputMapTransformer(stackOutput, input), nil
}

func OutputMapTransformer(stackOutput map[string]interface{}, input *code2cloudv1deployrdcstackk8smodel.RedisKubernetesStackInput) *code2cloudv1deployrdcmodel.RedisKubernetesStatusStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &code2cloudv1deployrdcmodel.RedisKubernetesStatusStackOutputs{}
	}
	return &code2cloudv1deployrdcmodel.RedisKubernetesStatusStackOutputs{
		Namespace:          backend.GetVal(stackOutput, outputs.GetNamespaceNameOutputName()),
		Service:            backend.GetVal(stackOutput, outputs.GetKubeServiceNameOutputName()),
		PortForwardCommand: backend.GetVal(stackOutput, outputs.GetKubePortForwardCommandOutputName()),
		KubeEndpoint:       backend.GetVal(stackOutput, outputs.GetKubeEndpointOutputName()),
		ExternalHostname:   backend.GetVal(stackOutput, outputs.GetExternalClusterHostnameOutputName()),
		InternalHostname:   backend.GetVal(stackOutput, outputs.GetInternalClusterHostnameOutputName()),
	}
}
