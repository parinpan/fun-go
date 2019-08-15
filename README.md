
# fun-go by Fachrin Aulia

Fun-go is the answer when you are sick of writing over and over again a looping block that seems giving much redundancy in your project. It's inspired by Javascript programming language to approach a better functional programming paradigm using Golang that at least currently doesn't support iteration methods.

Benchmark Result tested on **MacBook Pro (13-inch, 2016, Two Thunderbolt 3 ports) 2 GHz Intel Core i5 8 GB 1867 MHz LPDDR3** with 1.000.000 slice elements of struct can be found in the test file.
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
