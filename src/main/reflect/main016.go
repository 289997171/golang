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
//	var f, ok = typeof.FieldByName("A2")
//	fmt.Println(ok, f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
//	//>>true [2]  A2 []uint8 标签2 false
//}
