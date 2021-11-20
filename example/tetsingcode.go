package main

import (
	"fmt"
	"sync"
	"time"
)

type teststr struct {
	mu sync.RWMutex
	k  int
	r  int
}

func (t *teststr) ReadK() int {
	t.mu.Lock()
	fmt.Println("ReadK is locked")
	defer func() {
		fmt.Println("ReadK is unlocked")
		t.mu.Unlock()
	}()
	return t.k

}
func (t *teststr) Writek(j int) {
	t.mu.RLock()
	fmt.Println("Writek is locked")
	defer func() {
		fmt.Println("Writek is unlocked", j)
		t.mu.RUnlock()
	}()
	t.k = j
}

func main() {
	tst := new(teststr)
	for i := 0; i < 6; i++ {
		fmt.Println(i)
		go func(t *teststr) {
			t.Writek(i)

		}(tst)
		go func(t *teststr) {
			t.r = i + 5
			fmt.Println("R:", t.r)
		}(tst)
		go func(t *teststr) {
			fmt.Println("ReadK:", t.ReadK())
		}(tst)
	}
	time.Sleep(time.Second * 5)

}
