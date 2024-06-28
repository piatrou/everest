// everest
// Copyright (C) 2023 Percona LLC
//
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

package kubernetes

//go:generate ../../bin/ifacemaker -f cert_manager.go accounts.go -f  -f deployment.go -f install_plan.go -f kubernetes.go -f namespace.go -f operator.go -f tls.go -f jwt.go -f oidc.go -f secret.go -s Kubernetes -i KubernetesConnector -p kubernetes -o kubernetes_interface.go
//go:generate ../../bin/mockery --name=KubernetesConnector --case=snake --inpackage
