package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}
func (p *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.x+p2.x, 2) + math.Pow(p.y+p2.y, 2))

}
func main() {
	p1 := NewPoint(1, 5)
	p2 := NewPoint(4, 6)
	fmt.Printf("%.2f", p1.Distance(p2))
}
