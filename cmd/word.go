package cmd

import (
	"cr-tools/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	str  string
	mode int8
)

const (
	ModeUpper                      = iota + 1 //全部单词转换成大写
	ModeLower                                 //全部单词转换成小写
	ModeUnderScoreToUpperCamelCase            //下划线单词转换成大写驼峰
	ModeUnderScoreToLowerCamelCase            //下划线单词转换成小写驼峰
	ModeCamelCaseToUnderScore                 //驼峰单词转换成下划线
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下: ",
	"1: 全部单词转换成大写",
	"2: 全部单词转换成小写",
	"3: 下划线单词转换成大写驼峰",
	"4: 下划线单词转换成小写驼峰",
	"5: 驼峰单词转换成下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderScoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderScoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderScore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help wprd 查看帮助文档")
		}

		log.Printf("%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
