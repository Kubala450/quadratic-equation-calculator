package main

import (
	"fmt"
	"math"
)

func deltaFormula(a float64, b float64, c float64) *float64 {
	delta := math.Pow(b, 2) - 4*a*c
	if delta < 0 {
		return nil
	}
	return &delta
}

func calcRoots(delta float64, a float64, b float64) (*float64, *float64) {
	if delta <= 0 {
		return nil, nil
	}
	x1 := (-b - math.Sqrt(delta)) / (2 * a)
	x2 := (-b + math.Sqrt(delta)) / (2 * a)
	return &x1, &x2
}

func vertexes(delta float64, a float64, b float64) (*float64, *float64, *float64) {
	if delta <= 0 {
		return nil, nil, nil
	}
	x0 := -b / (2 * a)
	pVertex := -b / (2 * a)
	qVertex := -delta / (4 * a)
	return &x0, &pVertex, &qVertex
}

func printResults(label string, result *float64) {
	if result != nil {
		fmt.Printf("%s: %v\n", label, *result)
	} else {
		fmt.Printf("%s: nil\n", label)
	}
}

func main() {
	var a, b, c float64 = 1, -6, 2
	delta := deltaFormula(a, b, c)
	if delta == nil {
		fmt.Println("delta is less than 0!")
		return
	}

	fmt.Printf("Delta: %v\n", *delta)

	x1, x2 := calcRoots(*delta, a, b)
	printResults("x1", x1)
	printResults("x2", x2)

	x0, pVertex, qVertex := vertexes(*delta, a, b)
	printResults("x0", x0)
	printResults("pVertex", pVertex)
	printResults("qVertex", qVertex)
}
