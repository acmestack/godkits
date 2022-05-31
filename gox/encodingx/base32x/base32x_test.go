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

package base32x

import "testing"

// base32x encode to string test cases
func TestBase32EncodeToString(t *testing.T) {
	tests := []struct {
		msg  []byte
		want string
	}{
		{
			msg:  []byte("I am moremind"),
			want: "JEQGC3JANVXXEZLNNFXGI===",
		},
		{
			msg:  []byte("hello, OpeningO"),
			want: "NBSWY3DPFQQE64DFNZUW4Z2P",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.msg), func(t *testing.T) {
			result := Base32EncodeToString(tt.msg)
			if result != tt.want {
				t.Errorf("Base32EncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}

}

// base32x decode string to bytes
func TestBase32DecodeToBytes(t *testing.T) {
	tests := []struct {
		msg     string
		want    []byte
		wantErr bool
	}{
		{
			msg:     "JEQGC3JANVXXEZLNNFXGI===",
			want:    []byte("I am moremind"),
			wantErr: false,
		},
		{
			msg:     "NBSWY3DPFQQE64DFNZUW4Z2P",
			want:    []byte("hello, OpeningO"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			result, err := Base32DecodeToBytes(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base32DecodeToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != string(tt.want) {
				t.Errorf("Base32DecodeToBytes() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// base32x hex encode bytes to string
func TestBase32HexEncodeToString(t *testing.T) {
	tests := []struct {
		msg  []byte
		want string
	}{
		{
			msg:  []byte("I am moremind"),
			want: "94G62R90DLNN4PBDD5N68===",
		},
		{
			msg:  []byte("hello, OpeningO"),
			want: "D1IMOR3F5GG4US35DPKMSPQF",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.msg), func(t *testing.T) {
			result := Base32HexEncodeToString(tt.msg)
			if result != tt.want {
				t.Errorf("Base32HexEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// base32x hex decode test cases
func TestBase32HexDecodeToBytes(t *testing.T) {
	tests := []struct {
		msg     string
		want    []byte
		wantErr bool
	}{
		{
			msg:     "94G62R90DLNN4PBDD5N68===",
			want:    []byte("I am moremind"),
			wantErr: false,
		},
		{
			msg:     "D1IMOR3F5GG4US35DPKMSPQF",
			want:    []byte("hello, OpeningO"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			result, err := Base32HexDecodeToBytes(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base32HexDecodeToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != string(tt.want) {
				t.Errorf("Base32HexDecodeToBytes() result = %v, want %v", result, tt.want)
			}
		})
	}
}
