package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func isMagicNumber(value string) bool {
	// 判断一个数字是否为魔法数字
	// 这里可以自定义逻辑，比如过滤掉常用的0, 1等
	if value == "0" || value == "1" || value == "-1" {
		return false
	}
	// 正则表达式判断是否为数字
	matched, _ := regexp.MatchString(`^\d+(\.\d+)?$`, value)
	return matched
}

func findMagicNumbersInFile(filePath string) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// 遍历语法树
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.BasicLit:
			if x.Kind == token.INT || x.Kind == token.FLOAT {
				if isMagicNumber(x.Value) {
					fmt.Printf("Found magic number: %s at %s\n", x.Value, fileSet.Position(x.Pos()))
				}
			}
		}
		return true
	})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	root := os.Args[1]
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			fmt.Println("Scanning:", path)
			findMagicNumbersInFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}
