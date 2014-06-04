package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var a chan<- int
//	var typeA reflect.Type = reflect.TypeOf(a)
//	fmt.Println(typeA.ChanDir(), typeA.ChanDir() == reflect.SendDir)
//	//>>chan<- true
//
//	var b <-chan int
//	var typeB reflect.Type = reflect.TypeOf(b)
//	fmt.Println(typeB.ChanDir(), typeB.ChanDir() == reflect.RecvDir)
//	//>><-chan true
//
//	var c chan int
//	var typeC reflect.Type = reflect.TypeOf(c)
//	fmt.Println(typeC.ChanDir(), typeC.ChanDir() == reflect.BothDir)
//	//>>chan true
//
//}
//
