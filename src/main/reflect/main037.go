package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//
//	var a []int
//	var value reflect.Value = reflect.ValueOf(&a)
//
//	//判断指针是否指向内存地址
//	if !value.CanSet() {
//		value = value.Elem() //使指针指向内存地址
//	}
//
//	value = reflect.AppendSlice(value, reflect.ValueOf([]int{1, 2})) //支持切片
//	fmt.Println(value.Kind(), value.Slice(0, value.Len()).Interface())
//
//	value = reflect.AppendSlice(value, reflect.ValueOf([]int{3, 4, 5, 6, 7, 8, 9})) //支持切片
//	fmt.Println(value.Kind(), value.Slice(0, value.Len()).Interface())
//	//>>slice [1 2 3 4 5 6 7 8 9]
//
//}
