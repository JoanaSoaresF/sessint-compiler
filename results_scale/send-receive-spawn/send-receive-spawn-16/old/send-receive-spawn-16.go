package main
import ("fmt"
)
type _state_16 struct {
    c chan interface{}
  }
  func init_state_16(c chan interface{}) *_state_16 { return &_state_16{ c } } 
  func (x *_state_16) Send(v interface{}) { x.c <- v }
  func (x *_state_16) Recv() interface{} { return <-x.c }

  type _state_15 struct {
    c chan interface{}
    next *_state_16
  }
  
  func init_state_15(c chan interface{}) *_state_15 { return &_state_15{ c, nil } } 
  func (x *_state_15) Send(v int) *_state_16 { if x.next == nil { x.next = init_state_16(x.c) }; x.c <- v; return x.next}
  func (x *_state_15) Recv() (int, *_state_16) { if x.next == nil { x.next = init_state_16(x.c) }; return (<-x.c).(int),x.next}

  type _state_14 struct {
    c chan interface{}
    next *_state_15
  }
  
  func init_state_14(c chan interface{}) *_state_14 { return &_state_14{ c, nil } } 
  func (x *_state_14) Send(v int) *_state_15 { if x.next == nil { x.next = init_state_15(x.c) }; x.c <- v; return x.next}
  func (x *_state_14) Recv() (int, *_state_15) { if x.next == nil { x.next = init_state_15(x.c) }; return (<-x.c).(int),x.next}

  type _state_13 struct {
    c chan interface{}
    next *_state_14
  }
  
  func init_state_13(c chan interface{}) *_state_13 { return &_state_13{ c, nil } } 
  func (x *_state_13) Send(v int) *_state_14 { if x.next == nil { x.next = init_state_14(x.c) }; x.c <- v; return x.next}
  func (x *_state_13) Recv() (int, *_state_14) { if x.next == nil { x.next = init_state_14(x.c) }; return (<-x.c).(int),x.next}

  type _state_12 struct {
    c chan interface{}
    next *_state_13
  }
  
  func init_state_12(c chan interface{}) *_state_12 { return &_state_12{ c, nil } } 
  func (x *_state_12) Send(v int) *_state_13 { if x.next == nil { x.next = init_state_13(x.c) }; x.c <- v; return x.next}
  func (x *_state_12) Recv() (int, *_state_13) { if x.next == nil { x.next = init_state_13(x.c) }; return (<-x.c).(int),x.next}

  type _state_11 struct {
    c chan interface{}
    next *_state_12
  }
  
  func init_state_11(c chan interface{}) *_state_11 { return &_state_11{ c, nil } } 
  func (x *_state_11) Send(v int) *_state_12 { if x.next == nil { x.next = init_state_12(x.c) }; x.c <- v; return x.next}
  func (x *_state_11) Recv() (int, *_state_12) { if x.next == nil { x.next = init_state_12(x.c) }; return (<-x.c).(int),x.next}

  type _state_10 struct {
    c chan interface{}
    next *_state_11
  }
  
  func init_state_10(c chan interface{}) *_state_10 { return &_state_10{ c, nil } } 
  func (x *_state_10) Send(v int) *_state_11 { if x.next == nil { x.next = init_state_11(x.c) }; x.c <- v; return x.next}
  func (x *_state_10) Recv() (int, *_state_11) { if x.next == nil { x.next = init_state_11(x.c) }; return (<-x.c).(int),x.next}

  type _state_9 struct {
    c chan interface{}
    next *_state_10
  }
  
  func init_state_9(c chan interface{}) *_state_9 { return &_state_9{ c, nil } } 
  func (x *_state_9) Send(v int) *_state_10 { if x.next == nil { x.next = init_state_10(x.c) }; x.c <- v; return x.next}
  func (x *_state_9) Recv() (int, *_state_10) { if x.next == nil { x.next = init_state_10(x.c) }; return (<-x.c).(int),x.next}

  type _state_8 struct {
    c chan interface{}
    next *_state_9
  }
  
  func init_state_8(c chan interface{}) *_state_8 { return &_state_8{ c, nil } } 
  func (x *_state_8) Send(v int) *_state_9 { if x.next == nil { x.next = init_state_9(x.c) }; x.c <- v; return x.next}
  func (x *_state_8) Recv() (int, *_state_9) { if x.next == nil { x.next = init_state_9(x.c) }; return (<-x.c).(int),x.next}

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

  func send_ints(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send((n + 1))
c1 := c0.Send((n + 2))
c2 := c1.Send((n + 3))
c3 := c2.Send((n + 4))
c4 := c3.Send((n + 5))
c5 := c4.Send((n + 6))
c6 := c5.Send((n + 7))
c7 := c6.Send((n + 8))
c8 := c7.Send((n + 9))
c9 := c8.Send((n + 10))
c10 := c9.Send((n + 11))
c11 := c10.Send((n + 12))
c12 := c11.Send((n + 13))
c13 := c12.Send((n + 14))
c14 := c13.Send((n + 15))
c15 := c14.Send((n + 16))
c15.Send(nil)
}}
func main () {
m := init_state_16(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_16){
e := init_state_0(make(chan interface{}))
go send_ints(1)(e)
a1, e0 := e.Recv()
fmt.Printf("%v\n",a1)
a2, e1 := e0.Recv()
fmt.Printf("%v\n",a2)
a3, e2 := e1.Recv()
fmt.Printf("%v\n",a3)
a4, e3 := e2.Recv()
fmt.Printf("%v\n",a4)
a5, e4 := e3.Recv()
fmt.Printf("%v\n",a5)
a6, e5 := e4.Recv()
fmt.Printf("%v\n",a6)
a7, e6 := e5.Recv()
fmt.Printf("%v\n",a7)
a8, e7 := e6.Recv()
fmt.Printf("%v\n",a8)
a9, e8 := e7.Recv()
fmt.Printf("%v\n",a9)
a10, e9 := e8.Recv()
fmt.Printf("%v\n",a10)
a11, e10 := e9.Recv()
fmt.Printf("%v\n",a11)
a12, e11 := e10.Recv()
fmt.Printf("%v\n",a12)
a13, e12 := e11.Recv()
fmt.Printf("%v\n",a13)
a14, e13 := e12.Recv()
fmt.Printf("%v\n",a14)
a15, e14 := e13.Recv()
fmt.Printf("%v\n",a15)
a16, e15 := e14.Recv()
fmt.Printf("%v\n",a16)
e15.Recv()
m.Send(nil)
}(m)
}
