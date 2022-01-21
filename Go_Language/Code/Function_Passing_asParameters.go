package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Distance(p1 Point, p2 Point) (distance float64) {
	distance = math.Sqrt(math.Pow(p1.x-p2.x, 2.0) + math.Pow(p1.y-p2.y, 2.0))
	return //return distance value here
}

//d is represent what kind function will be passing in this parameter
func calc(p1 Point, p2 Point, d func(Point, Point) float64) float64 {
	return d(p1, p2)
}

func main() {
	a := Point{2., 4.}
	b := Point{5., 9.}

	//1.passing function that already defined
	dist := calc(a, b, Distance)
	fmt.Printf("result =%f\n", dist)

	//2. passing function that define right after
	dist = calc(a, b, func(p Point, q Point) float64 {
		return math.Abs(p.x-q.x) + math.Abs(p.y-q.y)
	})

	fmt.Printf("result =%f\n", dist) //%f represent the display value
}
