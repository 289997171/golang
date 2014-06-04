package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//
//	var a int32 //注意：类型是32位
//	var value reflect.Value = reflect.ValueOf(&a)
//
//	//判断指针是否指向内存地址
//	if !value.CanSet() {
//		value = value.Elem() //使指针指向内存地址
//	}
//
//	var i int64 = 2147483647
//	if value.OverflowInt(i) {
//		fmt.Println("溢出：超出范围")
//	}else {
//		fmt.Println("正常：范围之内")
//	}
//	//>>正常：范围之内
//
//	i = 2147483648
//	if value.OverflowInt(i) {
//		fmt.Println("溢出：超出范围")
//	}else {
//		fmt.Println("正常：范围之内")
//	}
//	//>>溢出：超出范围
//
//}
