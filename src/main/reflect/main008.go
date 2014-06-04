package main
//
//import (
//	"fmt"
//	. "model"
//	"unsafe"
//	"reflect"
//)
//
////type SliceHeader struct {
////	Data uintptr
////	Len  int
////	Cap  int
////}
//
//
//func main() {
//	var size int
//	var str string
//	size = int(unsafe.Sizeof(str))
//	fmt.Println(size) 	// 16
//
//
//	u := &User{} // User的所有成员变量都是private
//	fmt.Println("u :", u)
//	size= int(unsafe.Sizeof(*u))
//	fmt.Println(size)
//
//	// 可以将其理解为，指想固定地址，并有长度的[]byte
//	header := &reflect.SliceHeader {
//		Data : uintptr(unsafe.Pointer(u)),
//		Len	 : size,
//		Cap  : size,
//	}
//
//	b := *(*[]byte)(unsafe.Pointer(header))
//	println(len(b))
//	//12
//	fmt.Println(b)
//	//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
//	age := 27
//	size = int(unsafe.Sizeof(&age))
//	header = &reflect.SliceHeader {
//		Data : uintptr(unsafe.Pointer(&age)),
//		Len	 : size,
//		Cap  : size,
//	}
//	b2 := *(*[]byte)(unsafe.Pointer(header))
//	fmt.Println(b2)
//
//	b[16 + 16] = b2[0]
//	b[16 + 16 + 1] = b2[1]
//	b[16 + 16 + 2] = b2[2]
//	b[16 + 16 + 3] = b2[3]
//	b[16 + 16 + 4] = b2[4]
//	b[16 + 16 + 5] = b2[5]
//	b[16 + 16 + 5] = b2[6]
//	b[16 + 16 + 7] = b2[7]
//	fmt.Println(b)
//	fmt.Println("u :", u)
//
//
//}
