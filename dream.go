package main

import (
	"fmt"
)

// lottery stats index interface
type Statsor interface {
	redsum() int
	// redlocate() int
	// redarithprog() bool                  // include arithmetic progression or not
	// redintersectwithall() map[int]int    // intersection with former all issues
	// redintersectwith5union() map[int]int // intersection with former 5 issues' union
	// bluelocate() int
	// bluefreq() int
	// blueNtransfreq() map[int]int // last blue with former 1, 2, 3 transition probability
	// bluemissed() []int
}

type Ssq struct {
	red  []int
	blue int
}

type Dlt struct {
	red  []int
	blue []int
}

func redsum(red []int) int {
	sum := 0
	for i := 0; i < len(red); i++ {
		sum += red[i]
	}
	return sum
}

func (s Ssq) redsum() int {
	return redsum(s.red)
}

func (d Dlt) redsum() int {
	return redsum(d.red)
}

func measure(g Statsor) {
	fmt.Println(g)
	fmt.Println(g.redsum())
}

func main() {
	s := Ssq{red: []int{3, 7, 12, 18, 19, 33}, blue: 6}
	d := Dlt{red: []int{7, 15, 18, 29, 33}, blue: []int{6, 11}}
	measure(s)
	measure(d)
}
