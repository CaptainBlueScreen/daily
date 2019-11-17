// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

// Bonus: Can you do this in one pass?

package main

import "fmt"

func main() {
	fmt.Println(doPairsAddToK([]int{10, 15, 3, 7}, 17))
	fmt.Println(doPairsAddToK2([]int{10, 15, 3, 7}, 17))
}

func doPairsAddToK(numbers []int, k int) bool {
	for i, e1 := range numbers {
		for _, e2 := range numbers[i+1:] {
			if e1+e2 == k {
				return true
			}
		}
	}

	return false
}

func doPairsAddToK2(numbers []int, k int) bool {
	numberMap := make(map[int]bool)

	for _, e := range numbers {
		numberMap[e] = true
	}

	for _, e := range numbers {
		if k-e != e && numberMap[k-e] {
			return true
		}
	}

	return false
}
