package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
//	var value reflect.Value = reflect.ValueOf(a)
//	var vs = value.Slice(3, 8) //指定切片范围
//	fmt.Println(vs.Interface()) //做为接口类型读出
//	//>>[4 5 6 7 8]
//
//	var b = []interface{}{
//		"a",
//		'a',
//		1,
//		1.1,
//	}
//	var value1 reflect.Value = reflect.ValueOf(b)
//	var vs1 = value1.Slice(2, 3)
//	fmt.Println(vs1.Kind(), vs1.Interface())
//	//>>slice [1]
//
//}

