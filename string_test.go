package ut

import (
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				text: "123456",
			},
			want: "654321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.text); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				text: " 123 45 \n6\t",
			},
			want: "123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSpace(tt.args.text); got != tt.want {
				t.Errorf("TrimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLimitLength(t *testing.T) {
	type args struct {
		s        string
		length   int
		ellipses string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				s:        "你好 hello",
				length:   8,
				ellipses: "...",
			},
			want: "你好 hello",
		},
		{
			name: "normal test",
			args: args{
				s:        "你好 hello",
				length:   6,
				ellipses: "...",
			},
			want: "你好 ...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LimitLength(tt.args.s, tt.args.length, tt.args.ellipses); got != tt.want {
				t.Errorf("LimitLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
