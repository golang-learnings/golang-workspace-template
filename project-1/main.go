package main

import (
	"fmt"
	"github.com/golang-learnings/workspace/project-1/helpers"
	"strings"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(helpers.Sum(10, 10))
	fmt.Println(toUppercase("Hi"))
}

func toUppercase(str string) string {
	return strings.ToUpper(str)
}
