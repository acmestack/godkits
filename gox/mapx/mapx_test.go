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
)

func TestMap_ContainsKey(t *testing.T) {
	type fields struct {
		size int
		kv   map[any]any
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
				kv: map[any]any{
					"hello": "world",
				},
			},
			args: args{key: "hello"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				size: tt.fields.size,
				kv:   tt.fields.kv,
			}
			if got := m.ContainsKey(tt.args.key); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Empty(t *testing.T) {
	type fields struct {
		size int
		kv   map[any]any
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				size: 0,
				kv:   map[any]any{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				size: tt.fields.size,
				kv:   tt.fields.kv,
			}
			if got := m.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Get(t *testing.T) {
	type fields struct {
		size int
		kv   map[any]any
	}
	type args struct {
		key any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			fields: fields{
				size: 1,
				kv: map[any]any{
					"hello": "world",
				},
			},
			args: args{key: "hello"},
			want: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				size: tt.fields.size,
				kv:   tt.fields.kv,
			}
			if got := m.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Put(t *testing.T) {
	type fields struct {
		size int
		kv   map[any]any
	}
	type args struct {
		key   any
		value any
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
				kv:   map[any]any{},
			},
			args: args{
				key:   "hello",
				value: "world",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				size: tt.fields.size,
				kv:   tt.fields.kv,
			}
			if got := m.Put(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
			if value := m.Get(tt.args.key); value != tt.args.value {
				t.Errorf("Put() = %v, want %v", value, tt.args.value)
			}
		})
	}
}

func TestMap_Remove(t *testing.T) {
	type fields struct {
		size int
		kv   map[any]any
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
				kv: map[any]any{
					"hello": "world",
				},
			},
			args: args{key: "hello"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				size: tt.fields.size,
				kv:   tt.fields.kv,
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

func TestMap_Size(t *testing.T) {
	type fields struct {
		size int
		kv   map[any]any
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				size: 0,
				kv:   map[any]any{},
			},
			want: 0,
		},
		{
			fields: fields{
				size: 1,
				kv: map[any]any{
					"hello": "world",
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				size: tt.fields.size,
				kv:   tt.fields.kv,
			}
			if got := m.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Value(t *testing.T) {
	TestMap_Get(t)
}

func TestNewMap(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *Map
	}{
		{
			args: args{cap: 10},
			want: NewMap(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMap(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
