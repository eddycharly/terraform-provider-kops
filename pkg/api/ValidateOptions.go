package api

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ValidateOptions struct {
	Skip         bool
	Timeout      *metav1.Duration
	PollInterval *metav1.Duration
}
