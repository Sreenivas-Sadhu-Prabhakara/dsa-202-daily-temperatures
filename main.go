package main

import "fmt"

// DailyTemperatures returns, for each day, how many days until a strictly
// warmer one (0 if never). Next-greater-element in distance form: a
// monotonic stack of unresolved day indexes.
func DailyTemperatures(temps []int) []int {
	out := make([]int, len(temps))
	stack := []int{} // indexes with strictly decreasing temperatures
	for i, t := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < t {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out[j] = i - j
		}
		stack = append(stack, i)
	}
	return out
}

func main() {
	fmt.Println(DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
