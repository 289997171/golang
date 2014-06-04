package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a = func(a1 int, a2 string, a3 ...float64) (int, string, []float64){
//		fmt.Println(a1, a2, a3)
//		//>>123 456 [1.1 2.2 3.3 4.4 5 5]
//		return a1, a2, a3
//	}
//
//	var value reflect.Value = reflect.ValueOf(a)
//	var tt []reflect.Value
//	tt = append(tt, reflect.ValueOf(123))
//	tt = append(tt, reflect.ValueOf("456"))
//	tt = append(tt, reflect.ValueOf([]float64{1.1, 2.2, 3.3, 4.4, 5,5}))
//	var tt1 []reflect.Value = value.CallSlice(tt)
//	fmt.Println(tt1)
//	//>>[<int Value> 456 <[]float64 Value>]
//
//}

