package main

import (
	"fmt"
	"os"
	"similarity_compare/utils"
)

func main() {
	args := os.Args
	arg1 := args[1]
	arg2 := args[2]
	fmt.Println(utils.GetSimilarity(arg1, arg2))
}
