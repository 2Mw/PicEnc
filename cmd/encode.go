package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var encode = &cobra.Command{
	Use:   "en",
	Short: "Encode a picture",
	Long:  "Encode a picture",
	Run: func(cmd *cobra.Command, args []string) {
		picFile, data, factory, typ := examine()
		cryptor, err := factory.NewChaoCrypt(key, encMode, encMethod)
		if err != nil {
			log.Fatal(err)
		}

		ret := cryptor.Encrypt(data)

		_, err = picFile.ExportFile(ret, typ)
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	encode.Flags().Float64VarP(&key, "key", "k", 0.1, "The key to encrypt pictures. (Required)")
	encode.Flags().StringVarP(&file, "file", "f", "", "The filepath of picture. (Required)")
	encode.Flags().IntVarP(&mode, "mode", "m", 0, "The mode of encrypt. (Required, 1 - row, 2 - row and col)")
	encode.Flags().StringVarP(&method, "type", "t", "logistic", "The encrypt type. (Options: logistic)")
	_ = encode.MarkFlagRequired("file")
	_ = encode.MarkFlagRequired("key")
	_ = encode.MarkFlagRequired("mode")
}
