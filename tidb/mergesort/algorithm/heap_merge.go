/*
@Copyright:
*/
/*
@Time : 2020/3/12 23:59
@Author : teddy
@File : heap_merge.go

通过最小堆/最大堆实现多路归并merge
*/

package algorithm

import (
	"errors"
	"fmt"
	"log"
)

type SortableSlice interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Append(v int64)
	IndexOf(int) (int64, error)
	Pop() (int64, error)
	GetSlice() []int64
}

type Iterator struct {
	slice []int64
	index int
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.slice)-1
}

func (i *Iterator) Next() {
	i.index++
}

func (i *Iterator) Value() int64 {
	return i.slice[i.index]
}

type SortedSlice struct {
	slice []int64
	Iterator
}

func NewSortedSlice(slice []int64) *SortedSlice {
	return &SortedSlice{
		slice: slice,
		Iterator: Iterator{
			slice: slice,
			index: 0,
		},
	}
}



type HeapMerge struct {
	nodes []*SortedSlice
}

func NewHeapMerge(sources []*SortedSlice) *HeapMerge {
	// 需要保证
	return &HeapMerge{nodes: sources}
}

func (h *HeapMerge) Build() {
	for index := len(h.nodes) / 2; index >= 0; index-- {
		h.adjust(index, len(h.nodes))
	}
}

func (h *HeapMerge) Pop() (int64, error) {
	var value int64
	var err error

	if len(h.nodes) > 0 {
		value = h.nodes[0].Value()
		err = nil

		if h.nodes[0].HasNext() {
			h.nodes[0].Next() //不需要获取值
			h.adjust(0, len(h.nodes))
		} else { // 顶部的node(slice)已经为空
			if len(h.nodes) >= 1 {
				// 移除为已经合并完成的slice
				h.nodes = h.nodes[1:]
				//h.adjust(0, len(h.nodes))
				h.Build()
			} else {
				return 0, errors.New("merge complete")
			}
		}
	} else {
		return 0, errors.New("merge complete")
	}
	//h.Print()
	return value, err
}

func (h *HeapMerge) Print() {
	s := "heap merge:"
	for _, n := range h.nodes {
		s += fmt.Sprintf("%d ", n.Value())
	}
	log.Println(s)
}

func (h *HeapMerge) adjust(start, end int) {
	childIndex := 2*start + 1
	// 下标应该比长度小
	if childIndex >= end {
		return
	}
	if childIndex+1 < end && h.nodes[childIndex+1].Value() < h.nodes[childIndex].Value() {
		childIndex++
	}

	if h.nodes[childIndex].Value() < h.nodes[start].Value() {
		h.nodes[start], h.nodes[childIndex] = h.nodes[childIndex], h.nodes[start]

		// 一旦交换了之后，后面的节点要重新调整顺序
		h.adjust(childIndex, end)
	}
}

func (h *HeapMerge) Sort() []int64 {
	h.Build()
	length := 0

	for _, c := range h.nodes {
		length += len(c.slice)
	}

	mergeSlice := make([]int64, length, length)
	mergeSliceIndex := 0
loop:
	for {
		v, err := h.Pop()
		if err != nil {
			break loop
		} else {
			// 替换掉 mergeSlice = append(mergeSlice,v) 节省了大约10ms
			mergeSlice[mergeSliceIndex] = v
			//mergeSlice = append(mergeSlice, v)
		}
		mergeSliceIndex++
	}
	return mergeSlice

}
