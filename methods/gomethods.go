package main 

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y  )
}

func absTester() float64{
	v := &Vertex{3, 4}
	return v.Abs()
}
func main() {
	fmt.Println("absTester ", absTester())
}

