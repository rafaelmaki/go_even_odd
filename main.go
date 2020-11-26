package main

import "fmt"

func main() {
	s := []int{}

	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}

	for _, n := range s {
		evenOdd(n)
	}
}

func evenOdd(n int) {
	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}
