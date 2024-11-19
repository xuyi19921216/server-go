package main

type handlerFunc func()

// 协程池
type WorkerPool struct {
	WorkNum int
	tasks   chan handlerFunc
}

func NewWorkerPool(workNum int) *WorkerPool {
	w := &WorkerPool{WorkNum: workNum, tasks: make(chan handlerFunc)}
	w.start()
	return w
}
func (w *WorkerPool) start() {
	for i := 0; i < w.WorkNum; i++ {
		go func() {
			for task := range w.tasks {
				task()
			}
		}()
	}
}

func (w *WorkerPool) addTask(task handlerFunc) {
	w.tasks <- task
}
