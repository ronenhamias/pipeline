// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitoring

import (
	"fmt"
)

const (
	integratedServiceName            = "monitoring"
	prometheusOperatorReleaseName    = "monitor"
	prometheusPushgatewayReleaseName = "pushgateway"
	grafanaSecretTag                 = "app:grafana"
	prometheusSecretTag              = "app:prometheus"
	alertmanagerSecretTag            = "app:alertmanager"
	integratedServiceSecretTag       = "feature:monitoring"
	generatedSecretUsername          = "admin"
	alertManagerProviderConfigName   = "default-receiver"
	alertManagerNullReceiverName     = "null"

	ingressTypeGrafana      = "Grafana"
	ingressTypePrometheus   = "Prometheus"
	ingressTypeAlertmanager = "Alertmanager"

	pagerDutyIntegrationEventApiV2 = "eventsApiV2"
	pagerDutyIntegrationPrometheus = "prometheus"

	alertmanagerProviderSlack     = "slack"
	alertmanagerProviderPagerDuty = "pagerDuty"

	repoOperator         = "quay.io/coreos/prometheus-operator"
	tagOperator          = "v0.34.0"
	repoPrometheus       = "quay.io/prometheus/prometheus"
	tagPrometheus        = "v2.13.1"
	repoAlertmanager     = "quay.io/prometheus/alertmanager"
	tagAlertmanager      = "v0.19.0"
	repoGrafana          = "grafana/grafana"
	tagGrafana           = "6.5.2"
	repoKubeStateMetrics = "quay.io/coreos/kube-state-metrics"
	tagKubeStateMetrics  = "v1.9.3"
	repoNodeExporter     = "quay.io/prometheus/node-exporter"
	tagNodeExporter      = "v0.18.1"
	repoPushgateway      = "prom/pushgateway"
	tagPushgateway       = "v1.0.1"
)

func getClusterNameSecretTag(clusterName string) string {
	return fmt.Sprintf("cluster:%s", clusterName)
}

func getClusterUIDSecretTag(clusterUID string) string {
	return fmt.Sprintf("clusterUID:%s", clusterUID)
}

func getReleaseSecretTag() string {
	return fmt.Sprintf("release:%s", prometheusOperatorReleaseName)
}

func getPrometheusSecretName(clusterID uint) string {
	return fmt.Sprintf("cluster-%d-prometheus", clusterID)
}

func getAlertmanagerSecretName(clusterID uint) string {
	return fmt.Sprintf("cluster-%d-alertmanager", clusterID)
}

func getPushgatewaySecretName(clusterID uint) string {
	return fmt.Sprintf("cluster-%d-pushgateway", clusterID)
}

func getGrafanaSecretName(clusterID uint) string {
	return fmt.Sprintf("cluster-%d-grafana", clusterID)
}

func generateAnnotations(secretName string) map[string]interface{} {
	return map[string]interface{}{
		"traefik.ingress.kubernetes.io/auth-type":   "basic",
		"traefik.ingress.kubernetes.io/auth-secret": secretName,
	}
}
