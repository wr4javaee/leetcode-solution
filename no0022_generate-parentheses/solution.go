package no0022_generate_parentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	generate(0, 0, &res, "", n)
	return res
}


func generate(leftCnt int, rightCnt int, res *[]string, str string, n int) {
	if leftCnt < rightCnt {
		return
	}

	if leftCnt > n {
		return
	}

	if leftCnt == n && rightCnt == n {
		*res = append(*res, str)
		return
	}

	generate(leftCnt + 1, rightCnt, res, str + "(", n)
	generate(leftCnt, rightCnt + 1, res, str + ")", n)
}
