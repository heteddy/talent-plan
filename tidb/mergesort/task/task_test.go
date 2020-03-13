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
	s := make([]int64, 0, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cap(s); i++ {
		s = append(s, int64(r.Intn(10000)))
	}
	soredChan := make(chan *MinInt64Slice)
	defer close(soredChan)
	wg := sync.WaitGroup{}
	wg.Add(1)

	st := NewSortTask(s, soredChan)
	go func() {
		if err := st.Run(); err != nil {
			panic(err)
		}
		wg.Done()
	}()
	select {
	case ret := <-soredChan:
		log.Println(ret.Len())
	}
	wg.Wait()
}

func TestMergeTask_Run(t *testing.T) {
	mergedChan := make(chan []int64)
	slices := make([][]int64, 0, 2)
	slices = append(slices, []int64{2, 4, 6, 8, 10, 12, 200, 300})
	slices = append(slices, []int64{1, 3, 5, 7, 9, 11})
	mt := MergeTask{
		slices:  slices,
		retChan: mergedChan,
	}
	defer close(mergedChan)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := mt.Run(); err != nil {
			panic(err)
		}
		wg.Done()
	}()
	select {
	case ret := <-mergedChan:
		log.Println(len(ret))
	}
	wg.Wait()
}
