package main

//import (
//	"fmt"
//	"io"
//	"bufio"
//	"os"
//	"errors"
//	"time"
//)
//
//func main() {
//	ioPipeReader, ioPipeWriter := io.Pipe()
//	go func(reader *io.PipeReader) {
//		defer reader.Close()
//		buf := make([]byte, 64)
//		for {
//			n, err := reader.Read(buf)
//			if err != nil {
//				fmt.Println("error : ", err)
//				break
//			}
//			fmt.Println(string(buf[:n]))
//		}
//		fmt.Println("Reader Close !")
//	}(ioPipeReader)
//	reader := bufio.NewReader(os.Stdin)
//	if reader == nil {
//		os.Exit(1)
//	}
//	for {
//		// io.Copy(ioPipeWriter, os.Stdin)
//		line, _, err := reader.ReadLine()
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//
//		if string(line) == "exit" {
//			fmt.Println("Writer Close")
//			ioPipeWriter.CloseWithError(errors.New("exit"))
//			break
//		}
//		ioPipeWriter.Write(line)
//	}
//
//	time.Sleep(time.Second * 10)
//
//}

