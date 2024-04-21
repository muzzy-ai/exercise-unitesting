package testing_test

import (
	"math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"hitung/luas/testing"
)

var _ = Describe("Main", func() {
	Context("Square", func() {
        It("Should calculate the area of a square correctly", func() {
            side := 5.0
            expectedArea := 25.0
            Expect(testing.AreaOfSquare(side)).To(Equal(expectedArea))
        })
	

	})

	Context("Rectangle", func() {
        It("Should calculate the area of a rectangle correctly", func() {
            length := 5.0
            width := 4.0
            expectedArea := 20.0
            Expect(testing.AreaOfRectangle(length, width)).To(Equal(expectedArea))
        })
    })

	Context("Circle", func() {
        It("Should calculate the area of a circle correctly", func() {
            radius := 3.0
            expectedArea := math.Pi * radius * radius
            Expect(testing.AreaOfCircle(radius)).To(Equal(expectedArea))
        })
    })
})