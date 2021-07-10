package schemas

import (
	"reflect"
	"testing"

	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/google/go-cmp/cmp"
)

func TestExpandConfigAws(t *testing.T) {
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want config.Aws
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandConfigAws(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandConfigAws() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenConfigAwsInto(t *testing.T) {
	type args struct {
		in  config.Aws
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
			FlattenConfigAwsInto(tt.args.in, tt.args.out)
		})
	}
}

func TestFlattenConfigAws(t *testing.T) {
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
			want: map[string]interface{}{
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
		{
			name: "Profile - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.Profile = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "Region - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.Region = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "AccessKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.AccessKey = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "SecretKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.SecretKey = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "AssumeRole - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.AssumeRole = nil
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "S3Endpoint - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3Endpoint = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "S3Region - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3Region = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "S3AccessKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3AccessKey = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "S3SecretKey - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.S3SecretKey = ""
					return subject
				}(),
			},
			want: map[string]interface{}{
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
		{
			name: "SkipRegionCheck - default",
			args: args{
				in: func() config.Aws {
					subject := config.Aws{}
					subject.SkipRegionCheck = false
					return subject
				}(),
			},
			want: map[string]interface{}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenConfigAws(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FlattenConfigAws() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
