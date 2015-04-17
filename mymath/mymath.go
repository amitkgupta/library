package mymath

import "math"

func Pi() float64 {
    f := 0.0
    for k := 0; k <= 1000; k++ {
        f += 4 * math.Pow(-1, float64(k)) / (2*float64(k) + 1)
    }
    return f
}
