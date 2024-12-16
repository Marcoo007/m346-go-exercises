package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	a := 3.0
	b := 4.0
	hypotenuse := computeHypotenuse(a, b)
	fmt.Printf("Die Hypotenuse des Dreiecks mit den Seitenlängen %.2f und %.2f ist %.2f\n", a, b, hypotenuse)
}
