package main 

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println("	", s, time.Now().String())
	}
}

func routineRunner(){
	go say("hi")
	say("bye")
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Printf("	received=%v sum=%v time=%v\n", a, sum, time.Now())
	c <- sum
}

func channelTester() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
//	x, y := <-c, <-c
	x := <- c
	y := <- c
	
	
	fmt.Println("	", x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y, z := 0, 1, 0
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("	", x, time.Now())
		z = y
		y = x+y
		x = z
//		x, y = y, x+y
	}
	close(c)
}

func fibonacciTester() {
	c := make(chan int, 10)
	go fibonacci(cap(c) + 3, c)
	for i := range c {
		fmt.Println("		", i, time.Now())
	}
}


func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	
	for {
		select {
			case c <- x:
				fmt.Println("	", x, time.Now())
				x, y = y, x+y
			case <- quit:
				fmt.Println("	Quit!", time.Now())
				return
		}
	}
}

func fibonacciTester2() {
	c := make(chan int, 3)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("		", <- c, time.Now())
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
	
	

}

func main() {
	fmt.Println("routineRunner")
//	routineRunner()
	fmt.Println("channelTester")
	channelTester()
	fmt.Println("fibonacciTester")
	fibonacciTester()
	fmt.Println("fibonacciTester2")
	fibonacciTester2()
}

