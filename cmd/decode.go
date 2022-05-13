package cmd

import (
	"PicEnc/config"
	"PicEnc/service"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var decode = &cobra.Command{
	Use:   "de",
	Short: "Decode a picture",
	Long:  "Decode a picture",
	Run: func(cmd *cobra.Command, args []string) {
		picFile, data, factory, _ := examine()
		cryptor, err := factory.NewChaoCrypt(key, encMode, encMethod)
		if err != nil {
			log.Fatal(err)
		}

		ret := cryptor.Decrypt(data)

		_, err = picFile.ExportFile(ret, "png")
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

func examine() (*service.PicFile, [][]uint32, *service.ChaosCryptFactory, string) {
	path, _ := filepath.Abs(file)
	idx := strings.LastIndex(path, string(os.PathSeparator))
	picFile := service.PicFile{
		Dir:      path[:idx],
		Filename: path[idx+1:],
	}

	data, typ, err := picFile.ReadImage()
	if err != nil {
		log.Fatalf("Read image error: %v", err)
	}

	switch mode {
	case 1:
		encMode = config.Row
	case 2:
		encMode = config.RowAndLine
	default:
		log.Fatalf("Error encrypt mode, only 1 or 2")
	}

	switch method {
	case "logistic":
		encMethod = config.Logistic
	default:
		log.Fatalf("Error encrypt mode, only logistic or sigmoid")
	}

	factory := service.ChaosCryptFactory{}

	return &picFile, data, &factory, typ
}

func init() {
	decode.Flags().Float64VarP(&key, "key", "k", 0.1, "The key to decrypt pictures. (Required)")
	decode.Flags().StringVarP(&file, "file", "f", "", "The filepath of picture. (Required)")
	decode.Flags().IntVarP(&mode, "mode", "m", 0, "The mode of decrypt. (Required, 1 - row, 2 - row and col)")
	decode.Flags().StringVarP(&method, "type", "t", "logistic", "The decrypt method. (Options: logistic)")
	_ = decode.MarkFlagRequired("file")
	_ = decode.MarkFlagRequired("key")
	_ = decode.MarkFlagRequired("mode")
}
