package config

type Provider struct {
	// StateStore defines the state store used by kops
	StateStore string
	// Aws contains the aws configuration options
	Aws *Aws
	// RollingUpdateOptions contains the options used when doing a cluster rolling update
	RollingUpdate RollingUpdate
	// ValidateOptions contains the options used when validating the cluster
	Validate Validate
}
