package coloredPoint

import (
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
