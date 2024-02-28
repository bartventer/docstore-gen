package docstoregen

import (
	"reflect"
	"testing"
)

func Test_importPkgS_Add(t *testing.T) {
	type fields struct {
		paths []string
	}
	type args struct {
		paths []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *importPkgS
	}{
		{
			name: "Add non-empty paths",
			fields: fields{
				paths: []string{},
			},
			args: args{
				paths: []string{"path1", "path2"},
			},
			want: &importPkgS{
				paths: []string{`"path1"`, `"path2"`, ""},
			},
		},
		{
			name: "Add empty paths",
			fields: fields{
				paths: []string{},
			},
			args: args{
				paths: []string{"", ""},
			},
			want: &importPkgS{
				paths: []string{"", "", ""},
			},
		},
		{
			name: "Add duplicate paths",
			fields: fields{
				paths: []string{`"path1"`},
			},
			args: args{
				paths: []string{"path1", "path2"},
			},
			want: &importPkgS{
				paths: []string{`"path1"`, `"path2"`, ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := importPkgS{
				paths: tt.fields.paths,
			}
			if got := ip.Add(tt.args.paths...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importPkgS.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importPkgS_Paths(t *testing.T) {
	type fields struct {
		paths []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Paths",
			fields: fields{
				paths: []string{`"path1"`, `"path2"`, ""},
			},
			want: []string{`"path1"`, `"path2"`, ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := importPkgS{
				paths: tt.fields.paths,
			}
			if got := ip.Paths(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importPkgS.Paths() = %v, want %v", got, tt.want)
			}
		})
	}
}
