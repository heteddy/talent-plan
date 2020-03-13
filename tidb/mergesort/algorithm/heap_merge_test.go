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
	lens := []int{200, 200, 400, 300, 200, 200, 300, 40, 300, 700,}
	for i := 0; i < k; i++ {
		s := make([]int64, lens[i])
		prepareData(s)
		sortedSlices = append(sortedSlices, NewSortedSlice(s))
	}
	merge := NewHeapMerge(sortedSlices)
	merge.Sort()
	//merged2Slice := make([]int64, len(mergedSlice))
	//copy(merged2Slice, mergedSlice)
}

func BenchmarkHeapMerge_Sort(b *testing.B) {
	k := 6
	sortedSlices := make([]*SortedSlice, 0, k)
	lens := []int{2000, 2000, 4000, 3000, 2000, 2000, 3000, 4000, 3000, 7000,}
	for i := 0; i < k; i++ {
		s := make([]int64, lens[i])
		prepareData(s)
		sortedSlices = append(sortedSlices, NewSortedSlice(s))
	}
	for i := 0; i < k; i++ {
		s := make([]int64, lens[i])
		prepareData(s)
		sortedSlices = append(sortedSlices, NewSortedSlice(s))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		b.StartTimer()
		merge := NewHeapMerge(sortedSlices)
		merge.Sort()
	}
}
