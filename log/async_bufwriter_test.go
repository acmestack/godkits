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
	"time"
)

// TestNewAsyncBufferWriter test NewAsyncBufferWriter method
func TestNewAsyncBufferWriter(t *testing.T) {
	tests := []struct {
		writer io.Writer
		config Config
	}{
		{&bytes.Buffer{}, Config{
			FlushSize:     10240,
			BufferSize:    10240,
			FlushInterval: time.Second,
			Block:         true,
		}},
	}

	for _, tt := range tests {
		asyncBufferLogWriter := NewAsyncBufferWriter(tt.writer, tt.config)
		if reflect.TypeOf(asyncBufferLogWriter) != reflect.TypeOf(&AsyncBufferLogWriter{}) {
			t.Errorf("NewAsyncBufferWriter() method error, got = %v, want %v", reflect.TypeOf(asyncBufferLogWriter), reflect.TypeOf(&AsyncBufferLogWriter{}))
		}
	}
}

// TestAsyncBufferLogWriter_Write test AsyncBufferLogWriter.Write method
func TestAsyncBufferLogWriter_Write(t *testing.T) {
	type want struct {
		n   int
		err error
	}
	tests := []struct {
		asyncWriter *AsyncBufferLogWriter
		data        []byte
		wt          *want
	}{
		{
			NewAsyncBufferWriter(&bytes.Buffer{}),
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
			t.Errorf("AsyncBufferLogWriter Write() method error ! got: n=%v, error=%v want: n=%v, error=%v", n, er, tt.wt.n, tt.wt.err)
		}
	}
}

// TestAsyncBufferLogWriter_Close test AsyncBufferLogWriter.Close method
func TestAsyncBufferLogWriter_Close(t *testing.T) {
	tests := []struct {
		asyncBufferWriter *AsyncBufferLogWriter
		want              error
	}{
		{
			NewAsyncBufferWriter(&bytes.Buffer{}),
			nil,
		},
	}

	for _, tt := range tests {
		error := tt.asyncBufferWriter.Close()
		if error != tt.want {
			t.Errorf("AsyncBufferLogWriter Close() method error ! got=%v, want=%v", error, tt.want)
		}
	}
}
