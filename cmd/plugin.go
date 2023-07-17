package cmd

import (
	"github.com/zeromicro/goctl-plugin-cli/action"

	"github.com/spf13/cobra"
)

var PluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "生成插件命令",
	Long:  `生成插件命令`,
	RunE:  action.PluginCommand,
}
