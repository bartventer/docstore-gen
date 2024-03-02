package field

import (
	"reflect"
	"testing"

	"gocloud.dev/docstore"
	"gocloud.dev/docstore/driver"
)

func TestConvertMods(t *testing.T) {
	type args struct {
		mods []Mod
	}
	tests := []struct {
		name string
		args args
		want docstore.Mods
	}{
		{
			name: "Test case 1: ConvertMods with Set Mod",
			args: args{
				mods: []Mod{
					Set{
						Column: Column{Name: "testColumn1", path: []string{"testColumn1"}},
						Value:  "testValue1",
					},
				},
			},
			want: docstore.Mods{
				docstore.FieldPath("testColumn1"): "testValue1",
			},
		},
		{
			name: "Test case 2: ConvertMods with Unset Mod",
			args: args{
				mods: []Mod{
					Unset{
						Column: Column{Name: "testColumn2", path: []string{"testColumn2"}},
					},
				},
			},
			want: docstore.Mods{
				docstore.FieldPath("testColumn2"): nil,
			},
		},
		{
			name: "Test case 3: ConvertMods with Inc Mod",
			args: args{
				mods: []Mod{
					newInc(Column{Name: "testColumn3", path: []string{"testColumn3"}}, 5),
				},
			},
			want: docstore.Mods{
				docstore.FieldPath("testColumn3"): driver.IncOp{Amount: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertMods(tt.args.mods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertMods() = %v, want %v", got, tt.want)
			}
		})
	}
}
