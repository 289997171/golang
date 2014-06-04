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
//	b = []byte("1234567890")
//	advance, token, err = bufio.ScanBytes(b, false)
//	fmt.Println(advance, token, err)
//	//1 [49] <nil>
//	advance, token, err = bufio.ScanBytes(b, false)
//	fmt.Println(advance, token, err)
//	//1 [49] <nil>
//	advance, token, err = bufio.ScanBytes(b, true)
//	fmt.Println(advance, token, err)
//	//1 [49] <nil>
//	advance, token, err = bufio.ScanBytes(b, true)
//	fmt.Println(advance, token, err)
//	//1 [49] <nil>
//
//	b = []byte{}
//	advance, token, err = bufio.ScanBytes(b, true)
//	fmt.Println(advance, token, err)
//	//0 [] <nil>
//	advance, token, err = bufio.ScanBytes(b, false)
//	fmt.Println(advance, token, err)
//	//panic: runtime error: slice bounds out of range
//	//
//	//goroutine 1 [running]:
//	//main.main()
//	//  G:/学习/go/test.go:38 +0xa8e
//	//exit status 2
//
//}

