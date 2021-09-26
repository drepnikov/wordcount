package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	t := os.Args[1]

	fmt.Println(len(strings.Fields(t)))
}
