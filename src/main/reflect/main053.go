package main

//import (
//	"fmt"
//	"reflect"
//)
//
//type A struct {
//	A0 float64
//	A1 B
//}
//type B interface {}
//
//func main() {
//	a := A{3.14, nil}
//	var value reflect.Value = reflect.ValueOf(a)
//
//	fmt.Println(value.Field(1).InterfaceData())
//	//>>[0 0]
//
//	a.A1 = A{9999, nil}
//	fmt.Println(value.Field(1).InterfaceData())
//
//
//	fmt.Println("====================")
//
//	value = reflect.ValueOf(a)
//	fmt.Println(value.Field(1).InterfaceData())
//}
