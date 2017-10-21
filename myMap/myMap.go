//简单map操作，一个map分成多个map来用
//存取的时候根据key来获取key的哈希码，由哈希码确定子map的位置
//然后取值：子map[key]
package myMap

import (
	"math/rand"
	"sync"
	"time"
)

type SubMap struct {
	Entry        map[string]interface{}
	sync.RWMutex //读写锁嵌入字段，保证每个子map是线程安全的
}

type MyMap struct {
	size    uint32
	subMaps []*SubMap
}

type Pair struct {
	Key   string
	Value interface{}
}

//先设置随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

//创建
func CreateMapManager(size uint32) *MyMap {
	m := new(MyMap)
	m.size = size
	m.subMaps = make([]*SubMap, m.size)
	for i, _ := range m.subMaps {
		m.subMaps[i] = &SubMap{Entry: make(map[string]interface{})}
	}
	return m
}

//定位到子map
func (m *MyMap) locate(key string) *SubMap {
	return m.subMaps[bkdrHash(key)%m.size]
}

func (m *MyMap) Get(key string) (value interface{}, ok bool) {
	submap := m.locate(key)
	submap.RLock()
	value, ok = submap.Entry[key]
	submap.RUnlock()
	return
}

func (m *MyMap) Set(key string, value interface{}) {
	submap := m.locate(key)
	submap.Lock()
	submap.Entry[key] = value
	submap.Unlock()
}

func (m *MyMap) Delete(key string) {
	submap := m.locate(key)
	submap.Lock()
	delete(submap.Entry, key)
	submap.Unlock()
}

func (m *MyMap) Has(key string) bool {
	_, ok := m.Get(key)
	return ok
}

func (m *MyMap) GetPairs() []Pair {
	data := make([]Pair, 30)
	for _, submap := range m.subMaps {
		submap.RLock()
		for key, value := range submap.Entry {
			data = append(data, Pair{key, value})
		}
		submap.RUnlock()
	}
	return data
}

/*
func (m *MyMap) GetKeys() <-chan string {
	ch := make(chan string)
	go func() {
		for _, submap := range m.subMaps {
			submap.RLock()
			for key, _ := range submap.Entry {
				ch <- key
			}
			submap.RUnlock()
		}
		close(ch)
	}()
	return ch
}*/

const seed uint32 = 1313 //质数

//求得字符串对应的哈希值
func bkdrHash(str string) uint32 {
	var h uint32
	for _, c := range str {
		h = h*seed + uint32(c)
	}
	return h
}
