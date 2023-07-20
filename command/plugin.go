package command

import "github.com/spf13/cobra"

var (
	VarStringName string
	VarStringPath string
)

type Arg struct {
	Name string
	Path string
}

func PluginCommand(_ *cobra.Command, _ []string) error {
	arg := Arg{
		Name: VarStringName,
		Path: VarStringPath,
	}

	err := GenAction(arg)
	if err != nil {
		return err
	}
	return nil
}
