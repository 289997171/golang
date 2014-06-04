package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a int
//	var typeA reflect.Type = reflect.TypeOf(&a)
//	fmt.Println(typeA, typeA.Kind()) // 这里是指针，间接指向内存地址
//	//>>*int ptr
//
//	var b int
//	var typeB reflect.Type = reflect.TypeOf(&b)
//	fmt.Println(typeB.Elem(), typeB.Elem().Kind()) // 加上.Elem() 就可以直接指向内存地址。Elem()很有用，很多时候都用的到它。
//	//>>int int
//
//	fmt.Println("=================")
//
//	var c map[int]string
//	var typeC reflect.Type = reflect.TypeOf(c)
//	fmt.Println(typeC,"\t", typeC.Kind())
//	fmt.Println(typeC.Elem(),"\t", typeC.Elem().Kind())
//	//>>string string
//
//	fmt.Println("=================")
//
//	var d map[int]string
//	var typeD reflect.Type = reflect.TypeOf(&d)
//	fmt.Println(typeD,"\t", typeD.Elem(),"\t", typeD.Kind())
//	fmt.Println(typeD.Elem(),"\t", typeD.Elem().Elem(),"\t", typeD.Elem().Kind())
//	//>>map[int]string string map
//
//}

