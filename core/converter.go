package core

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

// GetBase64 returns base64 of named image
func GetBase64(imageName string) (imgBase64Str string) {
	imgFile, err := os.Open("D:/Projects/imaxel/core/images/" + imageName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str = base64.StdEncoding.EncodeToString(buf)

	return
}
