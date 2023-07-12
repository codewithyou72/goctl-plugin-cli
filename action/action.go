package action

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"zero-swagger/generate"
)

var (
	VarStringHost     string
	VarStringBasePath string
	VarStringFileName string
)

func Generator(_ *cobra.Command, _ []string) error {
	fileName := VarStringFileName

	if len(fileName) == 0 {
		fileName = "rest.swagger.json"
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	basepath := VarStringBasePath
	host := VarStringHost
	return generate.Do(fileName, host, basepath, p)
}
