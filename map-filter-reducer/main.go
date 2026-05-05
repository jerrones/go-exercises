package main

import "fmt"

type DataSlice []int

func main() {
	list := DataSlice{1, 2, 3, 4, 5}

	filtered := list.Filter(func(i int) bool {
		return i%2 == 0
	})

	filtered = filtered.Map(func(i int) int {
		return i * 100
	})

	fmt.Println(filtered)

	sum := filtered.Reduce(func(acc, i int) int {
		return acc + i
	}, 0)

	fmt.Println(sum)
}

func (d DataSlice) Filter(f func(int) bool) DataSlice {
	var result DataSlice
	for _, v := range d {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func (d DataSlice) Map(f func(int) int) DataSlice {
	result := make(DataSlice, len(d))
	for i, v := range d {
		result[i] = f(v)
	}

	return result
}

func (d DataSlice) Reduce(f func(int, int) int, initial int) int {
	result := initial
	for _, v := range d {
		result = f(result, v)
	}

	return result
}
