package main


//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a []int
//
//	//例子1
//	var value reflect.Value = reflect.ValueOf(&a)
//	fmt.Println(value.Kind())
//	//>>ptr
//
//	//判断指针是否指向内存地址
//	if !value.CanSet() {
//		value = value.Elem() //使指针指向内存地址
//	}
//	fmt.Println(value.Kind())
//	//>>slice
//
//	//例子2
//	var value1 reflect.Value = reflect.ValueOf(&a)
//	fmt.Println(value1.Kind())
//	//>>ptr
//
//	value1 = reflect.Indirect(value1) //功能等介于上面例子
//
//	fmt.Println(value1.Kind())
//	//>>slice
//
//}
