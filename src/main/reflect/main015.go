package main

//import (
//	"reflect"
//	"fmt"
//)
//
//type A struct {
//	A0 int "标签0"
//	A1 string "标签1"
//	A2 []byte "标签2"
//	A3 *A "指针指向自身"
//	B "匿名嵌套"
//}
//type B struct {
//	B0 int "B结构-标签B0"
//}
//
//func main(){
//	var a A
//	var typeof reflect.Type = reflect.TypeOf(a)
//
//	var r = []int{3,4}
//	var f = typeof.FieldByIndex(r)
//	fmt.Println(f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
//	//>>[4]  B main.B 匿名嵌套 true
//
//	r = []int{4,0}
//	f = typeof.FieldByIndex(r)
//	fmt.Println(f.Index, f.PkgPath, f.Name, f.Type, f.Tag, f.Anonymous)
//	//>>[0]  B0 int B结构-标签B0 false
//}
