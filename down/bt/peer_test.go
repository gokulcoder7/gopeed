package bt

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_decodeBitfield(t *testing.T) {
	fmt.Printf("%08b", 0xFF)
	type args struct {
		payload []byte
	}
	tests := []struct {
		name string
		args args
		want MessageBitfield
	}{
		{
			"case-11111111",
			args{
				[]byte{0xFF},
			},
			[]bool{true, true, true, true, true, true, true, true},
		},
		{
			"case-10101010",
			args{
				[]byte{0xAA},
			},
			[]bool{true, false, true, false, true, false, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeBitfield(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeBitfield() = %v, want %v", got, tt.want)
			}
		})
	}
}
