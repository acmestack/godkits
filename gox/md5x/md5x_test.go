package md5x

import "testing"

func TestMd5x(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{"hello"},
			want: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5x(tt.args.str); got != tt.want {
				t.Errorf("Md5x() = %v, want %v", got, tt.want)
			}
		})
	}
}
