package main

import "fmt"

// Preamble generation
type _state_1 struct {
	c chan interface{}
}

func init_state_1(c chan interface{}) *_state_1 { return &_state_1{c} }
func (x *_state_1) Send(v interface{})          { x.c <- v }
func (x *_state_1) Recv() interface{}           { return <-x.c }

type _state_0 struct {
	c    chan interface{}
	next *_state_1
}

func init_state_0(c chan interface{}) *_state_0 { return &_state_0{c, nil} }

type _multisend_type__state_0 struct {
	v0 int
	v1 int
	v2 int
}

func (x *_state_0) Send(v0 int, v1 int, v2 int) *_state_1 {
	if x.next == nil {
		x.next = init_state_1(x.c)
	}
	x.c <- _multisend_type__state_0{v0, v1, v2}
	return x.next
}
func (x *_state_0) Recv() (int, int, int, *_state_1) {
	if x.next == nil {
		x.next = init_state_1(x.c)
	}
	ll := <-x.c
	l := ll.(_multisend_type__state_0)
	return l.v0, l.v1, l.v2, x.next
}

//Declaration list compilation
func send_ints_optimized(n int) func(_x *_state_0) {
	return func(c *_state_0) {
		c0 := c.Send(n, (n + 1), (n + 2))
		c0.Send(nil)
	}
}

// Main compilation
func main() {
	m := init_state_1(make(chan interface{}))
	go func() {
		m.Recv()
	}()
	func(m *_state_1) {
		e := init_state_0(make(chan interface{}))
		go send_ints_optimized(1)(e)
		a1, a2, a3, e0 := e.Recv()
		fmt.Printf("%v\n", a1)
		fmt.Printf("%v\n", a2)
		fmt.Printf("%v\n", a3)
		e0.Recv()
		m.Send(nil)
	}(m)
}
