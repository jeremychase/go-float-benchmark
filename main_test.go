package main

import (
	"log"
	"strconv"
	"testing"
)

var Result float64

func BenchmarkParseFloatFloat32withFloat32Samples(b *testing.B) {
	data := load("float32.samples")
	var r float64
	var err error

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			r, err = strconv.ParseFloat(data[i], 32)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	Result = r
}

func BenchmarkParseFloatFloat32withFloat64Samples(b *testing.B) {
	data := load("float64.samples")
	var r float64
	var err error

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			r, err = strconv.ParseFloat(data[i], 32)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	Result = r
}

func BenchmarkParseFloatFloat64withFloat32Samples(b *testing.B) {
	data := load("float32.samples")
	var r float64
	var err error

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			r, err = strconv.ParseFloat(data[i], 64)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	Result = r
}

func BenchmarkParseFloatFloat64withFloat64Samples(b *testing.B) {
	data := load("float64.samples")
	var r float64
	var err error

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			r, err = strconv.ParseFloat(data[i], 64)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	Result = r
}
