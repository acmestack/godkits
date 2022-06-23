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

package base64x

import (
	"testing"
)

type any = interface{}

// test TestBase64EncodeToString
// base64("i am moremind") = "aSBhbSBtb3JlbWluZA=="
func TestBase64EncodeToString(t *testing.T) {
	tests := []struct {
		msg  []byte
		want string
	}{
		{
			msg:  []byte("i am moremind"),
			want: "aSBhbSBtb3JlbWluZA==",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.msg), func(t *testing.T) {
			result := Base64EncodeToString(tt.msg)
			if result != tt.want {
				t.Errorf("Base64EncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

// test TestBase64DecodeToBytes
func TestBase64DecodeToBytes(t *testing.T) {
	tests := []struct {
		msg     string
		want    string
		wantErr bool
	}{
		{
			msg:     "aSBhbSBtb3JlbWluZA==",
			want:    "i am moremind",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			result, err := Base64DecodeToBytes(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64DecodeToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != tt.want {
				t.Errorf("Base64DecodeToBytes() result = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestBase64URLEncodeToString(t *testing.T) {
	tests := []struct {
		msg  []byte
		want string
	}{
		{
			msg:  []byte("https://openingo.org/"),
			want: "aHR0cHM6Ly9vcGVuaW5nby5vcmcv",
		},
		{
			msg:  []byte("https://openingo.org?name=moremind"),
			want: "aHR0cHM6Ly9vcGVuaW5nby5vcmc_bmFtZT1tb3JlbWluZA==",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.msg), func(t *testing.T) {
			result := Base64URLEncodeToString(tt.msg)
			if result != tt.want {
				t.Errorf("Base64URLEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestBase64URLDecodeToByte(t *testing.T) {
	tests := []struct {
		msg     string
		want    []byte
		wantErr any
	}{
		{
			msg:     "aHR0cHM6Ly9vcGVuaW5nby5vcmcv",
			want:    []byte("https://openingo.org/"),
			wantErr: false,
		},
		{
			msg:     "aHR0cHM6Ly9vcGVuaW5nby5vcmc_bmFtZT1tb3JlbWluZA==",
			want:    []byte("https://openingo.org?name=moremind"),
			wantErr: false,
		},
		// TODO moremind add error test cases
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			result, err := Base64URLDecodeToByte(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64URLDecodeToByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != string(tt.want) {
				t.Errorf("Base64URLDecodeToByte() result = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestBase64RawEncodeToString(t *testing.T) {
	tests := []struct {
		msg  []byte
		want string
	}{
		{
			msg:  []byte("i am files bytes"),
			want: "aSBhbSBmaWxlcyBieXRlcw",
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.msg), func(t *testing.T) {
			result := Base64RawEncodeToString(tt.msg)
			if result != tt.want {
				t.Errorf("Base64RawEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestBase64RawDecodeToByte(t *testing.T) {
	tests := []struct {
		msg     string
		want    []byte
		wantErr bool
	}{
		{
			msg:  "aSBhbSBmaWxlcyBieXRlcw",
			want: []byte("i am files bytes"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			result, err := Base64RawDecodeToByte(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64URLDecodeToByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != string(tt.want) {
				t.Errorf("Base64RawEncodeToString() result = %v, want %v", result, tt.want)
			}
		})
	}
}
