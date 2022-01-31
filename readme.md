# Comparing strconv.ParseFloat using 32 and 64 bit  

```
Running tool: /home/jchase/sdk/go1.17.6/bin/go test -benchmem -run=^$ -bench ^(BenchmarkParseFloatFloat32withFloat32Samples|BenchmarkParseFloatFloat32withFloat64Samples|BenchmarkParseFloatFloat64withFloat32Samples|BenchmarkParseFloatFloat64withFloat64Samples)$ floaty

goos: linux
goarch: amd64
pkg: floaty
cpu: AMD Ryzen 5 PRO 3500U w/ Radeon Vega Mobile Gfx
BenchmarkParseFloatFloat32withFloat32Samples-8   	  174349	      6019 ns/op	       0 B/op	       0 allocs/op
BenchmarkParseFloatFloat32withFloat64Samples-8   	  126147	      8546 ns/op	       0 B/op	       0 allocs/op
BenchmarkParseFloatFloat64withFloat32Samples-8   	  220278	      5169 ns/op	       0 B/op	       0 allocs/op
BenchmarkParseFloatFloat64withFloat64Samples-8   	  144366	      7750 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	floaty	4.705s
```

### Setup

`float32.samples` and `float64.samples` were created using `go run .`