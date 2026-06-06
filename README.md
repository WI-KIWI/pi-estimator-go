# pi-estimator-go

Estimates π using a parallel Monte Carlo method, written in Go.

## How it works

Random points are sampled uniformly in the unit square `[0, 1) × [0, 1)`. A
point `(x, y)` is inside the quarter-circle of radius 1 when `x² + y² ≤ 1`.
The ratio of inside points to total points converges to π/4, so:

```
π ≈ 4 × (points inside quarter-circle) / (total points)
```

The work is split evenly across all logical CPUs using goroutines, with each
goroutine using its own independent RNG to avoid contention.

## Requirements

- Go 1.21+

## Run

```bash
go run .
```

## Build

```bash
go build -o pi-estimator .
./pi-estimator
```

## Example output

```
using 8 CPUs

samples:         1000  π ≈ 3.124000  error: 0.017593
samples:       100000  π ≈ 3.144360  error: 0.002767
samples:     10000000  π ≈ 3.141679  error: 0.000086
samples:    100000000  π ≈ 3.141576  error: 0.000017
```

## Documentation

```bash
go doc EstimatePi
```
