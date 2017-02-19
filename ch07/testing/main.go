package main

import (
	"fmt"

	"github.com/marklewin/go-web-succinctly/ch07/testing/math"
)

func main() {
	nums := []float64{1, 2, 3, 4}
	avg := math.Average(nums)
	fmt.Println(avg)
}
