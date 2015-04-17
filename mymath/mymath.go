package mymath

import "math"
import "runtime"

func Pi() float64 {
    runtime.GOMAXPROCS(runtime.NumCPU())

    ch := make(chan float64)
    for k := 0; k <= 1000; k++ {
        go term(ch, float64(k))
    }
    f := 0.0
    for k := 0; k <= 1000; k++ {
        f += <-ch
    }
    return f
}

func term(ch chan float64, k float64) {
    ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
