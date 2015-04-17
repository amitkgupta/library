package mymath_test

import (
	"github.com/amitkgupta/library/mymath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mymath", func() {
    Describe("Calculating primes", func() {
        It("calculates the first n primes", func() {
            Ω(mymath.FirstNPrimes(13)).Should(Equal([]int{
                2,
                3,
                5,
                7,
                11,
                13,
                17,
                19,
                23,
                29,
                31,
                37,
                41,
             }))
        })
    })
})
