package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//type A struct {
//	A0 int
//}
//
//func (f A) test() {}
//func (f A) test1() {}
//func (f A) test2() {}
//
//func main() {
//	var a A
//	var typeof reflect.Type = reflect.TypeOf(a)
//	fmt.Println(typeof.NumMethod())
//	//>>3
//	m0 := typeof.Method(0)
//	fmt.Println(A.test, m0.Index, m0.PkgPath, m0.Name, m0.Type, m0.Func)
//	m1 := typeof.Method(1)
//	fmt.Println(A.test, m1.Index, m1.PkgPath, m1.Name, m1.Type, m1.Func)
//	m2 := typeof.Method(2)
//	fmt.Println(A.test, m2.Index, m2.PkgPath, m2.Name, m2.Type, m2.Func)
//}
