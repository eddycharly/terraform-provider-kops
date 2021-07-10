package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceLeaderElectionConfiguration(t *testing.T) {
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
			if got := ExpandDataSourceLeaderElectionConfiguration(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandDataSourceLeaderElectionConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenDataSourceLeaderElectionConfigurationInto(t *testing.T) {
	_default := map[string]interface{}{
		"leader_elect":                         nil,
		"leader_elect_lease_duration":          nil,
		"leader_elect_renew_deadline_duration": nil,
		"leader_elect_resource_lock":           nil,
		"leader_elect_resource_name":           nil,
		"leader_elect_resource_namespace":      nil,
		"leader_elect_retry_period":            nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceLeaderElectionConfigurationInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLeaderElectionConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceLeaderElectionConfiguration(t *testing.T) {
	_default := map[string]interface{}{
		"leader_elect":                         nil,
		"leader_elect_lease_duration":          nil,
		"leader_elect_renew_deadline_duration": nil,
		"leader_elect_resource_lock":           nil,
		"leader_elect_resource_name":           nil,
		"leader_elect_resource_namespace":      nil,
		"leader_elect_retry_period":            nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceLeaderElectionConfiguration(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceLeaderElectionConfiguration() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
