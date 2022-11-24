package main

import (
	"fmt"
	"sync"
)

type taskFunc func(int, bool)

type Task struct {
	f taskFunc
	p bool
}

type WQueue interface {
	Init()
	Add(f taskFunc, p bool)
	Run()
}

type WaitQueue struct {
	queue []Task
	wg    *sync.WaitGroup
}

func (wq *WaitQueue) Init() {
	wq.queue = []Task{}
	wq.wg = new(sync.WaitGroup)
}

func (wq *WaitQueue) Add(f taskFunc, p bool) {
	wq.queue = append(wq.queue, Task{f, p})
}

func (wq *WaitQueue) Run() {
	wq.wg.Add(len(wq.queue))
	for i, task := range wq.queue {
		id := i + 1
		fmt.Printf("run task: %d\n", id)
		go func(idx int, task Task, wg *sync.WaitGroup) {
			defer wg.Done()
			task.f(idx, task.p)
		}(id, task, wq.wg)
	}
	wq.wg.Wait()
}
