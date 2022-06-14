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

package stringsx

import (
	"strings"
	"testing"
)

func TestBuilder(t *testing.T) {
	// go strings
	stringsBuilder := strings.Builder{}
	stringsBuilder.WriteString("hello")
	stringsBuilder.WriteByte(' ')
	stringsBuilder.WriteString("world")
	stringsBuilder.Grow(123)

	// gox stringsx
	stringxBuilder := &Builder{
		Builder: strings.Builder{},
	}
	stringxBuilder.JoinString("hello")
	stringxBuilder.JoinByte(' ')
	stringxBuilder.JoinString("world")
	stringxBuilder.Grow(123)

	if stringsBuilder.Cap() != stringxBuilder.Cap() {
		t.Errorf("Cap() is not eq, stringsBuilder.Cap() = %d, stringxsBuilder.Cap() = %d", stringsBuilder.Cap(), stringxBuilder.Cap())
	}

	if stringsBuilder.Len() != stringxBuilder.Len() {
		t.Errorf("Len() is not eq, stringsBuilder.Len() = %d, stringxsBuilder.Len() = %d", stringsBuilder.Len(), stringxBuilder.Len())
	}

	if stringsBuilder.String() != stringxBuilder.String() {
		t.Errorf("String() is not eq, stringsBuilder.String() = %s, stringxsBuilder.String() = %s", stringsBuilder.String(), stringxBuilder.String())
	}
}

func TestBuilder_JoinString(t *testing.T) {
	type fields struct {
		Builder strings.Builder
	}
	type args struct {
		strArray []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			fields:  fields{Builder: strings.Builder{}},
			args:    args{strArray: []string{"a", "b"}},
			want:    2,
			wantErr: false,
		},
		{
			fields:  fields{Builder: strings.Builder{}},
			args:    args{strArray: nil},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := &Builder{
				Builder: tt.fields.Builder,
			}
			got, err := builder.JoinString(tt.args.strArray...)
			if (err != nil) != tt.wantErr {
				t.Errorf("JoinString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JoinString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_JoinByte(t *testing.T) {
	type fields struct {
		Builder strings.Builder
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{

			fields:  fields{Builder: strings.Builder{}},
			args:    args{bytes: []byte{'a', 'b'}},
			want:    2,
			wantErr: false,
		},
		{
			fields:  fields{Builder: strings.Builder{}},
			args:    args{bytes: nil},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := &Builder{
				Builder: tt.fields.Builder,
			}
			got, err := builder.JoinByte(tt.args.bytes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("JoinByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JoinByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}
