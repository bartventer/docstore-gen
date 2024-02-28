package field

import (
	"reflect"
	"testing"

	"gocloud.dev/docstore"
)

func Test_expr_AddFieldPath(t *testing.T) {
	type fields struct {
		col Column
		e   Expression
	}
	type args struct {
		path docstore.FieldPath
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   expr
	}{
		{
			name: "AddFieldPath",
			fields: fields{
				col: Column{
					Name: "column",
					path: []string{"column"},
				},
				e: nil,
			},
			args: args{
				path: docstore.FieldPath("newPath"),
			},
			want: expr{
				col: Column{
					Name: "column",
					path: []string{"column", "newPath"},
				},
				e: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := expr{
				col: tt.fields.col,
				e:   tt.fields.e,
			}
			if got := e.AddFieldPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expr.AddFieldPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expr_ColumnName(t *testing.T) {
	type fields struct {
		col Column
		e   Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ColumnName",
			fields: fields{
				col: Column{
					Name: "column",
				},
				e: nil,
			},
			want: "column",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := expr{
				col: tt.fields.col,
				e:   tt.fields.e,
			}
			if got := e.ColumnName(); got != tt.want {
				t.Errorf("expr.ColumnName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expr_FieldPath(t *testing.T) {
	type fields struct {
		col Column
		e   Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   docstore.FieldPath
	}{
		{
			name: "FieldPath",
			fields: fields{
				col: Column{
					Name: "column",
					path: []string{"column", "path"},
				},
				e: nil,
			},
			want: docstore.FieldPath("column.path"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := expr{
				col: tt.fields.col,
				e:   tt.fields.e,
			}
			if got := e.FieldPath(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expr.FieldPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
