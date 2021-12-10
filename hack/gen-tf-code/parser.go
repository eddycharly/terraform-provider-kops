package main

import (
	"go/ast"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

type _field struct {
	reflect.StructField
	Owner   reflect.Type
	Default string
}

func (f *_field) QualifiedName() string {
	t := refresh(f.Type)
	return f.Type.String() + "/*" + t.PkgPath() + "*/"
}

type _nested struct {
	pack       string
	structName string
}

type _struct struct {
	doc      string
	fields   map[string]string
	defaults map[string]string
	nested   []_nested
}

func (s _struct) lookup(in string, c map[string]map[string]_struct) string {
	if s, ok := s.fields[in]; ok && s != "" {
		return s
	}
	for _, n := range s.nested {
		if p, ok := c[n.pack]; ok {
			if s, ok := p[n.structName]; ok {
				out := s.lookup(in, c)
				if out != "" {
					return out
				}
			}
		}
	}
	return ""
}

type parser struct {
	packages map[string]*packages.Package
	packs    map[string]map[string]_struct
}

func (p *parser) parseStruct(pack *packages.Package, typeSpec *ast.TypeSpec, doc string, file *ast.File) _struct {
	ret := _struct{
		doc:      doc,
		fields:   make(map[string]string),
		defaults: make(map[string]string),
	}
	structure, ok := typeSpec.Type.(*ast.StructType)
	if ok {
		importsLookup := buildImportsLookup(pack, file)
		for _, field := range structure.Fields.List {
			if len(field.Names) == 0 {
				switch t := field.Type.(type) {
				case *ast.Ident:
					ret.nested = append(ret.nested, _nested{
						pack:       pack.ID,
						structName: t.Name,
					})
				case *ast.SelectorExpr:
					x := t.X.(*ast.Ident)
					sel := t.Sel
					from := importsLookup[x.Name]
					ret.nested = append(ret.nested, _nested{
						pack:       from.ID,
						structName: sel.Name,
					})
				}
			} else {
				for _, name := range field.Names {
					ret.fields[name.Name] = field.Doc.Text()
					if field.Doc != nil {
						for _, line := range field.Doc.List {
							if strings.HasPrefix(line.Text, "// Default:") {
								ret.defaults[name.Name] = strings.TrimSpace(strings.ReplaceAll(line.Text, "// Default:", ""))
							}
						}
					}
				}
			}
		}
	}
	return ret
}

func (p *parser) parsePackage(pack *packages.Package) {
	p.packages[pack.PkgPath] = pack
	if _, ok := p.packs[pack.ID]; !ok {
		p.packs[pack.ID] = make(map[string]_struct)
		for _, v := range pack.Imports {
			p.parsePackage(v)
		}
		for _, node := range pack.Syntax {
			for _, decl := range node.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if ok {
					for _, spec := range genDecl.Specs {
						typeSpec, ok := spec.(*ast.TypeSpec)
						if ok {
							p.packs[pack.ID][typeSpec.Name.Name] = p.parseStruct(pack, typeSpec, genDecl.Doc.Text(), node)
						}
					}
				}
			}
		}
	}
}

func buildImportsLookup(pack *packages.Package, file *ast.File) map[string]*packages.Package {
	out := make(map[string]*packages.Package)
	for _, i := range file.Imports {
		path := strings.ReplaceAll(i.Path.Value, "\"", "")
		if i.Name != nil {
			out[i.Name.Name] = pack.Imports[path]
		} else {
			out[pack.Imports[path].Name] = pack.Imports[path]
		}
	}
	return out
}

func initParser(patterns ...string) (*parser, error) {
	config := packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	packs, err := packages.Load(&config, patterns...)
	if err != nil {
		return nil, err
	}
	parser := parser{
		packages: make(map[string]*packages.Package),
		packs:    make(map[string]map[string]_struct),
	}
	for _, pack := range packs {
		parser.parsePackage(pack)
	}
	return &parser, nil
}
