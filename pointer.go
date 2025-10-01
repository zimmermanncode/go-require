/** go-require: Runtime Go assertions w/ descriptive panics & a fluent API

Copyright (c) 2025 Stefan Zimmermann <user@zimmermann.co>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package require provides runtime Go assertions w/ descriptive panics & a fluent API.
package require

import (
	"github.com/samber/lo"
)

// NilPtr asserts that the given value is a nil pointer & returns it on success.
// It panics w/ a descriptive message based on given name if the pointer is not nil.
//
// Example:
//
//	value := require.NilPtr("My value", GetSomePointerValue())
//	// if not nil => panic: assertion failed: My value should be a nil pointer
func NilPtr[T any](name string, value *T) *T {
	lo.Assertf(value == nil, "%s should be a nil pointer", name)

	return value
}

// NotNilPtr asserts that the given value is no nil pointer & returns it on success.
// It panics w/ a descriptive message based on given name if the pointer is nil.
//
// Example:
//
//	value := require.NotNilPtr("My value", GetSomePointerValue())
//	// if nil => panic: assertion failed: My value should not be a nil pointer
func NotNilPtr[T any](name string, value *T) *T {
	lo.Assertf(value != nil, "%s should not be a nil pointer", name)

	return value
}
