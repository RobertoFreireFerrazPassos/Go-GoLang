package main

// go mod init basicExample

// go build main.go
// go run main.go

import (
	"bufio"
	"fmt"
	"os"
) // import others packages

func main() {
	sliceExample()
}

// base go types: string, integer, float, array, map

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sliceExample() {
	// array that can grow and shrink
	texts := []string{"A1", "C3"}
	texts = append(texts, "F3")
	fmt.Println(texts)

	for i, text := range texts {
		fmt.Println(i, text)
	}
}

func funcExample() {
	var variable = innerFuncExample("test")
	fmt.Println(variable)
}

func innerFuncExample(text string) string {
	return "Variable: " + text
}

func variableDeclaration() {
	card := "Variable 1" // := declaring a new variable and implicit understanding the type
	card = "Variable 2"
	fmt.Println(card)
}

func printFunc() {
	fmt.Println("Hi There!")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}
