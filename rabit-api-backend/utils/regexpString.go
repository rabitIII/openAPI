package utils

import (
	"fmt"
	"regexp"
)

// MatchString
//
// @Description: 验证前端返回的字符串是否符合某种（正则表达式）规则标准
// @param pattern: 正则表达式（验证规则）
// @param data: 待验证的数据（字符串类型）
// @return bool: true通过验证，false未通过验证
func MatchString(pattern, data string) bool {
	matched, err := regexp.MatchString(pattern, data)
	if !matched {
		return false
	}

	if err != nil {
		fmt.Println("[(UTILS)MATCHSTRING ERROR] -> ", err.Error())
		return false
	}
	return true
}
