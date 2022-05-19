package resources

type ApplyOptions struct {
	// Skip allows skipping cluster apply
	Skip bool
	// AllowKopsDowngrade permits applying with a kops version older than what was last used to apply to the cluster
	AllowKopsDowngrade bool
	// LifecycleOverrides is passed in to override the lifecycle for one of more tasks.
	// The key value is the task name such as InternetGateway and the value is the fi.Lifecycle
	// that is re-mapped.
	LifecycleOverrides map[string]string
}
