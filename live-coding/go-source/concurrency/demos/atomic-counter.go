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

func (sc *SafeCounter) incr() {
	sc.Mu.Lock()
	sc.Val++
	sc.Mu.Unlock()

}
func (sc *SafeCounter) decr() {
	sc.Mu.Lock()
	sc.Val--
	sc.Mu.Unlock()
}

func (sc *SafeCounter) Incr(iter int) {
	for i := 0; i < iter; i++ {
		sc.incr()
	}
	sc.WG.Done()
}
func (sc *SafeCounter) Decr(iter int) {
	for i := 0; i < iter; i++ {
		sc.decr()
	}
	sc.WG.Done()
}
