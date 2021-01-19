package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	currentSum := 0
	for i,v  := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue //To prevent the repeat
		}
		left := i+1
		right := len(nums)-1
		for left < right {
			currentSum = v + nums[left]+nums[right]
			if currentSum == 0 {
				result = append(result, []int{nums[i],nums[left],nums[right]})
				left++
				right--
			} else if currentSum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-2,0,0,2,2}))
}