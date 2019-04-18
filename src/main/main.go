package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

const Pi = 3.14

var (
	c, python, java bool
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum -x
	return
}

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is ", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Printf("Pi is about: %g\n ", math.Pi)

	fmt.Println("42 + 13 = ", add(42, 13))

	fmt.Println(swap("hello", "world"))

	fmt.Println(split(10))

	var i int
	fmt.Println(i, c, python, java)

	var c2, python2, java2 = true, false, "no"
	fmt.Println(c2, python2, java2)

	k := 3
	fmt.Println(k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	a, b := 3, 4
	f := math.Sqrt(float64(a*a + b*b))
	z := uint(f)
	fmt.Println(a, b, z)
	
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
