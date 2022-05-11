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
	"container/list"
	"testing"
)

func TestEmpty(t *testing.T) {
	type args struct {
		lst *list.List
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{lst: list.New()},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Empty(tt.args.lst); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListx_Empty(t *testing.T) {
	type fields struct {
		List list.List
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{List: *list.New()},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Listx{
				List: tt.fields.List,
			}
			if got := lst.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListx_NotEmpty(t *testing.T) {
	type fields struct {
		List list.List
	}
	lst := list.New()
	lst.PushBack("hello world")
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{List: *lst},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lst := &Listx{
				List: tt.fields.List,
			}
			if got := lst.NotEmpty(); got != tt.want {
				t.Errorf("NotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	type args struct {
		lst *list.List
	}
	lst := list.New()
	lst.PushBack("hello world")
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{lst: lst},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotEmpty(tt.args.lst); got != tt.want {
				t.Errorf("NotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
