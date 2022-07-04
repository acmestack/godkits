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

package log

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

// TestNewAsyncWriter test NewAsyncWriter method
//   @params t tests params
func TestNewAsyncWriter(t *testing.T) {
	tests := []struct {
		writer  io.Writer
		bufSize int
		block   bool
	}{
		{&bytes.Buffer{}, 0, false},
		{&bytes.Buffer{}, 1024, true},
	}

	for _, tt := range tests {
		asyncLogWriter := NewAsyncWriter(tt.writer, tt.bufSize, tt.block)
		if reflect.TypeOf(asyncLogWriter) != reflect.TypeOf(&AsyncLogWriter{}) {
			t.Errorf("NewAsyncWriter() method error, got = %v, want %v", reflect.TypeOf(asyncLogWriter), reflect.TypeOf(&AsyncLogWriter{}))
		}
	}
}

// TestAsyncLogWriter_Write test AsyncLogWriter.Write method
//   @params t tests params
func TestAsyncLogWriter_Write(t *testing.T) {
	type want struct {
		n   int
		err error
	}
	tests := []struct {
		asyncWriter *AsyncLogWriter
		data        []byte
		wt          *want
	}{
		{
			NewAsyncWriter(&bytes.Buffer{}, 0, false),
			[]byte("hello"),
			&want{
				n:   5,
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		n, er := tt.asyncWriter.Write(tt.data)
		if n != tt.wt.n || er != tt.wt.err {
			t.Errorf("AsyncLogWriter Write() method error ! got: n=%v, error=%v want: n=%v, error=%v", n, er, tt.wt.n, tt.wt.err)
		}
	}

}

// TestAsyncLogWriter_Close test AsyncLogWriter.Close method
func TestAsyncLogWriter_Close(t *testing.T) {
	tests := []struct {
		asyncWriter *AsyncLogWriter
		want        error
	}{
		{
			NewAsyncWriter(&bytes.Buffer{}, 0, false),
			nil,
		},
	}

	for _, tt := range tests {
		error := tt.asyncWriter.Close()
		if error != tt.want {
			t.Errorf("AsyncLogWriter Close() method error ! got=%v, want=%v", error, tt.want)
		}
	}
}
