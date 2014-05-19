// Copyright 2013 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package com

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// md5 hash string
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func sha(m hash.Hash, str string) string {
	io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

// sha1 hash string
func Sha1(str string) string {
	return sha(sha1.New(), str)
}

// sha256 hash string
func Sha256(str string) string {
	return sha(sha256.New(), str)
}

// trim space on left
func Ltrim(str string) string {
	return strings.TrimLeftFunc(str, unicode.IsSpace)
}

// trim space on right
func Rtrim(str string) string {
	return strings.TrimRightFunc(str, unicode.IsSpace)
}

// trim space in all string length
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// repeat string times
func StrRepeat(str string, times int) string {
	return strings.Repeat(str, times)
}

// replace find all occurs to string
func StrReplace(str string, find string, to string) string {
	return strings.Replace(str, find, to, -1)
}

// IsLetter returns true if the 'l' is an English letter.
func IsLetter(l uint8) bool {
	n := (l | 0x20) - 'a'
	if n >= 0 && n < 26 {
		return true
	}
	return false
}

// Expand replaces {k} in template with match[k] or subs[atoi(k)] if k is not in match.
func Expand(template string, match map[string]string, subs ...string) string {
	var p []byte
	var i int
	for {
		i = strings.Index(template, "{")
		if i < 0 {
			break
		}
		p = append(p, template[:i]...)
		template = template[i+1:]
		i = strings.Index(template, "}")
		if s, ok := match[template[:i]]; ok {
			p = append(p, s...)
		} else {
			j, _ := strconv.Atoi(template[:i])
			if j >= len(subs) {
				p = append(p, []byte("Missing")...)
			} else {
				p = append(p, subs[j]...)
			}
		}
		template = template[i+1:]
	}
	p = append(p, template...)
	return string(p)
}

// Reverse s string, support unicode
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
