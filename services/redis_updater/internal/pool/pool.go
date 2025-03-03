package pool

import (
	"log"
	"time"
)

var amount = 128

type Task func()

type Pool struct {
	task chan Task
	pool chan struct{}
}

func NewPool() *Pool {
	return &Pool{task: make(chan Task), pool: make(chan struct{}, amount)}
}

func (p *Pool) NewTask(t Task) {
	select {
	case p.task <- t:
	case p.pool <- struct{}{}:
		p.newWorker(t)
	}
	log.Printf("running workers: %d/%d", len(p.pool), cap(p.pool))
}

func (p *Pool) newWorker(t Task) {
	go func() {
		timer := time.NewTimer(time.Second * 30)
		var ok bool
	workerLoop:
		for {
			t()
			select {
			case t, ok = <-p.task:
				if !ok {
					break workerLoop
				}
			case <-timer.C:
				<-p.pool
				return
			}
		}
	}()
}
