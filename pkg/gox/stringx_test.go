package gox

import "testing"

func TestEmpty(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: []string{""}},
			want: true,
		},
		{
			args: args{str: []string{"hello world"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty(tt.args.str...); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: []string{""}},
			want: false,
		},
		{
			args: args{str: []string{"hello world"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotEmpty(tt.args.str...); got != tt.want {
				t.Errorf("NotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	type args struct {
		str        string
		defaultStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "", defaultStr: "hello world"},
			want: "hello world",
		},
		{
			args: args{str: "abc", defaultStr: "hello world"},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultIfEmpty(tt.args.str, tt.args.defaultStr); got != tt.want {
				t.Errorf("DefaultIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
