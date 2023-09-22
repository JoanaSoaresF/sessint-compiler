package main
import ("fmt"
)
type _state_3 struct {
    c chan interface{}
    next *_state_0
  }
  
  func init_state_3(c chan interface{}) *_state_3 { return &_state_3{ c, nil } } 
  func (x *_state_3) Send(v int) *_state_0 { if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next}
  func (x *_state_3) Recv() (int, *_state_0) { if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int),x.next}

  type _state_2 struct {
    c chan interface{}
    next *_state_3
  }
  
  func init_state_2(c chan interface{}) *_state_2 { return &_state_2{ c, nil } } 
  func (x *_state_2) Send(v int) *_state_3 { if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next}
  func (x *_state_2) Recv() (int, *_state_3) { if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int),x.next}

  type _state_1 struct {
    c chan interface{}
    next *_state_2
  }
  
  func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c, nil } } 
  func (x *_state_1) Send(v int) *_state_2 { if x.next == nil { x.next = init_state_2(x.c) }; x.c <- v; return x.next}
  func (x *_state_1) Recv() (int, *_state_2) { if x.next == nil { x.next = init_state_2(x.c) }; return (<-x.c).(int),x.next}

  type _state_4 struct {
    c chan interface{}
  }
  func init_state_4(c chan interface{}) *_state_4 { return &_state_4{ c } } 
  func (x *_state_4) Send(v interface{}) { x.c <- v }
  func (x *_state_4) Recv() interface{} { return <-x.c }

  type _state_0 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_0(c chan interface{}) *_state_0 { m := make(map[string]interface{})
	m["stop"] = init_state_4( c )
	m["div"] = init_state_1( c )
	m["mul"] = init_state_1( c )
	m["sub"] = init_state_1( c )
	m["add"] = init_state_1( c )
	return &_state_0{ c, m } }
  func (x *_state_0) Send(v string) { x.c <- v }
  func (x *_state_0) Recv() string  { return (<-x.c).(string) }

  func calc() func (_x *_state_0) {
 return func (c *_state_0){
label := c.Recv()
switch label {
case "add" :
c0 := c.ls["add"].(*_state_1)
a, c1 := c0.Recv()
b, c2 := c1.Recv()
c3 := c2.Send((a + b))
d := init_state_0(make(chan interface{}))
go calc()(d)
// FWD c d Start
for {
dc3 := c3.Recv()
d.Send(dc3)
switch dc3 {
case "add":
d0 := d.ls["add"].(*_state_1)
c4 := c3.ls["add"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "sub":
d0 := d.ls["sub"].(*_state_1)
c4 := c3.ls["sub"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "mul":
d0 := d.ls["mul"].(*_state_1)
c4 := c3.ls["mul"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "div":
d0 := d.ls["div"].(*_state_1)
c4 := c3.ls["div"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "stop":
d0 := d.ls["stop"].(*_state_4)
c4 := c3.ls["stop"].(*_state_4)
d0.Recv()
c4.Send(nil)
return
}
}
// FWD c d End
case "sub" :
c0 := c.ls["sub"].(*_state_1)
a, c1 := c0.Recv()
b, c2 := c1.Recv()
c3 := c2.Send((a - b))
d := init_state_0(make(chan interface{}))
go calc()(d)
// FWD c d Start
for {
dc3 := c3.Recv()
d.Send(dc3)
switch dc3 {
case "add":
d0 := d.ls["add"].(*_state_1)
c4 := c3.ls["add"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "sub":
d0 := d.ls["sub"].(*_state_1)
c4 := c3.ls["sub"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "mul":
d0 := d.ls["mul"].(*_state_1)
c4 := c3.ls["mul"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "div":
d0 := d.ls["div"].(*_state_1)
c4 := c3.ls["div"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "stop":
d0 := d.ls["stop"].(*_state_4)
c4 := c3.ls["stop"].(*_state_4)
d0.Recv()
c4.Send(nil)
return
}
}
// FWD c d End
case "mul" :
c0 := c.ls["mul"].(*_state_1)
a, c1 := c0.Recv()
b, c2 := c1.Recv()
c3 := c2.Send((a * b))
d := init_state_0(make(chan interface{}))
go calc()(d)
// FWD c d Start
for {
dc3 := c3.Recv()
d.Send(dc3)
switch dc3 {
case "add":
d0 := d.ls["add"].(*_state_1)
c4 := c3.ls["add"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "sub":
d0 := d.ls["sub"].(*_state_1)
c4 := c3.ls["sub"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "mul":
d0 := d.ls["mul"].(*_state_1)
c4 := c3.ls["mul"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "div":
d0 := d.ls["div"].(*_state_1)
c4 := c3.ls["div"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "stop":
d0 := d.ls["stop"].(*_state_4)
c4 := c3.ls["stop"].(*_state_4)
d0.Recv()
c4.Send(nil)
return
}
}
// FWD c d End
case "div" :
c0 := c.ls["div"].(*_state_1)
a, c1 := c0.Recv()
b, c2 := c1.Recv()
c3 := c2.Send((a / b))
d := init_state_0(make(chan interface{}))
go calc()(d)
// FWD c d Start
for {
dc3 := c3.Recv()
d.Send(dc3)
switch dc3 {
case "add":
d0 := d.ls["add"].(*_state_1)
c4 := c3.ls["add"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "sub":
d0 := d.ls["sub"].(*_state_1)
c4 := c3.ls["sub"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "mul":
d0 := d.ls["mul"].(*_state_1)
c4 := c3.ls["mul"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "div":
d0 := d.ls["div"].(*_state_1)
c4 := c3.ls["div"].(*_state_1)
d0c4, c5 := c4.Recv()
d1 := d0.Send(d0c4)
d1c5, c6 := c5.Recv()
d2 := d1.Send(d1c5)
c6d2, c6_d2 := d2.Recv()
d = c6_d2
c3 = c6.Send(c6d2)
case "stop":
d0 := d.ls["stop"].(*_state_4)
c4 := c3.ls["stop"].(*_state_4)
d0.Recv()
c4.Send(nil)
return
}
}
// FWD c d End
case "stop" :
c0 := c.ls["stop"].(*_state_4)
c0.Send(nil)
}
}}
func main () {
c := init_state_4(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_4){
d := init_state_0(make(chan interface{}))
go calc()(d)
d.Send("add")
d0 := d.ls["add"].(*_state_1)
d1 := d0.Send(3)
d2 := d1.Send(4)
v, d3 := d2.Recv()
fmt.Printf("%v\n",v)
d3.Send("stop")
d4 := d3.ls["stop"].(*_state_4)
d4.Recv()
c.Send(nil)
}(c)
}
