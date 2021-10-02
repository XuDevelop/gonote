package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// note.FileOperation()

	srcFilePath := "E:/develop/gonote/aa.txt"
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("打开文件失败！", err)
		return
	}
	defer func() {
		err = srcFile.Close()
		if err != nil {
			fmt.Println("关闭aa.txt文件出错！", err)
		}
	}()
	reader := bufio.NewReader(srcFile)
	destinationFilePath := "E:/develop/gonote/bb.txt"
	dstFile, err := os.OpenFile(destinationFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Println("创建文件失败！", err)
		return
	}
	defer func() {
		err = dstFile.Close()
		if err != nil {
			fmt.Println("关闭bb.txt文件出错！", err)
		}
	}()
	writer := bufio.NewWriter(dstFile)
	written, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println("拷贝文件失败！", err)
		return
	}
	fmt.Println("拷贝文件成功，拷贝了", written, "个字节")

}
