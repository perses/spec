// Copyright The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dashboard

myAnnotatedPanel: #Panel & {
	kind: "Panel"
	spec: {
		plugin: {
			kind: "TimeSeriesChart"
			spec: {}
		}
		annotations: [
			{
				display: {
					name:  "Deployments"
					color: "#EE6C6C"
				}
				plugin: {
					kind: "PrometheusPromQLAnnotation"
					spec: {
						expr:  "changes(kube_deployment_status_observed_generation[5m]) > 0"
						title: "Deployment of {{deployment}}"
						tags: ["deployment"]
					}
				}
			},
		]
	}
}
