# Merge Sort

## 问题描述
Go 语言实现一个16M的整数(int64)多路归并的数组排序

## 思路
将待排序数组分成多个组，利用多个goroutine实现各个组的并行排序；然后通过**Heap(最小堆)进行多路归并排序**；

## 实现
实现一个协程池实现任务的并行处理，将待排序切片分组并封装成SortTask放入协程池
运行，待全部执行完成后ConcurrentSorter收集排序结果，并封装成MergeTask放入协程池中进行合并。

+ 协程池pool.go
  
    - 配置最大协程数量
    - 按需创建协程
    - 空闲超时则回收协程
  
+ 合并有序切片algorithm.heap_merge.go
    通不采用2路循环合并，避免分配过多的内存碎片。通过堆实现多路的有序切片的合并，额外申请一倍的内存用于存放合并结果



### 归并算法
输入：n路待合并的有序slice
输出：有序slice

堆node定义为一个SortedSlice，实现了hasNext函数，用于迭代到当前slice的下一个元素；
```
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
```
堆的定义：
```
type HeapMerge struct {
	nodes []*SortedSlice
}
```

1. 构建一个n个元素的最小堆
2. 从每路slice中取首个元素组成数组，调整堆；每次从堆顶，取一个元素，放入合并后的slice中
     + 如果hasNext=true，执行当前node的Next()，重新调整当前的原因
     + 如果hasNext=false, 当前slice已经空了，因此剔除堆顶, 然后需要重建堆，原因是堆中的父子关系已经破坏。
```
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
```
### 代码结构

![截屏2020-03-1318.33.34.png](https://upload-images.jianshu.io/upload_images/9243349-0c28af6a9d9f4e20.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 性能测试
**并发8路排序的的情况下，性能大约提升三倍**，主要原因是分组排序之后需要进行多路的合并。测试结果如下：

![截屏2020-03-1317.11.58.png](https://upload-images.jianshu.io/upload_images/9243349-943ff777fc1fb4b5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


**内存消耗比直接排序增加了128M**，是因为合并排序结果过程申请了一块内存来暂存结果128M = 16M*8B

![截屏2020-03-1317.12.43.png](https://upload-images.jianshu.io/upload_images/9243349-23fa0062e2bd75db.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

cpu的消耗大多在排序过程，merge过程5%
![截屏2020-03-1317.18.11.png](https://upload-images.jianshu.io/upload_images/9243349-e36a7d74c8569678.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**merge过程中调用append(slice)消耗了290ms，直接改为修改slice的下标竟然减少了大约10ms**。

![截屏2020-03-1317.20.47.png](https://upload-images.jianshu.io/upload_images/9243349-cc55e3d816718d25.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![截屏2020-03-1319.05.17.png](https://upload-images.jianshu.io/upload_images/9243349-cae318f1eb460a1e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


 [github-mergesort源码](https://github.com/heteddy/talent-plan/tree/master/tidb/mergesort)
