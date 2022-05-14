package tests

type (
	BenchmarkTable[E any, S ~[]BenchmarkTableEntry[E]] interface {
		Name() string
		Benchmarks() S
	}

	benchmarkTable[E any, S ~[]BenchmarkTableEntry[E]] struct {
		name  string
		tests S
	}
)
