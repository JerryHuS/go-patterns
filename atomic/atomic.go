/**
 * @Author: alessonhu
 * @Description:
 * @File:  atomic.go
 * @Version: 1.0.0
 * @Date: 2022/4/14 11:09 AM
 */

package main

import (
	"sync/atomic"
)

type SafeMap struct {
	// 互斥锁mu，操作dirty需先获取mu
	mu Mutex

	// read是只读的数据结构，访问它无须加锁，sync.map的所有操作都优先读read
	// read中存储结构体readOnly，readOnly中存着真实数据---entry（详见1.3），read是dirty的子集
	// read中可能会存在脏数据：即entry被标记为已删除(详见1.3)
	read atomic.Value // readOnly

	// dirty是可以同时读写的数据结构，访问它要加锁，新添加的key都会先放到dirty中
	// dirty == nil的情况：1.被初始化 2.提升为read后，但它不能一直为nil，否则read和dirty会数据不一致。
	// 当有新key来时，会用read中的数据 (不是read中的全部数据，而是未被标记为已删除的数据，详见3.2)填充dirty
	// dirty != nil时它存着sync.map的全部数据（包括read中未被标记为已删除的数据和新来的数据）
	dirty map[interface{}]*entry

	// 统计访问read没有未命中然后穿透访问dirty的次数
	// 若miss等于dirty的长度，dirty会提升成read，提升后可以增加read的命中率，减少加锁访问dirty的次数
	misses int
}
