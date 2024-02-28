package parser

// Param parameters in method
type Param struct { // (user model.User)
	PkgPath   string // package's path: internal/model
	Package   string // package's name: model
	Name      string // param's name: user
	Type      string // param's type: User
	IsArray   bool   // is array or not
	IsPointer bool   // is pointer or not
}

// InMainPkg ...
func (p *Param) InMainPkg() bool {
	return p.Package == "main"
}
