package main
import ("fmt"
)
type _state_3 struct {
    c chan interface{}
  }
  func init_state_3(c chan interface{}) *_state_3 { return &_state_3{ c } } 
  func (x *_state_3) Send(v interface{}) { x.c <- v }
  func (x *_state_3) Recv() interface{} { return <-x.c }

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

  type _state_0 struct {
    c chan interface{}
    next *_state_1
  }
  
  func init_state_0(c chan interface{}) *_state_0 { return &_state_0{ c, nil } } 
  func (x *_state_0) Send(v int) *_state_1 { if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next}
  func (x *_state_0) Recv() (int, *_state_1) { if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int),x.next}

  type _state_6 struct {
    c chan interface{}
    next *_state_0
  }
  
  func init_state_6(c chan interface{}) *_state_6 { return &_state_6{ c, nil } } 
  func (x *_state_6) Send(v int) *_state_0 { if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next}
  func (x *_state_6) Recv() (int, *_state_0) { if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int),x.next}

  type _state_5 struct {
    c chan interface{}
    next *_state_6
  }
  
  func init_state_5(c chan interface{}) *_state_5 { return &_state_5{ c, nil } } 
  func (x *_state_5) Send(v int) *_state_6 { if x.next == nil { x.next = init_state_6(x.c) }; x.c <- v; return x.next}
  func (x *_state_5) Recv() (int, *_state_6) { if x.next == nil { x.next = init_state_6(x.c) }; return (<-x.c).(int),x.next}

  type _state_4 struct {
    c chan interface{}
    next *_state_5
  }
  
  func init_state_4(c chan interface{}) *_state_4 { return &_state_4{ c, nil } } 
  func (x *_state_4) Send(v int) *_state_5 { if x.next == nil { x.next = init_state_5(x.c) }; x.c <- v; return x.next}
  func (x *_state_4) Recv() (int, *_state_5) { if x.next == nil { x.next = init_state_5(x.c) }; return (<-x.c).(int),x.next}

  func send_three(n int) func (_x *_state_0) {
 return func (c *_state_0){
d := init_state_4(make(chan interface{}))
go func (d *_state_4){
v1, d0 := d.Recv()
v2, d1 := d0.Recv()
v3, d2 := d1.Recv()
d3 := d2.Send(v1)
d4 := d3.Send(v2)
d5 := d4.Send(v3)
d5.Send(nil)
}(d)
d0 := d.Send(n)
d1 := d0.Send((n + 1))
d2 := d1.Send((n + 2))
v1, d3 := d2.Recv()
v2, d4 := d3.Recv()
v3, d5 := d4.Recv()
c0 := c.Send(v1)
c1 := c0.Send(v2)
c2 := c1.Send(v3)
d5.Recv()
c2.Send(nil)
}}
func main () {
m := init_state_3(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_3){
e := init_state_0(make(chan interface{}))
go send_three(1)(e)
a1, e0 := e.Recv()
a2, e1 := e0.Recv()
a3, e2 := e1.Recv()
fmt.Printf("%v\n",a1)
fmt.Printf("%v\n",a2)
fmt.Printf("%v\n",a3)
e2.Recv()
m.Send(nil)
}(m)
}
