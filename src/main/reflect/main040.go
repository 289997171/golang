package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var swap = func(in []reflect.Value) []reflect.Value {
//		fmt.Println(in[0].Interface(), in[1].Interface())
//		//1 2
//		return []reflect.Value{in[1], in[0]}
//	}
//	var makeSwap = func(fptr interface{}) {
//		var valueOf reflect.Value = reflect.Indirect( reflect.ValueOf(fptr))
//		var v reflect.Value = reflect.MakeFunc(valueOf.Type(), swap)
//		valueOf.Set(v)
//
//		fmt.Println(valueOf)
//		//<func(int, int) (int, int) Value>
//		fmt.Println(v)
//		//<func(int, int) (int, int) Value>
//	}
//	var intSwap func(int, int) (int, int)
//	makeSwap(&intSwap)
//
//	fmt.Println(swap, intSwap)
//
//	fmt.Println(intSwap(1, 2))
//	//2 1
//
//}
