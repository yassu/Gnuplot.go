package main

import (
	"github.com/yassu/gnuplot.go"
)

func NumberOfUnderPrimes(n int) int {
	primes := []int{}
	for j := 2; j <= n; j++ {
		isPrime := true
		for _, p := range primes {
			if j%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, j)
		}
	}
	return len(primes)
}

func main() {
	fun := gnuplot.NewFunction2d()
	fun.SetF(func(x float64) float64 {
		return float64(NumberOfUnderPrimes(int(x + 0.49)))
	})
	fun.Configure("_xMin", []string{"0"})
	fun.Configure("_xMax", []string{"1000"})
	fun.Configure("w", []string{"dots"})

	graph := gnuplot.NewGraph2d()
	graph.Configure("angles", []string{"degrees"})
	graph.Configure("yrange", []string{"[0:1000]"})
	graph.Configure("key", []string{"false"})
	graph.AppendFunc(*fun)
	graph.Run()
}
