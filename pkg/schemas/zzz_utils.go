package schemas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Optional(in *schema.Schema) *schema.Schema {
	in.Optional = true
	return in
}

func Required(in *schema.Schema) *schema.Schema {
	in.Required = true
	return in
}

func Computed(in *schema.Schema) *schema.Schema {
	in.Computed = true
	in.MaxItems = 0
	return in
}

func Sensitive(in *schema.Schema) *schema.Schema {
	in.Sensitive = true
	return in
}

func ForceNew(in *schema.Schema) *schema.Schema {
	in.ForceNew = true
	return in
}

func SuppressDiff(in *schema.Schema) *schema.Schema {
	in.DiffSuppressFunc = func(_, _, _ string, _ *schema.ResourceData) bool {
		return true
	}
	return in
}

func Struct(in *schema.Resource) *schema.Schema {
	return Schema(schema.TypeList, in, false, false, false, 1)
}

func List(elem interface{}) *schema.Schema {
	return Schema(schema.TypeList, elem, false, false, false, 0)
}

func Nullable(in *schema.Schema) *schema.Schema {
	return Struct(
		&schema.Resource{
			Schema: map[string]*schema.Schema{
				"value": in,
			},
		},
	)
}

func Quantity() *schema.Schema {
	return String()
}

func Duration() *schema.Schema {
	return String()
}

func IntOrString() *schema.Schema {
	return String()
}

func Map(elem *schema.Schema) *schema.Schema {
	return Schema(schema.TypeMap, elem, false, false, false, 0)
}

func String() *schema.Schema {
	return Simple(schema.TypeString, false, false, false)
}

func FlattenString(in string) interface{} {
	return in
}

func Bool() *schema.Schema {
	return Simple(schema.TypeBool, false, false, false)
}

func FlattenBool(in bool) interface{} {
	return in
}

func Int() *schema.Schema {
	return Simple(schema.TypeInt, false, false, false)
}

func FlattenInt(in int) interface{} {
	return in
}

func Float() *schema.Schema {
	return Simple(schema.TypeFloat, false, false, false)
}

func FlattenFloat(in float64) interface{} {
	return in
}

// Diff

func CustomizeDiffRevision(_ context.Context, d *schema.ResourceDiff, m interface{}) error {
	// If anything changes, increment the revision
	if len(d.GetChangedKeysPrefix("")) > 0 {
		d.SetNew("revision", d.Get("revision").(int)+1)
	}
	return nil
}

// Tools

func Schema(t schema.ValueType, elem interface{}, required, optional, computed bool, maxItems int) *schema.Schema {
	return &schema.Schema{
		Type:     t,
		Optional: optional,
		Required: required,
		Computed: computed,
		MaxItems: maxItems,
		Elem:     elem,
	}
}

func Simple(t schema.ValueType, required, optional, computed bool) *schema.Schema {
	return Schema(t, nil, required, optional, computed, 0)
}

func SimpleRequired(t schema.ValueType) *schema.Schema {
	return Simple(t, true, false, false)
}

func SimpleOptional(t schema.ValueType) *schema.Schema {
	return Simple(t, false, true, false)
}

func SimpleComputed(t schema.ValueType) *schema.Schema {
	return Simple(t, false, false, true)
}

func SimpleOptionalComputed(t schema.ValueType) *schema.Schema {
	return Simple(t, false, true, true)
}

// // Set

// func RequiredSetList(elem interface{}) *schema.Schema {
// 	return Schema(schema.TypeSet, elem, true, false, false, 0)
// }

// func OptionalSetList(elem interface{}) *schema.Schema {
// 	return Schema(schema.TypeSet, elem, false, true, false, 0)
// }

// List

// func RequiredList(elem interface{}) *schema.Schema {
// 	return Schema(schema.TypeList, elem, true, false, false, 0)
// }

// func OptionalList(elem interface{}) *schema.Schema {
// 	return Schema(schema.TypeList, elem, false, true, false, 0)
// }

// func ComputedList(elem interface{}) *schema.Schema {
// 	return Schema(schema.TypeList, elem, false, false, true, 0)
// }

// func OptionalComputedList(elem interface{}) *schema.Schema {
// 	return Schema(schema.TypeList, elem, false, true, true, 0)
// }
