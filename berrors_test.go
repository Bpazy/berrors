package berrors

import (
	"reflect"
	"testing"
)

func TestMust(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "panic", args: args{nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Must(tt.args.err)
		})
	}
}

func TestUnwrap(t *testing.T) {
	type args struct {
		t   interface{}
		err error
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "string", args: args{"hello world", nil}, want: "hello world"},
		{name: "int", args: args{19941231, nil}, want: 19941231},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unwrap(tt.args.t, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
