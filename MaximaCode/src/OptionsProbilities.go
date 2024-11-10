/* ------------------------------------------------- */
/* This golang file is the equivalent of 3 cpp files */
/* OptionsProbabilities.h 							 */
/* OptionProbabilityExportedFunctions.cpp 			 */
/* OptionsProbabilities.cpp 						 */
/* ------------------------------------------------- */
/*
   The class is now a struct with methods The main func
   tests it run this by doing go run OptionsProbilities.go
*/
/* ------------------------------------------------- */
package Opt_test

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// OptionsProbabilities struct to hold data
type OptionsProbabilities struct {
	initialPrice  float64
	strike        float64
	avgStep       float64
	numDays       int
	numIterations int
}

// NewOptionsProbabilities creates a new instance of OptionsProbabilities
func NewOptionsProbabilities(initialPrice, strike, avgStep float64, nDays int) *OptionsProbabilities {
	return &OptionsProbabilities{
		initialPrice:  initialPrice,
		strike:        strike,
		avgStep:       avgStep,
		numDays:       nDays,
		numIterations: 1000, // Default number of iterations
	}
}

// SetNumIterations sets the number of iterations for probability calculations
func (o *OptionsProbabilities) SetNumIterations(n int) {
	o.numIterations = n
}

// ProbFinishAboveStrike calculates the probability of finishing above the strike price
func (o *OptionsProbabilities) ProbFinishAboveStrike() float64 {
	nAbove := 0
	for i := 0; i < o.numIterations; i++ {
		val := o.getLastPriceOfWalk()
		if val >= o.strike {
			nAbove++
		}
	}
	return float64(nAbove) / float64(o.numIterations)
}

// ProbFinishBelowStrike calculates the probability of finishing below the strike price
func (o *OptionsProbabilities) ProbFinishBelowStrike() float64 {
	nBelow := 0
	for i := 0; i < o.numIterations; i++ {
		val := o.getLastPriceOfWalk()
		if val <= o.strike {
			nBelow++
		}
	}
	return float64(nBelow) / float64(o.numIterations)
}

// ProbFinalPriceBetweenPrices calculates the probability of the final price being between two values
func (o *OptionsProbabilities) ProbFinalPriceBetweenPrices(lowPrice, highPrice float64) float64 {
	nBetween := 0
	for i := 0; i < o.numIterations; i++ {
		val := o.getLastPriceOfWalk()
		if val >= lowPrice && val <= highPrice {
			nBetween++
		}
	}
	return float64(nBetween) / float64(o.numIterations)
}

// getLastPriceOfWalk simulates a price walk and returns the last price
func (o *OptionsProbabilities) getLastPriceOfWalk() float64 {
	prev := o.initialPrice
	for i := 0; i < o.numDays; i++ {
		stepSize := o.gaussianValue(0, o.avgStep)
		if rand.Intn(2) == 0 {
			prev += stepSize * prev
		} else {
			prev -= stepSize * prev
		}
	}
	return prev
}

// gaussianValue generates a normally distributed random value with given mean and sigma
func (o *OptionsProbabilities) gaussianValue(mean, sigma float64) float64 {
	rand.Seed(time.Now().UnixNano())
	u1 := rand.Float64()
	u2 := rand.Float64()
	z := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return mean + sigma*z
}

// GetWalk simulates a price walk and returns all intermediate prices
func (o *OptionsProbabilities) GetWalk() []float64 {
	walk := make([]float64, o.numDays)
	prev := o.initialPrice
	walk[0] = prev
	for i := 1; i < o.numDays; i++ {
		stepSize := o.gaussianValue(0, o.avgStep)
		if rand.Intn(2) == 0 {
			prev += stepSize * prev
		} else {
			prev -= stepSize * prev
		}
		walk[i] = prev
	}
	return walk
}

func Opt_test() {
	optP := NewOptionsProbabilities(30, 35, 0.01, 100)
	fmt.Println("---------------")
	fmt.Println("OptionsProbilities Test")
	fmt.Println("---------------")
	fmt.Println("Above strike probability:", optP.ProbFinishAboveStrike())
	fmt.Println("Below strike probability:", optP.ProbFinishBelowStrike())
	fmt.Println("Between 28 and 32 probability:", optP.ProbFinalPriceBetweenPrices(28, 32))
	fmt.Println("---------------")
}
