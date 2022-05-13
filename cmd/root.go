package cmd

import (
	"PicEnc/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var root = &cobra.Command{
	Use:              "",
	Short:            "Short text for program",
	Long:             "Illustration of PicEnc",
	TraverseChildren: true, // 允许在主命令上使用flag
	Run: func(cmd *cobra.Command, args []string) {
		//log.Printf("%v", args)
		if len(args) == 0 {
			log.Printf("Please use `-h` to see the help text.")
		}
	},
}

var (
	key       float64
	file      string
	method    string
	mode      int
	encMode   config.EncryptMode
	encMethod config.EncryptMethodType
)

func Execute() {
	if err := root.Execute(); err != nil {
		time.Sleep(config.WaitToExit)
		os.Exit(0)
	}
}

func init() {
	// 关闭cmd提示文本
	root.AddCommand(encode, decode, version)
	cobra.MousetrapHelpText = ""
}
