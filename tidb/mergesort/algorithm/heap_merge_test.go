/*
@Copyright:
*/
/*
@Time : 2020/3/12 21:55
@Author : teddy
@File : heap_merge_test.go
*/

package algorithm

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func prepareData(src []int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(src); i++ {
		src[i] = int64(r.Int31n(100))
	}
	sort.Slice(src, func(i, j int) bool {
		return src[i] < src[j]
	})
}

func TestHeapMerge_Sort(t *testing.T) {
	k := 6
	sortedSlices := make([]*SortedSlice, 0, k)
	lens := []int{2, 2, 4, 3, 2, 2, 3, 4, 3, 7,}
	for i := 0; i < k; i++ {
		s := make([]int64, lens[i])
		prepareData(s)
		sortedSlices = append(sortedSlices, NewSortedSlice(s))
	}
	merge := NewHeapMerge(sortedSlices)
	mergedSlice := merge.Sort()
	merged2Slice := make([]int64, len(mergedSlice))
	copy(merged2Slice, mergedSlice)

	sort.Slice(merged2Slice, func(i, j int) bool { return merged2Slice[i] < merged2Slice[j] })
	for i := 0; i < len(mergedSlice); i++ {
		if mergedSlice[i] != merged2Slice[i] {
			t.Failed()
		}
	}

}
