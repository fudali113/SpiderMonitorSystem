package models

import (
	"fmt"
	"sync"
)

type SI64Map struct {
	sync.RWMutex
	m map[string]int64
}

func (this *SI64Map) Put(k string, v int64) {
	this.Lock()
	this.m[k] = v
	this.Unlock()
}

func (this *SI64Map) Get(k string) int64 {
	this.RLock()
	defer this.RUnlock()
	return this.m[k]
}

type S2Map struct {
	sync.RWMutex
	m map[string]string
}

func put() interface{} {

	var err interface{} = nil
	defer func() {
		fmt.Println("----------------------->")
		if e := recover(); e != nil {
			err = e
		}
		fmt.Println("----------------------->")
	}()

	const workers = 100

	var wg sync.WaitGroup
	wg.Add(workers)
	m := map[int]int{}
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return err
}

func myput() int64 {
	fmt.Println("----------------------->")
	const workers = 100

	var wg sync.WaitGroup
	wg.Add(workers)
	m := &SI64Map{m: make(map[string]int64)}
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
	fmt.Println(m.Get("oo"))
	return m.Get("oo")
}
