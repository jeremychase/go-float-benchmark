package main

import (
	"testing"
)

var res32 []float32
var res64 []float64

func BenchmarkFloat32with32Samples(b *testing.B) {
	data := load("float32.samples")
	var r []float32

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = fp32(data)
	}

	res32 = r
}

func BenchmarkFloat32with64Samples(b *testing.B) {
	data := load("float64.samples")
	var r []float32

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = fp32(data)
	}

	res32 = r
}

func BenchmarkFloat64with32Samples(b *testing.B) {
	data := load("float32.samples")
	var r []float64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = fp64(data)
	}

	res64 = r
}

func BenchmarkFloat64with64Samples(b *testing.B) {
	data := load("float64.samples")
	var r []float64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = fp64(data)
	}

	res64 = r
}
