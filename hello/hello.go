package main

import (
	"fmt"
	"github.com/aryantaheri/goexample/stringutil"
	"math/rand"
	"math"
	"math/cmplx"
	"runtime"
	)

func add(x int, y int) int {
	return x + y	
}

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return	
}

var c, p, j = true, false, "nope!"	
func vars(){
	var i, k int = 10, 13
	l, m := 15, "m-val"
	fmt.Println(i, k, l, m, c, p, j)
}


var (
	 ToBe bool = false
	 NotToBe = true
	 MaxInt uint64 = 1<<64-1
	 z complex128 = cmplx.Sqrt(-5 + 12i)
	)

func vars2(){
	fmt.Println(MaxInt)
	const f = "%T(%#v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

func conversion(){
	var x, y = 3, 5
	var f = math.Sqrt(float64(x*x + y*y))
	var i = int(f)
	fmt.Printf("%T(%v)", x, x)
	fmt.Printf("%T(%v)", f, f)
	fmt.Printf("%T(%v)", i, i)
}

func inference(){
	v := 42.1 + 12i
	const x = 2
	fmt.Printf("Type of %v is %T", v, v)
	fmt.Printf("Type of %v is %T", x, x)}

func myfor(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	} 
	fmt.Printf("sum is %v\n", sum)
}

func myfor2(){
	sum := 1
	for ; sum < 10;  {
		sum += sum
	} 
	fmt.Printf("sum is %v\n", sum)
}

func ifsqrt(x float64) string {
	if x < 0 {
		return ifsqrt(-x) + "i"
	}	
	return fmt.Sprint(math.Sqrt(x)) 
	}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := float64(1)
	return z
}

func myswitch() {
	switch os := runtime.GOOS; os {
		case "linux":
			fmt.Printf("OS is Linux ", os)
		default:
			fmt.Printf("%s.", os)
	}
}

func mydefer(){
	i := 0
	defer fmt.Println("world", i)
	i++
	defer fmt.Println("xxx", i)
	i++
	fmt.Println("hello", i)
}

func pointers(){
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println("i through pointer", *p) // read i through the pointer
	fmt.Println("i's  pointer", p)
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}

const (
	Big = 1 << 100
	Small 
	)
 
func main() {
	
	fmt.Printf("Hello Aryan.\n")
	fmt.Printf(stringutil.Reverse("\nHello Aryan."))
	fmt.Println("abc", rand.Intn(100))
	fmt.Printf("math %v\n", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
	
	fmt.Printf("%v + %v = %v", 2, 3, add(2, 3))
	fmt.Println()
	a, b := swap("a", "b")
	fmt.Println(a, b)
	
	fmt.Println("split func:")
	fmt.Println(split(17))
	a1, b1 := split(17)
	fmt.Println( a1, b1)
	
	fmt.Println("vars func:", c, p, j)
	vars()
	
	fmt.Println("vars2 func:")
	vars2()
	
	fmt.Println("conversion func:")
	conversion()
	
	fmt.Println("inferencefunc:")
	inference()
	
	fmt.Println("\nmyfor:")
	myfor()
	myfor2()
	
	fmt.Println("\nifsqrt:", ifsqrt(-2), ifsqrt(-16))
	
	fmt.Println("\npow:", pow(3, 2, 10), pow(3, 3, 20))
	
	fmt.Println("\nmyswitch:")
	myswitch()
	
	fmt.Println("\nmydefer:")
	mydefer()
	
	fmt.Println("\npointers:")
	pointers()
	
}

