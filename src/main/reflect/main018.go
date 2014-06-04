package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a = func(a1 int, a2 string, a3 uint, a4 interface{}) {}
//	var typeof reflect.Type = reflect.TypeOf(a)
//	fmt.Println(typeof.NumIn())
//	//>>4
//
//
//	for i:=0; i<typeof.NumIn(); i++ {
//		fmt.Println(typeof.In(i).Kind())
//		// 0 >>int
//		// 1 >>string
//		// 2 >>uint
//		// 3 >>interface
//	}
//
//	fmt.Println("=========================")
//	var a1 = func() (int, string, uint, float64) {
//		return 1, "1", 2, 1.2
//	}
//	var typeof2 reflect.Type = reflect.TypeOf(a1)
//	fmt.Println(typeof2.NumOut())
//	//>>4
//	for i:=0; i<typeof2.NumOut(); i++ {
//		fmt.Println(typeof2.Out(i).Kind())
//		// 0 >>int
//		// 1 >>string
//		// 2 >>uint
//		// 3 >>float64
//	}
//
//
//}

