package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var Version = "20220621"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本号",
	Long:  `版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		verSion := fmt.Sprintf("%s %s/%s", Version, runtime.GOOS, runtime.GOARCH)
		fmt.Println(verSion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
