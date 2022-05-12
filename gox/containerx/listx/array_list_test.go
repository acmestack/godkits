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

package listx

import (
	"reflect"
	"testing"
)

func TestArraylist_Add(t *testing.T) {
	emptyArrayList := Arraylist{
		elements: []any{},
		size:     0,
	}
	type fields struct {
		elements []any
		size     int
	}
	type args struct {
		element []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{
				elements: emptyArrayList.elements,
				size:     emptyArrayList.size,
			},
			args: args{element: []any{"hello", "world"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			lst.Add(tt.args.element...)
		})
	}
}

func TestArraylist_Contains(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type fields struct {
		elements []any
		size     int
	}
	type args struct {
		element any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int
	}{
		{
			fields: fields{
				elements: arrayList.elements,
				size:     arrayList.size,
			},
			args:  args{element: "hello"},
			want:  true,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			got, got1 := lst.Contains(tt.args.element)
			if got != tt.want {
				t.Errorf("Contains() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Contains() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArraylist_Empty(t *testing.T) {
	emptyArrayList := Arraylist{
		elements: []any{},
		size:     0,
	}
	type fields struct {
		elements []any
		size     int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				elements: emptyArrayList.elements,
				size:     emptyArrayList.size,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if got := lst.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraylist_Get(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type fields struct {
		elements []any
		size     int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			fields: fields{
				elements: arrayList.elements,
				size:     arrayList.size,
			},
			args: args{index: 0},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if got := lst.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraylist_IndexOf(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type fields struct {
		elements []any
		size     int
	}
	type args struct {
		element any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			fields: fields{
				elements: arrayList.elements,
				size:     arrayList.size,
			},
			args: args{"hello"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if got := lst.IndexOf(tt.args.element); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraylist_RemoveAtIndex(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type fields struct {
		elements []any
		size     int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{
			fields: fields{
				elements: arrayList.elements,
				size:     arrayList.size,
			},
			args: args{index: 0},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if got := lst.RemoveAtIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraylist_Size(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type fields struct {
		elements []any
		size     int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				elements: arrayList.elements,
				size:     arrayList.size,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if got := lst.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraylist_find(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type fields struct {
		elements []any
		size     int
	}
	type args struct {
		element any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int
	}{
		{
			fields: fields{
				elements: arrayList.elements,
				size:     arrayList.size,
			},
			args:  args{element: "hello"},
			want:  true,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			got, got1 := lst.find(tt.args.element)
			if got != tt.want {
				t.Errorf("find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNew(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type args struct {
		elements []any
	}

	tests := []struct {
		name string
		args args
		want *Arraylist
	}{
		{
			args: args{arrayList.elements},
			want: New("hello", "world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.elements...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraylist_Remove(t *testing.T) {
	arrayList := &Arraylist{
		elements: []any{"hello", "world"},
		size:     2,
	}
	type fields struct {
		elements []any
		size     int
	}
	type args struct {
		element any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				elements: arrayList.elements,
				size:     arrayList.size,
			},
			args: args{element: "hello"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Arraylist{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if got := lst.Remove(tt.args.element); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
