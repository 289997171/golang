package main

//import (
//	"fmt"
//	"reflect"
//)
//
//type A struct {
//	A0 int "这是A0"
//	A1 B "嵌套B结构"
//	B "匿名字段"
//}
//type B struct {
//	B0 string
//	B1 int
//}
//
//func main() {
//	var a A
//	var typeof reflect.Type = reflect.TypeOf(a)
//	var field reflect.StructField = typeof.Field(0) //0 表示是 a 结构中的第几位字段
//	fmt.Println("Index: ",field.Index)
//	fmt.Println("PkgPath: ",field.PkgPath)
//	fmt.Println("Name: ",field.Name)
//	fmt.Println("Offset: ",field.Offset)
//	fmt.Println("Type: ",field.Type)
//	fmt.Println("Tag: ",field.Tag)
//	fmt.Println("Anonymous: ",field.Anonymous)
//
//	fmt.Println("==========================")
//	field = typeof.Field(1)
//	fmt.Println("Index: ",field.Index)
//	fmt.Println("PkgPath: ",field.PkgPath)
//	fmt.Println("Name: ",field.Name)
//	fmt.Println("Offset: ",field.Offset)
//	fmt.Println("Type: ",field.Type)
//	fmt.Println("Tag: ",field.Tag)
//	fmt.Println("Anonymous: ",field.Anonymous)
//
//	fmt.Println("==========================")
//	field = typeof.Field(2)
//	fmt.Println("Index: ",field.Index)
//	fmt.Println("PkgPath: ",field.PkgPath)
//	fmt.Println("Name: ",field.Name)
//	fmt.Println("Offset: ",field.Offset)
//	fmt.Println("Type: ",field.Type)
//	fmt.Println("Tag: ",field.Tag)
//	fmt.Println("Anonymous: ",field.Anonymous)
//}
