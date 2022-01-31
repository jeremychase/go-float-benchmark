# Comparing strconv.ParseFloat using 32 and 64 bit  

```
Running tool: /home/jchase/sdk/go1.17.6/bin/go test -benchmem -run=^$ -bench ^(BenchmarkFloat32with32Samples|BenchmarkFloat32with64Samples|BenchmarkFloat64with32Samples|BenchmarkFloat64with64Samples)$ floaty

goos: linux
goarch: amd64
pkg: floaty
cpu: AMD Ryzen 5 PRO 3500U w/ Radeon Vega Mobile Gfx
BenchmarkFloat32with32Samples-8   	  137096	      8921 ns/op	    1016 B/op	       7 allocs/op
BenchmarkFloat32with64Samples-8   	  116844	     10729 ns/op	    1016 B/op	       7 allocs/op
BenchmarkFloat64with32Samples-8   	  123207	      8826 ns/op	    2040 B/op	       8 allocs/op
BenchmarkFloat64with64Samples-8   	  112838	     12489 ns/op	    2040 B/op	       8 allocs/op
PASS
ok  	floaty	6.378s
```

### Setup

`float32.samples` and `float64.samples` were created using `go run .`