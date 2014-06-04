package main

//import (
//	"reflect"
//	"fmt"
//)
//
//
//type User struct {
//	Username string `k:"v" k2:"v2"`
//}
//
//func main() {
//	u := User{}
//	typeof := reflect.TypeOf(u)
//	tag := typeof.Field(0).Tag
//	fmt.Println("tag == ", tag)
//	v := tag.Get("k")
//	fmt.Println("v == ", v)
//
//	v = tag.Get("k2")
//	fmt.Println("v2 == ", v)
//}
