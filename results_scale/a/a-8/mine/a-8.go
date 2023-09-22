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
func (x *_state_0) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_0) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  //Declaration list compilation
func plus_two(n int) func (_x *_state_0) {
 return func (c *_state_0){
plus_one((n + 1))(c)
}}
func plus_one(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send((n + 1))
c0.Send(nil)
}}
//Main compilation
func main () {
    m:= init_state_1(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_1){
d1d := init_state_0(make(chan interface{}))
go plus_two(1)(d1d)
a1, d1d0 := d1d.Recv()
fmt.Printf("%v\n",a1)
d1d0.Recv()
d2d := init_state_0(make(chan interface{}))
go plus_two(2)(d2d)
a2, d2d0 := d2d.Recv()
fmt.Printf("%v\n",a2)
d2d0.Recv()
d3d := init_state_0(make(chan interface{}))
go plus_two(3)(d3d)
a3, d3d0 := d3d.Recv()
fmt.Printf("%v\n",a3)
d3d0.Recv()
d4d := init_state_0(make(chan interface{}))
go plus_two(4)(d4d)
a4, d4d0 := d4d.Recv()
fmt.Printf("%v\n",a4)
d4d0.Recv()
d5d := init_state_0(make(chan interface{}))
go plus_two(5)(d5d)
a5, d5d0 := d5d.Recv()
fmt.Printf("%v\n",a5)
d5d0.Recv()
d6d := init_state_0(make(chan interface{}))
go plus_two(6)(d6d)
a6, d6d0 := d6d.Recv()
fmt.Printf("%v\n",a6)
d6d0.Recv()
d7d := init_state_0(make(chan interface{}))
go plus_two(7)(d7d)
a7, d7d0 := d7d.Recv()
fmt.Printf("%v\n",a7)
d7d0.Recv()
d8d := init_state_0(make(chan interface{}))
go plus_two(8)(d8d)
a8, d8d0 := d8d.Recv()
fmt.Printf("%v\n",a8)
d8d0.Recv()
m.Send(nil)
}(m)
}
