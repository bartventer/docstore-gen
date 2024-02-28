package template

const (
	// TableQueryStruct table query struct
	TableQueryStruct = createMethod + `
	type {{.QueryStructName}} struct {
		` + fields + `
	}
	` + tableMethod
)

const (
	createMethod = `
	// new{{.ModelStructName}} create new {{.ModelStructName}} query struct
    func new{{.ModelStructName}}() {{.QueryStructName}} {
        {{ $queryStructName := .QueryStructName -}}
        _{{$queryStructName}} := {{$queryStructName}}{}
        
		tableName := _{{$queryStructName}}.TableName()
		
        {{range .Fields -}}
		{{- if .ColumnName -}}_{{$queryStructName}}.{{.Name}} = field.New{{.GenType}}(tableName, "{{.ColumnName}}");{{- end -}}
        {{end}}

        return _{{$queryStructName}}
    }
    `

	fields = `
	{{range .Fields -}}
		{{- if .Name -}}{{.Name}} field.{{.GenType}};
		{{- end -}}
	{{end}}
	`

	tableMethod = `
	// TableName get table name
	func (q {{.QueryStructName}}) TableName() string {
		return "{{.TableName}}"
	}
	`
)
