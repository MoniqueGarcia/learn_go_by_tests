package main

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(arraysToSum ...[]int) []int {
	// Using append function
	result := []int{}
	for _, array := range arraysToSum {
		result = append(result, Sum(array))
	}

	// Using make func to create the result array
	// result := make([]int, len(arraysToSum))
	// for idx, array := range arraysToSum {
	// 	result[idx] = Sum(array)
	// }

	return result
}

func SumAllTails(numbersToSum ...[]int) (result []int) {
	var tail []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tail = []int{0}
		} else {
			tail = numbers[1:]
		}
		result = append(result, Sum(tail))
	}
	return 

}

func main() {
}