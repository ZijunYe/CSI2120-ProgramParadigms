package main

import "fmt"

type Point struct {
	x int
	y int
}

/*var (
	p1 = Point{1, 2}
	p2 = Point{2, 7}
	pp = &Point{3, 4} //pp is a pointer, points to the address
)*/

func main() {
	//dynamic allocation
	//ptr1 := new(Point)
	//ptr2 := &Point{9, 8}

	pt := Point{8, 1}
	complement(&pt)
	fmt.Printf("result = %d and %d\n", pt.x, pt.y)
}

func complement(p *Point) { //need to passing the address in the parameter
	p.x, p.y = -p.y, -p.x
}
