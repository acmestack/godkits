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

package aesx

import (
	"reflect"
	"testing"
)

const key = "1234567891234567"

func TestDecrypt(t *testing.T) {
	type args struct {
		encrypted []byte
		key       []byte
	}
	k := []byte(key)
	encrypt, err := Encrypt([]byte("zcq"), k)
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				encrypted: encrypt,
				key:       k,
			},
			want:    []byte("zcq"),
			wantErr: err != nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encrypted, tt.args.key)
			if (err != nil) != tt.wantErr {
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
	encrypt, err := Encrypt([]byte("zcq"), k)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{
				originData: []byte("zcq"),
				key:        k,
			},
			want:    encrypt,
			wantErr: err != nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.originData, tt.args.key)
			if (err != nil) != tt.wantErr {
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

	encrypt, err := Encrypt1([]byte("zcq"), key)

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
			got, err := Encrypt1(tt.args.originData, tt.args.strKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt1() got = %v, want %v", got, tt.want)
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

	encrypt, err := EncryptString1("zcq", key)

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
			got, err := EncryptString1(tt.args.str, tt.args.strKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptString1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncryptString1() got = %v, want %v", got, tt.want)
			}
		})
	}
}
