package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//type A struct {
//}
//
//func (A) test() {
//	fmt.Println("test")
//}
//
//func (this *A) test2(str string) {
//	fmt.Println("test2", str)
//}
//
//func main() {
//
//	var str string
//	var kind reflect.Kind = reflect.TypeOf(str).Kind()
//	fmt.Println(kind, kind == reflect.String, kind == reflect.Int)
//	//>>string true false
//
//	var a A
//	var method reflect.Method = reflect.TypeOf(a).Method(0) //0 表示是 a 结构中的第几位函数
//	fmt.Println(method.Index, method.Name, method.PkgPath, method.Type, method.Func)
//	//>>0 test main func(main.A) <func(main.A) Value>
//
//	method = reflect.TypeOf(&a).Method(1) // 注意是&a
//	fmt.Println(method.Index, method.Name, method.PkgPath, method.Type, method.Func)
//	//>>1 test2 main func(*main.A) <func(*main.A) Value>
//}
