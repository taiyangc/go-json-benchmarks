# Go Benchmarks for JSON encoding/decoding libraries

Because the standard Golang library sometimes isn't enough...

These benchmarks are aimed at providing the latest (Dec 2019) results of various popular drop-in replacements for the Golang encoding/json library.

* encoding/json: Golang standard library
* segmentio/json: https://github.com/segmentio/encoding/tree/master/json
* jsoniter: https://github.com/json-iterator/go
* easyjson: https://github.com/mailru/easyjson

Originally based on jsoniter's [go-benchmark](https://github.com/json-iterator/go-benchmark), these have been go-modules-modernized and improved.

Suggestions and contributions are welcome!

# Benchmark Results

go version go1.13.5 darwin/amd64

MacBook Pro (15-inch, 2018)

2.6 GHz 6-Core Intel Core i7

16 GB 2400 MHz DDR4

benchmark                                       | iter  | time/iter         | bytes/op       | allocs/op           
------------------------------------------------|-------|-------------------|----------------|--------------------
BenchmarkDecodeStdStructSmall-12                | 469744|         2334 ns/op|        272 B/op|         6 allocs/op
BenchmarkEncodeStdStructSmall-12                |2625330|          493 ns/op|        192 B/op|         2 allocs/op
BenchmarkDecodeSegmentioJsonStructSmall-12      |1284361|          889 ns/op|         64 B/op|         2 allocs/op
BenchmarkEncodeSegmentioJsonStructSmall-12      |3632870|          375 ns/op|        192 B/op|         2 allocs/op
BenchmarkDecodeJsoniterStructSmall-12           |1963063|          595 ns/op|         64 B/op|         2 allocs/op
BenchmarkEncodeJsoniterStructSmall-12           |2650510|          516 ns/op|        192 B/op|         2 allocs/op
BenchmarkDecodeEasyJsonSmall-12                 |1741574|          680 ns/op|         64 B/op|         2 allocs/op
BenchmarkEncodeEasyJsonSmall-12                 |4292940|          278 ns/op|        128 B/op|         1 allocs/op
BenchmarkDecodeStdStructMedium-12               |  63357|        19156 ns/op|        496 B/op|        11 allocs/op
BenchmarkEncodeStdStructMedium-12               |1625660|          749 ns/op|        224 B/op|         2 allocs/op
BenchmarkDecodeSegmentioJsonStructMedium-12     | 215605|         5494 ns/op|         80 B/op|         2 allocs/op
BenchmarkEncodeSegmentioJsonStructMedium-12     |2399955|          492 ns/op|        224 B/op|         2 allocs/op
BenchmarkDecodeJsoniterStructMedium-12          | 208089|         5739 ns/op|        384 B/op|        41 allocs/op
BenchmarkEncodeJsoniterStructMedium-12          |1600407|          757 ns/op|        224 B/op|         2 allocs/op
BenchmarkDecodeEasyJsonMedium-12                | 163221|         7059 ns/op|        160 B/op|         4 allocs/op
BenchmarkEncodeEasyJsonMedium-12                |2135684|          568 ns/op|        576 B/op|         3 allocs/op
BenchmarkDecodeStdStructLarge-12                |   5624|       213245 ns/op|        344 B/op|         7 allocs/op
BenchmarkEncodeStdStructLarge-12                |1372980|          853 ns/op|        160 B/op|         2 allocs/op
BenchmarkDecodeSegmentioJsonStructLarge-12      |  13791|        83988 ns/op|          0 B/op|         0 allocs/op
BenchmarkEncodeSegmentioJsonStructLarge-12      |2916681|          415 ns/op|        160 B/op|         2 allocs/op
BenchmarkDecodeJsoniterStructLarge-12           |  13924|        86726 ns/op|      12885 B/op|      1124 allocs/op
BenchmarkEncodeJsoniterStructLarge-12           |2094675|          581 ns/op|        160 B/op|         2 allocs/op
