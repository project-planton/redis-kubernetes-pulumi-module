package service

import (
	"fmt"
	"github.com/plantoncloud-inc/go-commons/kubernetes/network/dns"
)

func GetKubeServiceNameFqdn(redisKubernetesName, namespace string) string {
	return fmt.Sprintf("%s-master.%s.%s", GetKubeServiceName(redisKubernetesName), namespace, dns.DefaultDomain)
}

func GetKubeServiceName(redisKubernetesName string) string {
	return fmt.Sprintf(redisKubernetesName)
}
