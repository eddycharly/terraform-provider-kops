package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kops/pkg/apis/kops"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type options struct {
	exclude      []string
	required     []string
	computed     []string
	computedOnly []string
	sensitive    []string
	suppressDiff []string
}

func excluded(excluded ...string) func(o options) options {
	return func(o options) options {
		o.exclude = append(o.exclude, excluded...)
		return o
	}
}

func required(required ...string) func(o options) options {
	return func(o options) options {
		o.required = append(o.required, required...)
		return o
	}
}

func computed(computed ...string) func(o options) options {
	return func(o options) options {
		o.computed = append(o.computed, computed...)
		return o
	}
}

func computedOnly(computedOnly ...string) func(o options) options {
	return func(o options) options {
		o.computedOnly = append(o.computedOnly, computedOnly...)
		return o
	}
}

func sensitive(sensitive ...string) func(o options) options {
	return func(o options) options {
		o.sensitive = append(o.sensitive, sensitive...)
		return o
	}
}

func suppressDiff(suppressDiff ...string) func(o options) options {
	return func(o options) options {
		o.suppressDiff = append(o.suppressDiff, suppressDiff...)
		return o
	}
}

func schemaType(in reflect.Type) string {
	switch in.Kind() {
	case reflect.String:
		return "String"
	case reflect.Bool:
		return "Bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "Int"
	case reflect.Float32, reflect.Float64:
		return "Float"
	default:
		panic(fmt.Sprintf("unknown kind %v", in.Kind()))
	}
}

func fieldName(in string) string {
	in = strings.ReplaceAll(in, "CIDR", "Cidr")
	in = strings.ReplaceAll(in, "DNS", "Dns")
	in = strings.ReplaceAll(in, "IP", "Ip")
	in = strings.ReplaceAll(in, "SSH", "Ssh")
	in = strings.ReplaceAll(in, "API", "Api")
	in = strings.ReplaceAll(in, "SAN", "San")
	return in
}

func isValueType(in reflect.Type) bool {
	switch in.Kind() {
	case reflect.Ptr:
		return isValueType(in.Elem())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
		return true
	case reflect.Slice, reflect.Map:
		return false
	case reflect.Struct:
		return in.String() == "v1.Duration" || in.String() == "resource.Quantity" || in.String() == "intstr.IntOrString"
	default:
		panic(fmt.Sprintf("unknown kind %v", in.Kind()))
	}
}

func funcMap(o options) template.FuncMap {
	required := sets.NewString(o.required...)
	computed := sets.NewString(o.computed...)
	computedOnly := sets.NewString(o.computedOnly...)
	sensitive := sets.NewString(o.sensitive...)
	suppressDiff := sets.NewString(o.suppressDiff...)
	return template.FuncMap{
		"fields": func(t reflect.Type) []reflect.StructField {
			var ret []reflect.StructField
			for i := 0; i < t.NumField(); i++ {
				ret = append(ret, t.Field(i))
			}
			return ret
		},
		"schemaType": schemaType,
		"isRequired": func(in string) bool {
			return required.Has(in)
		},
		"isOptional": func(in string) bool {
			return !required.Has(in) && !computedOnly.Has(in)
		},
		"isComputed": func(in string) bool {
			return computed.Has(in) || computedOnly.Has(in)
		},
		"isSensitive": func(in string) bool {
			return sensitive.Has(in)
		},
		"suppressDiff": func(in string) bool {
			return suppressDiff.Has(in)
		},
		"fieldName":   fieldName,
		"isValueType": isValueType,
		"code": func(in string) string {
			return fmt.Sprintf("`%s`", in)
		},
		"isPtr": func(t reflect.Type) bool {
			return t.Kind() == reflect.Ptr
		},
		"isList": func(t reflect.Type) bool {
			return t.Kind() == reflect.Slice
		},
		"isStruct": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct
		},
		"isMap": func(t reflect.Type) bool {
			return t.Kind() == reflect.Map
		},
		"isDuration": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct && t.String() == "v1.Duration"
		},
		"isQuantity": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct && t.String() == "resource.Quantity"
		},
		"isIntOrString": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct && t.String() == "intstr.IntOrString"
		},
	}
}

func buildDoc(t reflect.Type, o options) {
	tmplString := `# kops_{{ fieldName .Type.Name | snakecase }}

{{- define "type" -}}
{{- if isPtr . -}}
{{- template "type" .Elem }}
{{- else if isList . -}}
List({{- template "type" .Elem }})
{{- else if isMap . -}}
Map({{- template "type" .Elem }})
{{- else if isDuration . -}}
Duration
{{- else if isQuantity . -}}
Quantity
{{- else if isIntOrString . -}}
IntOrString
{{- else if isStruct . -}}
[{{ .Name }}](./{{ .Name }}.md)
{{- else -}}
{{- schemaType . -}}
{{- end -}}
{{- end }}

{{- define "required" -}}
{{ if isRequired .Name }}:white_check_mark:{{ end }}
{{- end }}

{{- define "optional" -}}
{{ if isOptional .Name }}:white_check_mark:{{ end }}
{{- end }}

{{- define "computed" -}}
{{ if isComputed .Name }}:white_check_mark:{{ end }}
{{- end }}

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
{{- with .Type -}}
{{- range (fields .) -}}
{{- if not (has .Name $.Exclude) }}
| {{ fieldName .Name | snakecase | code }} | {{ template "type" .Type }} | {{ template "required" . }} | {{ template "optional" . }} | {{ template "computed" . }} |
{{- end -}}
{{- end -}}
{{- end }}
`
	tmpl := template.New("doc").Funcs(funcMap(o)).Funcs(sprig.TxtFuncMap())

	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}

	fileName := toSnakeCase(fieldName(t.Name())) + ".md"
	f, _ := os.Create("website/docs/resources/" + fileName)
	defer f.Close()

	if err := tmpl.Execute(f, map[string]interface{}{
		"Type":    t,
		"Exclude": o.exclude,
	}); err != nil {
		panic(err)
	}
}

func build(i interface{}, o ...func(o options) options) {
	opts := options{}
	for _, opt := range o {
		opts = opt(opts)
	}
	t := reflect.TypeOf(i)
	buildDoc(t, opts)
}

func main() {
	build(api.ProviderConfig{},
		required("StateStore"),
	)
	build(api.AwsConfig{})
	build(api.Cluster{},
		required("Name", "AdminSshKey", "CloudProvider", "Subnet", "NetworkID", "Topology", "EtcdCluster", "Networking", "InstanceGroup"),
		computed("MasterPublicName", "MasterInternalName", "ConfigBase", "NetworkCIDR", "NonMasqueradeCIDR", "IAM"),
		computedOnly("KubeConfig"),
		sensitive("AdminSshKey", "KubeConfig"),
		suppressDiff("RollingUpdateOptions", "ValidateOptions"),
	)
	build(api.RollingUpdateOptions{})
	build(api.ValidateOptions{})
	build(api.KubeConfig{},
		computedOnly("Server", "Context", "Namespace", "KubeBearerToken", "KubeUser", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
		sensitive("KubeBearerToken", "KubeUser", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
	)
	build(kops.AddonSpec{},
		required("Manifest"),
	)
	build(kops.EgressProxySpec{},
		required("HTTPProxy"),
	)
	build(kops.HTTPProxy{},
		required("Host", "Port"),
	)
	build(kops.ContainerdConfig{})
	build(kops.DockerConfig{})
	build(kops.KubeDNSConfig{})
	build(kops.KubeAPIServerConfig{})
	build(kops.KubeControllerManagerConfig{})
	build(kops.CloudControllerManagerConfig{})
	build(kops.KubeSchedulerConfig{})
	build(kops.KubeProxyConfig{})
	build(kops.KubeletConfigSpec{})
	build(kops.CloudConfiguration{})
	build(kops.ExternalDNSConfig{})
	build(kops.OpenstackConfiguration{})
	build(kops.OpenstackLoadbalancerConfig{})
	build(kops.OpenstackMonitor{})
	build(kops.OpenstackRouter{})
	build(kops.OpenstackBlockStorageConfig{})
	build(kops.LeaderElectionConfiguration{})
	build(kops.NodeLocalDNSConfig{})
	build(kops.AuthenticationSpec{})
	build(kops.AuthorizationSpec{})
	build(kops.NodeAuthorizationSpec{})
	build(kops.Assets{})
	build(kops.IAMSpec{})
	build(kops.KopeioAuthenticationSpec{})
	build(kops.AwsAuthenticationSpec{})
	build(kops.AlwaysAllowAuthorizationSpec{})
	build(kops.RBACAuthorizationSpec{})
	build(kops.NodeAuthorizerSpec{})
	build(api.InstanceGroup{},
		required("Name", "Role", "MinSize", "MaxSize", "MachineType", "Subnets"),
		computed("Image"),
	)
	build(kops.AccessSpec{})
	build(kops.DNSAccessSpec{})
	build(kops.LoadBalancerAccessSpec{},
		required("Type"),
	)
	build(kops.EtcdClusterSpec{},
		required("Name", "Members"),
	)
	build(kops.EtcdBackupSpec{},
		required("BackupStore", "Image"),
	)
	build(kops.EtcdManagerSpec{},
		required("Image"),
	)
	build(kops.EtcdMemberSpec{},
		required("Name", "InstanceGroup"),
	)
	build(kops.EnvVar{},
		required("Name"),
	)
	build(kops.ClusterSubnetSpec{},
		required("Name", "ProviderID", "Type", "Zone"),
		computed("CIDR"),
	)
	build(kops.TopologySpec{},
		required("Masters", "Nodes", "DNS"),
	)
	build(kops.BastionSpec{},
		required("BastionPublicName"),
	)
	build(kops.BastionLoadBalancerSpec{},
		required("AdditionalSecurityGroups"),
	)
	build(kops.DNSSpec{},
		required("Type"),
	)
	build(kops.NetworkingSpec{})
	build(kops.ClassicNetworkingSpec{})
	build(kops.KubenetNetworkingSpec{})
	build(kops.ExternalNetworkingSpec{})
	build(kops.CNINetworkingSpec{})
	build(kops.KopeioNetworkingSpec{})
	build(kops.WeaveNetworkingSpec{})
	build(kops.FlannelNetworkingSpec{})
	build(kops.CalicoNetworkingSpec{})
	build(kops.CanalNetworkingSpec{})
	build(kops.KuberouterNetworkingSpec{})
	build(kops.RomanaNetworkingSpec{})
	build(kops.AmazonVPCNetworkingSpec{})
	build(kops.CiliumNetworkingSpec{})
	build(kops.LyftVPCNetworkingSpec{})
	build(kops.GCENetworkingSpec{})
	build(kops.VolumeSpec{},
		required("Device"),
	)
	build(kops.VolumeMountSpec{},
		required("Device", "Filesystem", "Path"),
	)
	build(kops.MixedInstancesPolicySpec{})
	build(kops.UserData{},
		required("Name", "Type", "Content"),
	)
	build(kops.LoadBalancer{})
	build(kops.IAMProfileSpec{},
		required("Profile"),
	)
	build(kops.HookSpec{},
		required("Name"),
	)
	build(kops.ExecContainerAction{},
		required("Image"),
	)
	build(kops.FileAssetSpec{},
		required("Name", "Path", "Content"),
	)
	build(kops.RollingUpdate{})
}
