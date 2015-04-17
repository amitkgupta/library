package mymath_test

import (
    "math"

	"github.com/amitkgupta/library/mymath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mymath", func() {
    It("calculates π", func() {
        Ω(mymath.Pi()).Should(BeNumerically("~", math.Pi, 0.001))
    })
})
