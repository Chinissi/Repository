package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"

	"github.com/hamba/avro"
	"github.com/iancoleman/strcase"
)

type Conf struct {
	PackageName string
}

const outputTemplate = `package {{ .PackageName }}

{{ if len .Imports }}
import (
    {{- range .Imports }}
		"{{ . }}"
	{{- end }}
)
{{ end }}

{{- range .Typedefs }}
type {{ .Name }} struct {
	{{- range .Fields }}
		{{ .Name }} {{ .Type }} {{ .Tag }}
	{{- end }}
}
{{ end }}`

type data struct {
	PackageName string
	Imports     []string
	Typedefs    []typedef
}

type typedef struct {
	Name   string
	Fields []field
}

type field struct {
	Name string
	Type string
	Tag  string
}

var primitiveMappings = map[avro.Type]string{
	"string":  "string",
	"bytes":   "[]byte",
	"int":     "int",
	"long":    "int64",
	"float":   "float32",
	"double":  "float64",
	"boolean": "bool",
}

func Struct(s string, dst io.Writer, gc Conf) error {
	schema, err := avro.Parse(s)
	if err != nil {
		return err
	}

	rSchema, ok := schema.(*avro.RecordSchema)
	if !ok {
		return errors.New("can only generate Go code from Record Schemas")
	}

	td := data{PackageName: strcase.ToSnake(gc.PackageName)}
	_ = generateFrom(rSchema, &td)

	buf := &bytes.Buffer{}
	if err = writeCode(buf, &td); err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed formatting. %w", err)
	}

	_, err = dst.Write(formatted)
	return err
}

func generateFrom(schema avro.Schema, acc *data) string {
	switch t := schema.(type) {
	case *avro.RecordSchema:
		typeName := strcase.ToCamel(t.Name())
		fields := make([]field, len(t.Fields()))
		for i, f := range t.Fields() {
			fSchema := f.Type()
			fieldName := strcase.ToCamel(f.Name())
			typ := resolveType(fSchema, f.Prop("logicalType"), acc)
			tag := f.Name()
			fields[i] = newField(fieldName, typ, tag)
		}
		acc.Typedefs = append(acc.Typedefs, newType(typeName, fields))
		return typeName
	default:
		return resolveType(schema, nil, acc)
	}
}

func resolveType(fieldSchema avro.Schema, logicalType interface{}, acc *data) string {
	var typ string
	switch s := fieldSchema.(type) {
	case *avro.RefSchema:
		typ = resolveRefSchema(s)
	case *avro.RecordSchema:
		typ = generateFrom(s, acc)
	case *avro.PrimitiveSchema:
		typ = resolvePrimitiveLogicalType(logicalType, typ, s)
		if strings.Contains(typ, "time") {
			addImport(acc, "time")
		}
		if strings.Contains(typ, "big") {
			addImport(acc, "math/big")
		}
	case *avro.ArraySchema:
		typ = fmt.Sprintf("[]%s", generateFrom(s.Items(), acc))
	case *avro.EnumSchema:
		typ = "string"
	case *avro.FixedSchema:
		typ = fmt.Sprintf("[%d]byte", +s.Size())
	case *avro.MapSchema:
		typ = "map[string]" + resolveType(s.Values(), nil, acc)
	case *avro.UnionSchema:
		typ = resolveUnionTypes(s, acc)
	}
	return typ
}

func resolveRefSchema(s *avro.RefSchema) string {
	typ := ""
	switch sx := s.Schema().(type) {
	case *avro.RecordSchema:
		typ = sx.Name()
	case avro.NamedSchema:
		typ = sx.Name()
	}
	return strcase.ToCamel(typ)
}

func resolveUnionTypes(unionSchema *avro.UnionSchema, acc *data) string {
	nullIsAllowed := false
	typesInUnion := make([]string, 0)
	for _, elementSchema := range unionSchema.Types() {
		if _, ok := elementSchema.(*avro.NullSchema); ok {
			nullIsAllowed = true
		} else {
			typesInUnion = append(typesInUnion, generateFrom(elementSchema, acc))
		}
	}
	if nullIsAllowed && len(typesInUnion) == 1 {
		typ := typesInUnion[0]
		if strings.HasPrefix(typ, "[]") {
			return typ
		}
		return "*" + typ
	}
	return "interface{}"
}

func resolvePrimitiveLogicalType(logicalType interface{}, typ string, s avro.Schema) string {
	switch logicalType {
	case "", nil:
		typ = primitiveMappings[s.Type()]
	case "date", "timestamp-millis", "timestamp-micros":
		typ = "time.Time"
	case "time-millis", "time-micros":
		typ = "time.Duration"
	case "decimal":
		typ = "*big.Rat"
	}
	return typ
}

func newType(name string, fields []field) typedef {
	return typedef{
		Name:   name,
		Fields: fields,
	}
}

func newField(name string, typ string, tag string) field {
	return field{
		Name: name,
		Type: typ,
		Tag:  "`avro:\"" + tag + "\"`",
	}
}

func addImport(acc *data, statement string) {
	for _, k := range acc.Imports {
		if k == statement {
			return
		}
	}
	acc.Imports = append(acc.Imports, statement)
}

func writeCode(w io.Writer, data *data) error {
	parsed, err := template.New("out").Parse(outputTemplate)
	if err != nil {
		return err
	}

	return parsed.Execute(w, data)
}
