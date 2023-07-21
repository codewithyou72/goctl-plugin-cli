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

	PluginCmd.Flags().StringVar(&command.VarStringName, "name", "", "插件名称")
	PluginCmd.Flags().StringVar(&command.VarStringModule, "module", "", "module参数")
	PluginCmd.Flags().StringVar(&command.VarStringShort, "Short", "", "简短说明")
	PluginCmd.Flags().StringVar(&command.VarStringModule, "Long", "", "详细说明")
	PluginCmd.Flags().StringVar(&command.VarStringPath, "path", "", "路径")
	rootCmd.AddCommand(PluginCmd)
}
