package main

import (
	"fmt"
	"strconv"	
	"github.com/knoxcarey/supervisor"
)



func panic1() {
	panic("OMG")
}

func panic2(x int) {
	y := 4/(x - x)
	fmt.Println("Answer:", strconv.Itoa(y))
}

func nopanic() {
	fmt.Println("No biggie")
}




func main() {
	
	s0 := supervisor.New(supervisor.ONE_FOR_ONE, supervisor.RestartStrategy{Attempts: 1, Milliseconds: 1})
	s1 := supervisor.New(supervisor.ONE_FOR_ONE, supervisor.RestartStrategy{Attempts: 1, Milliseconds: 1})
	s1.Supervise(func() {panic2(37)})
	s0.Supervise(s1)
	s0.Start()
}
