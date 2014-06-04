package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a chan int
//	var value reflect.Value = reflect.ValueOf(&a)
//	fmt.Println(value.IsNil())
//	//>>false
//
//	var b []uintptr
//	var value1 reflect.Value = reflect.ValueOf(b)
//	fmt.Println(value1.IsNil())
//	//>>true
//
//	b = append(b, 123) //赋值 123
//	value1 = reflect.ValueOf(b)
//	fmt.Println(value1.IsNil()) // 赋值后，不再是空值
//	//>>false
//
//	var c func()
//	var value2 reflect.Value = reflect.ValueOf(c)
//	fmt.Println(value2.IsNil())
//	//>>true
//
//	var d map[int]string
//	var value3 reflect.Value = reflect.ValueOf(d)
//	fmt.Println(value3.IsNil())
//	//>>true
//
//}

