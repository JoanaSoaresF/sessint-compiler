package main

import (
	"fmt"
)

type _state_0 struct {
	c    chan interface{}
	next *_state_0
}

func init_state_0(c chan interface{}) *_state_0 { return &_state_0{c, nil} }
func (x *_state_0) Send(v int) *_state_0 {
	if x.next == nil {
		x.next = init_state_0(x.c)
	}
	x.c <- v
	return x
}
func (x *_state_0) Recv() (int, *_state_0) {
	if x.next == nil {
		x.next = init_state_0(x.c)
	}
	return (<-x.c).(int), x
}

type _state_1 struct {
	c chan interface{}
}

func init_state_1(c chan interface{}) *_state_1 { return &_state_1{c} }
func (x *_state_1) Send(v interface{})          { x.c <- v }
func (x *_state_1) Recv() interface{}           { return <-x.c }

func nats(n int) func(_x *_state_0) {
	return func(c *_state_0) {
		c0 := c.Send(n)
		fmt.Printf("%v\n", n)
		d := init_state_0(make(chan interface{}))
		go nats((n + 1))(d)
		// FWD c d Start
		for {
			c0d, c0_d := d.Recv()
			d = c0_d
			c0 = c0.Send(c0d)
		}
		// FWD c d End
	}
}
func main() {
	c := init_state_1(make(chan interface{}))
	go func() {
		c.Recv()
	}()
	func(c *_state_1) {
		c.Send(nil)
	}(c)
}
