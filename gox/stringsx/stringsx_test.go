/*
 * Copyright (c) 2022, OpeningO
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stringsx

import (
	"reflect"
	"testing"
)

func TestAnyEmpty(t *testing.T) {
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
		{
			args: args{str: nil},
			want: true,
		},
		{
			args: args{str: []string{" "}},
			want: true,
		},
		{
			args: args{str: []string{" ", "  "}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyEmpty(tt.args.str...); got != tt.want {
				t.Errorf("AnyEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoneEmpty(t *testing.T) {
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
			if got := NoneEmpty(tt.args.str...); got != tt.want {
				t.Errorf("NoneEmpty() = %v, want %v", got, tt.want)
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

func TestToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{str: "123"},
			want: 123,
		},
		{
			args: args{str: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.str); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{str: "123"},
			want: []byte("123"),
		},
		{
			args: args{str: ""},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBytes(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToComplex(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{
			args: args{str: "123"},
			want: 123,
		},
		{
			args: args{str: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToComplex(tt.args.str); got != tt.want {
				t.Errorf("ToComplex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{str: "0.001"},
			want: 0.001,
		},
		{
			args: args{str: ""},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.str); got != tt.want {
				t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: "1"},
			want: true,
		},
		{
			args: args{str: "0"},
			want: false,
		},
		{
			args: args{str: "t"},
			want: true,
		},
		{
			args: args{str: "f"},
			want: false,
		},
		{
			args: args{str: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBool(tt.args.str); got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUnit(t *testing.T) {
	type args struct {
		str  string
		base int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			args: args{
				str:  "10",
				base: 10,
			},
			want: 10,
		},
		{
			args: args{
				str:  "",
				base: 10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUnit(tt.args.str, tt.args.base); got != tt.want {
				t.Errorf("ToUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: "zcq"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotEmpty(tt.args.str); got != tt.want {
				t.Errorf("NotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty(tt.args.str); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceAll(t *testing.T) {
	type args struct {
		str string
		old string
		new string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "hth test", old: "test", new: "want"},
			want: "hth want",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAll(tt.args.str, tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("replaceAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripBlank(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "  ab c 	"},
			want: "ab c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripBlank(tt.args.str); got != tt.want {
				t.Errorf("StripBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrip(t *testing.T) {
	type args struct {
		str        string
		stripChars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "[{a12321a}]", stripChars: "[]"},
			want: "{a12321a}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strip(tt.args.str, tt.args.stripChars); got != tt.want {
				t.Errorf("Strip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripStart(t *testing.T) {
	type args struct {
		str        string
		stripChars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "yxabc  ", stripChars: "xyz"},
			want: "abc  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripStart(tt.args.str, tt.args.stripChars); got != tt.want {
				t.Errorf("StripStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripEnd(t *testing.T) {
	type args struct {
		str        string
		stripChars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "120.00", stripChars: ".0"},
			want: "12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripEnd(tt.args.str, tt.args.stripChars); got != tt.want {
				t.Errorf("StripEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
