package docstoregen

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/bartventer/docstore-gen/internal/generate"
	tmpl "github.com/bartventer/docstore-gen/internal/template"
	"github.com/bartventer/docstore-gen/internal/utils/pools"
	"golang.org/x/tools/imports"
)

var concurrent = runtime.NumCPU()

// ApplyInterface specifies .diy_method interfaces on structures, implment codes will be generated after calling g.Execute()
// eg: g.ApplyInterface(func(model.Method){}, model.User{}, model.Company{})
func (g *Generator) ApplyInterface(models ...interface{}) {
	structs, err := generate.ConvertStructs(models...)
	if err != nil {
		g.Logger.Error("check struct fail: %v", err)
		panic("check struct fail")
	}
	g.apply(structs)
}

func (g *Generator) apply(structs []*generate.QueryStructMeta) {

	for _, interfaceStructMeta := range structs {
		_, err := g.pushQueryStructMeta(interfaceStructMeta)
		if err != nil {
			g.Logger.Error("gen struct fail: %v", err)
			panic("gen struct fail")
		}
	}
}

func (g *Generator) pushQueryStructMeta(meta *generate.QueryStructMeta) (*genInfo, error) {
	structName := meta.ModelStructName
	if g.Data[structName] == nil {
		g.Data[structName] = &genInfo{QueryStructMeta: meta}
	}
	if g.Data[structName].Source != meta.Source {
		return nil, fmt.Errorf("cannot generate struct with the same name from different source:%s.%s and %s.%s",
			meta.StructInfo.Package, meta.ModelStructName, g.Data[structName].StructInfo.Package, g.Data[structName].ModelStructName)
	}
	return g.Data[structName], nil
}

// genInfo info about generated code
type genInfo struct {
	*generate.QueryStructMeta
}

// NewGenerator creates a new instance of Generator with the given configuration.
// It revises the configuration and panics if there is an error.
// If the LoggerHandler is not provided in the configuration, it uses a default JSON handler that writes to os.Stdout.
// Returns a pointer to the created Generator.
func NewGenerator(cfg Config) *Generator {
	if err := cfg.Revise(); err != nil {
		panic(fmt.Errorf("create generator fail: %w", err))
	}

	if cfg.LoggerHandler == nil {
		cfg.LoggerHandler = slog.NewJSONHandler(os.Stdout, nil)
	}

	return &Generator{
		Config: cfg,
		Data:   make(map[string]*genInfo),
		Logger: slog.New(cfg.LoggerHandler),
	}
}

// Revise revises the configuration by setting the absolute path for the output directory,
// generating a default output path if not provided, and setting the output file path.
// It also sets the query package name based on the output path.
// Returns an error if the outpath is invalid.
func (cfg *Config) Revise() (err error) {

	cfg.OutPath, err = filepath.Abs(cfg.OutPath)
	if err != nil {
		return fmt.Errorf("outpath is invalid: %w", err)
	}
	if cfg.OutPath == "" {
		cfg.OutPath = fmt.Sprintf(".%squery%s", string(os.PathSeparator), string(os.PathSeparator))
	}
	if cfg.OutFile == "" {
		cfg.OutFile = filepath.Join(cfg.OutPath, "gen.go")
	} else if !strings.Contains(cfg.OutFile, string(os.PathSeparator)) {
		cfg.OutFile = filepath.Join(cfg.OutPath, cfg.OutFile)
	}
	cfg.queryPkgName = filepath.Base(cfg.OutPath)

	return nil
}

// Generator generate code
type Generator struct {
	Config

	Logger *slog.Logger // logger

	Data map[string]*genInfo //gen query data
}

// Execute generates code based on the provided specifications.
// It first generates the query file and then logs the progress.
// If any error occurs during the generation process, it logs the error and panics.
// Finally, it logs the completion of code generation.
func (g *Generator) Execute() {
	g.Logger.Info("Start generating code.")

	if err := g.generateQueryFile(); err != nil {
		g.Logger.Error("generate query code fail: %s", err)
		panic("generate query code fail")
	}

	g.Logger.Info("Generate code done.")
}

// generateQueryFile generate query code and save to file
func (g *Generator) generateQueryFile() (err error) {
	if len(g.Data) == 0 {
		return nil
	}
	if err = os.MkdirAll(g.OutPath, os.ModePerm); err != nil {
		return fmt.Errorf("make dir outpath(%s) fail: %s", g.OutPath, err)
	}

	errChan := make(chan error)
	pool := pools.NewPool(concurrent)
	// generate query code for all struct
	for _, info := range g.Data {
		pool.Wait()
		go func(info *genInfo) {
			defer pool.Done()
			err := g.generateSingleQueryFile(info)
			if err != nil {
				errChan <- err
			}

		}(info)
	}
	select {
	case err = <-errChan:
		return err
	case <-pool.AsyncWaitAll():
	}

	// generate query file
	var buf bytes.Buffer
	err = render(tmpl.Header, &buf, map[string]interface{}{
		"Package":        g.queryPkgName,
		"ImportPkgPaths": importList.Add(g.importPkgPaths...).Paths(),
	})
	if err != nil {
		return err
	}

	err = render(tmpl.DefaultQuery, &buf, g)
	if err != nil {
		return err
	}

	err = render(tmpl.QueryMethod, &buf, g)
	if err != nil {
		return err
	}

	err = g.output(g.OutFile, buf.Bytes())
	if err != nil {
		return err
	}
	g.Logger.Info("generate query file: %s", g.OutFile, "")

	return nil
}

// render is a function that renders a template string using the provided data and writes the result to the given io.Writer.
// It returns an error if there was a problem parsing or executing the template.
func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

// generateSingleQueryFile generate query code and save to file
func (g *Generator) generateSingleQueryFile(data *genInfo) (err error) {
	var buf bytes.Buffer

	structPkgPath := data.StructInfo.PkgPath
	err = render(tmpl.Header, &buf, map[string]interface{}{
		"Package":        g.queryPkgName,
		"ImportPkgPaths": importList.Add(structPkgPath).Add(getImportPkgPaths(data)...).Paths(),
	})
	if err != nil {
		return err
	}

	structTmpl := tmpl.TableQueryStruct
	err = render(structTmpl, &buf, data.QueryStructMeta)
	if err != nil {
		return err
	}

	defer g.Logger.Info("generate query file: %s%s%s.gen.go", g.OutPath, string(os.PathSeparator), data.FileName, "")
	return g.output(fmt.Sprintf("%s%s%s.gen.go", g.OutPath, string(os.PathSeparator), data.TableName), buf.Bytes())
}

// getImportPkgPaths returns a slice of import package paths from the given genInfo data.
// It removes duplicate import paths and returns them in the order they were encountered.
func getImportPkgPaths(data *genInfo) []string {
	importPathMap := make(map[string]struct{})
	for _, path := range data.ImportPkgPaths {
		importPathMap[path] = struct{}{}
	}
	importPkgPaths := make([]string, 0, len(importPathMap))
	for importPath := range importPathMap {
		importPkgPaths = append(importPkgPaths, importPath)
	}
	return importPkgPaths
}

// output formats and writes the content to a file specified by fileName.
// It uses the imports.Process function to process the content and ensure proper formatting.
// If an error occurs during the formatting process, it logs the error and prints the surrounding lines of code.
// The function returns an error if it fails to format the file or write the content to the file.
func (g *Generator) output(fileName string, content []byte) error {
	result, err := imports.Process(fileName, content, nil)
	if err != nil {
		lines := strings.Split(string(content), "\n")
		errLine, _ := strconv.Atoi(strings.Split(err.Error(), ":")[1])
		startLine, endLine := errLine-5, errLine+5
		g.Logger.Error("Format fail: %s", err)
		if startLine < 0 {
			startLine = 0
		}
		if endLine > len(lines)-1 {
			endLine = len(lines) - 1
		}
		for i := startLine; i <= endLine; i++ {
			g.Logger.Error("%d: %s", strconv.Itoa(i), lines[i])
		}
		return fmt.Errorf("cannot format file: %w", err)
	}
	return os.WriteFile(fileName, result, 0640)
}
