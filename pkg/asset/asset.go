package asset

const (
	AssetPathSecrets            = "tls"
	AssetPathAdminKubeConfig    = "auth/kubeconfig"
	AssetPathManifests          = "manifests"
	AssetPathBootstrapManifests = "bootstrap-manifests"
)

var (
	BootstrapSecretsDir = "/etc/kubernetes/bootstrap-secrets" // Overridden for testing.
)
