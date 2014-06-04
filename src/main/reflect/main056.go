package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a int
//	var value reflect.Value = reflect.ValueOf(&a)
//	fmt.Println(value.CanSet())
//	//>>false
//
//	if !value.CanSet() {
//		value = value.Elem() //使指针指向内存地址
//	}
//
//	fmt.Println(value.CanSet())
//	//>>true
//
//	if (value.CanSet()) {
//		value.Set(reflect.ValueOf(9999))
//	}
//
//	fmt.Println(a)
//
//}

