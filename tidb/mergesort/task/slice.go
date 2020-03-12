/*
@Copyright:
*/
/*
@Time : 2020/3/11 21:30
@Author : teddy
@File : slice.go
*/

package task

import (
	"errors"
)

type MinInt64Slice struct {
	array []int64
}

func (s *MinInt64Slice) GetSlice() []int64 {
	return s.array
}
func (s *MinInt64Slice) Len() int {
	return len(s.array)
}

func (s *MinInt64Slice) Less(i, j int) bool {
	return s.array[i] < s.array[j]
}

func (s *MinInt64Slice) Swap(i, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

func (s *MinInt64Slice) Append(v int64) {
	s.array = append(s.array, v)
}
func (s *MinInt64Slice) IndexOf(i int) (int64, error) {
	if len(s.array) > i {
		return s.array[i], nil
	}
	return 0, errors.New("sorter is empty")
}

func (s *MinInt64Slice) Pop() (int64, error) {
	if len(s.array) > 0 {
		v := s.array[0]
		s.array = s.array[1:]
		return v, nil
	}
	return 0, errors.New("no element")
}
