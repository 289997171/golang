package main

//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	a := make(chan int,1)
//	var value reflect.Value = reflect.ValueOf(&a)
//
//	//判断指针是否指向内存地址
//	if !value.CanSet() {
//		value = value.Elem() //使指针指向内存地址
//	}
//	var gotest = func(c reflect.Value){
//		c.Send(reflect.ValueOf(3888)) //3>阻塞发送信息
//TRYSEND:
//		ok := c.TrySend(reflect.ValueOf(38881))
//		fmt.Println("非阻塞发送:", ok)
//		if !ok {
//			goto TRYSEND
//		}
//	}
//	go gotest(value) //1>异步
//
//	var c, ok = value.Recv() //2>阻塞等待
//
//	fmt.Println(c.Int(), ok) //4>读出数据
//	//>>3888 true
//
//TRYRECV:
//	c, ok = value.TryRecv() //2>非阻塞等待
//	fmt.Println("非阻塞接收:", ok)
//	if ok {
//		fmt.Println(c.Int(), ok) //4>读出数据
//	} else {
//		goto TRYRECV
//	}
//}

