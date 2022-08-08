package sort

import (
	"fmt"
	"sort"
	"testing"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func TestSort1(t *testing.T) {
	arr := []string{"aaa", "dd", "cccc"}
	sort.Sort(ByLength(arr))
	fmt.Println(arr)
	sort.Sort(sort.Reverse(ByLength(arr)))
	fmt.Println(arr)

	b := ByLength(arr)
	var s []string
	s = b
	fmt.Println("s:", s)
}
