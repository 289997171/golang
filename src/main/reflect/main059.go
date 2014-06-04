package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//
//	var a = map[string]int{
//		"a":1,
//		"b":2,
//		"c":3,
//		"d":4,
//		"e":5,
//	}
//	var value reflect.Value = reflect.ValueOf(&a)
//
//	//判断指针是否指向内存地址
//	if !value.CanSet() {
//		value = value.Elem() //使指针指向内存地址
//	}
//
//	keys := value.MapKeys()
//
//	fmt.Println(keys) //无顺序的返回Key
//	//>>[c e a d b]
//
//
//	for i := 0; i<len(keys); i++ {
//		mv := value.MapIndex(keys[i]) //使用Key名称指定返回值
//		fmt.Println(keys[i].String(), mv.Int())
//		// 0 >>c 3
//		// 1 >>d 4
//		// 2 >>e 5
//		// 3 >>a 1
//		// 4 >>b 2
//	}
//
//
//	value.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(123456)) //修改键值
//	value.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(nil)) //如果值是nil，则删除该键
//	value.SetMapIndex(reflect.ValueOf("c"), reflect.ValueOf(789)) //增加一个键
//
//	fmt.Println(value.Interface())
//	// map[a:123456 c:789 d:4 e:5]        b被删除
//
//}

