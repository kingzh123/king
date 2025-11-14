package goroutine

import (
	"fmt"
	"sync"
)

// RunSyncMap golang默认的map是非线程安全的
// 如果需要安全的map，官方提供了sync.Map
// 适用于以下两种场景：
//
//	写少读多：例如缓存，只写一次，读取多次。
//	多个 goroutine 操作不同 key：多个 goroutine 读取、写入和覆盖不相交的 key 集的条目。
//	因为大量写入的时候，会导致read map读不到数据而进一步加锁读取，同时dirty map也会一直晋升为read map，整体性能差，不如map + mutex
func RunSyncMap() {
	var m sync.Map
	// sync map 使用方法
	//Store(key, value any)：向 Map 中存储键值对。
	//Load(key any)：根据键获取值。
	//Delete(key any)：删除键值对。
	//LoadAndDelete(key any)：获取并删除键值对。
	//LoadOrStore(key, value any)：如果 key 已经存在，返回对应值，如果不存在，存储键值对。
	//Range(f func(key, value any) bool)：遍历 Map 中的键值对。
	m.Store("key", "value")
	m.Load("key")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
