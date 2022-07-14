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

package core

import (
	"bufio"
	"github.com/acmestack/godkits/gox/stringsx"
	"io"
	"os"
	"strings"
)

type Properties struct {
	m map[string]string
}

func NewProperties() *Properties {
	return &Properties{
		m: map[string]string{},
	}
}

// GetConfigFromFile get map from file
//  @param path file path
//  @return map[string]string
//  usage:
//  p := NewProperties()
//  p.GetConfigFromFile(your path)
func (p *Properties) GetConfigFromFile(path string) (err error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		_, _, err = p.put(key, value)
	}
	return nil
}

// Get load value by key from Properties
//  @receiver p properties
//  @param key the key
//  @return value the value
//  @return ok status
func (p *Properties) Get(key string) (value string, ok bool) {
	_, ok = p.m[key]
	if !ok {
		return "", false
	}
	return p.m[key], true
}

// put put key,value into properties
//  @receiver p properties
//  @param key key
//  @param value value
//  @return preValue pre
//  @return ok status
//  @return err error
func (p *Properties) put(key, value string) (preValue string, ok bool, err error) {
	if stringsx.Empty(key) {
		return "", false, nil
	}
	preValue, ok = p.Get(key)
	// check key and value
	if !ok {
		p.m[key] = value
		return preValue, true, nil
	}
	panic("key repeat!")
}
