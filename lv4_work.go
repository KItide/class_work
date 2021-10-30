package main

import (
	"fmt"
	"sort"
)

type Slice []int

func (s Slice) Len() int { return len(s) }
func (s Slice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s Slice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	a := []int{77,66,5,4,9,99,82,70,61}
	sort.Sort(Slice(a))
	fmt.Println("After sorted: ", a)
}