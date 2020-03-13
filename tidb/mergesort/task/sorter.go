/*
@Copyright:
*/
/*
@Time : 2020/3/10 18:24
@Author : teddy
@File : dispatcher.go
*/

package task

import (
	"pingcap/talentplan/tidb/mergesort/algorithm"
	"pingcap/talentplan/tidb/mergesort/pool"
	"runtime"
	"sync"
	"time"
)

const SORTING_ARRAY_THRESHOLD = 1 << 12

type Sorter interface {
	Run()
}

type SingleSorter struct {
	sortingArray []int64
}

func (s *SingleSorter) Run() {
	h := algorithm.NewQuick(&MinInt64Slice{s.sortingArray})
	h.Sort()
}

func NewSingle(src []int64) Sorter {
	return &SingleSorter{sortingArray: src}
}

type ConcurrentSorter struct {
	sortedChan   chan *MinInt64Slice
	sortingArray []int64
	pool         *pool.Pool
	taskNum      int
}

func NewConcurrent(src []int64) Sorter {
	// 拆分成子任务并行完成
	taskNum := runtime.NumCPU()
	return &ConcurrentSorter{
		sortedChan: make(chan *MinInt64Slice, 1),
		//mergeRetChan: make(chan []int64),
		sortingArray: src,
		pool: pool.NewPool(&pool.Config{
			QSize:   1,
			Workers: taskNum,
			MaxIdle: time.Second * 10,
		}),
		taskNum: taskNum,
	}
}

func (m *ConcurrentSorter) sort() {
	start := 0
	step := len(m.sortingArray) / m.taskNum
	// 不能整除，则最后一个task多处理一些
	count := 1
	for ; start < len(m.sortingArray); {
		end := (start + step) % len(m.sortingArray)
		// 最后一个任务
		if m.taskNum == count {
			end = len(m.sortingArray)
		}
		t := NewSortTask(m.sortingArray[start:end], m.sortedChan)
		start = end
		m.pool.Put(t)
		count++
	}
}

func (m *ConcurrentSorter) merge(mergedChan chan []int64) {
	sortedSlices := make([][]int64, 0, m.taskNum)
	sortedLen := 0
loop:
	for {
		select {
		case s := <-m.sortedChan:
			sortedLen += s.Len()
			sortedSlices = append(sortedSlices, s.GetSlice())
			// sort 阶段完成
			if sortedLen == len(m.sortingArray) {
				// 这里确保所有的task都已经退出，不然可能导致死锁
				// 死锁产生的场景，SortTask.Run()->SortedChan,如果该routine退出<-m.sortedChan;
				// 那么SortedTask无法退出；当前task m.pool.Put(MergeTask)就会阻塞
				break loop
			}
		}
	}
	mergeTask := NewMergeTask(sortedSlices, mergedChan)
	// 为避免死锁可以另外启动一个协程写入
	m.pool.Put(mergeTask)
}

func (m *ConcurrentSorter) Run() {
	mergedChan := make(chan []int64, 1)
	defer func() {
		m.pool.Close(true)
		close(m.sortedChan)
		close(mergedChan)
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		m.sort()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		m.merge(mergedChan)
		wg.Done()
	}()
	wg.Wait()

	resultSlice := <-mergedChan
	copy(m.sortingArray, resultSlice)
}
