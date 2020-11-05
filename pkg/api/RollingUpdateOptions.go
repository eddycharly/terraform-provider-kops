package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RollingUpdateOptions struct {
	Skip              bool
	MasterInterval    *metav1.Duration
	NodeInterval      *metav1.Duration
	BastionInterval   *metav1.Duration
	FailOnDrainError  bool
	FailOnValidate    bool
	PostDrainDelay    *metav1.Duration
	ValidationTimeout *metav1.Duration
	ValidateCount     *int
}
