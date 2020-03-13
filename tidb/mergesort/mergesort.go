package main

import (
	"pingcap/talentplan/tidb/mergesort/task"
)

//MergeSort performs the merge sort algorithm.
//Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {

	if len(src) > task.SORTING_ARRAY_THRESHOLD {
		sorter := task.NewConcurrent(src)
		sorter.Run()
	} else {
		sorter := task.NewSingle(src)
		sorter.Run()
	}
}
