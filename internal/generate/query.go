package generate

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/bartventer/docstore-gen/field"
	"github.com/bartventer/docstore-gen/internal/model"
	"github.com/bartventer/docstore-gen/internal/parser"
)

// QueryStructMeta struct info in generated code
type QueryStructMeta struct {
	Generated       bool   // whether to generate db model
	FileName        string // generated file name
	S               string // the first letter(lower case)of simple Name (receiver)
	QueryStructName string // internal query struct name
	ModelStructName string // origin/model struct name
	TableName       string // table name in db server
	StructInfo      parser.Param
	Fields          []*model.Field
	Source          model.SourceCode
	ImportPkgPaths  []string

	interfaceMode bool
}

func getDocstoreTag(reflectTag reflect.StructTag) field.DocstoreTag {
	tag := reflectTag.Get("docstore")
	if tag == "" {
		return nil
	}
	return strings.Split(tag, ",")
}

// parseStruct get all elements of struct with gorm's Parse, ignore unexported elements
func (b *QueryStructMeta) parseStruct(st interface{}) error {

	// using TableName method get table name
	if t, ok := st.(interface{ TableName() string }); ok {
		b.TableName = t.TableName()
	} else {
		return errors.New("not found TableName method")
	}

	// using reflect get all elements of struct
	v := reflect.ValueOf(st)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if f.PkgPath != "" {
			continue
		}
		docstoreTag := getDocstoreTag(f.Tag)
		if len(docstoreTag) == 0 || docstoreTag[0] == "-" || (len(docstoreTag) == 1 && docstoreTag[0] == "omitempty") {
			continue
		}
		b.appendOrUpdateField(&model.Field{
			Name:        f.Name,
			Type:        b.getFieldRealType(f.Type),
			ColumnName:  docstoreTag[0],
			DocstoreTag: getDocstoreTag(f.Tag),
		})
	}

	return nil
}

// ScanValuer interface for Field
type ScanValuer interface {
	Scan(src interface{}) error   // sql.Scanner
	Value() (driver.Value, error) // driver.Valuer
}

// getFieldRealType  get basic type of field
func (b *QueryStructMeta) getFieldRealType(f reflect.Type) string {
	scanValuer := reflect.TypeOf((*ScanValuer)(nil)).Elem()
	if f.Implements(scanValuer) || reflect.New(f).Type().Implements(scanValuer) {
		return "field"
	}

	if f.Kind() == reflect.Ptr {
		f = f.Elem()
	}
	if f.String() == "time.Time" {
		return "time.Time"
	}
	if f.String() == "[]uint8" || f.String() == "json.RawMessage" {
		return "bytes"
	}
	return f.Kind().String()
}

// check field if in BaseStruct update else append
func (b *QueryStructMeta) appendOrUpdateField(f *model.Field) {
	for i, m := range b.Fields {
		if m.Name == f.Name {
			b.Fields[i] = f
			return
		}
	}
	b.appendField(f)
}

func (b *QueryStructMeta) appendField(f *model.Field) { b.Fields = append(b.Fields, f) }

// check if struct is exportable and if struct in main package and if field's type is regular
func (b *QueryStructMeta) check() (err error) {
	if b.StructInfo.InMainPkg() {
		return fmt.Errorf("can't generated data object for struct in main package, ignore:%s", b.ModelStructName)
	}
	if !isCapitalize(b.ModelStructName) {
		return fmt.Errorf("can't generated data object for non-exportable struct, ignore:%s", b.QueryStructName)
	}
	return nil
}
