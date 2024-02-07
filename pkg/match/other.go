package match

import (
	"regexp"
	"sort"
	"strings"
)

// matchEqual 判断表达式性质 = / !=
func matchEqual(rule string, matchFunc bool) bool {
	re := regexp.MustCompile(`(\w+\s*)(!?=)"(.*?)"$`)
	match := re.FindStringSubmatch(rule)
	if strings.Contains(match[2], "!=") {
		return !matchFunc
	}
	return matchFunc
}

// CaseInsensitiveContains 全部转小写然后判断是否包含关系
func CaseInsensitiveContains(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

// unEscapeAndSpace 去除 \\ 换行符 两边空格
func unEscapeAndSpace(input string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(input, "\\", ""), "\n", ""))
}

// inSlice 判断切片中是否存在字符串
func inSlice(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}
