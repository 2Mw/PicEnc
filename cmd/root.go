package cmd

import (
	"PicEnc/config"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var root = &cobra.Command{}

func Execute() {
	if err := root.Execute(); err != nil {
		time.Sleep(config.WaitToExit)
		os.Exit(0)
	}
}
