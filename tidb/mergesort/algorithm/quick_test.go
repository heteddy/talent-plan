/*
@Copyright:
*/
/*
@Time : 2020/3/11 16:14
@Author : teddy
@File : quick_test.go
*/

package algorithm

import (
	"math/rand"
	"pingcap/talentplan/tidb/mergesort/task"
	"testing"
	"time"
)

func TestMinQuick_Sort(t *testing.T) {
	//random := rand.NewSource(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	s := make([]int64, 0, 30)

	for i := 0; i < cap(s); i++ {
		e := int64(rand.Intn(1000))
		s = append(s, e)
	}

	h := NewQuick(&task.MinInt64Slice{s})
	h.Sort()
	t.Log(s)
}
