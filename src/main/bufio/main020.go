package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	rb := bytes.NewBuffer([]byte{})
//	bufioReader := bufio.NewReader(rb)
//	b := make([]byte, 512)
//
//	n, err := bufioReader.Read(b)   //如果不是把 r 设置进来，其实是没有数据可读的
//	fmt.Println(string(b[:n]), err)
//
//	r := bufio.NewReader(bytes.NewBuffer([]byte("1234567890")))
//	bufioReader.Reset(r) // 设置 r 为 bufioReader 的可读数据对象
//
//	n, err = bufioReader.Read(b)   //如果不是把 r 设置进来，其实是没有数据可读的
//	fmt.Println(string(b[:n]), err)
//	//1234567890 <nil>
//
//
//
//	r = bufio.NewReader(bytes.NewBuffer([]byte("abcdefg")))
//	bufioReader.Reset(r) // 设置 r 为 bufioReader 的可读数据对象
//
//	n, err = bufioReader.Read(b)   //如果不是把 r 设置进来，其实是没有数据可读的
//	fmt.Println(string(b[:n]), err)
//
//
//	fmt.Println("----------------------")
//
//	r = bufio.NewReader(bytes.NewBuffer([]byte("xyz")))
//	bufioReader.Reset(r) // 设置 r 为 bufioReader 的可读数据对象
//
//	r1 := bufio.NewReader(bytes.NewBuffer([]byte("lmn")))
//	bufioReader.Reset(r1) // 设置 r 为 bufioReader 的可读数据对象
//
//	n, err = bufioReader.Read(b)   //如果不是把 r 设置进来，其实是没有数据可读的
//	fmt.Println(string(b[:n]), err)
//	// lmn <nil> 只会打印 lmn
//
//}

