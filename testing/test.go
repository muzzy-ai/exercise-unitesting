package testing

import "math"

//menghitung luas bangun datar persegi , persegi panjang, dan lingkaran

func AreaOfSquare(side float64) float64 {
    return side * side
}

func AreaOfRectangle(length, width float64) float64 {
    return length * width
}

func AreaOfCircle(radius float64) float64 {
    return math.Pi * radius * radius
}