// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package functions stores the function name in a map for future reference.
package functions

import (
	"github.com/fcavani/e"
	"github.com/fcavani/util/types"
)

var functions map[string]interface{}

func init() {
	functions = make(map[string]interface{})
}

// Store stores a function.
func Store(f interface{}) {
	name := types.Name(f)
	if _, found := functions[name]; !found {
		functions[name] = f
	}
}

// Get returns a function by its name. Panic if function not found.
func Get(name string) interface{} {
	f, found := functions[name]
	if !found {
		panic("function not found: " + name)
	}
	return f
}

// GetFunc returns a function by its name.
func GetFunc(name string) (interface{}, error) {
	f, found := functions[name]
	if !found {
		return nil, e.New("function not found: %v", name)
	}
	return f, nil
}

// Dump shows the stored functions.
func Dump() {
	for key := range functions {
		println(key)
	}
}
