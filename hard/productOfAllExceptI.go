// Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

// For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

// Follow-up: what if you can't use division?

package main

import "fmt"

func main() {
	fmt.Println("Testing [1, 2, 3, 4, 5]")
	fmt.Println("Expected", []int{120, 60, 40, 30, 24})
	result := productOfAllExceptI([]int{1, 2, 3, 4, 5})
	fmt.Println("Acutal", result)

	fmt.Println("Testing [3, 2, 1]")
	fmt.Println("Expected", []int{2, 3, 6})
	result = productOfAllExceptI([]int{3, 2, 1})
	fmt.Println("Acutal", result)

	fmt.Println("Testing []")
	fmt.Println("Expected", []int{})
	result = productOfAllExceptI([]int{})
	fmt.Println("Acutal", result)

	fmt.Println("Testing [1, 2, 3, 4, 5]")
	fmt.Println("Expected", []int{120, 60, 40, 30, 24})
	result = productOfAllExceptIWithoutDivision([]int{1, 2, 3, 4, 5})
	fmt.Println("Acutal", result)

	fmt.Println("Testing [3, 2, 1]")
	fmt.Println("Expected", []int{2, 3, 6})
	result = productOfAllExceptIWithoutDivision([]int{3, 2, 1})
	fmt.Println("Acutal", result)

	fmt.Println("Testing []")
	fmt.Println("Expected", []int{})
	result = productOfAllExceptIWithoutDivision([]int{})
	fmt.Println("Acutal", result)
}

// O(n)
func productOfAllExceptI(numbers []int) []int {
	total := productOfAll(numbers)

	ret := make([]int, len(numbers))
	for i, e := range numbers {
		ret[i] = total / e
	}

	return ret
}

func productOfAll(numbers []int) int {
	total := 1
	for _, e := range numbers {
		total = total * e
	}

	return total
}

// (n-1)(n-2)...(1)  = O(n^2/2) = O(n^2)
// can be recuded further but I don't see a way to do this in less than O(n^2)

// attempted to make log(n) subsets then multiply by the subset totals but you still end up with O(n * n/log(n) * log(n)) = O(n^2)
func productOfAllExceptIWithoutDivision(numbers []int) []int {
	ret := make([]int, len(numbers))
	runningTotal := 1
	for i, e1 := range numbers {
		ret[i] = runningTotal * productOfAll(numbers[i+1:])
		runningTotal = runningTotal * e1
	}
	return ret
}

func productOfAllExceptIRec(numbers []int, i int, headTotal int) ([]int, int) {
	if i == len(numbers) {
		return make([]int, len(numbers)), 1
	}

	products, tailTotal := productOfAllExceptIRec(numbers, i+1, headTotal*numbers[i])
	products[i] = headTotal * tailTotal

	return products, tailTotal * numbers[i]
}
