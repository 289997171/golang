package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a = func(a1 int, a2 string, a3 float64) (int, string, float64) {
//		fmt.Println(a1, a2, a3)
//		//>>123 456 123.456
//		return a1, a2, a3
//	}
//
//	var value reflect.Value = reflect.ValueOf(a)
//	fmt.Println("参数：", value.Type().NumIn())
//	fmt.Println("返回：", value.Type().NumOut())
//	var tt []reflect.Value
//	tt = append(tt, reflect.ValueOf(123))
//	tt = append(tt, reflect.ValueOf("456"))
//	tt = append(tt, reflect.ValueOf(123.456))
//	var tt1 []reflect.Value = value.Call(tt)
//	fmt.Println(tt1[0], tt1[1], tt1[2])
//	//>><int Value> 456 <float64 Value>
//
//}
