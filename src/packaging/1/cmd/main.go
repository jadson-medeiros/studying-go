package main

import (
	"fmt"

	"github.com/jadson-medeiros/studying-go/src/packaging/1/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}