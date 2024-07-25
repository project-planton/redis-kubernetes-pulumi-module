package pkg

var vars = struct {
	RedisPort        int
	HelmChartVersion string
	HelmChartName    string
	HelmChartRepoUrl string
}{
	RedisPort:        6379,
	HelmChartVersion: "17.10.1",
	HelmChartName:    "redis",
	HelmChartRepoUrl: "https://charts.bitnami.com/bitnami",
}
