package main

//import (
//	"bytes"
//	"bufio"
//	"fmt"
//)
//
//func main() {
//	var newBuffer_a *bytes.Buffer = bytes.NewBuffer([]byte("1234 5678 1234567901234567890 abcd ABCD"))
//	var newScanner_a *bufio.Scanner = bufio.NewScanner(newBuffer_a)
//	newScanner_a.Split(bufio.ScanWords)
//
//	for newScanner_a.Scan() { //Scan 每次都调用splitFunc函数
//		fmt.Println(newScanner_a.Text())
//	}
//	//1234
//	//5678
//	//1234567901234567890
//	//abcd
//	//ABCD
//
//	fmt.Println(newScanner_a.Err())
//	//<nil>
//
//}

