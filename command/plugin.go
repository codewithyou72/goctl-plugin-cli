package command

import "github.com/spf13/cobra"

var (
	//		VarStringSrc string
	//VarStringDir        string
	VarProject          string // 所在项目名
	VarApiVersion       string // 指定 API 版本，默认为 v0
	VarBoolCache        bool
	VarBoolIdea         bool
	VarStringURL        string
	VarApiFileDir       string //api文件名
	VarProtoFileDir     string //proto文件名,对应是proto3格式
	VarApiDir           string //api代码路径
	VarRpcDir           string //rpc代码路径
	VarStringSliceTable []string
	//	VarStringTable             string
	//这里使用Slice主要是偷懒
	//VarStringPrimarySliceTable []string //主表
	//VarStringSubSliceTable     []string //子表
	//VarStringModeType          string   //模型关系
	VarStringStyle string
	//		VarStringDatabase string
	//		VarStringSchema string
	VarStringHome string
	//VarStringRemote             string
	//VarStringBranch             string
	VarBoolStrict               bool
	VarBoolComment              bool
	VarBoolForce                bool
	VarBoolValidationForce      bool
	VarStringTagType            string
	VarStringValidation         string
	VarStringFieldStyle         string //前端返回风格
	VarStringSliceIgnoreColumns []string
)

func PluginCommand(_ *cobra.Command, _ []string) error {
	return nil
}
