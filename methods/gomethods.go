package main 

import (
	"fmt"
	"math"
	"os"
	"time"
	"io"
	"strings"
	"image"
	"golang.org/x/tour/pic"
)

type MyFloat float64

type IPAddr [4]byte

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64	
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y  )
}

func (f MyFloat) Abs() float64 {
	return math.Abs(float64(f)) 
}

/**
	If the pointer is used (v *Vertex) the changes (scaling) will be reflected in the parent value
	Otherwise (v Vertex), a copy of values are passed and no changes are made to the parent variables
**/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f	
}

/**
	Implement Stringer's String interface for Vertex
**/

func (v Vertex) String() string {
	return fmt.Sprintf("(Long=%f, Lat=%f)", v.X, v.Y)
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func absTester() float64 {
	v := &Vertex{3, 4}
	return v.Abs()
}

func absScaleTester() {
	v := &Vertex{3, 4}
	fmt.Printf("	Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("	After scaling: %+v, Abs: %v\n", v, v.Abs())
}


func interfaceTester() {
	var a Abser
	f := MyFloat(-4)
	v := Vertex{3, 4}
	
	a = f
	fmt.Println("	MyFloat Abser interface.Abs()", a.Abs())
	a = &v
	fmt.Println("	*Vertex Abser interface.Abs()", a.Abs())
}

func interfaceTester2() {
	var w Writer
	
	w = os.Stdout
	fmt.Fprintf(w, "heheheh\n")
}

func ipStringerTester() {
	addrs := map[string]IPAddr {
		"loopback": IPAddr{127, 0, 0, 1},
		"googleDNS": IPAddr{8, 8, 8, 8},
	}
	
	for name, addr := range addrs {
		fmt.Printf("	%v: %v\n", name, addr)
	}
}

func errorRunner() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errorTester() {
		if err := errorRunner(); err != nil {
		fmt.Println("	", err)
	}
}

func stringReader() {
	reader := strings.NewReader("Hello, 123 456 678 10")
	
	b := make([]byte, 8)
	for {
		number, err := reader.Read(b)
		fmt.Printf("	n = %v error = %v read-bytes = %v\n", number, err, b)
		fmt.Printf("	b[:number] = %q \n", b[:number])
		if err == io.EOF {
			break
		}
	}
}

func imageTester() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println("Bounds = ", m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	pic.ShowImage(m)
}

func main() {
	fmt.Println("absTester ", absTester())
	fmt.Println("absScaleTester ")
	absScaleTester()
	fmt.Println("interfaceTester ")
	interfaceTester()
	fmt.Println("interfaceTester2 ")
	interfaceTester2()
	fmt.Println("ipStringerTester ")
	ipStringerTester()
	fmt.Println("errorTester ")
	errorTester()
	fmt.Println("stringReader ")
	stringReader()
	fmt.Println("imageTester ")
	imageTester()
	
}

