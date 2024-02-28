package generate

import (
	"fmt"
	"reflect"

	"github.com/bartventer/docstore-gen/internal/model"
	"github.com/bartventer/docstore-gen/internal/parser"
)

// ConvertStructs convert to base structures
func ConvertStructs(structs ...interface{}) (metas []*QueryStructMeta, err error) {
	for _, st := range structs {
		if isNil(st) {
			continue
		}
		if base, ok := st.(*QueryStructMeta); ok {
			metas = append(metas, base)
			continue
		}
		if !isStructType(reflect.ValueOf(st)) {
			return nil, fmt.Errorf("%s is not a struct", reflect.TypeOf(st).String())
		}

		structType := reflect.TypeOf(st)
		name := getStructName(structType.String())
		newStructName := name
		if st, ok := st.(interface{ GenInternalDoName() string }); ok {
			newStructName = st.GenInternalDoName()
		}

		meta := &QueryStructMeta{
			S:               getPureName(name),
			ModelStructName: name,
			QueryStructName: uncaptialize(newStructName),
			StructInfo:      parser.Param{PkgPath: structType.PkgPath(), Type: name, Package: getPackageName(structType.String())},
			Source:          model.Struct,
		}
		if err := meta.parseStruct(st); err != nil {
			return nil, fmt.Errorf("transform struct [%s.%s] error:%s", meta.StructInfo.Package, name, err)
		}
		if err := meta.check(); err != nil {
			fmt.Println(err.Error())
			continue
		}

		metas = append(metas, meta)
	}
	return
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	// if v is not ptr, return false(i is not nil)
	// if v is ptr, return v.IsNil()
	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Ptr && v.IsNil()
}

func isStructType(data reflect.Value) bool {
	return data.Kind() == reflect.Struct ||
		(data.Kind() == reflect.Ptr && data.Elem().Kind() == reflect.Struct)
}
