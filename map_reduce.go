/**
 * @Author: alessonhu
 * @Description:
 * @File:  map_reduce.go
 * @Version: 1.0.0
 * @Date: 2022/3/29 10:56 上午
 */

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i, _ := range list {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i, _ := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i, _ := range list {
		sum += fn(&list[i])
	}
	return sum
}

// Generic Map
func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

// 健壮版map
func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {
	//check the <code data-enlighter-language="raw" class="EnlighterJSRAW">slice</code> type is Slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	//check the function signature
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()

}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	//Check it is a funciton
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

// 健壮版reduce
func Reduce(slice, pairFunc, zero interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}

	elemType := sliceInType.Type().Elem()
	fn := reflect.ValueOf(pairFunc)
	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}

	var ins [2]reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)
	out := fn.Call(ins[:])[0]

	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}

// 健壮版filter
func Filter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

var boolType = reflect.ValueOf(true).Type()

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {

	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType

	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}

	return out.Interface(), len(which)
}

func main() {
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)
	//old people: 2

	noVacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no vacation: %v\n", noVacation)
	//People no vacation: [{Hao 44 0 8000} {Jack 26 0 4000} {Marry 29 0 6000}]

	names := []string{"pony", "tony"}
	upCase := func(name string) string {
		return strings.ToUpper(name)
	}
	fmt.Println(Map(names, upCase))
	//[PONY TONY]

	res, num := filter(list, func(e Employee) bool {
		if e.Salary > 4000 {
			return true
		}
		return false
	}, false)
	fmt.Println(res, num)
}
