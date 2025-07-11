package queue_test

import (
	"sync"
	"testing"

	"github.com/hsdodo/sen-tools/queue"
)

func BenchmarkQueue(b *testing.B) {

	q := queue.NewQueue()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < b.N; i++ {
			q.Put(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < b.N; i++ {
			_, _ = q.Get()
		}
		wg.Done()
	}()

	wg.Wait()
	b.Log(q.Len())
}
