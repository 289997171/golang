package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a map[int]string
//	var value reflect.Value = reflect.ValueOf(&a)
//	value = reflect.Indirect(value) //使指针指向内存地址
//	value = reflect.MakeMap(value.Type())
//	value.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf("a"))
//	value.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf("b"))
//	fmt.Println(value.Kind(), value.Interface())
//	//>>map map[1:a 2:b]
//}
//
