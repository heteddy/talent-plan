/*
@Copyright:
*/
/*
@Time : 2020/3/11 15:41
@Author : teddy
@File : quick.go
*/

package algorithm

import "sort"

type Sorter interface {
	Sort()
	GetSortable() SortableSlice
}

type Quick struct {
	slice SortableSlice
}

func NewQuick(p SortableSlice) *Quick {
	return &Quick{slice: p}
}

func (q *Quick) Sort() {
	if q.slice.Len() <= 1 {
		return
	}
	//q.quick_sort(0, q.slice.Len()-1)
	s := q.slice.GetSlice()
	// 原来sort slice那么快
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func (q *Quick) GetSortable() SortableSlice {
	return q.slice
}

func (q *Quick) quick_sort(left, right int) {
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right; i++ {
		if q.slice.Less(i, left) {
			//分割位定位++
			explodeIndex++;
			if i != explodeIndex {
				q.slice.Swap(i, explodeIndex)
			}
		}
	}
	//起始位和分割位交换
	q.slice.Swap(left, explodeIndex)
	q.quick_sort(left, explodeIndex-1)
	q.quick_sort(explodeIndex+1, right)
}
