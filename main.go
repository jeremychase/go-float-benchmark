package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

// fp32 is for testing parsing to 32 bit
func fp32(data []string) (result []float32) {

	for i := 0; i < len(data); i++ {
		f, err := strconv.ParseFloat(data[i], 32)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, float32(f))
	}

	return
}

// fp64 is for testing parsing to 64 bit
func fp64(data []string) (result []float64) {

	for i := 0; i < len(data); i++ {
		f, err := strconv.ParseFloat(data[i], 64)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, f)
	}

	return
}

// main is only used to generate the sample inputs
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
		result = append(result, fmt.Sprintf("%v\n", rand.Float32()))
	}
	return
}

func genFloat64() (result []string) {
	for i := 0; i < 100; i++ {
		result = append(result, fmt.Sprintf("%v\n", rand.Float64()))
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
