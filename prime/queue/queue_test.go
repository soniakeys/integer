package queue_test

import (
	"testing"

	"github.com/soniakeys/integer/prime/queue"
)

func TestPQLimit(t *testing.T) {
	if (queue.PQueue{}).Limit() != 1<<64-1 {
		t.Error("queue.PQueue Limit() fail")
	}
}

func Benchmark1e4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.PQueue{}.Iterate(1, 1e4, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.PQueue{}.Iterate(1, 1e5, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.PQueue{}.Iterate(1, 1e6, func(uint64) (terminate bool) {
			return
		})
	}
}

func Benchmark1e7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue.PQueue{}.Iterate(1, 1e7, func(uint64) (terminate bool) {
			return
		})
	}
}
