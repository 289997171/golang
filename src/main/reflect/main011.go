package main

//import (
//	"fmt"
//	"unsafe"
//	"model"
//	"reflect"
//)
//
//func main() {
//	u2 := model.User2{}
//	u2.Username.Strcpy("Vicky")
//	u2.Password.Strcpy("12345")
//	fmt.Println(unsafe.Sizeof(u2))
//	fmt.Println(u2.Username.String(), u2.Username)
//	fmt.Println(u2.Password.String(), u2.Password)
//
//	var typeof reflect.Type = reflect.TypeOf(u2)
//	field, found := typeof.FieldByName("Username")
//	if found {
//		fmt.Println("Index: ",field.Index)
//		fmt.Println("PkgPath: ",field.PkgPath)
//		fmt.Println("Name: ",field.Name)
//		fmt.Println("Offset: ",field.Offset)
//		fmt.Println("Type: ",field.Type)
//		fmt.Println("Tag: ",field.Tag)
//		fmt.Println("Anonymous: ",field.Anonymous)
//	}
//
//
//	fmt.Println("========================")
//
//	field, found = typeof.FieldByName("Password")
//	if found {
//		fmt.Println("Index: ", field.Index)
//		fmt.Println("PkgPath: ", field.PkgPath)
//		fmt.Println("Name: ", field.Name)
//		fmt.Println("Offset: ", field.Offset)
//		fmt.Println("Type: ", field.Type)
//		fmt.Println("Tag: ", field.Tag)
//		fmt.Println("Anonymous: ", field.Anonymous)
//	}
//}
