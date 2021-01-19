package main

import "fmt"

type StockSpanner struct {
	container [][]int
	size  int
}

func(s *StockSpanner) isEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}
func(s *StockSpanner) Push(i int, result int) {
	s.container = append(s.container, []int{i, result})
	s.size++
}

func(s *StockSpanner) Pop() {
	if s.size == 0 {
		return
	}
	s.container = s.container[:len(s.container)-1]
	s.size--
}

func(s *StockSpanner) Top() int {
	return s.container[len(s.container)-1][0]
}

func Constructor() StockSpanner {
	return StockSpanner{}
}


func (this *StockSpanner) Next(price int) int {
	result := 1
	if this.isEmpty() {
		this.Push(price, result)
	} else if !this.isEmpty() && this.Top() > price {
		this.Push(price, result)
	} else if !this.isEmpty() && this.Top() <= price {
		for !this.isEmpty() && this.Top() <= price {
			result += this.container[len(this.container)-1][1]
			this.Pop()
		}
		if this.isEmpty() {
			this.Push(price, result)
		} else if this.Top() > price {
			this.Push(price, result)
		}
	}
	return result
}

func main() {
	var result []int
	s := StockSpanner{}
	price := []int{100,80,60,70,60,75,85}
	for _, v := range price {
		result = append(result, s.Next(v))
	}
	fmt.Println(result)
}
