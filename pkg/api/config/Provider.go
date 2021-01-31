package config

type Provider struct {
	// StateStore defines the state store used by kops
	StateStore string
	// Aws contains the aws configuration options
	Aws *Aws
}
