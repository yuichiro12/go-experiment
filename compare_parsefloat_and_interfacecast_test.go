package tests

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_interfaceCastOrParseFloat(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		var f interface{}
		f = rand.Float64()
		_ = f.(float64)
	}
	elapsed := time.Since(start)
	fmt.Println("interface assignment and interface cast:", elapsed.String())

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		var f string
		f = fmt.Sprintf("%f", rand.Float64())
		_, _ = strconv.ParseFloat(f, 64)
	}
	elapsed = time.Since(start)
	fmt.Println("fmt.Sprintf and ParseFloat:", elapsed.String())

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		var f interface{}
		f = rand.Float64()
		_ = f
	}
	elapsed = time.Since(start)
	fmt.Println("interface assignment:", elapsed.String())

	start = time.Now()
	for i := 0; i < 1000000; i++ {
		var f string
		f = fmt.Sprintf("%f", rand.Float64())
		_ = f
	}
	elapsed = time.Since(start)
	fmt.Println("fmt.Sprintf:", elapsed.String())

	var fi interface{}
	fi = rand.Float64()
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_ = fi.(float64)
	}
	elapsed = time.Since(start)
	fmt.Println("interface cast:", elapsed.String())

	var fs string
	fs = fmt.Sprintf("%f", rand.Float64())
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_, _ = strconv.ParseFloat(fs, 64)
	}
	elapsed = time.Since(start)
	fmt.Println("ParseFloat:", elapsed.String())
}
