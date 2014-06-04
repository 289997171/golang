package main

//import (
//	"fmt"
//	"bufio"
//)
//
//func main() {
//	var (
//		b []byte
//		advance int
//		token []byte
//		err error
//	)
//	b = []byte("中国人说中国话\r\n123456789")
//	advance, token, err = bufio.ScanLines(b, false)
//	fmt.Println(advance, string(token), err)
//	//23 中国人说中国话 <nil>
//	advance, token, err = bufio.ScanLines(b, false)
//	fmt.Println(advance, string(token), err)
//	//23 中国人说中国话 <nil>
//	advance, token, err = bufio.ScanLines(b, true)
//	fmt.Println(advance, string(token), err)
//	//23 中国人说中国话 <nil>
//	advance, token, err = bufio.ScanLines(b, true)
//	fmt.Println(advance, string(token), err)
//	//23 中国人说中国话 <nil>
//
//	b = []byte("")
//	fmt.Println(len(b))
//	advance, token, err = bufio.ScanLines(b, true)
//	fmt.Println(advance, string(token), err)
//	//0  <nil>
//	advance, token, err = bufio.ScanLines(b, false)
//	fmt.Println(advance, string(token), err)
//	//0  <nil>
//
//}

