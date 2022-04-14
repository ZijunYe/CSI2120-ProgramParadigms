package main

import (
	"fmt"
)

type Accumulator interface {
	SetInput([]float64)
	GetResult() float64
}

type Product struct {
	input
	factor float64
	Result float64
}

type input struct {
	inputSlice []float64
}

func (p *Product) GetResult() float64 {
	p.Result = p.factor
	for _, x := range p.input.inputSlice {
		p.Result = p.Result * x
	}
	return p.Result
}
func (i *input) SetInput(inSlide []float64) {
	i.inputSlice = make([]float64, len(inSlide))
	copy(inSlide, i.inputSlice)
	return
}

func main() {
	accu := []Accumulator{&Product{input{[]float64{1.0, 2.0, 4.0}},
		2.0, 0.0}}
	for i, ac := range accu {
		fmt.Printf("%d. Accumulater %f\n", i+1, ac.GetResult())
	}
	return
}
