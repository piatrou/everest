// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	"context"

	cmVersioned "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	v1 "github.com/operator-framework/api/pkg/operators/v1"
	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	versioned "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	packagev1 "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	everestv1alpha1 "github.com/percona/everest-operator/api/v1alpha1"
)

// KubeClientConnector ...
type KubeClientConnector interface {
	// CreateBackupStorage creates an backupStorage.
	CreateBackupStorage(ctx context.Context, storage *everestv1alpha1.BackupStorage) error
	// UpdateBackupStorage updates an backupStorage.
	UpdateBackupStorage(ctx context.Context, storage *everestv1alpha1.BackupStorage) error
	// GetBackupStorage returns the backupStorage.
	GetBackupStorage(ctx context.Context, namespace, name string) (*everestv1alpha1.BackupStorage, error)
	// ListBackupStorages returns the backupStorage.
	ListBackupStorages(ctx context.Context, namespace string, options metav1.ListOptions) (*everestv1alpha1.BackupStorageList, error)
	// DeleteBackupStorage deletes the backupStorage.
	DeleteBackupStorage(ctx context.Context, namespace, name string) error
	// CertManager returns CertManager client set.
	//
	//nolint:ireturn
	CertManager() cmVersioned.Interface
	// GetConfigMap returns config map by name and namespace.
	GetConfigMap(ctx context.Context, namespace, name string) (*corev1.ConfigMap, error)
	// CreateConfigMap creates the provided ConfigMap.
	CreateConfigMap(ctx context.Context, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error)
	// UpdateConfigMap updates the provided ConfigMap.
	UpdateConfigMap(ctx context.Context, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error)
	// Config returns restConfig to the pkg/kubernetes.Kubernetes client.
	Config() *rest.Config
	// Clientset returns the k8s clientset.
	//
	//nolint:ireturn
	Clientset() kubernetes.Interface
	// ClusterName returns the name of the k8s cluster.
	ClusterName() string
	// Namespace returns the namespace of the k8s cluster.
	Namespace() string
	// GetSecretsForServiceAccount returns secret by given service account name.
	GetSecretsForServiceAccount(ctx context.Context, accountName string) (*corev1.Secret, error)
	// GenerateKubeConfigWithToken generates kubeconfig with a user and token provided as a secret.
	GenerateKubeConfigWithToken(user string, secret *corev1.Secret) ([]byte, error)
	// GetServerVersion returns server version.
	GetServerVersion() (*version.Info, error)
	// ApplyObject applies object.
	ApplyObject(obj runtime.Object) error
	// DeleteObject deletes object from the k8s cluster.
	DeleteObject(obj runtime.Object) error
	// ListObjects lists objects by provided group, version, kind.
	ListObjects(gvk schema.GroupVersionKind, into runtime.Object) error
	// GetObject retrieves an object by provided group, version, kind and name.
	GetObject(gvk schema.GroupVersionKind, name string, into runtime.Object) error
	// GetLogs returns logs for pod.
	GetLogs(ctx context.Context, pod, container string) (string, error)
	// GetEvents retrieves events from a pod by a name.
	GetEvents(ctx context.Context, name string) (string, error)
	// ApplyFile accepts manifest file contents, parses into []runtime.Object
	// and applies them against the cluster.
	ApplyFile(fileBytes []byte) error
	// ApplyManifestFile accepts manifest file contents, parses into []runtime.Object
	// and applies them against the cluster.
	ApplyManifestFile(fileBytes []byte, namespace string, ignoreObjects ...client.Object) error
	// DeleteManifestFile accepts manifest file contents, parses into []runtime.Object
	// and deletes them from the cluster.
	DeleteManifestFile(fileBytes []byte, namespace string) error
	// DoCSVWait waits until for a CSV to be applied.
	DoCSVWait(ctx context.Context, key types.NamespacedName) error
	// GetSubscriptionCSV retrieves a subscription CSV.
	GetSubscriptionCSV(ctx context.Context, subKey types.NamespacedName) (types.NamespacedName, error)
	// DoRolloutWait waits until a deployment has been rolled out susccessfully or there is an error.
	DoRolloutWait(ctx context.Context, key types.NamespacedName) error
	// GetOperatorGroup retrieves an operator group details by namespace and name.
	GetOperatorGroup(ctx context.Context, namespace, name string) (*v1.OperatorGroup, error)
	// CreateOperatorGroup creates an operator group to be used as part of a subscription.
	CreateOperatorGroup(ctx context.Context, namespace, name string, targetNamespaces []string) (*v1.OperatorGroup, error)
	// CreateSubscription creates an OLM subscription.
	CreateSubscription(ctx context.Context, namespace string, subscription *v1alpha1.Subscription) (*v1alpha1.Subscription, error)
	// UpdateSubscription updates an OLM subscription.
	UpdateSubscription(ctx context.Context, namespace string, subscription *v1alpha1.Subscription) (*v1alpha1.Subscription, error)
	// CreateSubscriptionForCatalog creates an OLM subscription.
	CreateSubscriptionForCatalog(ctx context.Context, namespace, name, catalogNamespace, catalog, packageName, channel, startingCSV string, approval v1alpha1.Approval) (*v1alpha1.Subscription, error)
	// GetSubscription retrieves an OLM subscription by namespace and name.
	GetSubscription(ctx context.Context, namespace, name string) (*v1alpha1.Subscription, error)
	// ListSubscriptions all the subscriptions in the namespace.
	ListSubscriptions(ctx context.Context, namespace string) (*v1alpha1.SubscriptionList, error)
	// DoPackageWait for the package to be available in OLM.
	DoPackageWait(ctx context.Context, namespace, name string) error
	// GetPackageManifest returns a package manifest by given name.
	GetPackageManifest(ctx context.Context, namespace, name string) (*packagev1.PackageManifest, error)
	// ListCRDs returns a list of CRDs.
	ListCRDs(ctx context.Context, labelSelector *metav1.LabelSelector) (*apiextv1.CustomResourceDefinitionList, error)
	// ListCRs returns a list of CRs.
	ListCRs(ctx context.Context, namespace string, gvr schema.GroupVersionResource, labelSelector *metav1.LabelSelector) (*unstructured.UnstructuredList, error)
	// GetClusterServiceVersion retrieve a CSV by namespaced name.
	GetClusterServiceVersion(ctx context.Context, key types.NamespacedName) (*v1alpha1.ClusterServiceVersion, error)
	// ListClusterServiceVersion list all CSVs for the given namespace.
	ListClusterServiceVersion(ctx context.Context, namespace string) (*v1alpha1.ClusterServiceVersionList, error)
	// UpdateClusterServiceVersion updates a CSV and returns the updated CSV.
	UpdateClusterServiceVersion(ctx context.Context, csv *v1alpha1.ClusterServiceVersion) (*v1alpha1.ClusterServiceVersion, error)
	// DeleteClusterServiceVersion deletes a CSV by namespaced name.
	DeleteClusterServiceVersion(ctx context.Context, key types.NamespacedName) error
	// DeleteFile accepts manifest file contents parses into []runtime.Object
	// and deletes them from the cluster.
	DeleteFile(fileBytes []byte) error
	// GetService returns k8s service by provided namespace and name.
	GetService(ctx context.Context, namespace, name string) (*corev1.Service, error)
	// GetClusterRoleBinding returns cluster role binding by given name.
	GetClusterRoleBinding(ctx context.Context, name string) (*rbacv1.ClusterRoleBinding, error)
	// GetCRD gets a CustomResourceDefinition by name.
	// Provided name should be of the format <resourcename>.<apiGroup>.
	// Example: installplans.operators.coreos.com .
	GetCRD(ctx context.Context, name string) (*apiextensionsv1.CustomResourceDefinition, error)
	// ListDatabaseClusters returns list of managed database clusters.
	ListDatabaseClusters(ctx context.Context, namespace string, options metav1.ListOptions) (*everestv1alpha1.DatabaseClusterList, error)
	// GetDatabaseCluster returns database clusters by provided name.
	GetDatabaseCluster(ctx context.Context, namespace, name string) (*everestv1alpha1.DatabaseCluster, error)
	// ListDatabaseClusterBackups returns list of managed database cluster backups.
	ListDatabaseClusterBackups(ctx context.Context, namespace string, options metav1.ListOptions) (*everestv1alpha1.DatabaseClusterBackupList, error)
	// GetDatabaseClusterBackup returns database cluster backups by provided name.
	GetDatabaseClusterBackup(ctx context.Context, namespace, name string) (*everestv1alpha1.DatabaseClusterBackup, error)
	// UpdateDatabaseClusterBackup updates the provided database cluster backup.
	UpdateDatabaseClusterBackup(ctx context.Context, backup *everestv1alpha1.DatabaseClusterBackup) (*everestv1alpha1.DatabaseClusterBackup, error)
	// ListDatabaseClusterRestores returns list of managed database clusters.
	ListDatabaseClusterRestores(ctx context.Context, namespace string, options metav1.ListOptions) (*everestv1alpha1.DatabaseClusterRestoreList, error)
	// GetDatabaseClusterRestore returns database clusters by provided name.
	GetDatabaseClusterRestore(ctx context.Context, namespace, name string) (*everestv1alpha1.DatabaseClusterRestore, error)
	// ListDatabaseEngines returns list of managed database clusters.
	ListDatabaseEngines(ctx context.Context, namespace string) (*everestv1alpha1.DatabaseEngineList, error)
	// GetDatabaseEngine returns database clusters by provided name.
	GetDatabaseEngine(ctx context.Context, namespace, name string) (*everestv1alpha1.DatabaseEngine, error)
	// UpdateDatabaseEngine updates a database engine and returns the updated object.
	UpdateDatabaseEngine(ctx context.Context, namespace string, engine *everestv1alpha1.DatabaseEngine) (*everestv1alpha1.DatabaseEngine, error)
	// GetDeployment returns deployment by name.
	GetDeployment(ctx context.Context, name string, namespace string) (*appsv1.Deployment, error)
	// ListDeployments returns deployment by name.
	ListDeployments(ctx context.Context, namespace string) (*appsv1.DeploymentList, error)
	// UpdateDeployment updates a deployment and returns the updated object.
	UpdateDeployment(ctx context.Context, deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	// GetInstallPlan retrieves an OLM install plan by namespace and name.
	GetInstallPlan(ctx context.Context, namespace string, name string) (*v1alpha1.InstallPlan, error)
	// ListInstallPlans lists install plans.
	ListInstallPlans(ctx context.Context, namespace string) (*v1alpha1.InstallPlanList, error)
	// UpdateInstallPlan updates the existing install plan in the specified namespace.
	UpdateInstallPlan(ctx context.Context, namespace string, installPlan *v1alpha1.InstallPlan) (*v1alpha1.InstallPlan, error)
	// DeleteAllMonitoringResources deletes all resources related to monitoring from k8s cluster.
	DeleteAllMonitoringResources(ctx context.Context, namespace string) error
	// CreateMonitoringConfig creates an monitoringConfig.
	CreateMonitoringConfig(ctx context.Context, config *everestv1alpha1.MonitoringConfig) error
	// UpdateMonitoringConfig updates an monitoringConfig.
	UpdateMonitoringConfig(ctx context.Context, config *everestv1alpha1.MonitoringConfig) error
	// GetMonitoringConfig returns the monitoringConfig.
	GetMonitoringConfig(ctx context.Context, namespace, name string) (*everestv1alpha1.MonitoringConfig, error)
	// ListMonitoringConfigs returns the monitoringConfig.
	ListMonitoringConfigs(ctx context.Context, namespace string) (*everestv1alpha1.MonitoringConfigList, error)
	// DeleteMonitoringConfig deletes the monitoringConfig.
	DeleteMonitoringConfig(ctx context.Context, namespace, name string) error
	// CreateNamespace creates the given namespace.
	CreateNamespace(ctx context.Context, namespace *corev1.Namespace) (*corev1.Namespace, error)
	// GetNamespace returns a namespace.
	GetNamespace(ctx context.Context, name string) (*corev1.Namespace, error)
	// DeleteNamespace deletes a namespace.
	DeleteNamespace(ctx context.Context, name string) error
	// ListNamespaces returns a list of namespaces.
	ListNamespaces(ctx context.Context, opts metav1.ListOptions) (*corev1.NamespaceList, error)
	// UpdateNamespace updates the given namespace.
	UpdateNamespace(ctx context.Context, namespace *corev1.Namespace, opts metav1.UpdateOptions) (*corev1.Namespace, error)
	// OLM returns OLM client set.
	//
	//nolint:ireturn
	OLM() versioned.Interface
	// GetNodes returns list of nodes.
	GetNodes(ctx context.Context) (*corev1.NodeList, error)
	// GetPods returns list of pods.
	GetPods(ctx context.Context, namespace string, labelSelector *metav1.LabelSelector) (*corev1.PodList, error)
	// ListPods lists pods.
	ListPods(ctx context.Context, namespace string, options metav1.ListOptions) (*corev1.PodList, error)
	// DeletePod deletes a pod by given name in the given namespace.
	DeletePod(ctx context.Context, namespace, name string) error
	// ListSecrets returns secrets.
	ListSecrets(ctx context.Context, namespace string) (*corev1.SecretList, error)
	// GetSecret returns secret by name.
	GetSecret(ctx context.Context, namespace, name string) (*corev1.Secret, error)
	// UpdateSecret updates k8s Secret.
	UpdateSecret(ctx context.Context, secret *corev1.Secret) (*corev1.Secret, error)
	// CreateSecret creates k8s Secret.
	CreateSecret(ctx context.Context, secret *corev1.Secret) (*corev1.Secret, error)
	// DeleteSecret deletes the k8s Secret.
	DeleteSecret(ctx context.Context, namespace, name string) error
	// GetStorageClasses returns all storage classes available in the cluster.
	GetStorageClasses(ctx context.Context) (*storagev1.StorageClassList, error)
	// GetPersistentVolumes returns Persistent Volumes available in the cluster.
	GetPersistentVolumes(ctx context.Context) (*corev1.PersistentVolumeList, error)
}
