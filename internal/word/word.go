package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线单词转为大写驼峰单词（HelloWord）
func UnderScoreToUpperCamelCase(s string) string {
	//现将下划线替换成空格
	s = strings.Replace(s, "_", " ", -1)
	//再将带空格的字符串转换为首字母大写
	s = strings.Title(s)
	//最后将空格去掉
	return strings.Replace(s, " ", "", -1)
}

//下划线单词转为小写驼峰单词（helloWord）
func UnderScoreToLowerCamelCase(s string) string {
	//先将其转换为大写驼峰单词
	s = UnderScoreToUpperCamelCase(s)
	//再将首字母转为小写字母
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//驼峰单词转下划线单词
func CamelCaseToUnderScore(s string) string {
	var result []rune
	for i, value := range s {
		if i == 0 {
			result = append(result, unicode.ToLower(value))
			continue
		}
		if unicode.IsUpper(value) {
			result = append(result, '_', unicode.ToLower(value))
		} else {
			result = append(result, value)
		}
	}
	return string(result)
}
