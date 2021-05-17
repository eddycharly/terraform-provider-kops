package resources

type ApplyOptions struct {
	// Skip allows skipping cluster apply
	Skip bool
	// AllowKopsDowngrade permits applying with a kops version older than what was last used to apply to the cluster
	AllowKopsDowngrade bool
}
