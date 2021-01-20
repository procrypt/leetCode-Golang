package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	max := func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i:=1; i<len(nums); i++ {
		maxEndingHere = max(nums[i], nums[i]+maxEndingHere)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}