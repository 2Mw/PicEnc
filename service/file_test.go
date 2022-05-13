package service

import (
	"log"
	"testing"
)

var entry = &PicFile{
	dir:      "C:\\Users\\张三\\Desktop",
	filename: "2.jpg",
}

func TestReadImage(t *testing.T) {
	data, typ, err := entry.ReadImage()
	if err != nil {
		log.Print(err)
		return
	}

	_, err = entry.ExportFile(data, typ)
	if err != nil {
		return
	}

}
