package main

import "fmt"

func main() {
	length := 0
	fmt.Print("Count of inputs: ")
	fmt.Scanln(&length)
	if length > 10000 {
		return
	}
	fmt.Println("Inputs: ")
	numbers := make([]int, length)

	//tekan enter setiap memasukkan nilai
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
		if numbers[i] != 1 && numbers[i] != 0 {
			fmt.Printf("wrong input")
			return
		}
	}
	fmt.Println("Array: ", numbers)
	fmt.Println("Output: ", countOfConsecutiveOnes(numbers))

}

func countOfConsecutiveOnes(nums []int) int {
	total, currentCount := 0, 0
	for _, v := range nums {
		if v == 1 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > total {
			total = currentCount
		}
	}
	return total
}
