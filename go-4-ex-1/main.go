package main

import "fmt"

func computeGrade(gotPoints, maxPoints float64) float64 {
    if maxPoints == 0 {
        return 1.0 
    }
    return (gotPoints / maxPoints) * 5 + 1
}

func main() {
    gotPoints := 0.0
    maxPoints := 100.0
    grade := computeGrade(gotPoints, maxPoints)
    fmt.Printf("The grade is: %.1f\n", grade)
}
