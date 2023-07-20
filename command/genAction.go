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

	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	filename := filepath.Join(dirAbs, "action", "action.go")

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
			//"apiFields":             apiString,
			//"apiFormFields":         apiFormString,
			////"apiRequestString":      apiRequestString,
			//"proto3fields":          proto3String,
			//"itemStr":               itemStr,
			//"itemStrWithoutUnix":    itemStrWithoutUnix,
			//"inStr":                 inStr,
			//"reqStr":                reqStr,
			//"data":                  table,
			//"tableComment":          table.TableComment,
			//"tagType":               Arg.TagType,
			//"pkgName":               pkgName,
			//"logic":                 fmt.Sprintf("%sLogic", item.Function),
			//"function":              item.Function,
			//"id":                    table.PrimaryKey.Field.Name.ToCamel(),
			//"serviceModule":         ServiceModule,
			//"projectPath":           Arg.ProjectPath,
			//"lowerStartProjectName": Arg.LowerStartProjectName,
			//"upperStartProjectName": Arg.UpperStartProjectName,
			//"apiVersion":            Arg.ApiVersion,
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
