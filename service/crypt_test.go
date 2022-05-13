package service

import (
	"PicEnc/config"
	"log"
	"testing"
)

var filename = "2.jpg"
var mode = config.RowAndLine

func TestEncrypt(t *testing.T) {
	var file = &PicFile{
		Dir:      "C:\\Users\\张三\\Desktop",
		Filename: filename,
	}
	factory := ChaosCryptFactory{}
	c, _ := factory.NewChaoCrypt(0.6, mode, config.Logistic)

	data, _, err := file.ReadImage()
	if err != nil {
		log.Print(err)
		return
	}

	data = c.Encrypt(data)
	//data = c.Decrypt(data)

	_, err = file.ExportFile(data, "png")
	if err != nil {
		return
	}
}

func TestDecrypt(t *testing.T) {
	var file = &PicFile{
		Dir:      "C:\\Users\\张三\\Desktop",
		Filename: "o_" + filename,
	}
	factory := ChaosCryptFactory{}
	c, _ := factory.NewChaoCrypt(0.6, mode, config.Logistic)

	data, typ, err := file.ReadImage()
	if err != nil {
		log.Print(err)
		return
	}

	//data = c.Encrypt(data)
	data = c.Decrypt(data)

	_, err = file.ExportFile(data, typ)
	if err != nil {
		return
	}
}
