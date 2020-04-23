package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int { return len(p) }

func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {

	studyGroup := people{"Hello", "Javier", "Zeno", "Al", "John"}

	fmt.Println(studyGroup)

	sort.Sort(studyGroup)

	fmt.Println(studyGroup)
	s := []string{"Hello", "Javier", "Zeno", "Al", "John"}

	fmt.Println()
	fmt.Println(s)

	//sort.Sort(sort.Strings(s))
	sort.Strings(s)

	fmt.Println(s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println()
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)

	sr := []string{"Hello", "Javier", "Zeno", "Al", "John"}

	fmt.Println()

	fmt.Println(sr)
	sort.Sort(sort.Reverse(sort.StringSlice(sr)))
	fmt.Println(sr)
}
