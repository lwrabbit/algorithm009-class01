package main

import "fmt"

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	generate(0, 0, 3, "")
	return res
}

func generate(left, right, n int, s string) {
	if left == n && right == n {
		res = append(res, s)
	}
	if left < n {
		generate(left+1, right, n, s+"(")
	}
	if left > right {
		generate(left, right+1, n, s+")")
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
