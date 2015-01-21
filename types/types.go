// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2010-08-11
// Last modification:	2010-08-

// Types have functions to create an instatiation of one type from the type name.
package types

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var typemap map[string]reflect.Type

func init() {
	typemap = make(map[string]reflect.Type, 100)
	Insert(errors.New(""))
	InsertName("os.errorString", errors.New(""))
	Insert("")
	Insert(new(string))
	Insert(int(0))
	Insert(new(int))
	Insert(false)
	Insert(new(bool))
	Insert(float64(3.14))
	Insert(new(float64))
	Insert(time.Time{})
	Insert(&time.Time{})
}

// Dump the name and the type from the type base.
func Dump() {
	//fmt.Println(typemap)
	for key, t := range typemap {
		fmt.Println(key, t)
	}
}

func pkgname() (name string) {
	pc, _, _, ok := runtime.Caller(4)
	if !ok {
		return
	}
	f := runtime.FuncForPC(pc)
	s := strings.SplitN(f.Name(), ".", 2)
	if len(s) != 2 {
		return
	}
	name = s[0]
	//fmt.Println("pkgname", s)
	return
}

func findpkgname(t reflect.Type) (name string) {
	switch t.Kind() {
	case reflect.Ptr:
		name = findpkgname(t.Elem())
	//case *reflect.InterfaceType:
	case reflect.Array, reflect.Slice:
		name = findpkgname(t.Elem())
	case reflect.Chan:
		name = findpkgname(t.Elem())
	case reflect.Map:
		name = findpkgname(t.Elem())
	default:
		name = t.PkgPath()
	}
	return
}

func replacepkgname(in string, t reflect.Type) (out string) {
	//pkg := pkgname()
	pkg := findpkgname(t)
	s := strings.Split(pkg, "/")
	if len(s) <= 0 {
		return
	}
	out = strings.Replace(in, s[len(s)-1], pkg, 1)
	return
}

func nameof(t reflect.Type) (name string) {
	n := t.Name()
	if n == "" {
		name = replacepkgname(t.String(), t)
		if name == "" {
			name = t.String()
		}
	} else {
		pkg := t.PkgPath()
		if pkg == "" {
			name = n
		} else {
			name = pkg + "." + n
		}
	}
	//fmt.Printf("\"%v\", \"%v\", \"%v\" = %v\n", t.Name(), t.String(), t.PkgPath(), name)
	return
}

// NameOf returns the package name and the name of the type
func NameOf(t reflect.Type) string {
	return nameof(t)
}

func FuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// Insert type for future instantiation.
// Do this in the same package where the type was declared.
// The use of  init function is advised.
func InsertType(t reflect.Type) {
	tname := nameof(t)
	//fmt.Println(tname)
	if _, found := typemap[tname]; !found {
		typemap[tname] = t
	}
}

// Insert type for future instantiation.
// Do this in the same package where the type was declared.
// The use of  init function is advised.
func Insert(i interface{}) {
	t := reflect.ValueOf(i).Type()
	tname := nameof(t)
	//fmt.Println("Insert type:", tname)
	if _, found := typemap[tname]; !found {
		typemap[tname] = t
	}
}

// InsertName inserts a new type with the name.
func InsertName(tname string, i interface{}) {
	t := reflect.ValueOf(i).Type()
	if _, found := typemap[tname]; !found {
		typemap[tname] = t
	}
}

// Type returns the Type from the type name.
func Type(tname string) reflect.Type {
	//fmt.Printf("%#v\n", typemap)
	if t, found := typemap[tname]; found {
		return t
	}
	panic("Type not found: " + tname)
}

// IsEqualName compares the value type name with one name.
func IsEqualName(val reflect.Value, tname string) bool {
	//println("IsEqualName:", nameof(val.Type()), tname)
	return nameof(val.Type()) == tname
}

// MakeZero creates a zero value type for the type name.
func MakeZero(tname string) reflect.Value {
	return reflect.Zero(Type(tname))
}

// MakeNew create a new value from the type's name
func MakeNew(tname string, bufcap int) (val reflect.Value) {
	t := Type(tname)
	val = MakeNewType(t, bufcap)
	return
}

// MakeNewType creates a new value with type t.
func MakeNewType(t reflect.Type, bufcap int) (val reflect.Value) {
	switch t.Kind() {
	case reflect.Ptr:
		//println("**** ptr")
		val = reflect.New(t).Elem()
		val.Set(reflect.New(val.Type().Elem()))
	case reflect.Chan:
		//println("**** Chan")
		val = reflect.New(t).Elem()
		val.Set(reflect.MakeChan(t, bufcap)) //TODO: set buf?
	case reflect.Slice:
		//println("**** Slice")
		val = reflect.New(t).Elem()
		val.Set(reflect.MakeSlice(t, 0, bufcap))
		val.SetLen(bufcap)
	case reflect.Map:
		//println("**** Map")
		val = reflect.New(t).Elem()
		val.Set(reflect.MakeMap(t))
	default:
		//println("**** default")
		val = reflect.New(t).Elem()
	}
	return
}

//AllocStructPtrs find pointers in a struct and alloc than recursivily.
func AllocStructPtrs(v reflect.Value) {
	val := reflect.Indirect(v)
	t := val.Type()

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := val.Field(i)
			switch field.Type().Kind() {
			case reflect.Ptr:
				v := MakeNewType(field.Type(), 0)
				AllocStructPtrs(v)
				if field.CanSet() {
					field.Set(v)
				}
			case reflect.Slice:
				v := MakeNewType(field.Type(), 0)
				if field.CanSet() {
					field.Set(v)
				}
			default:
				continue
			}
		}
	}
	return
}
