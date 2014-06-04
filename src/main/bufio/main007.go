package main

//import (
//	"os"
//	"fmt"
//	"bufio"
//)
//
//func main() {
//	var openFile,_ = os.OpenFile("study.iml", os.O_RDONLY, 0777)
//	var newReaderSize = bufio.NewReaderSize(openFile, 0)
//	fmt.Println(newReaderSize)
//	buf, _, err := newReaderSize.ReadLine()
//	fmt.Println(string(buf), err) // NewReaderSize 最小缓冲区为16字节，这里虽然调用的是readLine但读完16字节后就已经结束了，其缓冲区不会动态增加的
//}

