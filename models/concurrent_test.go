package models

import (
	"fmt"
	"testing"
)

func Test_SI64Map(t *testing.T) {
	m := &SI64Map{m: make(map[string]int64)}
	m.Put("oo", 1)
	m.Put("oo", 2)
	if m.Get("oo") != 2 {
		t.Error("oo")
	}
}

func Test_S2Map(t *testing.T) {
	m := &SI64Map{m: make(map[string]string)}
	m.Put("oo", "o")
	m.Put("oo", "oo")
	if m.Get("oo") != "oo" {
		t.Error("oo")
	}
}

func Test_myput(t *testing.T) {
	if a := myput(); a == 100 {
		fmt.Println(a)
		t.Error("ooooo")
	}
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
