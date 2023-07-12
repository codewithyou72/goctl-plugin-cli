package cmd

import (
	"zero-swagger/action"

	"github.com/spf13/cobra"
)

var swaggerCmd = &cobra.Command{
	Use:   "swagger",
	Short: "generates swagger.json",
	Long:  `generates swagger.json 完整内容`,
	RunE:  action.Generator,
}

func init() {
	rootCmd.AddCommand(swaggerCmd)

	swaggerCmd.Flags().StringVar(&action.VarStringHost, "host", "", `api request address`)
	swaggerCmd.Flags().StringVar(&action.VarStringBasePath, "basepath", "", `url request prefix`)
	swaggerCmd.Flags().StringVar(&action.VarStringFileName, "filename", "", `swagger save file name`)
}
