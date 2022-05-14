package tests

func NewTestTable[G, W any, E TestTableEntry[G, W], S ~[]E](name string) TestTable[G, W, E, S] {
	return &testTable[G, W, E, S]{name, nil}
}
func NewTestTableEntry[E any]() TestTableEntry[E] {}

func NewBenchmarkTable[E any, S ~[]TestTableEntry[E]]() BenchmarkTable[E, S] {}
func NewBenchmarkTableEntry[E any]() BenchmarkTableEntry[E]                  {}
