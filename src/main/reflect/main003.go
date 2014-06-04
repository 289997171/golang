package main
//
//import (
//	"fmt"
//)
//
//func main() {
//
//	var c = make(chan int, 1)
//	go func(c chan int) {
//		for i := 0; i < 10; i++ {
//			c<-i
//		}
//		close(c)
//	}(c)
//
//L:
//	for {
//		select {
//		case val, ok := <-c:
//			if ok {
//				fmt.Printf("接收: %d\n", val)
//			}else {
//				fmt.Println("信道关闭了")
//				break L
//			}
//		}
//	}
//
//}
