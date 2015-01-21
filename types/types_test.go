// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2010-09-21
// Last modification:	2010-

package types

import (
	"reflect"
	"testing"
)

type testitem struct {
	t    interface{}
	name string
	val  interface{}
}

type testvec0 []int
type testvec1 []testitem
type teststr string

var boolean bool = true
var ptrbool *bool = &boolean
var interger int = 42
var ptrint *int = &interger
var interger32 int32 = 42
var ptrint32 *int32 = &interger32
var interger8 uint8 = 42
var ptrint8 *uint8 = &interger8
var cadeia string = "outrastring"
var ptrcadeia *string = &cadeia

var tests []testitem = []testitem{
	{true, "bool", true},
	{new(bool), "*bool", ptrbool},
	{int(0), "int", 42},
	{new(int), "*int", ptrint},
	{int32(0), "int32", int32(64)},
	{new(int32), "*int32", ptrint32},
	{byte(0), "uint8", uint8(1)},
	{new(byte), "*uint8", ptrint8},
	{"string", "string", "istoeumastring"},
	{new(string), "*string", ptrcadeia},
	{struct{}{}, "struct {}", struct{}{}},
	{&struct{}{}, "*struct {}", &struct{}{}},
	{testitem{}, "github.com/fcavani/util/types.testitem", testitem{int(0), "int", 0}},
	{&testitem{}, "*github.com/fcavani/util/types.testitem", &testitem{int(0), "int", 1}},
	{[3]int{}, "[3]int", [3]int{1, 2, 3}},
	{[]int{}, "[]int", []int{1, 2, 3}},
	{[2]testitem{}, "[2]github.com/fcavani/util/types.testitem", [2]testitem{{int(9), "int", 0}, {int(9), "int", 2}}},
	{[]testitem{}, "[]github.com/fcavani/util/types.testitem", []testitem{{int(7), "int", 0}, {int(8), "int", 2}}},
	{testvec0{}, "github.com/fcavani/util/types.testvec0", testvec0{1, 2, 3, 4}},
	{map[string]string{}, "map[string]string", map[string]string{"a": "1", "b": "2"}},
	{map[string]string{}, "map[string]string", map[string]string{}},
	//{&testvec0{}, "*serialization/types.testvec0", &testvec0{1,2,3,4}},
	{teststr("oi"), "github.com/fcavani/util/types.teststr", teststr("oi")},
}

func TestNameOf(t *testing.T) {
	//print("\nTestNameOf\n\n")
	for i, test := range tests {
		s := NameOf(reflect.ValueOf(test.t).Type())
		if s != test.name {
			t.Fatalf("NameOf %v failed: %v != %v", i, test.name, s)
		}
	}
}

func TestIsert(t *testing.T) {
	for _, typ := range tests {
		Insert(typ.t)
	}
}

func TestMake(t *testing.T) {
	for i, typ := range tests {
		val := MakeNew(typ.name, 0)
		//println("can set:", val.CanSet())
		name := NameOf(val.Type())
		//println(name)
		if name != typ.name {
			t.Fatalf("type name differ in %v: %v != %v", i, name, typ.name)
		}
		//println(i)
		val.Set(reflect.ValueOf(typ.val))
		if !reflect.DeepEqual(val.Interface(), typ.val) {
			t.Fatalf("not equal: %v", i)
		}
	}
}
