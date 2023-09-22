package main
import "fmt"
//Preamble generation
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

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
type _multisend_type__state_2 struct {
v0 int
v1 int
v2 int
v3 int
}
func (x *_state_2) Send(v0 int, v1 int, v2 int, v3 int) *_state_3 {
   if x.next == nil { x.next = init_state_3(x.c) };
 x.c <- _multisend_type__state_2{v0, v1, v2, v3}
return x.next }
func (x *_state_2) Recv() (int, int, int, int, *_state_3) {
   if x.next == nil { x.next = init_state_3(x.c) };ll := <- x.c
l := ll.(_multisend_type__state_2)
return l.v0, l.v1, l.v2, l.v3, x.next }

type _state_1 struct {
   c chan interface{}
   next *_state_2
}

func init_state_1(c chan interface{}) *_state_1 { return &_state_1{c, nil} }
func (x *_state_1) Send(v int) *_state_2 {
   if x.next == nil { x.next = init_state_2(x.c) }; x.c <- v; return x.next }
func (x *_state_1) Recv() (int, *_state_2) {
   if x.next == nil { x.next = init_state_2(x.c) }; return (<-x.c).(int), x.next }

  type _state_0 struct {
   c chan interface{}
   next *_state_1
}

func init_state_0(c chan interface{}) *_state_0 { return &_state_0{c, nil} }
type _multisend_type__state_0 struct {
v0 int
v1 int
v2 int
}
func (x *_state_0) Send(v0 int, v1 int, v2 int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) };
 x.c <- _multisend_type__state_0{v0, v1, v2}
return x.next }
func (x *_state_0) Recv() (int, int, int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) };ll := <- x.c
l := ll.(_multisend_type__state_0)
return l.v0, l.v1, l.v2, x.next }

type _state_11 struct {
   c chan interface{}
   next *_state_3
}

func init_state_11(c chan interface{}) *_state_11 { return &_state_11{c, nil} }
func (x *_state_11) Send(v int) *_state_3 {
   if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next }
func (x *_state_11) Recv() (int, *_state_3) {
   if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int), x.next }

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

  //Declaration list compilation
func send_three_optimized(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send(n, (n + 1), (n + 2))
a1, c1 := c0.Recv()
c2 := c1.Send(a1, (a1 - 1), (a1 - 2), (a1 - 3))
c2.Send(nil)
}}
func send_three(n int) func (_x *_state_4) {
 return func (c *_state_4){
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
//Main compilation
func main () {
    m:= init_state_3(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_3){
e := init_state_0(make(chan interface{}))
go send_three_optimized(1)(e)
a1, a2, a3, e0 := e.Recv()
fmt.Printf("%v\n",a1)
fmt.Printf("%v\n",a2)
fmt.Printf("%v\n",a3)
e1 := e0.Send(500)
a4, a5, a6, a7, e2 := e1.Recv()
fmt.Printf("%v\n",a4)
fmt.Printf("%v\n",a5)
fmt.Printf("%v\n",a6)
fmt.Printf("%v\n",a7)
e2.Recv()
m.Send(nil)
}(m)
}
