package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last := 0
	secondLast := 0
	return func() int {
		if last == 0 && secondLast == 0 {
			last = 1
			return 0
		} else if last == 1 && secondLast == 0 {
			secondLast = 1
			return 1
		}

		actual := last + secondLast
		secondLast = last
		last = actual

		return actual
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
