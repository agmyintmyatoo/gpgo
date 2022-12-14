package main

import (
	"fmt"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, temp := range temperatures {
		frequency[temp]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}

