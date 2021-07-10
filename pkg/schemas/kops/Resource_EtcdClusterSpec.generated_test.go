package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceEtcdClusterSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdClusterSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceEtcdClusterSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceEtcdClusterSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceEtcdClusterSpecInto(t *testing.T) {
	type args struct {
		in  kops.EtcdClusterSpec
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
			FlattenResourceEtcdClusterSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceEtcdClusterSpec(t *testing.T) {
	type args struct {
		in kops.EtcdClusterSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.EtcdClusterSpec{},
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "Provider - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.Provider = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "Member - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.Members = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "EnableEtcdTLS - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.EnableEtcdTLS = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "EnableTLSAuth - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.EnableTLSAuth = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "LeaderElectionTimeout - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.LeaderElectionTimeout = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "HeartbeatInterval - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.HeartbeatInterval = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "Image - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.Image = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "Backups - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.Backups = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "Manager - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.Manager = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":                    "",
				"provider":                "",
				"member":                  func() []interface{} { return nil }(),
				"enable_etcd_tls":         false,
				"enable_tls_auth":         false,
				"version":                 "",
				"leader_election_timeout": nil,
				"heartbeat_interval":      nil,
				"image":                   "",
				"backups":                 nil,
				"manager":                 nil,
				"memory_request":          nil,
				"cpu_request":             nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceEtcdClusterSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceEtcdClusterSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
