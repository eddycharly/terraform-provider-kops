package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceHookSpec(t *testing.T) {
	_default := kops.HookSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.HookSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceHookSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceHookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceHookSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"name":             "",
		"disabled":         false,
		"roles":            func() []interface{} { return nil }(),
		"requires":         func() []interface{} { return nil }(),
		"before":           func() []interface{} { return nil }(),
		"exec_container":   nil,
		"manifest":         "",
		"use_raw_manifest": false,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceHookSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceHookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceHookSpec(t *testing.T) {
	_default := map[string]interface{}{
		"name":             "",
		"disabled":         false,
		"roles":            func() []interface{} { return nil }(),
		"requires":         func() []interface{} { return nil }(),
		"before":           func() []interface{} { return nil }(),
		"exec_container":   nil,
		"manifest":         "",
		"use_raw_manifest": false,
	}
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
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
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceHookSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceHookSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
