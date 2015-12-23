package sort

import (
	"sort"
)

type sorter struct{
	len func() int
	less func(i, j int) bool
	swap func(i, j int)
}

func (s *sorter) Len() int {
	return s.len()
}

func (s *sorter) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *sorter) Swap(i, j int) {
	s.swap(i, j)
}

func Sort(len func() int, less func(i, j int) bool, swap func(i, j int)) {
	sort.Sort(&sorter{
		len,
		less,
		swap,
	})
}