package main

//import (
//	"fmt"
//	"io"
//	"bufio"
//	"os"
//	_ "errors"
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
//			if string(buf[:n]) == "exit" {
//				// reader.CloseWithError(errors.New("close by exit"))
//				reader.Close()
//				break
//			}
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
//		_, err = ioPipeWriter.Write(line)
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//	}
//
//	time.Sleep(time.Second * 10)
//
//}

