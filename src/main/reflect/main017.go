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
//func main(){
//	var a A
//	var typeof reflect.Type = reflect.TypeOf(a)
//
//	var fun = func(s string) bool {
//		if s == "A3" {
//			return true
//		}
//		return false
//	}
//	var f, ok = typeof.FieldByNameFunc(fun)
//	fmt.Println(ok, f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
//	//>>true [3]  A3 *main.A 指针指向自身 false
//}
//
