package command

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

func Gen() error {
	text, err := pathx.LoadTemplate(item.Category, item.TemplateFile, item.DefaultTemplateFile)
	if err != nil {
		return err
	}

	nameRune := []rune(item.TemplateFile)
	name := string(nameRune[0 : len(nameRune)-4])

	output, err := util.With(name).
		Parse(text).
		Execute(map[string]any{
			"gitUser":   gitUtil.GetGitName(),
			"gitEmail":  gitUtil.GetGitEmail(),
			"withCache": withCache,
			//"method":                methods,
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"apiFields":             apiString,
			"apiFormFields":         apiFormString,
			//"apiRequestString":      apiRequestString,
			"proto3fields":          proto3String,
			"itemStr":               itemStr,
			"itemStrWithoutUnix":    itemStrWithoutUnix,
			"inStr":                 inStr,
			"reqStr":                reqStr,
			"data":                  table,
			"tableComment":          table.TableComment,
			"tagType":               Arg.TagType,
			"pkgName":               pkgName,
			"logic":                 fmt.Sprintf("%sLogic", item.Function),
			"function":              item.Function,
			"id":                    table.PrimaryKey.Field.Name.ToCamel(),
			"serviceModule":         ServiceModule,
			"projectPath":           Arg.ProjectPath,
			"lowerStartProjectName": Arg.LowerStartProjectName,
			"upperStartProjectName": Arg.UpperStartProjectName,
			"apiVersion":            Arg.ApiVersion,
		})
	if err != nil {
		return err
	}

	item.FileCode = output.String()

	return nil
}
