package schemas

import (
	"context"
	"log"
)

func upgradeNullable(name string, rawState map[string]interface{}) map[string]interface{} {
	log.Printf("UPGRADING %v", name)
	if rawState[name] == nil {
		return rawState
	}
	if _, ok := rawState[name].([]interface{}); !ok {
		rawState[name] = []interface{}{map[string]interface{}{"value": rawState[name]}}
	}
	return rawState
}

func UpgradeDataSourceKubeAPIServerConfig0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState = upgradeNullable("anonymous_auth", rawState)
	return rawState, nil
}

func UpgradeDataSourceKubeletConfigSpec0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState = upgradeNullable("anonymous_auth", rawState)
	return rawState, nil
}

func UpgradeDataSourceMixedInstancesPolicySpec0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState = upgradeNullable("on_demand_base", rawState)
	rawState = upgradeNullable("on_demand_above_base", rawState)
	return rawState, nil
}

func UpgradeResourceKubeAPIServerConfig0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState = upgradeNullable("anonymous_auth", rawState)
	return rawState, nil
}

func UpgradeResourceKubeletConfigSpec0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState = upgradeNullable("anonymous_auth", rawState)
	return rawState, nil
}

func UpgradeResourceMixedInstancesPolicySpec0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState = upgradeNullable("on_demand_base", rawState)
	rawState = upgradeNullable("on_demand_above_base", rawState)
	return rawState, nil
}

func UpgradeResourceCluster0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Printf("UPGRADING UpgradeResourceCluster0")
	return rawState, nil
}
