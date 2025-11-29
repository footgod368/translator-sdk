package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	youdao "github.com/footgod368/translator-sdk"
)

func main() {
	ctx := context.Background()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("查询：")
		if !scanner.Scan() {
			break
		}
		query := scanner.Text()
		result, err := youdao.Query(ctx, query)
		if err != nil {
			fmt.Println(err)
		}
		youdao.PrintResult(query, result)
		fmt.Println(strings.Repeat("-", 70))
	}
}
