package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	id int
	sync.Mutex
}

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, host chan bool) {
	for eatCount := 1; eatCount <= 3; eatCount++ {
		host <- true

		if p.leftCS.id < p.rightCS.id {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}

		fmt.Println(fmt.Sprintf("starting to eat [%d]", p.id))
		time.Sleep(500 * time.Millisecond)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host

		fmt.Println(fmt.Sprintf("finished eating [%d]", p.id))
		fmt.Println(fmt.Sprintf("eat count for philosopher[%d]: %d", p.id, eatCount))
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	host := make(chan bool, 2)

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = &ChopS{i, sync.Mutex{}}
	}

	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, host)
	}


	wg.Wait()
}
