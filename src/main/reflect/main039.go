package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a chan int
//	fmt.Println(reflect.ValueOf(a).Cap())
//	var value reflect.Value = reflect.ValueOf(&a)
//	value = reflect.Indirect(value) //使指针指向内存地址
//	value = reflect.MakeChan(value.Type(), 99) // 相当与设置其为缓冲区
//	fmt.Println(value.Kind(), value.Cap())
//	//>>chan 99
//
//}
