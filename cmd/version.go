package cmd

import (
	"fmt"
	"goctl-plugin-cli/variable"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本号",
	Long:  `版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		verSion := fmt.Sprintf("%s %s/%s", variable.BuildVersion, runtime.GOOS, runtime.GOARCH)
		fmt.Println(verSion)
	},
}
