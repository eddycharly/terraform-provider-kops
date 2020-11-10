package api

type ProviderConfig struct {
	// StateStore defines the state store used by kops
	StateStore string
	// Aws contains the aws configuration options
	Aws *AwsConfig
	// RollingUpdateOptions contains the options used when doing a cluster rolling update
	RollingUpdateOptions RollingUpdateOptions
	// ValidateOptions contains the options used when validating the cluster
	ValidateOptions ValidateOptions
}
