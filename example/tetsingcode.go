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

	fmt.Println("ReadK is locked")
	defer func() {
		fmt.Println("ReadK is unlocked")

	}()
	return t.k

}
func (t *teststr) Writek(j int) {

	fmt.Println("Writek is locked")
	defer func() {
		fmt.Println("Writek is unlocked", j)

	}()
	t.k += j
}

func main() {
	tst := new(teststr)
	for i := 0; i < 6; i++ {
		fmt.Println(i)
		go func(t *teststr) {
			t.mu.Lock()
			t.Writek(i)
			time.Sleep(time.Second * 3)
			t.mu.Unlock()
		}(tst)
		go func(t *teststr) {
			t.mu.Lock()
			defer t.mu.Unlock()
			t.r = i + 5
			fmt.Println("R:", t.r)
		}(tst)
		go func(t *teststr) {
			t.mu.Lock()
			defer t.mu.Unlock()
			fmt.Println("ReadK:", t.ReadK())
		}(tst)
		time.Sleep(time.Second * 1)
	}

}
