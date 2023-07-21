package command

import "github.com/spf13/cobra"

var (
	VarStringName   string
	VarStringModule string
	VarStringPath   string
	VarStringShort  string
	VarStringLong   string
)

type Arg struct {
	Name   string
	Path   string
	Module string
}

func PluginCommand(_ *cobra.Command, _ []string) error {
	arg := Arg{
		Name:   VarStringName,
		Path:   VarStringPath,
		Module: VarStringModule,
	}

	err := GenAction(arg)
	if err != nil {
		return err
	}
	err = GenRoot(arg)
	if err != nil {
		return err
	}
	err = GenUsage(arg)
	if err != nil {
		return err
	}

	return nil
}
