package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

func Validate(s *schema.Schema, v schema.SchemaValidateFunc) *schema.Schema {
	s.ValidateFunc = v
	return s
}

func Simple(t schema.ValueType, required, optional, computed bool) *schema.Schema {
	return Schema(t, nil, required, optional, computed, 0)
}

func SimpleRequired(t schema.ValueType) *schema.Schema {
	return Simple(t, true, false, false)
}

func SimpleOptional(t schema.ValueType) *schema.Schema {
	s := Simple(t, false, true, false)
	s.DefaultFunc = func() (interface{}, error) { return nil, nil }
	return s
}

func SimpleOptionalComputed(t schema.ValueType) *schema.Schema {
	return Simple(t, false, true, true)
}

// Quantity

func QuantityOptionalComputed() *schema.Schema {
	return StringOptionalComputed()
}

// Duration

func DurationOptionalComputed() *schema.Schema {
	return StringOptionalComputed()
}

// IntOrString

func IntOrStringOptionalComputed() *schema.Schema {
	return StringOptionalComputed()
}

// Map

func MapOptionalComputed(elem *schema.Schema) *schema.Schema {
	return Schema(schema.TypeMap, elem, false, true, true, 0)
}

// Struct

func StructOptionalComputed(elem *schema.Resource) *schema.Schema {
	s := Schema(schema.TypeList, elem, false, true, true, 1)
	// s.ConfigMode = schema.SchemaConfigModeAttr
	return s
}

func StructRequired(elem *schema.Resource) *schema.Schema {
	s := Schema(schema.TypeList, elem, true, false, false, 1)
	// s.ConfigMode = schema.SchemaConfigModeAttr
	return s
}

// List

func ListOptional(elem interface{}) *schema.Schema {
	return Schema(schema.TypeList, elem, false, true, false, 0)
}

func ListRequired(elem interface{}) *schema.Schema {
	return Schema(schema.TypeList, elem, true, false, false, 0)
}

// String

func String() *schema.Schema {
	return Simple(schema.TypeString, false, false, false)
}

func StringRequired() *schema.Schema {
	return SimpleRequired(schema.TypeString)
}

func StringOptional() *schema.Schema {
	return SimpleOptional(schema.TypeString)
}

func StringOptionalComputed() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeString)
}

func StringEnumRequired(valid ...string) *schema.Schema {
	return Validate(SimpleRequired(schema.TypeString), validation.StringInSlice(valid, false))
}

func StringEnumOptionalComputed(valid ...string) *schema.Schema {
	return Validate(SimpleOptionalComputed(schema.TypeString), validation.StringInSlice(valid, false))
}

// Bool

func Bool() *schema.Schema {
	return Simple(schema.TypeBool, false, false, false)
}

func BoolRequired() *schema.Schema {
	return SimpleRequired(schema.TypeBool)
}

func BoolOptionalComputed() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeBool)
}

// Int

func Int() *schema.Schema {
	return Simple(schema.TypeInt, false, false, false)
}

func IntRequired() *schema.Schema {
	return SimpleRequired(schema.TypeInt)
}

func IntOptional() *schema.Schema {
	return SimpleOptional(schema.TypeInt)
}

func IntOptionalComputed() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeInt)
}

// Float

func Float() *schema.Schema {
	return Simple(schema.TypeFloat, false, false, false)
}

func FloatRequired() *schema.Schema {
	return SimpleRequired(schema.TypeFloat)
}

func FloatOptionalComputed() *schema.Schema {
	return SimpleOptionalComputed(schema.TypeFloat)
}
