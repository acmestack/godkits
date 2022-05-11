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

package defaultx

import (
	"testing"

	"github.com/openingo/godkits/gox/errorsx"
)

func TestDefaultComplexIfError(t *testing.T) {
	type args struct {
		err          error
		value        complex128
		defaultValue complex128
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{
			args: args{
				err:          nil,
				value:        123,
				defaultValue: 123,
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultComplexIfError(tt.args.err, tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("DefaultComplexIfError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFloat64IfError(t *testing.T) {
	type args struct {
		err          error
		value        float64
		defaultValue float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				err:          errorsx.Err("error"),
				value:        1.12,
				defaultValue: 0.123,
			},
			want: 0.123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultFloat64IfError(tt.args.err, tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("DefaultFloat64IfError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultIntIfError(t *testing.T) {
	type args struct {
		err          error
		value        int
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				err:          nil,
				value:        123,
				defaultValue: 123,
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultIntIfError(tt.args.err, tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("DefaultIntIfError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultUint64IfError(t *testing.T) {
	type args struct {
		err          error
		value        uint64
		defaultValue uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			args: args{
				err:          nil,
				value:        123,
				defaultValue: 123,
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultUint64IfError(tt.args.err, tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("DefaultUint64IfError() = %v, want %v", got, tt.want)
			}
		})
	}
}
