package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

// EstimatePi estimates π using a parallel Monte Carlo method.
// It spawns one goroutine per logical CPU, each independently sampling
// random points in the unit square and counting how many fall inside
// the quarter-circle of radius 1. The ratio of inside points to total
// points converges to π/4, so the estimate is 4 * inside / total.
// Accuracy improves with larger sample counts at the cost of runtime.
func EstimatePi(samples int) float64 {
	cpus := runtime.NumCPU()
	chunk := samples / cpus
	var inside atomic.Int64

	done := make(chan struct{}, cpus)
	for i := 0; i < cpus; i++ {
		go func(n int) {
			rng := rand.New(rand.NewSource(rand.Int63()))
			count := int64(0)
			for j := 0; j < n; j++ {
				x, y := rng.Float64(), rng.Float64()
				if x*x+y*y <= 1.0 {
					count++
				}
			}
			inside.Add(count)
			done <- struct{}{}
		}(chunk)
	}
	for i := 0; i < cpus; i++ {
		<-done
	}

	return 4.0 * float64(inside.Load()) / float64(chunk*cpus)
}

// main runs EstimatePi across increasing sample sizes and prints the
// estimated value of π, the absolute error, and the elapsed time for each.
func main() {
	fmt.Printf("using %d CPUs\n\n", runtime.NumCPU())
	for _, n := range []int{1_000, 100_000, 10_000_000, 100_000_000, 100_000_000_000} {
		t := time.Now()
		pi := EstimatePi(n)
		elapsed := time.Since(t)
		fmt.Printf("samples: %12d  π ≈ %.6f  error: %.6f  elapsed: %v\n", n, pi, math.Abs(pi-math.Pi), elapsed)
	}
}
