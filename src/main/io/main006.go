package main

//import (
//	"io"
//	"fmt"
//)
//
//func main(){
//	ioPipeReader, ioPipeWriter := io.Pipe()
//	go func(ioPipeReader *io.PipeReader){
//		//读取数据
//		b := make([]byte, 50)
//		n, err := ioPipeReader.Read(b)
//		fmt.Println(n, err)
//		//16 <nil>
//		ioPipeReader.Close()
//		fmt.Println(b)
//		//[229 134 153 229 133 165 49 50 51 52 53 54 55 56 57 48 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
//	}(ioPipeReader)
//
//	//写入数据
//	data := []byte("写入1234567890")
//	ioPipeWriter.Write(data)
//	ioPipeWriter.Close()
//}
//
