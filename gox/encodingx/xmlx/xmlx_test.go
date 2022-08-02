/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
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

package xmlx

import (
	"reflect"
	"testing"
)

type user struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

var u = user{
	Name: "renzhuyan",
	Age:  123,
}

func TestXmlBytesToAny(t *testing.T) {
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
			name: "xml bytes to any",
			args: args{
				bytes: []byte("<user><name>renzhuyan</name><age>123</age></user>"),
				v:     &user{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := XmlBytesToAny(tt.args.bytes, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("XmlBytesToAny() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.v, &u) {
				t.Errorf("xmlStringToAny() got = %v, want %v", tt.args.v, &u)
			}
		})
	}
}

func TestXmlStringToAny(t *testing.T) {
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
			name: "xml string to any",
			args: args{
				str: "<user><name>renzhuyan</name><age>123</age></user>",
				v:   &user{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := XmlStringToAny(tt.args.str, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("xmlStringToAny() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.v, &u) {
				t.Errorf("xmlStringToAny() got = %v, want %v", tt.args.v, &u)
			}
		})
	}
}

func TestToXmlBytes(t *testing.T) {
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
			name:    "any to xml bytes",
			args:    args{v: u},
			want:    []byte("<user><name>renzhuyan</name><age>123</age></user>"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToXmlBytes(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJXmlBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToXmlBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToXmlString(t *testing.T) {
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
			name:    "any to xml string",
			args:    args{v: u},
			want:    "<user><name>renzhuyan</name><age>123</age></user>",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToXmlString(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToXmlString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToXmlString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
