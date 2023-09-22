package main
import ("fmt"
)
type _state_32 struct {
    c chan interface{}
  }
  func init_state_32(c chan interface{}) *_state_32 { return &_state_32{ c } } 
  func (x *_state_32) Send(v interface{}) { x.c <- v }
  func (x *_state_32) Recv() interface{} { return <-x.c }

  type _state_31 struct {
    c chan interface{}
    next *_state_32
  }
  
  func init_state_31(c chan interface{}) *_state_31 { return &_state_31{ c, nil } } 
  func (x *_state_31) Send(v int) *_state_32 { if x.next == nil { x.next = init_state_32(x.c) }; x.c <- v; return x.next}
  func (x *_state_31) Recv() (int, *_state_32) { if x.next == nil { x.next = init_state_32(x.c) }; return (<-x.c).(int),x.next}

  type _state_30 struct {
    c chan interface{}
    next *_state_31
  }
  
  func init_state_30(c chan interface{}) *_state_30 { return &_state_30{ c, nil } } 
  func (x *_state_30) Send(v int) *_state_31 { if x.next == nil { x.next = init_state_31(x.c) }; x.c <- v; return x.next}
  func (x *_state_30) Recv() (int, *_state_31) { if x.next == nil { x.next = init_state_31(x.c) }; return (<-x.c).(int),x.next}

  type _state_29 struct {
    c chan interface{}
    next *_state_30
  }
  
  func init_state_29(c chan interface{}) *_state_29 { return &_state_29{ c, nil } } 
  func (x *_state_29) Send(v int) *_state_30 { if x.next == nil { x.next = init_state_30(x.c) }; x.c <- v; return x.next}
  func (x *_state_29) Recv() (int, *_state_30) { if x.next == nil { x.next = init_state_30(x.c) }; return (<-x.c).(int),x.next}

  type _state_28 struct {
    c chan interface{}
    next *_state_29
  }
  
  func init_state_28(c chan interface{}) *_state_28 { return &_state_28{ c, nil } } 
  func (x *_state_28) Send(v int) *_state_29 { if x.next == nil { x.next = init_state_29(x.c) }; x.c <- v; return x.next}
  func (x *_state_28) Recv() (int, *_state_29) { if x.next == nil { x.next = init_state_29(x.c) }; return (<-x.c).(int),x.next}

  type _state_27 struct {
    c chan interface{}
    next *_state_28
  }
  
  func init_state_27(c chan interface{}) *_state_27 { return &_state_27{ c, nil } } 
  func (x *_state_27) Send(v int) *_state_28 { if x.next == nil { x.next = init_state_28(x.c) }; x.c <- v; return x.next}
  func (x *_state_27) Recv() (int, *_state_28) { if x.next == nil { x.next = init_state_28(x.c) }; return (<-x.c).(int),x.next}

  type _state_26 struct {
    c chan interface{}
    next *_state_27
  }
  
  func init_state_26(c chan interface{}) *_state_26 { return &_state_26{ c, nil } } 
  func (x *_state_26) Send(v int) *_state_27 { if x.next == nil { x.next = init_state_27(x.c) }; x.c <- v; return x.next}
  func (x *_state_26) Recv() (int, *_state_27) { if x.next == nil { x.next = init_state_27(x.c) }; return (<-x.c).(int),x.next}

  type _state_25 struct {
    c chan interface{}
    next *_state_26
  }
  
  func init_state_25(c chan interface{}) *_state_25 { return &_state_25{ c, nil } } 
  func (x *_state_25) Send(v int) *_state_26 { if x.next == nil { x.next = init_state_26(x.c) }; x.c <- v; return x.next}
  func (x *_state_25) Recv() (int, *_state_26) { if x.next == nil { x.next = init_state_26(x.c) }; return (<-x.c).(int),x.next}

  type _state_24 struct {
    c chan interface{}
    next *_state_25
  }
  
  func init_state_24(c chan interface{}) *_state_24 { return &_state_24{ c, nil } } 
  func (x *_state_24) Send(v int) *_state_25 { if x.next == nil { x.next = init_state_25(x.c) }; x.c <- v; return x.next}
  func (x *_state_24) Recv() (int, *_state_25) { if x.next == nil { x.next = init_state_25(x.c) }; return (<-x.c).(int),x.next}

  type _state_23 struct {
    c chan interface{}
    next *_state_24
  }
  
  func init_state_23(c chan interface{}) *_state_23 { return &_state_23{ c, nil } } 
  func (x *_state_23) Send(v int) *_state_24 { if x.next == nil { x.next = init_state_24(x.c) }; x.c <- v; return x.next}
  func (x *_state_23) Recv() (int, *_state_24) { if x.next == nil { x.next = init_state_24(x.c) }; return (<-x.c).(int),x.next}

  type _state_22 struct {
    c chan interface{}
    next *_state_23
  }
  
  func init_state_22(c chan interface{}) *_state_22 { return &_state_22{ c, nil } } 
  func (x *_state_22) Send(v int) *_state_23 { if x.next == nil { x.next = init_state_23(x.c) }; x.c <- v; return x.next}
  func (x *_state_22) Recv() (int, *_state_23) { if x.next == nil { x.next = init_state_23(x.c) }; return (<-x.c).(int),x.next}

  type _state_21 struct {
    c chan interface{}
    next *_state_22
  }
  
  func init_state_21(c chan interface{}) *_state_21 { return &_state_21{ c, nil } } 
  func (x *_state_21) Send(v int) *_state_22 { if x.next == nil { x.next = init_state_22(x.c) }; x.c <- v; return x.next}
  func (x *_state_21) Recv() (int, *_state_22) { if x.next == nil { x.next = init_state_22(x.c) }; return (<-x.c).(int),x.next}

  type _state_20 struct {
    c chan interface{}
    next *_state_21
  }
  
  func init_state_20(c chan interface{}) *_state_20 { return &_state_20{ c, nil } } 
  func (x *_state_20) Send(v int) *_state_21 { if x.next == nil { x.next = init_state_21(x.c) }; x.c <- v; return x.next}
  func (x *_state_20) Recv() (int, *_state_21) { if x.next == nil { x.next = init_state_21(x.c) }; return (<-x.c).(int),x.next}

  type _state_19 struct {
    c chan interface{}
    next *_state_20
  }
  
  func init_state_19(c chan interface{}) *_state_19 { return &_state_19{ c, nil } } 
  func (x *_state_19) Send(v int) *_state_20 { if x.next == nil { x.next = init_state_20(x.c) }; x.c <- v; return x.next}
  func (x *_state_19) Recv() (int, *_state_20) { if x.next == nil { x.next = init_state_20(x.c) }; return (<-x.c).(int),x.next}

  type _state_18 struct {
    c chan interface{}
    next *_state_19
  }
  
  func init_state_18(c chan interface{}) *_state_18 { return &_state_18{ c, nil } } 
  func (x *_state_18) Send(v int) *_state_19 { if x.next == nil { x.next = init_state_19(x.c) }; x.c <- v; return x.next}
  func (x *_state_18) Recv() (int, *_state_19) { if x.next == nil { x.next = init_state_19(x.c) }; return (<-x.c).(int),x.next}

  type _state_17 struct {
    c chan interface{}
    next *_state_18
  }
  
  func init_state_17(c chan interface{}) *_state_17 { return &_state_17{ c, nil } } 
  func (x *_state_17) Send(v int) *_state_18 { if x.next == nil { x.next = init_state_18(x.c) }; x.c <- v; return x.next}
  func (x *_state_17) Recv() (int, *_state_18) { if x.next == nil { x.next = init_state_18(x.c) }; return (<-x.c).(int),x.next}

  type _state_16 struct {
    c chan interface{}
    next *_state_17
  }
  
  func init_state_16(c chan interface{}) *_state_16 { return &_state_16{ c, nil } } 
  func (x *_state_16) Send(v int) *_state_17 { if x.next == nil { x.next = init_state_17(x.c) }; x.c <- v; return x.next}
  func (x *_state_16) Recv() (int, *_state_17) { if x.next == nil { x.next = init_state_17(x.c) }; return (<-x.c).(int),x.next}

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
c16 := c15.Send((n + 17))
c17 := c16.Send((n + 18))
c18 := c17.Send((n + 19))
c19 := c18.Send((n + 20))
c20 := c19.Send((n + 21))
c21 := c20.Send((n + 22))
c22 := c21.Send((n + 23))
c23 := c22.Send((n + 24))
c24 := c23.Send((n + 25))
c25 := c24.Send((n + 26))
c26 := c25.Send((n + 27))
c27 := c26.Send((n + 28))
c28 := c27.Send((n + 29))
c29 := c28.Send((n + 30))
c30 := c29.Send((n + 31))
c31 := c30.Send((n + 32))
c31.Send(nil)
}}
func main () {
m := init_state_32(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_32){
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
a17, e16 := e15.Recv()
fmt.Printf("%v\n",a17)
a18, e17 := e16.Recv()
fmt.Printf("%v\n",a18)
a19, e18 := e17.Recv()
fmt.Printf("%v\n",a19)
a20, e19 := e18.Recv()
fmt.Printf("%v\n",a20)
a21, e20 := e19.Recv()
fmt.Printf("%v\n",a21)
a22, e21 := e20.Recv()
fmt.Printf("%v\n",a22)
a23, e22 := e21.Recv()
fmt.Printf("%v\n",a23)
a24, e23 := e22.Recv()
fmt.Printf("%v\n",a24)
a25, e24 := e23.Recv()
fmt.Printf("%v\n",a25)
a26, e25 := e24.Recv()
fmt.Printf("%v\n",a26)
a27, e26 := e25.Recv()
fmt.Printf("%v\n",a27)
a28, e27 := e26.Recv()
fmt.Printf("%v\n",a28)
a29, e28 := e27.Recv()
fmt.Printf("%v\n",a29)
a30, e29 := e28.Recv()
fmt.Printf("%v\n",a30)
a31, e30 := e29.Recv()
fmt.Printf("%v\n",a31)
a32, e31 := e30.Recv()
fmt.Printf("%v\n",a32)
e31.Recv()
m.Send(nil)
}(m)
}
