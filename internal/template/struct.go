package template

const (
	// TableQueryStruct table query struct
	TableQueryStruct = createMethod + `
	type {{.QueryStructName}} struct {
		{{.QueryStructName}}CollectionDo
		` + fields + `
	}
	` + modelMethod + entities + collectionMethod + actionListMethod + queryMethod + documentIteratorMethod
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

	modelMethod = `
	// TableName get table name
	func ({{.S}} {{.QueryStructName}}) TableName() string {
		return "{{.TableName}}"
	}

	// WithCollection use new collection.
	// This method is used to set the collection for the query struct and 
	// is typically used right after initializing the query struct.
	func ({{.S}} *{{.QueryStructName}}) WithCollection(coll *docstore.Collection) *{{.QueryStructName}} {
		{{.S}}.{{.QueryStructName}}CollectionDo = {{.QueryStructName}}CollectionDo{coll}
		return {{.S}}
	}
	`

	entities = `
	type (
		// {{.QueryStructName}}CollectionDo collection data object.
		// Embeds the [docstore.Collection] type.
		{{.QueryStructName}}CollectionDo struct {*docstore.Collection}
		
		// {{.QueryStructName}}ActionListDo action list data object.
		// Embeds the [docstore.ActionList] type.
		{{.QueryStructName}}ActionListDo struct {*docstore.ActionList}

		// {{.QueryStructName}}QueryDo query data object.
		// Embeds the [docstore.Query] type.
		{{.QueryStructName}}QueryDo struct {*docstore.Query}

		// {{.QueryStructName}}DocumentIteratorDo document iterator data object.
		// Embeds the [docstore.DocumentIterator] type.
		{{.QueryStructName}}DocumentIteratorDo struct {*docstore.DocumentIterator}
	)
	`

	collectionMethod = `
	// ================================= CollectionDO =============================

	{{ $structPackageType := printf "%s.%s" .StructInfo.Package .ModelStructName -}}
	// Actions get action list
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Actions() *{{.QueryStructName}}ActionListDo {
		return &{{.QueryStructName}}ActionListDo{ {{.S}}.Collection.Actions() }
	}

	// Query get query
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Query() *{{.QueryStructName}}QueryDo {
		return &{{.QueryStructName}}QueryDo{ {{.S}}.Collection.Query() }
	}

	// Create create new {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Create(ctx context.Context, doc *{{$structPackageType}}) error {
		return {{.S}}.Collection.Create(ctx, doc)
	}

	// Delete delete {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Delete(ctx context.Context, doc *{{$structPackageType}}) error {
		return {{.S}}.Collection.Delete(ctx, doc)
	}

	// Get get {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Get(ctx context.Context, doc *{{$structPackageType}}, fps ...docstore.FieldPath) error {
		return {{.S}}.Collection.Get(ctx, doc, fps...)
	}

	// Put put {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Put(ctx context.Context, doc *{{$structPackageType}}) error {
		return {{.S}}.Collection.Put(ctx, doc)
	}

	// Replace replace {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Replace(ctx context.Context, doc *{{$structPackageType}}) error {
		return {{.S}}.Collection.Replace(ctx, doc)
	}

	// Update update {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}CollectionDo) Update(ctx context.Context, doc *{{$structPackageType}}, mods ...field.Mod) error {
		return {{.S}}.Collection.Update(ctx, doc, field.ConvertMods(mods))
	}
	`

	actionListMethod = `
	// ================================= ActionListDO =============================
	
	{{ $structPackageType := printf "%s.%s" .StructInfo.Package .ModelStructName -}}
	// Create create new {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}ActionListDo) Create(doc *{{$structPackageType}}) *{{.QueryStructName}}ActionListDo {
		{{.S}}.ActionList = {{.S}}.ActionList.Create(doc)
		return {{.S}}
	}

	// Delete delete {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}ActionListDo) Delete(doc *{{$structPackageType}}) *{{.QueryStructName}}ActionListDo {
		{{.S}}.ActionList = {{.S}}.ActionList.Delete(doc)
		return {{.S}}
	}

	// Get get {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}ActionListDo) Get(doc *{{$structPackageType}}, fps ...docstore.FieldPath) *{{.QueryStructName}}ActionListDo {
		{{.S}}.ActionList = {{.S}}.ActionList.Get(doc, fps...)
		return {{.S}}
	}

	// Put put {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}ActionListDo) Put(doc *{{$structPackageType}}) *{{.QueryStructName}}ActionListDo {
		{{.S}}.ActionList = {{.S}}.ActionList.Put(doc)
		return {{.S}}
	}

	// Replace replace {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}ActionListDo) Replace(doc *{{$structPackageType}}) *{{.QueryStructName}}ActionListDo {
		{{.S}}.ActionList = {{.S}}.ActionList.Replace(doc)
		return {{.S}}
	}

	// Update update {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}ActionListDo) Update(doc *{{$structPackageType}}, mods ...field.Mod) *{{.QueryStructName}}ActionListDo {
		{{.S}}.ActionList = {{.S}}.ActionList.Update(doc, field.ConvertMods(mods))
		return {{.S}}
	}

	// BeforeDo before do
	func ({{.S}} *{{.QueryStructName}}ActionListDo) BeforeDo(f func(asFunc func(interface{}) bool) error) *{{.QueryStructName}}ActionListDo {
		{{.S}}.ActionList = {{.S}}.ActionList.BeforeDo(f)
		return {{.S}}
	}
	`

	queryMethod = `
	// =================================== QueryDO ================================

	// Get get {{.ModelStructName}}
	func ({{.S}} *{{.QueryStructName}}QueryDo) Get(ctx context.Context, fps ...docstore.FieldPath) *{{.QueryStructName}}DocumentIteratorDo {
		return &{{.QueryStructName}}DocumentIteratorDo{ {{.S}}.Query.Get(ctx, fps...) }
	}

	// BeforeQuery before query
	func ({{.S}} *{{.QueryStructName}}QueryDo) BeforeQuery(f func(asFunc func(interface{}) bool) error) *{{.QueryStructName}}QueryDo {
		{{.S}}.Query = {{.S}}.Query.BeforeQuery(f)
		return {{.S}}
	}

	// Limit limit
	func ({{.S}} *{{.QueryStructName}}QueryDo) Limit(n int) *{{.QueryStructName}}QueryDo {
		{{.S}}.Query = {{.S}}.Query.Limit(n)
		return {{.S}}
	}

	// Offset offset
	func ({{.S}} *{{.QueryStructName}}QueryDo) Offset(n int) *{{.QueryStructName}}QueryDo {
		{{.S}}.Query = {{.S}}.Query.Offset(n)
		return {{.S}}
	}

	// OrderBy order by
	func ({{.S}} *{{.QueryStructName}}QueryDo) OrderBy(orderBys ...field.OrderByExpression) *{{.QueryStructName}}QueryDo {
		for _, orderBy := range orderBys {
			field, direction := orderBy.BuildOrderBy()
			{{.S}}.Query = {{.S}}.Query.OrderBy(field, direction)
		}
		return {{.S}}
	}

	// Where where
	func ({{.S}} *{{.QueryStructName}}QueryDo) Where(conds ...field.Expr) *{{.QueryStructName}}QueryDo {
		for _, cond := range conds {
			fieldPath, op, value := cond.Build()
			{{.S}}.Query = {{.S}}.Query.Where(fieldPath, op, value)
		}
		return {{.S}}
	}
	`

	documentIteratorMethod = `
	// ============================== DocumentIteratorDO ==========================
	
	{{ $structPackageType := printf "%s.%s" .StructInfo.Package .ModelStructName -}}
	// Next next
	func ({{.S}} *{{.QueryStructName}}DocumentIteratorDo) Next(ctx context.Context, doc *{{$structPackageType}}) error {
		return {{.S}}.DocumentIterator.Next(ctx, doc)
	}
	`
)
