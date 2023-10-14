package demos

import "sync"

type SafeCounter struct {
	Mu  sync.Mutex
	WG  sync.WaitGroup
	Val int
}

func NewSafeCounter() *SafeCounter {
	sc := SafeCounter{}
	sc.Mu = sync.Mutex{}
	sc.Val = 0
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
