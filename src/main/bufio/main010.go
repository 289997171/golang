package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	var rb = bytes.NewBuffer([]byte("1234567890"))
//	var newReadSize = bufio.NewReaderSize(rb, 1000000)
//	var (
//		buf = make([]byte, 3)
//		n int
//		err error
//	)
//
//	for  {
//		n, err = newReadSize.Read(buf)
//		fmt.Println(n, err, string(buf[:n]))
//		if err != nil {
//			break
//		}
//	}
//	//3 <nil> 123
//	//3 <nil> 456
//	//3 <nil> 789
//	//1 <nil> 0
//	//0 EOF
//
//}

