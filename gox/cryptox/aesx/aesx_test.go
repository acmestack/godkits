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

package aesx

import (
	"crypto/aes"
	"reflect"
	"testing"
)

const key = "1234567891234567"

type any = interface{}

func TestDecrypt0(t *testing.T) {
	type args struct {
		encrypted []byte
		strKey    string
	}

	encrypt, err := EncryptString("zcq", []byte(key))

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				encrypted: encrypt,
				strKey:    key,
			},
			want:    []byte("zcq"),
			wantErr: err != nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt0(tt.args.encrypted, tt.args.strKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt0() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt0() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		encrypted []byte
		key       []byte
	}

	k := []byte(key)
	encrypt, _ := Encrypt([]byte("zcq"), k)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr any
	}{
		{
			args: args{
				encrypted: encrypt,
				key:       k,
			},
			want:    []byte("zcq"),
			wantErr: nil,
		},
		{
			args: args{
				encrypted: encrypt,
				key:       nil,
			},
			want:    nil,
			wantErr: aes.KeySizeError(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encrypted, tt.args.key)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Decrypt() err = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		originData []byte
		key        []byte
	}

	k := []byte(key)
	encrypt, _ := Encrypt([]byte("zcq"), k)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr any
	}{
		{
			args: args{
				originData: []byte("zcq"),
				key:        k,
			},
			want:    encrypt,
			wantErr: nil,
		},
		{
			args: args{
				originData: []byte("zcq"),
				key:        nil,
			},
			want:    nil,
			wantErr: aes.KeySizeError(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.originData, tt.args.key)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypt1(t *testing.T) {
	type args struct {
		originData []byte
		strKey     string
	}

	encrypt, err := Encrypt0([]byte("zcq"), key)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				originData: []byte("zcq"),
				strKey:     key,
			},
			want:    encrypt,
			wantErr: err != nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt0(tt.args.originData, tt.args.strKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt0() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt0() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptString(t *testing.T) {
	type args struct {
		str string
		key []byte
	}

	encrypt, err := EncryptString("zcq", []byte(key))

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				str: "zcq",
				key: []byte(key),
			},
			want:    encrypt,
			wantErr: err != nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptString(tt.args.str, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptString1(t *testing.T) {
	type args struct {
		str    string
		strKey string
	}

	encrypt, err := EncryptString0("zcq", key)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				str:    "zcq",
				strKey: key,
			},
			want:    encrypt,
			wantErr: err != nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptString0(tt.args.str, tt.args.strKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptString0() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptString0() got = %v, want %v", got, tt.want)
			}
		})
	}
}
