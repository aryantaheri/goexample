package main 

import (
	"fmt"
	"strings"
//	"golang.org/x/tour/wc"
	//"golang.org/x/tour/pic"
)

type Vertex struct {
	X, Y int
}

func makeVertex(x, y int) Vertex {
	v := Vertex{1, 2}
	v.X = 4
	return v
}

func makeVertexPtr(x, y int) Vertex {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(p)
	return v	
}

func makeVertexLiteral() {
	v1 := Vertex{1, 2}
	v2 := Vertex{Y: 4, X: 3}
	v3 := Vertex{}
	pv4 := &Vertex{5, 6}
	
	fmt.Println("makeVertexLiteral", v1, v2, v3, pv4)
}

func myarrays() [5]string {
	//[n]T
	var a [5]string 
	a[0] = "a"
	a[1] = "b"
	fmt.Println(a)
	array := [3]string{"x", "y", "z"}
	fmt.Printf("array: %T %v\n", array, array)
	fmt.Println(array)
	return a
}

func myslices() {
	//[]T
	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("slice type:%T value:%v\n", slice, slice)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("	slice[%d] == %s\n", i, slice[i])
	} 
	
	fmt.Println("	slice[2:4] == ", slice[2:4])
	fmt.Println("	slice[:4] == ", slice[:4])
	fmt.Println("	slice[3:] == ", slice[3:])
	
	slice2 := make([]int, 3, 5)
	fmt.Println("	slice2 == ", slice2)
	slice2[2] = 2
	slice2 = append(slice2, 3, 4, 5)
	fmt.Println("	slice2 == ", slice2, len(slice2), cap(slice2))
	
	for i, v := range slice {
		fmt.Printf("i=%d v=%s pow=%d \n", i, v, 1<<uint(i))
	}
}

type VertexFloat64 struct {
	Lat, Long float64
}

func makeMap() map[string]VertexFloat64 {
	var m map[string]VertexFloat64
	m = make(map[string]VertexFloat64)
	m["a"] = VertexFloat64{1.1, 2.2}	
	m["b"] = VertexFloat64{3.3, 4.4}
	
	var m2 = map[string]VertexFloat64{
		"x": VertexFloat64{10.1, 20.2},
		"y": VertexFloat64{30.1, 40.2},
	}
	fmt.Println("m2", m2)
	var m3 = map[string]VertexFloat64{
		"x": {10.1, 20.2},
		"y": {30.1, 40.2},
	}
	fmt.Println("m3", m3)
	m3["x1"] = VertexFloat64{-1, -1}
	v, present := m3["x2"]
	fmt.Println("m3", m3, v, present)

	return m
}

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	fields := strings.Fields(s)
	for i, field := range fields {
		fmt.Printf("	field %d %s\n", i, field)
		wc[field] += 1 
	}
	return wc
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func adderTester() {
	pos, neg := adder(), adder()
	fmt.Printf("addTester %T %v\n", pos, pos)
	fmt.Printf("addTester %T %v\n", neg, neg)
	for i := 0; i < 10; i++ {
		fmt.Println("	", i, pos(i), neg(-i))
	} 
}

func fibonacci() func() int {
	f := make([]int, 2, 100)
	f[0] = 0 
	f[1] = 1
	i := -1
	return func() int {
		i++
		if i > 1 {
			f = append(f, f[i-1] + f[i-2])
		}
		return f[i]
	}
}

func fibonacci2() func() int {
	f := 0
	f0 := 0 
	f1 := 1
	i := -1
	return func() int {
		i++
		 if i == 0 {
			return f0
		} else if i == 1 {
			return f1
		}
		f = f0 + f1
		f0 = f1
		f1 = f
		return f
	}
}

func fibonacciTest() {
	f := fibonacci()
	f2 := fibonacci2()
	for i := 0; i < 10; i++ {
		fmt.Println("Fibonacci ", i, f(), f2())
	}
}

func main() {
	fmt.Println("makeVertex", makeVertex(1, 2))
	fmt.Println("makeVertexPtr", makeVertexPtr(1, 2))
	makeVertexLiteral()
	myarrays()
	myslices()
	mymap := makeMap()
	fmt.Printf("makeMap %T %v\n", mymap, mymap)
	fmt.Println("WordCount:", WordCount("a bc d1 x4 zy a d1"))
//	wc.Test(WordCount)
	adderTester()
	fibonacciTest()
}

