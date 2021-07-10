package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceLeaderElectionConfiguration(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.LeaderElectionConfiguration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceLeaderElectionConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceLeaderElectionConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceLeaderElectionConfigurationInto(t *testing.T) {
	type args struct {
		in  kops.LeaderElectionConfiguration
		out map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlattenResourceLeaderElectionConfigurationInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceLeaderElectionConfiguration(t *testing.T) {
	type args struct {
		in kops.LeaderElectionConfiguration
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.LeaderElectionConfiguration{},
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
		{
			name: "LeaderElect - default",
			args: args{
				in: func() kops.LeaderElectionConfiguration {
					subject := kops.LeaderElectionConfiguration{}
					subject.LeaderElect = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
		{
			name: "LeaderElectLeaseDuration - default",
			args: args{
				in: func() kops.LeaderElectionConfiguration {
					subject := kops.LeaderElectionConfiguration{}
					subject.LeaderElectLeaseDuration = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
		{
			name: "LeaderElectRenewDeadlineDuration - default",
			args: args{
				in: func() kops.LeaderElectionConfiguration {
					subject := kops.LeaderElectionConfiguration{}
					subject.LeaderElectRenewDeadlineDuration = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
		{
			name: "LeaderElectResourceLock - default",
			args: args{
				in: func() kops.LeaderElectionConfiguration {
					subject := kops.LeaderElectionConfiguration{}
					subject.LeaderElectResourceLock = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
		{
			name: "LeaderElectResourceName - default",
			args: args{
				in: func() kops.LeaderElectionConfiguration {
					subject := kops.LeaderElectionConfiguration{}
					subject.LeaderElectResourceName = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
		{
			name: "LeaderElectResourceNamespace - default",
			args: args{
				in: func() kops.LeaderElectionConfiguration {
					subject := kops.LeaderElectionConfiguration{}
					subject.LeaderElectResourceNamespace = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
		{
			name: "LeaderElectRetryPeriod - default",
			args: args{
				in: func() kops.LeaderElectionConfiguration {
					subject := kops.LeaderElectionConfiguration{}
					subject.LeaderElectRetryPeriod = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"leader_elect":                         nil,
				"leader_elect_lease_duration":          nil,
				"leader_elect_renew_deadline_duration": nil,
				"leader_elect_resource_lock":           nil,
				"leader_elect_resource_name":           nil,
				"leader_elect_resource_namespace":      nil,
				"leader_elect_retry_period":            nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceLeaderElectionConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceLeaderElectionConfiguration() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
