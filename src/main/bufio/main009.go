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
//	var peek, err = newReadSize.Peek(2) // .Peek() 不会改变缓冲区的数据大小
//	fmt.Println(err, string(peek), newReadSize.Buffered())
//	//<nil> 12 10
//
//	peek, err = newReadSize.Peek(1000)
//	fmt.Println(err, string(peek), newReadSize.Buffered())
//	//EOF 1234567890 10
//
//}

