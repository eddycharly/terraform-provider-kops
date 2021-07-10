package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceHookSpec(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.HookSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandResourceHookSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandResourceHookSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenResourceHookSpecInto(t *testing.T) {
	type args struct {
		in  kops.HookSpec
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
			FlattenResourceHookSpecInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenResourceHookSpec(t *testing.T) {
	type args struct {
		in kops.HookSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.HookSpec{},
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "Name - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.Name = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "Disabled - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.Disabled = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "Roles - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.Roles = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "Requires - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.Requires = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "Before - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.Before = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "ExecContainer - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.ExecContainer = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "Manifest - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.Manifest = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
		{
			name: "UseRawManifest - default",
			args: args{
				in: func() kops.HookSpec {
					subject := kops.HookSpec{}
					subject.UseRawManifest = false
					return subject
				}(),
			},
			want: map[string]interface{}{
				"name":             "",
				"disabled":         false,
				"roles":            func() []interface{} { return nil }(),
				"requires":         func() []interface{} { return nil }(),
				"before":           func() []interface{} { return nil }(),
				"exec_container":   nil,
				"manifest":         "",
				"use_raw_manifest": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenResourceHookSpec(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenResourceHookSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
