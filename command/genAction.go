package command

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/goctl-plugin-cli/tpl"
	"os"
	"path/filepath"
)

var ActionFileName = "action"
var ActionDir = "action"

func GenAction(arg Arg) error {

	dirAbs, err := filepath.Abs(arg.Path)
	if err != nil {
		return err
	}

	dirAbs = filepath.Join(dirAbs, ActionDir)

	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	filename := filepath.Join(dirAbs, ActionFileName+".go")

	text, err := pathx.LoadTemplate(ActionDir, ActionFileName+".tpl", tpl.ActionTpl)
	if err != nil {
		return err
	}
	fmt.Println(arg)

	output, err := util.With(ActionFileName + ".tpl").
		Parse(text).
		Execute(map[string]any{
			"module": arg.Module,
		})
	if err != nil {
		return err
	}

	str := output.String()

	err = os.WriteFile(filename, []byte(str), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

var RootFileName = "root"
var CmdDir = "cmd"

func GenRoot(arg Arg) error {

	dirAbs, err := filepath.Abs(arg.Path)
	if err != nil {
		return err
	}

	dirAbs = filepath.Join(dirAbs, CmdDir)

	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	filename := filepath.Join(dirAbs, RootFileName+".go")

	text, err := pathx.LoadTemplate(CmdDir, RootFileName+".tpl", tpl.RootTpl)
	if err != nil {
		return err
	}

	output, err := util.With(RootFileName + ".tpl").
		Parse(text).
		Execute(map[string]any{
			"name": arg.Name,
		})
	if err != nil {
		return err
	}

	str := output.String()

	err = os.WriteFile(filename, []byte(str), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

var UsageFileName = "usage"

func GenUsage(arg Arg) error {

	dirAbs, err := filepath.Abs(arg.Path)
	if err != nil {
		return err
	}

	dirAbs = filepath.Join(dirAbs, CmdDir)

	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	filename := filepath.Join(dirAbs, UsageFileName+".go")

	text, err := pathx.LoadTemplate(CmdDir, UsageFileName+".tpl", tpl.CmdUsageTpl)
	if err != nil {
		return err
	}

	output, err := util.With(UsageFileName + ".tpl").
		Parse(text).
		Execute(map[string]any{
			"name": arg.Name,
		})
	if err != nil {
		return err
	}

	str := output.String()

	err = os.WriteFile(filename, []byte(str), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
