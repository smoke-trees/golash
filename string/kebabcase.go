/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package string

import "strings"

//ToScreamingDelimited converts a string to SCREAMING.DELIMITED.SNAKE.CASE (in this case `del = '.'; screaming = true`) or delimited.snake.case (in this case `del = '.'; screaming = false`)
func ToScreamingDelimited(s string, del uint8, screaming bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextCaseIsChanged := false
		if i+1 < len(s) {
			next := s[i+1]
			nextCaseIsChanged = isNextCaseIsChanged(v, next, nextCaseIsChanged)
		}

		n = toDelimited(i, n, del, nextCaseIsChanged, v)
	}

	if screaming {
		n = strings.ToUpper(n)
	} else {
		n = strings.ToLower(n)
	}
	return n
}

func toDelimited(i int, n string, del uint8, nextCaseIsChanged bool, v int32) string {
	if i > 0 && n[len(n)-1] != del && nextCaseIsChanged {
		// add underscore if next letter case type is changed
		n = addUnderScore(v, n, del)
	} else if v == ' ' || v == '_' || v == '-' {
		// replace spaces/underscores with delimiters
		n += string(del)
	} else {
		n = n + string(v)
	}
	return n
}

func addUnderScore(v int32, n string, del uint8) string {
	if v >= 'A' && v <= 'Z' {
		n += string(del) + string(v)
	} else if v >= 'a' && v <= 'z' {
		n += string(v) + string(del)
	}
	return n
}

func isNextCaseIsChanged(v int32, next uint8, nextCaseIsChanged bool) bool {
	if (v >= 'A' && v <= 'Z' && next >= 'a' && next <= 'z') || (v >= 'a' && v <= 'z' && next >= 'A' && next <= 'Z') {
		nextCaseIsChanged = true
	}
	return nextCaseIsChanged
}

// ToDelimited converts a string to delimited.snake.case (in this case `del = '.'`)
func ToDelimited(s string, del uint8) string {
	return ToScreamingDelimited(s, del, false)
}

// KebabCase converts a string to kebab-case
func KebabCase(s string) string {
	return ToDelimited(s, '-')
}

// ToScreamingSnake converts a string to SCREAMING_SNAKE_CASE
func ToScreamingSnake(s string) string {
	return ToScreamingDelimited(s, '_', true)
}

// SnakeCase converts a string to snake_case
func SnakeCase(s string) string {
	return ToDelimited(s, '_')
}
