package tests

func NewTestTable[G any, W comparable, S ~[]TestTableEntry[G, W]](name string) TestTable[G, W, S] {
	return &testTable[G, W, S]{name, nil}
}
func NewTestTableEntry[G any, W comparable]() TestTableEntry[G, W] { return &testTableEntry[G, W]{} }

func NewBenchmarkTable[G any, W comparable, S ~[]BenchmarkTableEntry[G, W]]() BenchmarkTable[G, W, S] {
	return &benchmarkTable[G, W, S]{}
}

func NewBenchmarkTableEntry[G any, W comparable]() BenchmarkTableEntry[G, W] {
	return &benchmarkTableEntry[G, W]{}
}
