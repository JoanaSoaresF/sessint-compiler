package main

import (
	"fmt"
)

type _state_1 struct {
	c chan interface{}
}

func init_state_1(c chan interface{}) *_state_1 { return &_state_1{c} }
func (x *_state_1) Send(v interface{})          { x.c <- v }
func (x *_state_1) Recv() interface{}           { return <-x.c }

type _state_3 struct {
	c    chan interface{}
	next *_state_1
}

func init_state_3(c chan interface{}) *_state_3 { return &_state_3{c, nil} }
func (x *_state_3) Send(v interface{}) *_state_1 {
	if x.next == nil {
		x.next = init_state_1(x.c)
	}
	x.c <- v
	return x.next
}
func (x *_state_3) Recv() (interface{}, *_state_1) {
	if x.next == nil {
		x.next = init_state_1(x.c)
	}
	return (<-x.c).(interface{}), x.next
}

type _state_2 struct {
	c    chan interface{}
	next *_state_3
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
func (x *_state_2) Send(v int) *_state_3 {
	if x.next == nil {
		x.next = init_state_3(x.c)
	}
	x.c <- v
	return x.next
}
func (x *_state_2) Recv() (int, *_state_3) {
	if x.next == nil {
		x.next = init_state_3(x.c)
	}
	return (<-x.c).(int), x.next
}

type _state_0 struct {
	c  chan interface{}
	ls map[string]interface{}
}

func init_state_0(c chan interface{}) *_state_0 {
	m := make(map[string]interface{})
	m["cons"] = init_state_2(c)
	m["null"] = init_state_1(c)
	return &_state_0{c, m}
}
func (x *_state_0) Send(v string) { x.c <- v }
func (x *_state_0) Recv() string  { return (<-x.c).(string) }

func null() func(_x *_state_0) {
	return func(c *_state_0) {
		c.Send("null")
		c0 := c.ls["null"].(*_state_1)
		c0.Send(nil)
	}
}
func cons(v int) func(_x *_state_0, l *_state_0) {
	return func(c *_state_0, l *_state_0) {
		c.Send("cons")
		c0 := c.ls["cons"].(*_state_2)
		c1 := c0.Send(v)
		fmt.Printf("%v\n", v)
		l_ := init_state_0(make(chan interface{}))
		go func(l_ *_state_0, l *_state_0) {
			// FWD l_ l Start
			for {
				l_l := l.Recv()
				l_.Send(l_l)
				switch l_l {
				case "null":
					l0 := l.ls["null"].(*_state_1)
					l_0 := l_.ls["null"].(*_state_1)
					l0.Recv()
					l_0.Send(nil)
					return
				case "cons":
					l0 := l.ls["cons"].(*_state_2)
					l_0 := l_.ls["cons"].(*_state_2)
					l_0l0, l1 := l0.Recv()
					l_1 := l_0.Send(l_0l0)
					l_1l1, l2 := l1.Recv()
					l_2 := l_1.Send(l_1l1)
					l2.Recv()
					l_2.Send(nil)
					return
				}
			}
			// FWD l_ l End
		}(l_, l)
		c2 := c1.Send(l_)
		c2.Send(nil)
	}
}
func dealloc() func(_x *_state_1, l *_state_0) {
	return func(c *_state_1, l *_state_0) {
		label := l.Recv()
		switch label {
		case "cons":
			l0 := l.ls["cons"].(*_state_2)
			_, l1 := l0.Recv()
			l_, l2 := l1.Recv()
			l_0 := l_.(*_state_0)
			l2.Recv()
			d := init_state_1(make(chan interface{}))
			go dealloc()(d, l_0)
			// FWD c d Start
			d.Recv()
			c.Send(nil)
			return
		// FWD c d End
		case "null":
			l0 := l.ls["null"].(*_state_1)
			l0.Recv()
			c.Send(nil)
		}
	}
}
func main() {
	c := init_state_1(make(chan interface{}))
	go func() {
		c.Recv()
	}()
	func(c *_state_1) {
		e0 := init_state_0(make(chan interface{}))
		go null()(e0)
		e1 := init_state_0(make(chan interface{}))
		go cons(1)(e1, e0)
		e2 := init_state_0(make(chan interface{}))
		go cons(2)(e2, e1)
		e3 := init_state_0(make(chan interface{}))
		go cons(3)(e3, e2)
		e4 := init_state_1(make(chan interface{}))
		go dealloc()(e4, e3)
		e4.Recv()
		c.Send(nil)
	}(c)
}
