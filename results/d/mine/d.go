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
fmt.Printf("%v\n",n)
fmt.Printf("%v\n",n)
fmt.Printf("%v\n",n)
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
d := init_state_0(make(chan interface{}))
go plus_two(1)(d)
a, d0 := d.Recv()
fmt.Printf("%v\n",a)
d0.Recv()
e := init_state_0(make(chan interface{}))
go plus_two(a)(e)
a1, e0 := e.Recv()
fmt.Printf("%v\n",a1)
e0.Recv()
m.Send(nil)
}(m)
}
