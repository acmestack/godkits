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

package stringx

import (
	"strings"
	"testing"
)

func TestBuilderX_JoinString(t *testing.T) {
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

func TestBuilderX_JoinByte(t *testing.T) {
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
