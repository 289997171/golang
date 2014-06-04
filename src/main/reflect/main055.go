package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a  int
//	a = 999
//	var value reflect.Value = reflect.ValueOf(&a)
//	fmt.Println(reflect.ValueOf(&a), reflect.ValueOf(a))
//	fmt.Println(value.CanAddr(), value.Elem().CanAddr())
//	//>>false true
//
//	if value.Elem().CanAddr() {
//		fmt.Println(value.Elem().Addr(), value.Elem().Kind(), value.Elem().Int())
//
//	}
//}
