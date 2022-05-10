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

package errorsx

import "testing"

func TestErr(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args:    args{message: "error message"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Err(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Err() error = %v, wantErr %v, logError %v", err, tt.wantErr, LogError(err))
			}
		})
	}
}
