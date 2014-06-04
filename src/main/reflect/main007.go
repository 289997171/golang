package main
//
//import (
//	"fmt"
//	"unsafe"
//	"reflect"
//)
//
//func main() {
//	a := "abc"    // 字符串 "abc"
//	b := (*uintptr)(unsafe.Pointer(&a))  // b 存储 a 的地址
//
//	var c []byte // d 将 c 的结构用 reflect.SliceHeader 表示
//	d := (*reflect.SliceHeader)((unsafe.Pointer(&c)))
//	d.Cap = len(a)
//	d.Len = len(a)
//	d.Data = *b  // *b 存储字符串首元素地址
//	fmt.Println(c)
//	// [97 98 99]
//	fmt.Println(string(c))
//	// abc
//}
