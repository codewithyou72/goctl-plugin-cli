package cmd

import (
	"github.com/zeromicro/goctl-plugin-cli/command"

	"github.com/spf13/cobra"
)

var PluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "生成插件命令",
	Long:  `生成插件命令`,
	RunE:  command.PluginCommand,
}

func init() {

	//go run sql3c.go proto create --package="demo" --go-package="demo_pb" --service-name="Demo" --source-paths="admin-gateway/demo1.proto" --source-paths="admin-gateway/demo2.proto" --path="admin-gateway/demo.proto"
	//proto create --package="demo" --go-package="demo_pb" --service-name="Demo" --source-paths="admin-gateway/demo1.proto" --source-paths="admin-gateway/demo2.proto" --path="admin-gateway/demo.proto"

	PluginCmd.Flags().StringVar(&command.VarStringSyntax, "syntax", "proto3", "syntax版本，默认proto3")
	PluginCmd.Flags().StringVar(&command.VarStringPackage, "package", "", "包名")
	PluginCmd.Flags().StringVar(&command.VarStringGoPackage, "go-package", "", "go-package名称")
	PluginCmd.Flags().StringVar(&command.VarStringServiceName, "service-name", "", "服务名称")
	PluginCmd.Flags().StringSliceVar(&command.VarStringSliceSourcePath, "source-paths", []string{}, "需要合并的proto文件")
	PluginCmd.Flags().StringVar(&command.VarStringPath, "path", "", "最终生成的文件名")
	rootCmd.AddCommand(PluginCmd)
}
