package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// main is used to generate the sample inputs
func main() {

	f64, err := os.Create("float64.samples")
	if err != nil {
		log.Fatal(err)
	}
	defer f64.Close()
	for _, v := range genFloat64() {
		f64.WriteString(v)
	}

	f32, err := os.Create("float32.samples")
	if err != nil {
		log.Fatal(err)
	}
	defer f32.Close()
	for _, v := range genFloat32() {
		f32.WriteString(v)
	}

}

func genFloat32() (result []string) {
	for i := 0; i < 100; i++ {
		result = append(result, fmt.Sprintf("%g\n", rand.Float32()))
	}
	return
}

func genFloat64() (result []string) {
	for i := 0; i < 100; i++ {
		result = append(result, fmt.Sprintf("%g\n", rand.Float64()))
	}
	return
}

func load(file string) (data []string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scnnr := bufio.NewScanner(f)
	for scnnr.Scan() {
		data = append(data, scnnr.Text())
	}

	return
}
