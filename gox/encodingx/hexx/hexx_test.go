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

package hexx

import (
	"testing"
)

// TestHexEncodeToString test hex encode to string.
func TestHexEncodeToString(t *testing.T) {
	tests := []struct {
		msg  []byte
		want string
	}{
		{
			msg:  []byte("I am moremind"),
			want: "4920616d206d6f72656d696e64",
		},
		{
			msg:  []byte("hello,OpeningO"),
			want: "68656c6c6f2c4f70656e696e674f",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.msg), func(t *testing.T) {
			result := EncodeBytesToString(tt.msg)
			if result != tt.want {
				t.Errorf("HexEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestEncodeToBytes base method testcase, encode to bytes
func TestEncodeToBytes(t *testing.T) {
	tests := []struct {
		dst      []byte
		src      []byte
		hextable string
		want     int
	}{
		{
			dst:      []byte("4920616d206d6f72656d696e64"),
			src:      []byte("I am moremind"),
			hextable: "0123456789abcdef",
			want:     26,
		},
		{
			dst:      []byte("4920616D206D6F72656D696E64"),
			src:      []byte("I am moremind"),
			hextable: "0123456789ABCDEF",
			want:     26,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.src), func(t *testing.T) {
			result := EncodeToBytes(tt.dst, tt.src, tt.hextable)
			if result != tt.want {
				t.Errorf("HexEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestEncodeToCaseBytes test encode bytes to upper or lower case bytes
func TestEncodeToCaseBytes(t *testing.T) {
	tests := []struct {
		src         []byte
		toLowerCase bool
		want        []byte
	}{
		{
			src:         []byte("I am moremind"),
			toLowerCase: true,
			want:        []byte("4920616d206d6f72656d696e64"),
		},
		{
			src:         []byte("I am moremind"),
			toLowerCase: false,
			want:        []byte("4920616D206D6F72656D696E64"),
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.src), func(t *testing.T) {
			result := EncodeToCaseBytes(tt.src, tt.toLowerCase)
			if string(result) != string(tt.want) {
				t.Errorf("HexEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestEncodeToCaseString test encode bytes to upper or lower case string
func TestEncodeToCaseString(t *testing.T) {
	tests := []struct {
		src         []byte
		toLowerCase bool
		want        string
	}{
		{
			src:         []byte("I am moremind"),
			toLowerCase: true,
			want:        "4920616d206d6f72656d696e64",
		},
		{
			src:         []byte("I am moremind"),
			toLowerCase: false,
			want:        "4920616D206D6F72656D696E64",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.src), func(t *testing.T) {
			result := EncodeToCaseString(tt.src, tt.toLowerCase)
			if result != tt.want {
				t.Errorf("HexEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// TestEncodeToHexString encode string to hex string
func TestEncodeToHexString(t *testing.T) {
	tests := []struct {
		src         string
		toLowerCase bool
		want        string
	}{
		{
			src:         "I am moremind",
			toLowerCase: true,
			want:        "4920616d206d6f72656d696e64",
		},
		{
			src:         "I am moremind",
			toLowerCase: false,
			want:        "4920616D206D6F72656D696E64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.src, func(t *testing.T) {
			result := EncodeToHexString(tt.src, tt.toLowerCase)
			if result != tt.want {
				t.Errorf("HexEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}

}

// TestHexDecodeToString test hex decode to string.
func TestHexDecodeToString(t *testing.T) {
	tests := []struct {
		msg     string
		want    []byte
		wantErr bool
	}{
		{
			msg:     "4920616d206d6f72656d696e64",
			want:    []byte("I am moremind"),
			wantErr: false,
		},
		{
			msg:     "68656c6c6f2c4f70656e696e674f",
			want:    []byte("hello,OpeningO"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			result, err := HexDecodeToBytes(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexDecodeToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != string(tt.want) {
				t.Errorf("HexDecodeToBytes() result = %v, want %v", result, tt.want)
			}
		})
	}
}
