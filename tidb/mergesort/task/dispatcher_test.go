/*
@Copyright:
*/
/*
@Time : 2020/3/12 09:10
@Author : teddy
@File : dispatcher_test.go
*/

package task

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func constructConcurrentSorter() *ConcurrentSorter {
	slice := make([]int64, 0, 100)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cap(slice); i++ {
		slice = append(slice, int64(r.Intn(10000)))
	}
	return NewConcurrent(slice).(*ConcurrentSorter)
}

func TestConcurrentSorter_Run(t *testing.T) {
	c := constructConcurrentSorter()
	c.Run()
	log.Println(c.sortingArray)
}
