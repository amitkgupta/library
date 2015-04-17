package mymath_test

import (
    "math"

	"github.com/amitkgupta/library/mymath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mymath", func() {
    Describe("Calculating π", func() {
        var result float64
        var n int

        JustBeforeEach(func() {
            result = mymath.Pi(n)
        })

        Context("with small n", func() {
            It("calculates π, with possibly low precision", func() {
                Ω(mymath.Pi(100)).Should(BeNumerically("~", math.Pi, 0.2))
            })
        })

        Context("with large n", func() {
            It("calculates π, with high low precision", func() {
                Ω(mymath.Pi(10000)).Should(BeNumerically("~", math.Pi, 0.00001))
            })
        })
    })
})
