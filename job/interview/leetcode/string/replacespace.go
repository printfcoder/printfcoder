package main

import (
	"fmt"
)

func main() {
	out := replaceOneTime("Hello World !")
	fmt.Printf("%s\n", out)
}

func replaceOneTime(input string) (out string) {
	// 获取所有空格长度
	sCount := 0
	for _, si := range input {
		if si == ' ' {
			sCount++
		}
	}

	// 总体要挪sCount*2位
	sCount *= 2
	// 需要先扩容，用最后一个字符扩容
	for i := 0; i < sCount; i++ {
		input += string(input[len(input)-1])
	}
	fmt.Printf("扩容后：%s \n", input)

	l := len(input) - sCount
	for i := l - 1; i > 0 && sCount > 0; i-- {
		si := input[i]
		if si != ' ' {
			// 挪到后面去
			input = input[:i+sCount] + string(si) + input[i+sCount+1:]
		} else {
			// 是空格，0按正常的挪
			input = input[:i+sCount] + string('0') + input[i+sCount+1:]
			// 2挪sCount-1个
			input = input[:i+sCount-1] + string('2') + input[i+sCount:]
			input = input[:i+sCount-2] + string('%') + input[i+sCount-1:]
			// 挪完要减掉刚刚新增的%20
			sCount -= 2
		}
	}

	return input
}
