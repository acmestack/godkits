/*
 * Copyright (c) 2022, AcmeStack
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

package bytesx

import (
	"reflect"
	"testing"
)

func TestSetAll(t *testing.T) {
	type args struct {
		b    []byte
		v    byte
		want []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"nil", args{nil, 127, nil}},
		{"empty", args{[]byte{}, 127, []byte{}}},
		{"one", args{[]byte{0}, 127, []byte{127}}},
		{"multi", args{[]byte{0, 1, 2, 3, 4, 5, 6, 7}, 127, []byte{127, 127, 127, 127, 127, 127, 127, 127}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetAll(tt.args.b, tt.args.v)
			if !reflect.DeepEqual(tt.args.b, tt.args.want) {
				t.Errorf("SetAll() = %v, want %v", tt.args.b, tt.args.want)
			}
		})
	}
}

func TestZeroAll(t *testing.T) {
	type args struct {
		b    []byte
		want []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"nil", args{nil, nil}},
		{"empty", args{[]byte{}, []byte{}}},
		{"one", args{[]byte{0}, []byte{0}}},
		{"multi", args{[]byte{0, 1, 2, 3, 4, 5, 6, 7}, []byte{0, 0, 0, 0, 0, 0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ZeroAll(tt.args.b)
			if !reflect.DeepEqual(tt.args.b, tt.args.want) {
				t.Errorf("ZeroAll() = %v, want %v", tt.args.b, tt.args.want)
			}
		})
	}
}

func TestFromUint16LE(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"zero", args{0}, []byte{0, 0}},
		{"one", args{1}, []byte{1, 0}},
		{"random", args{0x1234}, []byte{0x34, 0x12}},
		{"max", args{65535}, []byte{255, 255}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint16LE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint16LE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16LE(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"nil", args{nil}, 0},
		{"empty", args{[]byte{}}, 0},
		{"one", args{[]byte{1}}, 1},
		{"random", args{[]byte{0x34, 0x12}}, 0x1234},
		{"max", args{[]byte{255, 255}}, 65535},
		{"max+1", args{[]byte{255, 255, 1}}, 65535},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16LE(tt.args.b); got != tt.want {
				t.Errorf("ToUint16LE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint16BE(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"zero", args{0}, []byte{0, 0}},
		{"one", args{1}, []byte{0, 1}},
		{"random", args{0x1234}, []byte{0x12, 0x34}},
		{"max", args{65535}, []byte{255, 255}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint16BE(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint16BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16BE(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"nil", args{nil}, 0},
		{"empty", args{[]byte{}}, 0},
		{"one", args{[]byte{1}}, 1},
		{"random", args{[]byte{0x12, 0x34}}, 0x1234},
		{"max", args{[]byte{255, 255}}, 65535},
		{"max+1", args{[]byte{255, 255, 1}}, 65535},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16BE(tt.args.b); got != tt.want {
				t.Errorf("ToUint16BE() = %v, want %v", got, tt.want)
			}
		})
	}
}
