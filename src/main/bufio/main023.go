package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	var rb = bytes.NewBuffer([]byte("1234567890,abcd,Abcd,ABCD,ABcd"))
//	var newWriterSize = bufio.NewWriterSize(rb, 10)
//	fmt.Println(newWriterSize.Available())
//	//10
//	var (
//		b = []byte("1234567890")
//	)
//	fmt.Println(newWriterSize.Buffered())
//	//0
//	newWriterSize.Write(b)
//	newWriterSize.Flush()
//	fmt.Println(newWriterSize.Buffered())
//	//0
//	fmt.Println(newWriterSize.Available())
//	//10
//	fmt.Println(rb)
//	//1234567890,abcd,Abcd,ABCD,ABcd1234567890
//
//}

