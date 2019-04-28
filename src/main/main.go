package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	_ "golang.org/x/tour/wc"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
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
	v1                     = Vertex{1, 2}  // has type Vertex
	v2                     = Vertex{X: 1}  // Y:0 is implicit
	v3                     = Vertex{}      // X:0 and Y:0
	r                      = &Vertex{1, 2} // has type *Vertex
	m               map[string]Vertex2
	n               = map[string]Vertex2{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
)

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	Lat, Long float64
}

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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	previous := 2.0
	for i := 0; i < 50; i++ {
		previous = z
		z -= (z*z - x) / (2 * z)
		if (previous-z) < 0.0000001 && (z-previous) < 0.0000001 {
			fmt.Println("Number of iterations needed:", i)
			return z
		}
	}
	return z
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dx, dx)

	for x := 0; x < dx; x++ {
		temp := make([]uint8, dy, dy)
		for y := 0; y < dy; y++ {
			temp[y] = uint8(y)
		}
		result[x] = temp
	}

	return result
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for _, v := range words {
		result[v]++
	}
	return result
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

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(Sqrt(2))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday ?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	/*	defer fmt.Println("world")
		fmt.Println("hello")

		fmt.Println("counting")
		for i := 0; i < 10; i++ {
			defer fmt.Println(i)
		}
		fmt.Println("done")*/

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	q := &v
	q.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, r, v2, v3)

	var array [2]string
	array[0] = "Hello"
	array[1] = "World"

	fmt.Println(array[0], array[1])
	fmt.Println(array)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	namesA := names[0:2]
	namesB := names[1:3]
	fmt.Println(namesA, namesB)

	namesB[0] = "XXX"
	fmt.Println(namesA, namesB)
	fmt.Println(names)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = make([]int, 5)
	printSlice2("a", s)

	s = make([]int, 0, 5)
	printSlice2("b", s)

	s = s[:2]
	printSlice2("c", s)

	s = s[2:5]
	printSlice2("d", s)

	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	pic.Show(Pic)

	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(n)

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	value, ok := m["Answer"]
	fmt.Println("The value:", value, "Present?", ok)

	wc.Test(WordCount)

}
