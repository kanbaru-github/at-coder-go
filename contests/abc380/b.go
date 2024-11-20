package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	charSlice := strings.Split(input, "|")

	a := make([]int, len(charSlice))
	for i := 0; i < len(charSlice); i++ {
		a[i] = strings.Count(charSlice[i], "-")
	}

	for _, v := range a {
		if v != 0 {
			fmt.Print(v, " ")
		}
	}
}
