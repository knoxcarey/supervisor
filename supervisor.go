package supervisor

import (
	"fmt"
)

type exit struct {
	ok    bool
	err   error
}


type SupervisionStrategy int8

const (
	ONE_FOR_ALL   SupervisionStrategy = iota + 1
	ONE_FOR_ONE
	REST_FOR_ONE
)

type RestartStrategy struct {
	Attempts     int              // Try to restart this many times...
	Milliseconds int              // in this many milliseconds before giving up
}


// Actors can be supervisors or processes
type actor interface {
	Start()
}

// Supervisor
type Supervisor struct {
	children []actor
	supst    SupervisionStrategy
	rest     RestartStrategy
}

// Process
type process struct {
	function func()
}


func (p process) Start() {
	done := make(chan exit)

	spawn(p.function, done)

	result := <- done
	if result.ok {
		fmt.Println("Exited normally")
	} else {
		fmt.Println("Exited with error:", result.err)
	}	
}



func New(ss SupervisionStrategy, rs RestartStrategy) *Supervisor {
	return &Supervisor {
		children: make([]actor, 0),
		supst: ss,
		rest: rs,
	}
}


func (s *Supervisor) Supervise(a interface{}) {
	switch v := a.(type) {
	case func():
		s.children = append(s.children, &process{function: v})
	case *Supervisor:
		s.children = append(s.children, v)
	default:
	}
}


func (s *Supervisor) Start() {
	for _, a := range s.children {
		a.Start()
	}
}


func spawn(f func(), done chan<- exit) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				done <- exit{ok: false, err: r.(error)}
			}

		}()
		f()
		done <- exit{ok: true, err: nil}
	}()
}



