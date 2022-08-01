package core

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

type Image struct {
	Name  string
	Bytes []byte
}

func (image Image) GetBase64() (base64Str string, err error) {

	if image.Bytes == nil {
		err = errors.New("ERROR: Image.Bytes is nil")

		return
	}

	// convert the bytes to base64 string
	base64Str = base64.StdEncoding.EncodeToString(image.Bytes)

	return
}

func (image Image) WithBytes() (result Image) {

	if image.Bytes != nil {
		result = image

		return
	}

	// get image from core/images/ path
	imgFile, err := os.Open("D:/Projects/imaxel/core/images/" + image.Name)

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

	result = Image{
		Name:  image.Name,
		Bytes: buf,
	}

	return
}
