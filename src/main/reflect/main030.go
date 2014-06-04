package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var i int
//	var recvDir reflect.ChanDir = reflect.RecvDir
//	var chanOf reflect.Type = reflect.ChanOf(recvDir, reflect.TypeOf(i))
//	fmt.Println(chanOf.Kind(), chanOf.ChanDir(), chanOf.String())
//	//>>chan chan <-chan int
//
//	var i1 int
//	var recvDir1 reflect.ChanDir = reflect.SendDir
//	var chanOf1 reflect.Type = reflect.ChanOf(recvDir1, reflect.TypeOf(i1))
//	fmt.Println(chanOf1.Kind(), chanOf1.ChanDir(), chanOf1.String())
//	//>>chan chan chan<- int
//
//	var i2 int
//	var recvDir2 reflect.ChanDir = reflect.BothDir
//	var chanOf2 reflect.Type = reflect.ChanOf(recvDir2, reflect.TypeOf(i2))
//	fmt.Println(chanOf2.Kind(), chanOf2.ChanDir(), chanOf2.String())
//	//>>chan chan chan int
//
//	var i3 string
//	var recvDir3 reflect.ChanDir = reflect.BothDir
//	var chanOf3 reflect.Type = reflect.ChanOf(recvDir3, reflect.TypeOf(i3))
//	fmt.Println(chanOf3.Kind(), chanOf3.ChanDir(), chanOf3.String())
//	//>>chan chan chan string
//
//
//	var c chan string
//	fmt.Println("xxx", reflect.ChanOf(reflect.TypeOf(c).ChanDir(), reflect.TypeOf("")))
//}
//
