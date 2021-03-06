package say_chat

import (
	"reflect"
	"testing"
)

func Test_decodeCmd(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *SayChatCmd
	}{
		{
			"test promote cmd",
			args{"hi"},
			&SayChatCmd{"hi", map[string]string{}},
		},
		{
			"test promote cmd blanks",
			args{"  hi key=value"},
			&SayChatCmd{"hi", map[string]string{}},
		},
		{
			"test promote cmd blanks",
			args{"  hi    key=value   "},
			&SayChatCmd{"hi", map[string]string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeCmd(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitMultiBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test split",
			args{"  43 test   key=value"},
			[]string{"43", "test", "key=value"},
		},
		{
			"test split 2",
			args{"  43 test   key=value  "},
			[]string{"43", "test", "key=value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitMultiBlank(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitMultiBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}
