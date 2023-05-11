package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetDistance(a *Point) float64 {
	first := math.Pow(float64(p.x-a.x), 2)
	second := math.Pow(float64(p.y-a.y), 2)
	
	return math.Sqrt(first + second)
}

func main() {
	point1 := NewPoint(3, 1)
	point2 := NewPoint(2, 0)

	fmt.Println(point1.GetDistance(point2))
}