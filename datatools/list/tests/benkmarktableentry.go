package tests

type (
	BenchmarkTableEntry[T any] interface {
		Name() string
		Got() T
	}

	benchmarkTableEntry[E any, S ~[]TestTableEntry[E]] struct {
		name  string
		tests S
	}
)
