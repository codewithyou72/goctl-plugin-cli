package action

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"goctl-plugin-cli/generate"
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

	//读取调试文件内容===========================开始===============
	/*	file, err := os.Open("./demo.json")
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("err=", err)
			}
		}(file)

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		var result plugin.Plugin
		err = json.Unmarshal(data, &result)
		if err != nil {
			return err
		}
		p := &result*/
	//读取调试文件内容===========================结束===============

	basepath := VarStringBasePath
	host := VarStringHost
	return generate.Do(fileName, host, basepath, p)
}

func Demo() {
	//新增插入JSON格式文件======================开始=================
	//debug命令：
	//api plugin -p goctl-swagger="swagger -filename gateway.json" --api="./gateway.api" --dir .

	//插入代码路径
	//tools/goctl/plugin/plugin.go
	//content, err := execx.Run(bin+" "+args, filepath.Dir(ex), bytes.NewBuffer(transferData)) 代码前面插入
	/*	file, err := os.Create("demo.json")
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("文件打开错误")
			}
		}(file)

		_, err = file.Write(transferData)
		if err != nil {
			// 处理写入文件错误
			return err
		}

		err = file.Sync()
		if err != nil {
			return err
		}*/
	//新增插入JSON格式文件======================结束=================
}
