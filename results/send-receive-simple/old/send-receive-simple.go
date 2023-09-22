package main
import ("fmt"
)
type _state_0 struct {
    c chan interface{}
  }
  func init_state_0(c chan interface{}) *_state_0 { return &_state_0{ c } } 
  func (x *_state_0) Send(v interface{}) { x.c <- v }
  func (x *_state_0) Recv() interface{} { return <-x.c }

  type _state_20 struct {
    c chan interface{}
    next *_state_0
  }
  
  func init_state_20(c chan interface{}) *_state_20 { return &_state_20{ c, nil } } 
  func (x *_state_20) Send(v int) *_state_0 { if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next}
  func (x *_state_20) Recv() (int, *_state_0) { if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int),x.next}

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

  func main () {
c := init_state_0(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_0){
d := init_state_1(make(chan interface{}))
go func (d *_state_1){
v1, d0 := d.Recv()
v2, d1 := d0.Recv()
v3, d2 := d1.Recv()
v4, d3 := d2.Recv()
v5, d4 := d3.Recv()
v6, d5 := d4.Recv()
v7, d6 := d5.Recv()
v8, d7 := d6.Recv()
v9, d8 := d7.Recv()
v10, d9 := d8.Recv()
d10 := d9.Send(v1)
d11 := d10.Send(v2)
d12 := d11.Send(v3)
d13 := d12.Send(v4)
d14 := d13.Send(v5)
d15 := d14.Send(v6)
d16 := d15.Send(v7)
d17 := d16.Send(v8)
d18 := d17.Send(v9)
d19 := d18.Send(v10)
d19.Send(nil)
}(d)
d0 := d.Send(1)
d1 := d0.Send(2)
d2 := d1.Send(3)
d3 := d2.Send(4)
d4 := d3.Send(5)
d5 := d4.Send(6)
d6 := d5.Send(7)
d7 := d6.Send(8)
d8 := d7.Send(9)
d9 := d8.Send(10)
v1, d10 := d9.Recv()
fmt.Printf("%v\n",v1)
v2, d11 := d10.Recv()
fmt.Printf("%v\n",v2)
v3, d12 := d11.Recv()
fmt.Printf("%v\n",v3)
v4, d13 := d12.Recv()
fmt.Printf("%v\n",v4)
v5, d14 := d13.Recv()
fmt.Printf("%v\n",v5)
v6, d15 := d14.Recv()
fmt.Printf("%v\n",v6)
v7, d16 := d15.Recv()
fmt.Printf("%v\n",v7)
v8, d17 := d16.Recv()
fmt.Printf("%v\n",v8)
v9, d18 := d17.Recv()
fmt.Printf("%v\n",v9)
v10, d19 := d18.Recv()
fmt.Printf("%v\n",v10)
fmt.Printf("%v\n",0)
d19.Recv()
c.Send(nil)
}(c)
}
