package ut

import (
	"reflect"
	"testing"
)

func TestItemInSlice(t *testing.T) {
	type args struct {
		i    interface{}
		list interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "int in slice test 1",
			args: args{
				i:    1,
				list: []int{1, 2},
			},
			want: true,
		},
		{
			name: "int in slice test 2",
			args: args{
				i:    1,
				list: []int{2, 3},
			},
			want: false,
		},
		{
			name: "string test 1",
			args: args{
				i:    "hello",
				list: []string{"2", "hello"},
			},
			want: true,
		},
		{
			name: "mix test 1",
			args: args{
				i:    3,
				list: []string{"2", "3"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemInSlice(tt.args.i, tt.args.list); got != tt.want {
				t.Errorf("ItemInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal test",
			args: args{
				min: 1,
				max: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "normal test",
			args: args{
				min: 2,
				max: 2,
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Range(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
