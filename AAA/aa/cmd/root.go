package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zeromicro/goctl-plugin-cli/tpl"
	"github.com/zeromicro/goctl-plugin-cli/variable"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

const (
	codeFailure = 1
	dash        = "-"
	doubleDash  = "--"
	assign      = "="
)

var rootCmd = &cobra.Command{
	Use:   "git-version",
	Short: "根据goctl生成swagger",
	Long:  `根据.api文件生成swagger`,
}

func Execute() {
	os.Args = supportGoStdFlag(os.Args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}

func init() {
	cobra.AddTemplateFuncs(template.FuncMap{
		"blue":    blue,
		"green":   green,
		"rpadx":   rpadx,
		"rainbow": rainbow,
	})

	rootCmd.SetUsageTemplate(tpl.UsageTpl)
	cobra.OnInitialize(initConfig) //初始化配置文件

	rootCmd.Version = fmt.Sprintf("%s %s/%s", variable.BuildVersion, runtime.GOOS, runtime.GOARCH)
	rootCmd.AddCommand(versionCmd)
	//rootCmd.AddCommand(PluginCmd)

}

func supportGoStdFlag(args []string) []string {
	copyArgs := append([]string(nil), args...)
	parentCmd, _, err := rootCmd.Traverse(args[:1])
	if err != nil { // ignore it to let cobra handle the error.
		return copyArgs
	}

	//处理copyArgs[0:]的第2个参数之后的数据
	for idx, arg := range copyArgs[0:] {
		parentCmd, _, err = parentCmd.Traverse([]string{arg})
		if err != nil { // ignore it to let cobra handle the error.
			break
		}
		if !strings.HasPrefix(arg, dash) { //没有-符号直接不处理这个参数
			continue
		}

		flagExpr := strings.TrimPrefix(arg, doubleDash)
		flagExpr = strings.TrimPrefix(flagExpr, dash)
		flagName, flagValue := flagExpr, ""
		assignIndex := strings.Index(flagExpr, assign)
		if assignIndex > 0 {
			flagName = flagExpr[:assignIndex]
			flagValue = flagExpr[assignIndex:]
		}

		//跳过version和help命令
		if !isBuiltin(flagName) {
			// The method Flag can only match the user custom flags.
			f := parentCmd.Flag(flagName)
			if f == nil {
				continue
			}
			if f.Shorthand == flagName {
				continue
			}
		}

		goStyleFlag := doubleDash + flagName //命令行参数=之前的东西
		if assignIndex > 0 {                 //=命令行之后
			goStyleFlag += flagValue
		}

		copyArgs[idx] = goStyleFlag
	}
	return copyArgs
}

// 判断是否是其中2个命令
func isBuiltin(name string) bool {
	return name == "version" || name == "help"
}

var ConfigFile string

// 初始化viper读取常用配置信息
func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, variable.ViperConfigFileVariableName))

		viper.SetConfigType(variable.ViperConfigSuffixFileVariableName)
		viper.SetConfigName(variable.ViperConfigDirVariableName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(aurora.Green(fmt.Sprintf("读取初始化配置文件成功:%s", viper.ConfigFileUsed())))
	} else {
		//没有参数就设置默认参数
		fmt.Println(aurora.Green(fmt.Sprintf("【非报错内容】没有设置配置文件err=%+v直接使用默认值::%s=%+v", err, variable.ViperStyleVariableName, variable.VarStringStyle)))
		viper.SetDefault(variable.ViperStyleVariableName, variable.VarStringStyle)
	}
}
