package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tianzuoan/cobra-demmo/internal/word"
	"log"
	"strconv"
	"strings"
)

const (
	MODE_UPPER = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var wordString string
var mode int8

var desc = strings.Join([]string{
	"此子命令用户各种单词格式转换，mode参数具体使用如下：",
	strconv.Itoa(MODE_UPPER) + "：" + getModeDesc(MODE_UPPER),
	strconv.Itoa(MODE_LOWER) + "：" + getModeDesc(MODE_LOWER),
	strconv.Itoa(MODE_UNDERSCORE_TO_UPPER_CAMELCASE) + "：" + getModeDesc(MODE_UNDERSCORE_TO_UPPER_CAMELCASE),
	strconv.Itoa(MODE_UNDERSCORE_TO_LOWER_CAMELCASE) + "：" + getModeDesc(MODE_UNDERSCORE_TO_LOWER_CAMELCASE),
	strconv.Itoa(MODE_CAMELCASE_TO_UNDERSCORE) + "：" + getModeDesc(MODE_CAMELCASE_TO_UNDERSCORE),
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(wordString)
		case MODE_LOWER:
			content = word.ToLower(wordString)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderScoreToUpperCamelCase(wordString)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderScoreToLowerCamelCase(wordString)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderScore(wordString)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word 查看帮助文档！")
		}
		log.Printf("输入的单词是：%s, 转换规则是:【%d:%s】,转换之后的结果是：%s", wordString, mode, getModeDesc(mode), content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&wordString, "str", "s", "", "请输入要转换的单词！")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式！")
}

func getModeDesc(m int8) string {
	var modeDesc string
	switch m {
	case MODE_UPPER:
		modeDesc = "全部单词转为大写"
	case MODE_LOWER:
		modeDesc = "全部单词转为小写"
	case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
		modeDesc = "下划线单词转大写驼峰单词"
	case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
		modeDesc = "下划线单词转小写驼峰单词"
	case MODE_CAMELCASE_TO_UNDERSCORE:
		modeDesc = "驼峰单词转下划线单词"
	default:
		modeDesc = "暂不支持该模式"
	}
	return modeDesc
}
