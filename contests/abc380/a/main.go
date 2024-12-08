package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mem := make([]int, 10)

	for n < 0 {
		digit := n % 10
		mem[digit]++
		n /= 10
	}

	if mem[1] == 1 && mem[2] == 2 && mem[3] == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
