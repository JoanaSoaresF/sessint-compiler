package main
import "fmt"
//Preamble generation
type _state_1 struct {
    c chan interface{}
}
func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c } }
func (x *_state_1) Send(v interface{}) { x.c <- v }
func (x *_state_1) Recv() interface{} { return <-x.c }

  type _state_0 struct {
   c chan interface{}
   next *_state_1
}

func init_state_0(c chan interface{}) *_state_0 { return &_state_0{c, nil} }
type _multisend_type__state_0 struct {
v0 int
v1 int
v2 int
v3 int
v4 int
v5 int
v6 int
v7 int
v8 int
v9 int
v10 int
v11 int
v12 int
v13 int
v14 int
v15 int
}
func (x *_state_0) Send(v0 int, v1 int, v2 int, v3 int, v4 int, v5 int, v6 int, v7 int, v8 int, v9 int, v10 int, v11 int, v12 int, v13 int, v14 int, v15 int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) };
 x.c <- _multisend_type__state_0{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15}
return x.next }
func (x *_state_0) Recv() (int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) };ll := <- x.c
l := ll.(_multisend_type__state_0)
return l.v0, l.v1, l.v2, l.v3, l.v4, l.v5, l.v6, l.v7, l.v8, l.v9, l.v10, l.v11, l.v12, l.v13, l.v14, l.v15, x.next }

type _state_17 struct {
   c chan interface{}
   next *_state_1
}

func init_state_17(c chan interface{}) *_state_17 { return &_state_17{c, nil} }
func (x *_state_17) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_17) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_16 struct {
   c chan interface{}
   next *_state_17
}

func init_state_16(c chan interface{}) *_state_16 { return &_state_16{c, nil} }
func (x *_state_16) Send(v int) *_state_17 {
   if x.next == nil { x.next = init_state_17(x.c) }; x.c <- v; return x.next }
func (x *_state_16) Recv() (int, *_state_17) {
   if x.next == nil { x.next = init_state_17(x.c) }; return (<-x.c).(int), x.next }

  type _state_15 struct {
   c chan interface{}
   next *_state_16
}

func init_state_15(c chan interface{}) *_state_15 { return &_state_15{c, nil} }
func (x *_state_15) Send(v int) *_state_16 {
   if x.next == nil { x.next = init_state_16(x.c) }; x.c <- v; return x.next }
func (x *_state_15) Recv() (int, *_state_16) {
   if x.next == nil { x.next = init_state_16(x.c) }; return (<-x.c).(int), x.next }

  type _state_14 struct {
   c chan interface{}
   next *_state_15
}

func init_state_14(c chan interface{}) *_state_14 { return &_state_14{c, nil} }
func (x *_state_14) Send(v int) *_state_15 {
   if x.next == nil { x.next = init_state_15(x.c) }; x.c <- v; return x.next }
func (x *_state_14) Recv() (int, *_state_15) {
   if x.next == nil { x.next = init_state_15(x.c) }; return (<-x.c).(int), x.next }

  type _state_13 struct {
   c chan interface{}
   next *_state_14
}

func init_state_13(c chan interface{}) *_state_13 { return &_state_13{c, nil} }
func (x *_state_13) Send(v int) *_state_14 {
   if x.next == nil { x.next = init_state_14(x.c) }; x.c <- v; return x.next }
func (x *_state_13) Recv() (int, *_state_14) {
   if x.next == nil { x.next = init_state_14(x.c) }; return (<-x.c).(int), x.next }

  type _state_12 struct {
   c chan interface{}
   next *_state_13
}

func init_state_12(c chan interface{}) *_state_12 { return &_state_12{c, nil} }
func (x *_state_12) Send(v int) *_state_13 {
   if x.next == nil { x.next = init_state_13(x.c) }; x.c <- v; return x.next }
func (x *_state_12) Recv() (int, *_state_13) {
   if x.next == nil { x.next = init_state_13(x.c) }; return (<-x.c).(int), x.next }

  type _state_11 struct {
   c chan interface{}
   next *_state_12
}

func init_state_11(c chan interface{}) *_state_11 { return &_state_11{c, nil} }
func (x *_state_11) Send(v int) *_state_12 {
   if x.next == nil { x.next = init_state_12(x.c) }; x.c <- v; return x.next }
func (x *_state_11) Recv() (int, *_state_12) {
   if x.next == nil { x.next = init_state_12(x.c) }; return (<-x.c).(int), x.next }

  type _state_10 struct {
   c chan interface{}
   next *_state_11
}

func init_state_10(c chan interface{}) *_state_10 { return &_state_10{c, nil} }
func (x *_state_10) Send(v int) *_state_11 {
   if x.next == nil { x.next = init_state_11(x.c) }; x.c <- v; return x.next }
func (x *_state_10) Recv() (int, *_state_11) {
   if x.next == nil { x.next = init_state_11(x.c) }; return (<-x.c).(int), x.next }

  type _state_9 struct {
   c chan interface{}
   next *_state_10
}

func init_state_9(c chan interface{}) *_state_9 { return &_state_9{c, nil} }
func (x *_state_9) Send(v int) *_state_10 {
   if x.next == nil { x.next = init_state_10(x.c) }; x.c <- v; return x.next }
func (x *_state_9) Recv() (int, *_state_10) {
   if x.next == nil { x.next = init_state_10(x.c) }; return (<-x.c).(int), x.next }

  type _state_8 struct {
   c chan interface{}
   next *_state_9
}

func init_state_8(c chan interface{}) *_state_8 { return &_state_8{c, nil} }
func (x *_state_8) Send(v int) *_state_9 {
   if x.next == nil { x.next = init_state_9(x.c) }; x.c <- v; return x.next }
func (x *_state_8) Recv() (int, *_state_9) {
   if x.next == nil { x.next = init_state_9(x.c) }; return (<-x.c).(int), x.next }

  type _state_7 struct {
   c chan interface{}
   next *_state_8
}

func init_state_7(c chan interface{}) *_state_7 { return &_state_7{c, nil} }
func (x *_state_7) Send(v int) *_state_8 {
   if x.next == nil { x.next = init_state_8(x.c) }; x.c <- v; return x.next }
func (x *_state_7) Recv() (int, *_state_8) {
   if x.next == nil { x.next = init_state_8(x.c) }; return (<-x.c).(int), x.next }

  type _state_6 struct {
   c chan interface{}
   next *_state_7
}

func init_state_6(c chan interface{}) *_state_6 { return &_state_6{c, nil} }
func (x *_state_6) Send(v int) *_state_7 {
   if x.next == nil { x.next = init_state_7(x.c) }; x.c <- v; return x.next }
func (x *_state_6) Recv() (int, *_state_7) {
   if x.next == nil { x.next = init_state_7(x.c) }; return (<-x.c).(int), x.next }

  type _state_5 struct {
   c chan interface{}
   next *_state_6
}

func init_state_5(c chan interface{}) *_state_5 { return &_state_5{c, nil} }
func (x *_state_5) Send(v int) *_state_6 {
   if x.next == nil { x.next = init_state_6(x.c) }; x.c <- v; return x.next }
func (x *_state_5) Recv() (int, *_state_6) {
   if x.next == nil { x.next = init_state_6(x.c) }; return (<-x.c).(int), x.next }

  type _state_4 struct {
   c chan interface{}
   next *_state_5
}

func init_state_4(c chan interface{}) *_state_4 { return &_state_4{c, nil} }
func (x *_state_4) Send(v int) *_state_5 {
   if x.next == nil { x.next = init_state_5(x.c) }; x.c <- v; return x.next }
func (x *_state_4) Recv() (int, *_state_5) {
   if x.next == nil { x.next = init_state_5(x.c) }; return (<-x.c).(int), x.next }

  type _state_3 struct {
   c chan interface{}
   next *_state_4
}

func init_state_3(c chan interface{}) *_state_3 { return &_state_3{c, nil} }
func (x *_state_3) Send(v int) *_state_4 {
   if x.next == nil { x.next = init_state_4(x.c) }; x.c <- v; return x.next }
func (x *_state_3) Recv() (int, *_state_4) {
   if x.next == nil { x.next = init_state_4(x.c) }; return (<-x.c).(int), x.next }

  type _state_2 struct {
   c chan interface{}
   next *_state_3
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
func (x *_state_2) Send(v int) *_state_3 {
   if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next }
func (x *_state_2) Recv() (int, *_state_3) {
   if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int), x.next }

  //Declaration list compilation
func send_ints_optimized(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send((n + 1), (n + 2), (n + 3), (n + 4), (n + 5), (n + 6), (n + 7), (n + 8), (n + 9), (n + 10), (n + 11), (n + 12), (n + 13), (n + 14), (n + 15), (n + 16))
c0.Send(nil)
}}
func send_ints(n int) func (_x *_state_2) {
 return func (c *_state_2){
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
//Main compilation
func main () {
    m:= init_state_1(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_1){
e := init_state_0(make(chan interface{}))
go send_ints_optimized(1)(e)
a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, e0 := e.Recv()
fmt.Printf("%v\n",a1)
fmt.Printf("%v\n",a2)
fmt.Printf("%v\n",a3)
fmt.Printf("%v\n",a4)
fmt.Printf("%v\n",a5)
fmt.Printf("%v\n",a6)
fmt.Printf("%v\n",a7)
fmt.Printf("%v\n",a8)
fmt.Printf("%v\n",a9)
fmt.Printf("%v\n",a10)
fmt.Printf("%v\n",a11)
fmt.Printf("%v\n",a12)
fmt.Printf("%v\n",a13)
fmt.Printf("%v\n",a14)
fmt.Printf("%v\n",a15)
fmt.Printf("%v\n",a16)
e0.Recv()
m.Send(nil)
}(m)
}
