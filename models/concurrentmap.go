package models

import (
	"sync"
)

type SI64Map struct {
	sync.RWMutex
	M map[string]int64
}

func (this *SI64Map) Put(k string, v int64) {
	this.Lock()
	this.M[k] = v
	this.Unlock()
}

func (this *SI64Map) Get(k string) int64 {
	this.RLock()
	defer this.RUnlock()
	return this.M[k]
}

func (this *SI64Map) Delete(k string) {
	this.Lock()
	delete(this.M, k)
	this.Unlock()
}

type S2Map struct {
	sync.RWMutex
	M map[string]string
}

func (this *S2Map) Put(k string, v string) {
	this.Lock()
	this.M[k] = v
	this.Unlock()
}

func (this *S2Map) Get(k string) string {
	this.RLock()
	defer this.RUnlock()
	return this.M[k]
}

func myput() int64 {
	const workers = 100
	var wg sync.WaitGroup
	wg.Add(workers)
	m := &SI64Map{M: make(map[string]int64)}
	m.Put("oo", 0)
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m.Put("oo", m.Get("oo")+1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return m.Get("oo")
}
