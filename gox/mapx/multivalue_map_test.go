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

package mapx

import (
	"reflect"
	"testing"

	"github.com/acmestack/godkits/gox/containerx/listx"
)

func TestMultiValueMap_ContainsKey(t *testing.T) {
	list := listx.New("hello", "world")
	type fields struct {
		size int
		kvs  map[any]*listx.Arraylist
	}
	type args struct {
		key any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				size: 1,
				kvs: map[any]*listx.Arraylist{
					"hello": list,
				},
			},
			args: args{key: "hello"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MultiValueMap{
				size: tt.fields.size,
				kvs:  tt.fields.kvs,
			}
			if got := m.ContainsKey(tt.args.key); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiValueMap_Empty(t *testing.T) {
	type fields struct {
		size int
		kvs  map[any]*listx.Arraylist
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				size: 0,
				kvs:  map[any]*listx.Arraylist{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MultiValueMap{
				size: tt.fields.size,
				kvs:  tt.fields.kvs,
			}
			if got := m.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiValueMap_Get(t *testing.T) {
	list := listx.New("hello", "world")
	type fields struct {
		size int
		kvs  map[any]*listx.Arraylist
	}
	type args struct {
		key any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *listx.Arraylist
	}{
		{
			fields: fields{
				size: 1,
				kvs: map[any]*listx.Arraylist{
					"hello": list,
				},
			},
			args: args{key: "hello"},
			want: list,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MultiValueMap{
				size: tt.fields.size,
				kvs:  tt.fields.kvs,
			}
			if got := m.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiValueMap_Put(t *testing.T) {
	type fields struct {
		size int
		kvs  map[any]*listx.Arraylist
	}
	type args struct {
		key   any
		value []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				size: 0,
				kvs:  map[any]*listx.Arraylist{},
			},
			args: args{
				key:   "hello",
				value: []any{"hello", "world"},
			},
			want: true,
		},
	}
	list := listx.New("hello", "world")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MultiValueMap{
				size: tt.fields.size,
				kvs:  tt.fields.kvs,
			}
			if got := m.Put(tt.args.key, tt.args.value...); got != tt.want {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
			if value := m.Get(tt.args.key); !reflect.DeepEqual(value, list) {
				t.Errorf("Put() = %v, want %v", value, list)
			}
		})
	}
}

func TestMultiValueMap_Remove(t *testing.T) {
	list := listx.New("hello", "world")
	type fields struct {
		size int
		kvs  map[any]*listx.Arraylist
	}
	type args struct {
		key any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				size: 1,
				kvs: map[any]*listx.Arraylist{
					"hello": list,
				},
			},
			args: args{key: "hello"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MultiValueMap{
				size: tt.fields.size,
				kvs:  tt.fields.kvs,
			}
			if got := m.Remove(tt.args.key); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
			if value := m.Get(tt.args.key); value != nil {
				t.Errorf("Remove() = %v, want %v", value, nil)
			}
		})
	}
}

func TestMultiValueMap_Size(t *testing.T) {
	type fields struct {
		size int
		kvs  map[any]*listx.Arraylist
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				size: 0,
				kvs:  map[any]*listx.Arraylist{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MultiValueMap{
				size: tt.fields.size,
				kvs:  tt.fields.kvs,
			}
			if got := m.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiValueMap_Values(t *testing.T) {
	TestMultiValueMap_Get(t)
}

func TestNewMultiValueMap(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *MultiValueMap
	}{
		{
			args: args{cap: 10},
			want: NewMultiValueMap(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMultiValueMap(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
