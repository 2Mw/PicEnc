package service

import (
	"log"
	"path/filepath"
	"strings"
	"testing"
)

var entry = &PicFile{
	Dir:      "C:\\Users\\张三\\Desktop",
	Filename: "2.jpg",
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

func TestSplit(t *testing.T) {
	s := "asdsadsad"
	log.Printf("%v", strings.Split(s, "kkk"))
}

func TestFile(t *testing.T) {
	name := "C:\\Users\\张三\\Desktop\\2.jpg"
	//file, err := os.Open(name)
	//if err != nil {
	//	return
	//}
	//log.Printf("Name: %v", file.Name())
	//filepath.Abs(name)
	path, err := filepath.Abs(name)
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%v", path)
}
