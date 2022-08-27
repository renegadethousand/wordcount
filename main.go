package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fields := strings.Fields(os.Args[1])
	fmt.Println(len(fields))
}
