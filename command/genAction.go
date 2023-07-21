package command

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/goctl-plugin-cli/tpl"
	"os"
	"path/filepath"
)

func GenAction(arg Arg) error {

	dirAbs, err := filepath.Abs(arg.Path)
	if err != nil {
		return err
	}

	dirAbs = filepath.Join(dirAbs, "action")

	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	filename := filepath.Join(dirAbs, "action.go")

	text, err := pathx.LoadTemplate("action", "action.tpl", tpl.ActionTpl)
	if err != nil {
		return err
	}
	fmt.Println(arg)

	output, err := util.With("action.tpl").
		Parse(text).
		Execute(map[string]any{
			//"gitUser":   gitUtil.GetGitName(),
			//"gitEmail":  gitUtil.GetGitEmail(),
			//"withCache": withCache,
			//"method":                methods,
			//"upperStartCamelObject": table.Name.ToCamel(),
			//"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
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
