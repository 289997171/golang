package main

//import (
//	"fmt"
//	"reflect"
//)
//
//type A struct {
//	A0 int "标签0"
//	A1 string "标签1"
//	A2 []byte "标签2"
//	A3 *A "指针指向自身"
//}
//
//func main() {
//	{
//		var a int
//		var typeof reflect.Type = reflect.TypeOf(a)
//		fmt.Println(typeof.Align())
//		//>>4
//	}
//
//	var a A
//	var typeof reflect.Type = reflect.TypeOf(a.A0)
//	fmt.Println(typeof.FieldAlign())
//	//>>4
//	typeof = reflect.TypeOf(a.A1)
//	fmt.Println(typeof.FieldAlign())
//	//>>4
//	typeof = reflect.TypeOf(a.A2)
//	fmt.Println(typeof.FieldAlign())
//	//>>4
//
//
//}
//
