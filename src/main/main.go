package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

const (
	Pi = 3.14
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var (
	c, python, java bool
	ToBe                   = false
	MaxInt          uint64 = 1<<64 - 1
	z                      = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
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

	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
