package docstoregen

import "log/slog"

// Config represents the configuration options for generating query code.
type Config struct {
	OutPath        string       // OutPath specifies the path where the query code will be generated. Example: "/path/to/project/pkg/query"
	OutFile        string       // OutFile specifies the name of the query code file. The default value is "gen.go".
	queryPkgName   string       // queryPkgName specifies the name of the package where the query code will be generated.
	importPkgPaths []string     // importPkgPaths specifies the import paths of the packages that the generated query code will import.
	LoggerHandler  slog.Handler // LoggerHandler specifies the handler for the logger used by the generator. Example: slog.NewJSONHandler(os.Stdout, nil)
}
