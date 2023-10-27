package demos

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeCounter struct {
	Mu    sync.Mutex
	WG    sync.WaitGroup
	Val   int
	Print bool
}

func NewSafeCounter() *SafeCounter {
	sc := SafeCounter{}
	sc.Mu = sync.Mutex{}
	sc.Val = 0
	sc.Print = false
	return &sc
}

func (sc *SafeCounter) increment() {
	sc.Mu.Lock()
	sc.Val++
	sc.Mu.Unlock()
}
func (sc *SafeCounter) decrement() {
	sc.Mu.Lock()
	sc.Val--
	sc.Mu.Unlock()
}

func (sc *SafeCounter) Increment(iter int) {
	for i := 0; i < iter; i++ {
		sc.increment()
	}
	sc.WG.Done()
}
func (sc *SafeCounter) Decrement(iter int) {
	for i := 0; i < iter; i++ {
		sc.decrement()
	}
	sc.WG.Done()
}
func (sc *SafeCounter) CompareAndSwap(expected int, new int) int {
	sc.Mu.Lock()
	actual := sc.Val
	if actual == expected {
		if sc.Print {
			fmt.Println("Val ", sc.Val, "-->", new, "Good =", new == (sc.Val+1))
		}
		sc.Val = new
	}

	sc.Mu.Unlock()

	return actual
}
func (sc *SafeCounter) TestAndSet(new int) int {
	sc.Mu.Lock()
	old := sc.Val
	if sc.Print {
		fmt.Println("Val ", sc.Val, "-->", new, "Good =", new == (sc.Val+1))
	}

	sc.Val = new
	sc.Mu.Unlock()
	return old
}

func (sc *SafeCounter) RunDemoCounterWithCAS() {
	sc.Print = true
	start := sc.Val

	f := func(initial int) {
		expected := initial
		for i := 0; i < 10; i++ {
			new := expected + 1
			actual := sc.CompareAndSwap(expected, new)
			expected = actual
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
		sc.WG.Done()
	}

	sc.WG.Add(2)
	go f(start)
	go f(start)

	sc.WG.Wait()
	fmt.Println("Consistent counter")
}

func (sc *SafeCounter) RunDemoCounterWithTAS() {
	sc.Print = true
	start := sc.Val

	f := func(initial int) {
		old := initial
		new := old + 1
		for i := 0; i < 10; i++ {
			old = sc.TestAndSet(new)
			new = new + 1
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
		sc.WG.Done()
	}

	sc.WG.Add(2)
	go f(start)
	go f(start)

	sc.WG.Wait()
	fmt.Println("Inconsistent counter")

}
