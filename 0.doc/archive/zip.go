package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

/**
https://golang.org/pkg/archive/zip

Package zip provides support for reading and writing ZIP archives.

为了向后兼容，文件头包含32位和64位字段。64位字段总是包含正确的值，对于普通存档，这两个字段是相同的。对于需要ZIP64格式的文件，32位字段将是0xffffffff，必须使用64位字段。
*/

// 打开一个 zip 文件, 并循环里面文件和文件内容
func main() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("test.zip")
	checkError(err)
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	// TODO: 为什么 r.File 写成 r.Reader.File 也是一样
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)

		rc, err := f.Open()
		checkError(err)

		// TODO: 为什么会打印? 多个文件如何读取? 中文乱码怎么解决?
		_, err = io.CopyN(os.Stdin, rc, 68)
		fmt.Println()
		checkError(err)

		rc.Close()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
