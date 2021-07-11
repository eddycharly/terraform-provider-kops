package main

import (
	"reflect"
	"text/template"
)

func reflectFuncs() template.FuncMap {
	return template.FuncMap{
		"isPtr":         isPtr,
		"isBool":        isBool,
		"isInt":         isInt,
		"isString":      isString,
		"isFloat":       isFloat,
		"isList":        isSlice,
		"isStruct":      isStruct,
		"isMap":         isMap,
		"isDuration":    isDuration,
		"isQuantity":    isQuantity,
		"isIntOrString": isIntOrString,
		"isValueType":   isValueType,
		"fields":        getFields,
	}
}

func optionFuncs(dataSource bool, optionsMap map[reflect.Type]*options) template.FuncMap {
	return template.FuncMap{
		"needsSchema": func(t reflect.Type) bool {
			return !optionsMap[t].noSchema
		},
		"hasVersion": func(t reflect.Type) bool {
			return optionsMap[t].version != nil
		},
		"schemaVersion": func(t reflect.Type) int {
			return *(optionsMap[t].version)
		},
		"isExcluded": func(in _field) bool {
			return optionsMap[in.Owner].exclude.Has(in.Name)
		},
		"isRequired": func(in _field) bool {
			return optionsMap[in.Owner].required.Has(in.Name)
		},
		"isOptional": func(in _field) bool {
			if dataSource {
				return optionsMap[in.Owner].computed.Has(in.Name)
			}
			return !optionsMap[in.Owner].required.Has(in.Name) && !optionsMap[in.Owner].computedOnly.Has(in.Name)
		},
		"isNullable": func(in _field) bool {
			return optionsMap[in.Owner].nullable.Has(in.Name)
		},
		"isComputed": func(in _field) bool {
			if optionsMap[in.Owner].required.Has(in.Name) {
				return false
			}
			return dataSource || (optionsMap[in.Owner].computed.Has(in.Name) || optionsMap[in.Owner].computedOnly.Has(in.Name))
		},
		"isSensitive": func(in _field) bool {
			return optionsMap[in.Owner].sensitive.Has(in.Name)
		},
		"forceNew": func(in _field) bool {
			return optionsMap[in.Owner].forceNew.Has(in.Name)
		},
		"suppressDiff": func(in _field) bool {
			return optionsMap[in.Owner].suppressDiff.Has(in.Name)
		},
		"fieldName": func(in _field) string {
			if optionsMap[in.Owner].rename[in.Name] != "" {
				return fieldName(optionsMap[in.Owner].rename[in.Name])
			}
			return fieldName(in.Name)
		},
	}
}

func docFuncs(parser *parser, optionsMap map[reflect.Type]*options) template.FuncMap {
	return template.FuncMap{
		"resourceComment": func(t reflect.Type) string {
			return getResourceComment(t.PkgPath(), t.Name(), parser.packs)
		},
		"attributeComment": func(f _field) string {
			return getAttributeComment(f.Owner.PkgPath(), f.Owner.Name(), f.Name, parser.packs)
		},
		"subResources": func(t reflect.Type) []reflect.Type {
			return getSubResources(t, map[reflect.Type]bool{}, func(in _field) bool {
				return optionsMap[in.Owner].exclude.Has(in.Name)
			})[1:]
		},
	}
}
