package worker

import (
	"log"
	"time"
)

type JobI interface {
	Do() error
	Stop()
}

type Worker struct {
	job        JobI
	timeTicker *time.Ticker
	stop       chan struct{}
}

func NewWorker(job JobI, duration time.Duration) Worker {
	return Worker{
		job:        job,
		timeTicker: time.NewTicker(duration),
		stop:       make(chan struct{}),
	}
}

func (w Worker) Do() {
	go func() {
		for {
			select {
			case <-w.timeTicker.C:
				err := w.job.Do()
				if err != nil {
					log.Printf("WARNING, job error: %v", err)
				}
			case <-w.stop:
				w.job.Stop()
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	w.stop <- struct{}{}
}
