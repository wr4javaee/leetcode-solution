package no0020_valid_parentheses

import "container/list"

func isValid(s string) bool {
	// 映射字典
	strMap := map[string]int{ "(" : 0, "{" : 1, "[" : 2, ")" : 0, "}" : 1, "]" : 2 }

	l := list.New()
	// 遇到左括号入栈, 遇到右括号出栈并判断是否匹配
	for _, v := range s {
		str := string(v)
		if str == "(" || str == "{" || str == "[" {
			strVal, _ := strMap[str]
			l.PushBack(strVal)
			continue
		}

		// str == ")" || str == "}" || str == "]"
		// 若栈为空, 则结果为不匹配
		if l.Len() == 0 {
			return false
		}

		// 出栈
		tmpStrValElement := l.Back()
		l.Remove(tmpStrValElement)
		tmpStrVal := tmpStrValElement.Value.(int)
		strVal, _ := strMap[str]
		if tmpStrVal != strVal {
			return false
		}
	}

	// 如果栈不为空返回false
	if l.Len() > 0 {
		return false
	}
	return true
}
