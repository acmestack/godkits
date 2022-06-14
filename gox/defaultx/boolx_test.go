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

package defaultx

import (
	"testing"

	"github.com/acmestack/godkits/gox/errorsx"
)

func TestDefaultIfError(t *testing.T) {
	type args struct {
		value bool
		err   error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				value: true,
				err:   nil,
			},
			want: true,
		},
		{
			args: args{
				value: true,
				err:   errorsx.Err("error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultIfError(tt.args.value, tt.args.err); got != tt.want {
				t.Errorf("DefaultIfError() = %v, want %v", got, tt.want)
			}
		})
	}
}
