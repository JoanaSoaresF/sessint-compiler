package main
import ("fmt"
)
type _state_8 struct {
    c chan interface{}
  }
  func init_state_8(c chan interface{}) *_state_8 { return &_state_8{ c } } 
  func (x *_state_8) Send(v interface{}) { x.c <- v }
  func (x *_state_8) Recv() interface{} { return <-x.c }

  type _state_7 struct {
    c chan interface{}
    next *_state_8
  }
  
  func init_state_7(c chan interface{}) *_state_7 { return &_state_7{ c, nil } } 
  func (x *_state_7) Send(v int) *_state_8 { if x.next == nil { x.next = init_state_8(x.c) }; x.c <- v; return x.next}
  func (x *_state_7) Recv() (int, *_state_8) { if x.next == nil { x.next = init_state_8(x.c) }; return (<-x.c).(int),x.next}

  type _state_6 struct {
    c chan interface{}
    next *_state_7
  }
  
  func init_state_6(c chan interface{}) *_state_6 { return &_state_6{ c, nil } } 
  func (x *_state_6) Send(v int) *_state_7 { if x.next == nil { x.next = init_state_7(x.c) }; x.c <- v; return x.next}
  func (x *_state_6) Recv() (int, *_state_7) { if x.next == nil { x.next = init_state_7(x.c) }; return (<-x.c).(int),x.next}

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

  type _state_3 struct {
    c chan interface{}
    next *_state_4
  }
  
  func init_state_3(c chan interface{}) *_state_3 { return &_state_3{ c, nil } } 
  func (x *_state_3) Send(v int) *_state_4 { if x.next == nil { x.next = init_state_4(x.c) }; x.c <- v; return x.next}
  func (x *_state_3) Recv() (int, *_state_4) { if x.next == nil { x.next = init_state_4(x.c) }; return (<-x.c).(int),x.next}

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

  func send_three(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send(n)
c1 := c0.Send((n + 1))
c2 := c1.Send((n + 2))
a1, c3 := c2.Recv()
c4 := c3.Send(a1)
c5 := c4.Send((a1 - 1))
c6 := c5.Send((a1 - 2))
c7 := c6.Send((a1 - 3))
c7.Send(nil)
}}
func main () {
m := init_state_8(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_8){
e := init_state_0(make(chan interface{}))
go send_three(1)(e)
a1, e0 := e.Recv()
fmt.Printf("%v\n",a1)
a2, e1 := e0.Recv()
fmt.Printf("%v\n",a2)
a3, e2 := e1.Recv()
fmt.Printf("%v\n",a3)
e3 := e2.Send(500)
a4, e4 := e3.Recv()
fmt.Printf("%v\n",a4)
a5, e5 := e4.Recv()
fmt.Printf("%v\n",a5)
a6, e6 := e5.Recv()
fmt.Printf("%v\n",a6)
a7, e7 := e6.Recv()
fmt.Printf("%v\n",a7)
e7.Recv()
m.Send(nil)
}(m)
}
