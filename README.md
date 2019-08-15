# fun-go


Benchmark Result
```
goos: darwin
goarch: amd64
pkg: github.com/parinpan/fun-go
BenchmarkFunctionalIteratorMap/functionalIterator-4             2000000000               0.29 ns/op
BenchmarkFunctionalIteratorMap/normalIterator-4                 2000000000               0.13 ns/op
BenchmarkFunctionalIteratorFilter/functionalIterator-4          2000000000               0.14 ns/op
BenchmarkFunctionalIteratorFilter/normalIterator-4              2000000000               0.05 ns/op
BenchmarkFunctionalIteratorForEach/functionalIterator-4         2000000000               0.03 ns/op
BenchmarkFunctionalIteratorForEach/normalIterator-4             2000000000               0.00 ns/op
BenchmarkFunctionalIteratorReduce/functionalIterator-4          2000000000               0.03 ns/op
BenchmarkFunctionalIteratorReduce/normalIterator-4              2000000000               0.00 ns/op
```
