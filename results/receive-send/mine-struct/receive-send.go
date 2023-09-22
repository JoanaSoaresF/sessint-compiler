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
func (x *_state_2) Send(v int) *_state_3 {
   if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next }
func (x *_state_2) Recv() (int, *_state_3) {
   if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int), x.next }

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
func (x *_state_0) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_0) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_4 struct {
   c chan interface{}
   next *_state_3
}

func init_state_4(c chan interface{}) *_state_4 { return &_state_4{c, nil} }
type _multisend_type__state_4 struct {
v0 int
v1 int
v2 int
}
func (x *_state_4) Send(v0 int, v1 int, v2 int) *_state_3 {
   if x.next == nil { x.next = init_state_3(x.c) };
 x.c <- _multisend_type__state_4{v0, v1, v2}
return x.next }
func (x *_state_4) Recv() (int, int, int, *_state_3) {
   if x.next == nil { x.next = init_state_3(x.c) };ll := <- x.c
l := ll.(_multisend_type__state_4)
return l.v0, l.v1, l.v2, x.next }

//Declaration list compilation
func recv_three(n int) func (_x *_state_0) {
 return func (c *_state_0){
a1, c0 := c.Recv()
a2, c1 := c0.Recv()
a3, c2 := c1.Recv()
fmt.Printf("%v\n",(a1 + n))
fmt.Printf("%v\n",a2)
fmt.Printf("%v\n",a3)
c2.Send(nil)
}}
func recv_three_optimized(n int) func (_x *_state_4) {
 return func (c *_state_4){
a1, a2, a3, c0 := c.Recv()
fmt.Printf("%v\n",(a1 + n))
fmt.Printf("%v\n",a2)
fmt.Printf("%v\n",a3)
c0.Send(nil)
}}
//Main compilation
func main () {
    m:= init_state_3(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_3){
e := init_state_4(make(chan interface{}))
go recv_three_optimized(1)(e)
e0 := e.Send(1, 2, 3)
e0.Recv()
m.Send(nil)
}(m)
}
