package main

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		diff := target - v
		if _, ok := hashTable[diff]; ok {
			return []int{hashTable[diff], i}
		} else {
			hashTable[v] = i
		}
	}
	return []int{}
}