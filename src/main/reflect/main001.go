package main
//
//import (
//	"fmt"
//	"reflect"
//	"strconv"
//)
//
//type A struct {
//	A0 []int
//	A1 []int
//}
//
//
//type Obj struct {
//	Name  string
//	Id    int
//}
//
//func (this *Obj) String() string {
//	return this.Name + " " + strconv.Itoa(this.Id)
//}
//
//func main() {
//	var a A
//	a.A0 = append(a.A0, 1, 2, 3, 4, 5, 6, 7)
//	a.A1 = append(a.A1, 9, 8, 7, 6)
//	fmt.Println(a)
//	// Copy 复制src的内容复制到dst，直到dst已被填补满，或src已经耗尽。它返回复制的元素的数量。
//	// 每个 dst 和 src 的 Kind（样）都必须切片(Slice)“或”数组(Array)，dst和src必须具有相同的元素类型。
//	var n = reflect.Copy(reflect.ValueOf(a.A0), reflect.ValueOf(a.A1))
//	fmt.Println(n, "\t", a)
//	n = reflect.Copy(reflect.ValueOf(a.A1), reflect.ValueOf(a.A0))
//	fmt.Println(n, "\t", a)
//
//	fmt.Println("------------------------")
//	var o1, o2, o3, o4, o5 Obj
//	o1.Id = 1
//	o1.Name = "o1"
//	o2.Id = 2
//	o2.Name = "o2"
//	o3.Id = 3
//	o3.Name = "o3"
//	o4.Id = 4
//	o4.Name = "o"
//	o5.Id = 5
//	o5.Name = "o5"
//
//	s1 := make([]*Obj, 0)
//	s2 := make([]*Obj, 0)
//	s1 = append(s1, &o1, &o2, &o3)
//	s2 = append(s2, &o4, &o5)
//
//	fmt.Println(s1)
//	fmt.Println(s2)
//
//	fmt.Println("------------------------")
//	o1.Id = 999
//	fmt.Println(s1)
//	fmt.Println(s2)
//	fmt.Println("------------------------")
//
//	reflect.Copy(reflect.ValueOf(s2), reflect.ValueOf(s1))
//	fmt.Println(s1)
//	fmt.Println(s2)
//
//	s2[0].Id = 99999
//	fmt.Println(s1)
//	fmt.Println(s2)
//}
