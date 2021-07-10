package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceWeaveNetworkingSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.WeaveNetworkingSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceWeaveNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceWeaveNetworkingSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceWeaveNetworkingSpecInto(t *testing.T) {
	type args struct {
		in  kops.WeaveNetworkingSpec
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
			FlattenResourceWeaveNetworkingSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceWeaveNetworkingSpec(t *testing.T) {
	type args struct {
		in kops.WeaveNetworkingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.WeaveNetworkingSpec{},
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "MTU - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.MTU = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "ConnLimit - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.ConnLimit = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "NoMasqLocal - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.NoMasqLocal = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "MemoryRequest - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.MemoryRequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "CPURequest - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.CPURequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "MemoryLimit - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.MemoryLimit = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "CPULimit - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.CPULimit = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "NetExtraArgs - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.NetExtraArgs = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "NPCMemoryRequest - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.NPCMemoryRequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "NPCCPURequest - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.NPCCPURequest = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "NPCMemoryLimit - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.NPCMemoryLimit = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "NPCCPULimit - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.NPCCPULimit = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "NPCExtraArgs - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.NPCExtraArgs = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.WeaveNetworkingSpec {
					subject := kops.WeaveNetworkingSpec{}
					subject.Version = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"mtu":                nil,
				"conn_limit":         nil,
				"no_masq_local":      nil,
				"memory_request":     nil,
				"cpu_request":        nil,
				"memory_limit":       nil,
				"cpu_limit":          nil,
				"net_extra_args":     "",
				"npc_memory_request": nil,
				"npccpu_request":     nil,
				"npc_memory_limit":   nil,
				"npccpu_limit":       nil,
				"npc_extra_args":     "",
				"version":            "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceWeaveNetworkingSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceWeaveNetworkingSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
