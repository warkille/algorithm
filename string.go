package main

import "fmt"

//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。
//需保证 返回结果的字典序最小。
//要求不能打乱其他字符的相对位置)。
//输入：s = “cbacdcbc”。
//输出：”acdb”。
//转自链接：https://learnku.com/articles/84707

func removeDuplicateLetters(s string) string {

	cnts := make([]int, 26)
	enter := make([]bool, 26)

	for _, ch := range s {
		cnts[ch-'a']++
	}

	stack := make([]byte, 26)
	size := 0

	for i := 0; i < len(s); i++ {
		cur := s[i]
		if !enter[cur-'a'] {
			for size > 0 && stack[size-1] > cur && cnts[stack[size-1]-'a'] > 0 {
				enter[stack[size-1]-'a'] = false
				size--
			}
			stack[size] = cur
			size++
			enter[cur-'a'] = true
		}
		cnts[cur-'a']--
	}

	return string(stack[:size])
}

func main() {
	s := "cbacdcbc"
	result := removeDuplicateLetters(s)
	fmt.Println(result)
}
