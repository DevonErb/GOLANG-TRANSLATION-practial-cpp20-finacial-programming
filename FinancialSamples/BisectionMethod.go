package financialsamples

import (
	"fmt"
	"math"
)

const defaultError = 0.001

// MathFunction defines an interface for functions that take a float64 input and return a float64
type MathFunction interface {
	Calculate(value float64) float64
}

// BisectionMethod represents the method for finding roots using the bisection method
type BisectionMethod struct {
	f     MathFunction
	error float64
}

// NewBisectionMethod creates a new BisectionMethod with a given function
func NewBisectionMethod(f MathFunction) *BisectionMethod {
	return &BisectionMethod{
		f:     f,
		error: defaultError,
	}
}

// GetRoot calculates the root of the function between x1 and x2
func (bm *BisectionMethod) GetRoot(x1, x2 float64) float64 {
	var root float64
	for math.Abs(x1-x2) > bm.error {
		x3 := (x1 + x2) / 2
		root = x3
		fmt.Printf("root is %f\n", x3) // This line just for demonstration
		if bm.f.Calculate(x1)*bm.f.Calculate(x3) < 0 {
			x2 = x3
		} else {
			x1 = x3
		}
		if bm.f.Calculate(x1)*bm.f.Calculate(x2) > 0 {
			fmt.Println("function does not converge")
			break
		}
	}
	return root
}

// Example function implementation for MathFunction interface
type F1 struct{}

// Calculate defines the function (x - 1)^3
func (f F1) Calculate(x float64) float64 {
	return (x - 1) * (x - 1) * (x - 1)
}

func bi_test() {
	fmt.Println("---------")
	fmt.Println("BisectionMethod Test")
	fmt.Println("---------")

	f := F1{}
	bm := NewBisectionMethod(f)
	fmt.Printf("The root of the function is %f\n", bm.GetRoot(-100, 100))
}
