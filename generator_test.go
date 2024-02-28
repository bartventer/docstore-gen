package docstoregen

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/bartventer/docstore-gen/internal/generate"
	"github.com/bartventer/docstore-gen/internal/parser"
)

type User struct {
	Name string `docstore:"name"`
}

func (User) TableName() string {
	return "users"
}

func TestGenerator_ApplyInterface(t *testing.T) {
	type args struct {
		models []interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name:      "invalid: struct has no TableName method",
			args:      args{models: []interface{}{struct{}{}}},
			wantPanic: true,
		},
		{
			name: "valid",
			args: args{models: []interface{}{&User{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("Generator.ApplyInterface() recover = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()
			g, cleanup, err := setupGenerator(nil)
			if err != nil {
				t.Errorf("setupGenerator() error = %v", err)
			}
			t.Cleanup(cleanup)
			g.ApplyInterface(tt.args.models...)
		})
	}
}

func TestConfig_Revise(t *testing.T) {
	type fields struct {
		OutPath        string
		OutFile        string
		queryPkgName   string
		importPkgPaths []string
		LoggerHandler  slog.Handler
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid OutPath and OutFile",
			fields: fields{
				OutPath: "testpath",
				OutFile: "testfile.go",
			},
			wantErr: false,
		},
		{
			name: "Empty OutPath and OutFile",
			fields: fields{
				OutPath: "",
				OutFile: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				OutPath:        tt.fields.OutPath,
				OutFile:        tt.fields.OutFile,
				queryPkgName:   tt.fields.queryPkgName,
				importPkgPaths: tt.fields.importPkgPaths,
				LoggerHandler:  tt.fields.LoggerHandler,
			}
			if err := cfg.Revise(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Revise() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_generateQueryFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "No data",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, cleanup, err := setupGenerator(nil)
			if err != nil {
				t.Errorf("setupGenerator() error = %v", err)
			}
			t.Cleanup(cleanup)
			g.Data = map[string]*genInfo{
				"test": {
					QueryStructMeta: &generate.QueryStructMeta{
						Generated:       false,
						FileName:        "",
						S:               "t",
						QueryStructName: "TestQuery",
						ModelStructName: "TestModel",
						TableName:       "test_table",
						StructInfo:      parser.Param{PkgPath: "", Package: "", Name: "", Type: "", IsArray: false, IsPointer: false},
						Source:          0,
						ImportPkgPaths:  []string{"fmt", "os"},
					},
				},
			}
			if err := g.generateQueryFile(); (err != nil) != tt.wantErr {
				t.Errorf("Generator.generateQueryFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_render(t *testing.T) {
	type args struct {
		tmpl string
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantWr  string
		wantErr bool
	}{
		{
			name: "Valid template with no data",
			args: args{
				tmpl: "Hello, World!",
				data: nil,
			},
			wantWr:  "Hello, World!",
			wantErr: false,
		},
		{
			name: "Valid template with data",
			args: args{
				tmpl: "Hello, {{.Name}}!",
				data: map[string]string{"Name": "John"},
			},
			wantWr:  "Hello, John!",
			wantErr: false,
		},
		{
			name: "Invalid template",
			args: args{
				tmpl: "Hello, {{.Name!",
				data: map[string]string{"Name": "John"},
			},
			wantWr:  "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &bytes.Buffer{}
			if err := render(tt.args.tmpl, wr, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWr := wr.String(); gotWr != tt.wantWr {
				t.Errorf("render() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}

func TestGenerator_generateSingleQueryFile(t *testing.T) {
	type args struct {
		data *genInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid data",
			args: args{data: &genInfo{
				QueryStructMeta: &generate.QueryStructMeta{
					Generated:       false,
					FileName:        "",
					S:               "",
					QueryStructName: "TestQuery",
					ModelStructName: "",
					TableName:       "test_table",
					StructInfo:      parser.Param{PkgPath: "", Package: "", Name: "", Type: "", IsArray: false, IsPointer: false},
					Source:          0,
					ImportPkgPaths:  []string{"fmt", "os"},
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, cleanup, err := setupGenerator(nil)
			if err != nil {
				t.Errorf("setupGenerator() error = %v", err)
			}
			t.Cleanup(cleanup)
			if err := g.generateSingleQueryFile(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Generator.generateSingleQueryFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getImportPkgPaths(t *testing.T) {
	type args struct {
		data *genInfo
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "No imports",
			args: args{
				data: &genInfo{
					&generate.QueryStructMeta{ImportPkgPaths: []string{}},
				},
			},
			want: []string{},
		},
		{
			name: "Single import",
			args: args{
				data: &genInfo{
					&generate.QueryStructMeta{ImportPkgPaths: []string{"fmt"}},
				},
			},
			want: []string{"fmt"},
		},
		{
			name: "Multiple imports",
			args: args{
				data: &genInfo{
					&generate.QueryStructMeta{ImportPkgPaths: []string{"fmt", "os", "net/http"}},
				},
			},
			want: []string{"fmt", "os", "net/http"},
		},
		{
			name: "Duplicate imports",
			args: args{
				data: &genInfo{
					&generate.QueryStructMeta{ImportPkgPaths: []string{"fmt", "os", "fmt"}},
				},
			},
			want: []string{"fmt", "os"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getImportPkgPaths(tt.args.data)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getImportPkgPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setupGenerator(config *Config) (*Generator, func(), error) {
	dirPath, err := os.MkdirTemp("./", "__temp__")
	if err != nil {
		return nil, nil, err
	}
	if config == nil {
		config = &Config{}
	}
	config.OutPath = dirPath
	g := NewGenerator(*config)
	return g, func() {
		os.RemoveAll(dirPath)
	}, nil
}

func TestGenerator_output(t *testing.T) {

	type args struct {
		fileName string
		content  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Output with valid content",
			args: args{
				fileName: "testfile.go",
				content:  []byte("package main\n\nfunc main() {}\n"),
			},
			wantErr: false,
		},
		{
			name: "Output with invalid content",
			args: args{
				fileName: "testfile.go",
				content:  []byte("package main\n\nfunc main() {\n"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, cleanup, err := setupGenerator(nil)
			if err != nil {
				t.Errorf("setupGenerator() error = %v", err)
			}
			t.Cleanup(cleanup)
			if err := g.output(fmt.Sprintf("%s/%s", g.Config.OutPath, tt.args.fileName), tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Generator.output() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
