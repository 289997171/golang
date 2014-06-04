package main

//import (
//	"fmt"
//	"reflect"
//)
//
//type A struct {
//	C
//}
//type B struct {
//	C
//	D
//	E
//}
//type C interface {
//	test()
//}
//type D interface {
//	test1()
//}
//type E interface {}
//
//func main() {
//	var a A
//	var b B
//	var typeof reflect.Type = reflect.TypeOf(a)
//
//	//判断a结构是否存在b结构中的C接口
//	var booL = typeof.Implements(reflect.TypeOf(b).Field(0).Type)
//	fmt.Println(booL)
//	//>>true
//
//	//判断a结构是否存在b结构中的D接口
//	booL = typeof.Implements(reflect.TypeOf(b).Field(1).Type)
//	fmt.Println(booL)
//	//>>false
//
//	//b结构中的E因为是空接口，所以也返回true。
//	booL = typeof.Implements(reflect.TypeOf(b).Field(2).Type)
//	fmt.Println(booL)
//	//>>true
//}
//
