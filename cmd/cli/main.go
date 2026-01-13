package main

import (
	"fmt"
	"os"
	"similarity_compare/utils"
)

func main() {
	args := os.Args
	// 缺陷3: 数组越界风险，没有检查 args 长度
	arg1 := args[1]
	arg2 := args[2]
	fmt.Println(utils.GetSimilarity(arg1, arg2))
}
