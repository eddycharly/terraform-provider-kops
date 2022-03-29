package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceEtcdClusterSpec(t *testing.T) {
	_default := kops.EtcdClusterSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.EtcdClusterSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"name":                    "",
					"provider":                "",
					"member":                  func() []interface{} { return nil }(),
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceEtcdClusterSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceEtcdClusterSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdClusterSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":                    "",
		"provider":                "",
		"member":                  func() []interface{} { return nil }(),
		"version":                 "",
		"leader_election_timeout": nil,
		"heartbeat_interval":      nil,
		"image":                   "",
		"backups":                 nil,
		"manager":                 nil,
		"memory_request":          nil,
		"cpu_request":             nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceEtcdClusterSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdClusterSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceEtcdClusterSpec(t *testing.T) {
	_default := map[string]interface{}{
		"name":                    "",
		"provider":                "",
		"member":                  func() []interface{} { return nil }(),
		"version":                 "",
		"leader_election_timeout": nil,
		"heartbeat_interval":      nil,
		"image":                   "",
		"backups":                 nil,
		"manager":                 nil,
		"memory_request":          nil,
		"cpu_request":             nil,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
		{
			name: "CpuRequest - default",
			args: args{
				in: func() kops.EtcdClusterSpec {
					subject := kops.EtcdClusterSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceEtcdClusterSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceEtcdClusterSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
