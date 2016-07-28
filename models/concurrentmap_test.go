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
