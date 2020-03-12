/*
@Copyright:
*/
/*
@Time : 2020/3/10 21:47
@Author : teddy
@File : task_test.go
*/

package task

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSortTask_Run(t *testing.T) {
	s := make([]int64, 0, 10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cap(s); i++ {
		s = append(s, int64(r.Intn(10000)))
	}

	retChan := make(chan *MinInt64Slice)
	defer close(retChan)
	wg := sync.WaitGroup{}
	wg.Add(1)

	//st := SortTask{
	//	sorter:  algorithm.NewQuick(&MinInt64Slice{array: s}),
	//	sortedChan: sortedChan,
	//}
	st := NewSortTask(s, retChan)
	go func() {
		st.Run()
		wg.Done()
	}()
	select {
	case ret := <-retChan:
		log.Println(*ret.pSlice)
	}
	wg.Wait()
}

func TestMergeTask_Run(t *testing.T) {
	retChan := make(chan *MemRecycleSlice)

	mt := MergeTask{
		slice1:  NewMemRecycleSlice([]int64{2, 4, 6, 8, 10, 12, 200, 300}),
		slice2:  NewMemRecycleSlice([]int64{1, 3, 5, 7, 9, 11}),
		retChan: retChan,
	}
	defer close(retChan)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		mt.Run()
		wg.Done()
	}()
	select {
	case ret := <-retChan:
		log.Println(*ret.pSlice)
	}
	wg.Wait()
}
