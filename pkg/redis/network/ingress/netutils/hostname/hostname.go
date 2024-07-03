package hostname

import "fmt"

func GetInternalHostname(redisKubernetesId, environmentName, endpointDomainName string) string {
	return fmt.Sprintf("%s.%s-internal.%s", redisKubernetesId, environmentName, endpointDomainName)
}

func GetExternalHostname(redisKubernetesId, environmentName, endpointDomainName string) string {
	return fmt.Sprintf("%s.%s.%s", redisKubernetesId, environmentName, endpointDomainName)
}
