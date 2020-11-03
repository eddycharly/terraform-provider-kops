package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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

// Quantity

func OptionalQuantity() *schema.Schema {
	return OptionalString()
}

// Duration

func OptionalDuration() *schema.Schema {
	return OptionalString()
}

// IntOrString

func OptionalIntOrString() *schema.Schema {
	return OptionalString()
}

// Map

func RequiredMap(elem *schema.Schema) *schema.Schema {
	return Schema(schema.TypeMap, elem, true, false, false, 0)
}

func OptionalMap(elem *schema.Schema) *schema.Schema {
	return Schema(schema.TypeMap, elem, false, true, false, 0)
}

func OptionalComputedMap(elem *schema.Schema) *schema.Schema {
	return Schema(schema.TypeMap, elem, false, true, true, 0)
}

// Struct

func RequiredStruct(elem *schema.Resource) *schema.Schema {
	return Schema(schema.TypeList, elem, true, false, false, 1)
}

func OptionalStruct(elem *schema.Resource) *schema.Schema {
	return Schema(schema.TypeList, elem, false, true, false, 1)
}

func OptionalComputedStruct(elem *schema.Resource) *schema.Schema {
	return Schema(schema.TypeList, elem, false, true, true, 1)
}

// List

func List(elem interface{}) *schema.Schema {
	return Schema(schema.TypeList, elem, false, false, false, 0)
}

func RequiredList(elem interface{}) *schema.Schema {
	return Schema(schema.TypeList, elem, true, false, false, 0)
}

func OptionalList(elem interface{}) *schema.Schema {
	return Schema(schema.TypeList, elem, false, true, false, 0)
}

func OptionalComputedList(elem interface{}) *schema.Schema {
	return Schema(schema.TypeList, elem, false, true, true, 0)
}

// String

func String() *schema.Schema {
	return Simple(schema.TypeString, false, false, false)
}

func RequiredString() *schema.Schema {
	return SimpleRequired(schema.TypeString)
}

func OptionalString() *schema.Schema {
	return SimpleOptional(schema.TypeString)
}

func ComputedString() *schema.Schema {
	return SimpleOptional(schema.TypeString)
}

func OptionalComputedString() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeString)
}

// Bool

func Bool() *schema.Schema {
	return Simple(schema.TypeBool, false, false, false)
}

func RequiredBool() *schema.Schema {
	return SimpleRequired(schema.TypeBool)
}

func OptionalBool() *schema.Schema {
	return SimpleOptional(schema.TypeBool)
}

func ComputedBool() *schema.Schema {
	return SimpleComputed(schema.TypeBool)
}

func OptionalComputedBool() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeBool)
}

// Int

func Int() *schema.Schema {
	return Simple(schema.TypeInt, false, false, false)
}

func RequiredInt() *schema.Schema {
	return SimpleRequired(schema.TypeInt)
}

func OptionalInt() *schema.Schema {
	return SimpleOptional(schema.TypeInt)
}

func ComputedInt() *schema.Schema {
	return SimpleComputed(schema.TypeInt)
}

func OptionalComputedInt() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeInt)
}

// Float

func Float() *schema.Schema {
	return Simple(schema.TypeFloat, false, false, false)
}

func RequiredFloat() *schema.Schema {
	return SimpleRequired(schema.TypeFloat)
}

func OptionalFloat() *schema.Schema {
	return SimpleOptional(schema.TypeFloat)
}

func ComputedFloat() *schema.Schema {
	return SimpleComputed(schema.TypeFloat)
}

func OptionalComputedFloat() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeFloat)
}
