package main

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < 2; i++ {
		sum += numbers[i]
	}
	return sum
}

//refactored code with range
//	func Sum(numbers [5]int) int {
//		sum := 0
//		for _, number := range numbers {
//			sum += number
//		}
//		return sum
//	}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int

	for _, number := range numbersToSum {
		if len(number) == 0 {
			sum = append(sum, 0)
		} else {
			tail := number[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}
