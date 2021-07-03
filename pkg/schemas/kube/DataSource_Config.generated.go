package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
)

var _ = Schema

func ExpandDataSourceConfig(in map[string]interface{}) kube.Config {
	if in == nil {
		panic("expand Config failure, in is nil")
	}
	out := kube.Config{}
	if in, ok := in["server"]; ok && in != nil {
		out.Server = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["context"]; ok && in != nil {
		out.Context = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["namespace"]; ok && in != nil {
		out.Namespace = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kube_user"]; ok && in != nil {
		out.KubeUser = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["kube_password"]; ok && in != nil {
		out.KubePassword = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["ca_cert"]; ok && in != nil {
		out.CaCert = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["client_cert"]; ok && in != nil {
		out.ClientCert = func(in interface{}) string { return string(in.(string)) }(in)
	}
	if in, ok := in["client_key"]; ok && in != nil {
		out.ClientKey = func(in interface{}) string { return string(in.(string)) }(in)
	}
	return out
}
