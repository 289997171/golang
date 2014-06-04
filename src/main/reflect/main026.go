package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	// IsVariadic 返回true，如果一个函数类型的最后一个输入参数是一个“...”参数。
//	// 如果是这样，t.In(t.NumIn() - 1)返回的实际类型参数的隐式 []T。
//	var a func(a1 int, a2 string, a3...interface{})
//	var typeA reflect.Type = reflect.TypeOf(a)
//	fmt.Println(typeA.IsVariadic())
//	//>>true
//
//	var b func(a1 int, a2 string, a3 interface{})
//	var typeB reflect.Type = reflect.TypeOf(b)
//	fmt.Println(typeB.IsVariadic())
//	//>>false
//
//}
