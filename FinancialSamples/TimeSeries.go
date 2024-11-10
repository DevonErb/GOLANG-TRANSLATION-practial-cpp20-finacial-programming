package financialsamples

import (
	"fmt"
	"math"
)

/*
   TimeSeries();
   TimeSeries(const TimeSeries &);
   TimeSeries &operator=(const TimeSeries &);
   ~TimeSeries();

   void addValue(double val);
   double stdDev();
   double mean();
   size_t size();
   double elem(int i);
*/
//
type TimeSeries struct {
	m_values []float64
}

func newTimeSeries() *TimeSeries {
	return &TimeSeries{
		m_values: []float64{},
	}
}

func (o *TimeSeries) operatorEquals(ts *TimeSeries) {
	if o != ts {
		o.m_values = make([]float64, len(ts.m_values))
		copy(o.m_values, ts.m_values)
	}
}

func (o *TimeSeries) addValue(val float64) {
	o.m_values = append(o.m_values, val)
}

func (o *TimeSeries) mean() float64 {
	var sum float64
	sum = 0

	for i := 0; i < len(o.m_values); i++ {
		sum += o.m_values[i]
	}
	fmt.Printf(" average is %v \n", (sum / float64(len(o.m_values))))
	return sum / float64(len(o.m_values))
}

func (o *TimeSeries) stdDev() float64 {
	var stddev float64
	stddev = 0

	mean := o.mean()

	for i := 0; i < len(o.m_values); i++ {
		stddev += math.Pow((o.m_values[i] - mean), 2)
	}

	return math.Sqrt(stddev / float64((len(o.m_values) - 1)))
}

func (o *TimeSeries) size() int {
	return len(o.m_values)
}

func (o *TimeSeries) elem(pos int64) float64 {
	return o.m_values[pos]
}

func TimeSeries_test() {
	fmt.Println("---------------")
	fmt.Println("TimeSeries Test")
	fmt.Println("---------------")
	oTime := newTimeSeries()

	oTime.addValue(1)
	oTime.addValue(2)
	oTime.addValue(3)
	oTime.addValue(4)
	oTime.addValue(5)
	oTime.addValue(6)

	fmt.Printf(" mean  %v \n", oTime.mean())

	fmt.Printf(" stdDev  %v \n", oTime.stdDev())

	fmt.Printf(" size  %v \n", oTime.size())

	fmt.Printf(" element at 3  %v \n", oTime.elem(3))

	oTime2 := newTimeSeries()

	oTime2.operatorEquals(oTime)

	fmt.Printf(" element at 3  %v \n", oTime2.elem(3))

	fmt.Println("---------------")
}
