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

import "testing"

// test hex encode to string.
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
			result := HexEncodeToString(tt.msg)
			if result != tt.want {
				t.Errorf("HexEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// test hex decode to string.
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

// test hex dump.
// please see ascii code with result.
func TestHexDump(t *testing.T) {
	tests := []struct {
		msg  []byte
		want string
	}{
		{
			msg: []byte("moremind"),
			// 6d -> m, 6f -> o, 72 -> o, 65 -> e, 69 -> i, 6e -> n, 64 -> d
			want: "00000000  6d 6f 72 65 6d 69 6e 64                           |moremind|\n",
		},
		{
			msg:  []byte("hello,OpeningO"),
			want: "00000000  68 65 6c 6c 6f 2c 4f 70  65 6e 69 6e 67 4f        |hello,OpeningO|\n",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.msg), func(t *testing.T) {
			result := HexDump(tt.msg)
			if result != tt.want {
				t.Errorf("HexDump() result = %v, want %v", result, tt.want)
			}
		})
	}
}
