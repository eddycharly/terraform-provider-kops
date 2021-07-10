package schemas

import (
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigAws(t *testing.T) {
	_default := config.Aws{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Aws
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"profile":           "",
					"region":            "",
					"access_key":        "",
					"secret_key":        "",
					"assume_role":       nil,
					"s3_endpoint":       "",
					"s3_region":         "",
					"s3_access_key":     "",
					"s3_secret_key":     "",
					"skip_region_check": false,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandConfigAws(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandConfigAws() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigAwsInto(t *testing.T) {
	_default := map[string]interface{}{
		"profile":           "",
		"region":            "",
		"access_key":        "",
		"secret_key":        "",
		"assume_role":       nil,
		"s3_endpoint":       "",
		"s3_region":         "",
		"s3_access_key":     "",
		"s3_secret_key":     "",
		"skip_region_check": false,
	}
	type args struct {
		in config.Aws
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.Aws{},
			},
			want: _default,
		},
		{
			name: "Profile - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.Profile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Region - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AccessKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.AccessKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.SecretKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AssumeRole - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.AssumeRole = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3Endpoint - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3Endpoint = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3Region - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3AccessKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3AccessKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3SecretKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3SecretKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipRegionCheck - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.SkipRegionCheck = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenConfigAwsInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigAws() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenConfigAws(t *testing.T) {
	_default := map[string]interface{}{
		"profile":           "",
		"region":            "",
		"access_key":        "",
		"secret_key":        "",
		"assume_role":       nil,
		"s3_endpoint":       "",
		"s3_region":         "",
		"s3_access_key":     "",
		"s3_secret_key":     "",
		"skip_region_check": false,
	}
	type args struct {
		in config.Aws
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: config.Aws{},
			},
			want: _default,
		},
		{
			name: "Profile - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.Profile = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Region - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AccessKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.AccessKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SecretKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.SecretKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "AssumeRole - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.AssumeRole = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3Endpoint - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3Endpoint = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3Region - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3Region = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3AccessKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3AccessKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "S3SecretKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3SecretKey = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "SkipRegionCheck - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.SkipRegionCheck = false
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenConfigAws(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenConfigAws() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
