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

package jsonx

import (
	"reflect"
	"testing"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var u = user{
	Name: "qicz",
	Age:  123,
}

func TestJsonBytesToAny(t *testing.T) {
	type args struct {
		bytes []byte
		v     any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "json bytes to any",
			args: args{
				bytes: []byte("{\"name\":\"qicz\",\"age\":123}"),
				v:     &user{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JsonBytesToAny(tt.args.bytes, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JsonBytesToAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJsonStringToAny(t *testing.T) {
	type args struct {
		str string
		v   any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "json string to any",
			args: args{
				str: "{\"name\":\"qicz\",\"age\":123}",
				v:   &u,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JsonStringToAny(tt.args.str, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JsonStringToAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToJsonBytes(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "any to json bytes",
			args:    args{v: u},
			want:    []byte("{\"name\":\"qicz\",\"age\":123}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJsonBytes(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJsonBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToJsonBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToJsonString(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "any to json string",
			args:    args{v: u},
			want:    "{\"name\":\"qicz\",\"age\":123}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJsonString(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJsonString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJsonString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
