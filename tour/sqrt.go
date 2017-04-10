package main

import (
	"fmt"
	"math"
)

// z(n+1) = z(n) - z(n)*z(n) -x / 2* z(n)
// see http://www.guokr.com/question/461510/ to figure the logic behind the
// caculation.
func _sqrt(x float64) float64 {

	var oz, z, delta float64 = 1.0, 1.0, 0.0
	const limit float64 = 0.0000000001

	for i := 0; i < 10; i++ {
		oz = z
		z = z/2 + x/(2*z)

		delta = math.Abs(oz - z)

		fmt.Printf("z = %v, delta = %v\n", z, delta)

		if delta < limit {
			fmt.Printf("%v times reiteration\n", i)
			fmt.Printf("final result is %v\n", z)
			break
		}
	}
	return z
}

func main() {
	fmt.Println(_sqrt(3.0))
}
