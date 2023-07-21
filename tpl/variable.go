package tpl

import (
	_ "embed"
)

var (
	//go:embed usage.tpl
	UsageTpl string

	//go:embed action/action.tpl
	ActionTpl string

	//go:embed cmd/root.tpl
	RootTpl string

	//go:embed cmd/usage.tpl
	CmdUsageTpl string
)
