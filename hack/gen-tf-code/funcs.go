package main

import (
	"fmt"
	"log"
	"reflect"
	"text/template"
)

func reflectFuncs(baseType reflect.Type, mappings map[string]string, parser *parser) template.FuncMap {
	return template.FuncMap{
		"qualifiedName": func(in reflect.Type) string {
			return resolve(in, mappings)
		},
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
		"imports": func() map[string]string {
			out := make(map[string]string)
			for i, v := range parser.packages[baseType.PkgPath()].Imports {
				if v, ok := mappings[i]; ok {
					out["github.com/eddycharly/terraform-provider-kops/pkg/schemas/"+v] = v + "schemas"
				}
				for i2 := range parser.packages[v.PkgPath].Imports {
					if v, ok := mappings[i2]; ok {
						out["github.com/eddycharly/terraform-provider-kops/pkg/schemas/"+v] = v + "schemas"
					}
				}
			}
			return out
		},
		"mapping": func(t reflect.Type) string {
			if baseType.PkgPath() == t.PkgPath() {
				return ""
			}
			return mappings[t.PkgPath()] + "schemas."
		},
	}
}

func optionFuncs(dataSource bool, optionsMap map[reflect.Type]*options, parser *parser) template.FuncMap {
	isRequired := func(in _field) bool {
		req := optionsMap[in.Owner].required.Has(in.Name)
		if req {
			return req
		}
		// attempt to detect invalid config
		if !dataSource {
			pkg := in.Owner.PkgPath()
			str := in.Owner.Name()
			fld := in.Name
			if parser.packs[pkg][str].defaults[fld] == "true" {
				log.Printf("WARNING: %s/%s/%s is not marked required but it seems to have a non zero value default", pkg, str, fld)
			}
		}
		return false
	}
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
		"isRequired": isRequired,
		"isOptional": func(in _field) bool {
			if dataSource {
				return optionsMap[in.Owner].computed.Has(in.Name)
			}
			return !isRequired(in) && !optionsMap[in.Owner].computedOnly.Has(in.Name)
		},
		"isNullable": func(in _field) bool {
			return optionsMap[in.Owner].nullable.Has(in.Name)
		},
		"isComputed": func(in _field) bool {
			if isRequired(in) {
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

func schemaFuncs(scope string) template.FuncMap {
	return template.FuncMap{
		"scope": func() string {
			return scope
		},
		"schemaType": schemaType,
		"schemaName": func(in string) string {
			return fieldName(in)
		},
	}
}

func docFuncs(header, footer string, parser *parser, optionsMap map[reflect.Type]*options) template.FuncMap {
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
		"header": func() string {
			return header
		},
		"footer": func() string {
			return footer
		},
		"code": func(in string) string {
			return fmt.Sprintf("`%s`", in)
		},
	}
}
