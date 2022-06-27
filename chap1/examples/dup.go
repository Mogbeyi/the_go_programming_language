package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) 
	input := bufio.NewScanner(os.Stdin) 

	for input.Scan() {
		counts[input.Text()]++ 
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Printf("Enter a value: ")
// 	scanner.Scan()
// 	input := scanner.Text()
// 	fmt.Printf("You typed %v", input)
// }

