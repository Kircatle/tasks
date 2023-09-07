package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	var result = make([]int, 2)
	for index, value := range nums {
		_, ok := m[target-value]
		if ok {
			result[0] = m[target-value]
			result[1] = index
			return result
		}
		m[value] = index
	}
	return result
}

func main() {
	inputArray := []int{2, 7, 11, 15}
	var result []int
	var target int = 9
	result = twoSum(inputArray, target)
	fmt.Println(result)

}
