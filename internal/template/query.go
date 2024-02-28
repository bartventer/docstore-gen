package template

// DefaultQuery default query
const DefaultQuery = `
var (
	// Q default query
	Q =new(Query)
	{{range $name,$d :=.Data -}}
	// {{$d.ModelStructName}} default query
	{{$d.ModelStructName}} *{{$d.QueryStructName}}
	{{end -}}
)

// Initialize initialize docstore query
func Initialize() {
	*Q = *initialize()
	{{range $name,$d :=.Data -}}
	{{$d.ModelStructName}} = &Q.{{$d.ModelStructName}}
	{{end -}}
}

`

// QueryMethod query method template
const QueryMethod = `
// initialize query
func initialize() *Query {
	return &Query{
		{{range $name,$d :=.Data -}}
		{{$d.ModelStructName}}: new{{$d.ModelStructName}}(),
		{{end -}}
	}
}

// Query query struct
type Query struct{
	{{range $name,$d :=.Data -}}
	{{$d.ModelStructName}} {{$d.QueryStructName}}
	{{end}}
}

`
