package main

import "fmt"

func MapSlice__int(f func(int) int, list []int) []int {
	result := make([]int, len(list))
	for i, val := range list {
		result[i] = f(val)
	}
	return result
}
func MapSlice__string(f func(string) string, list []string) []string {
	result := make([]string, len(list))
	for i, val := range list {
		result[i] = f(val)
	}
	return result
}

func main() {
	mapSliceInt := MapSlice__int(func(i int) int { return 2 * i }, []int{1, 2, 3, 4, 5})
	fmt.Println(mapSliceInt)

	mapSliceString := MapSlice__string(func(s string) string { return s + s }, []string{"hoge", "fuga", "yura"})
	fmt.Println(mapSliceString)
}
