package main
import "fmt"
//Preamble generation
type _state_2 struct {
    c chan interface{}
}
func init_state_2(c chan interface{}) *_state_2 { return &_state_2{ c } }
func (x *_state_2) Send(v interface{}) { x.c <- v }
func (x *_state_2) Recv() interface{} { return <-x.c }

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

  //Declaration list compilation
func foo()func (_x *_state_0) {
 return func (c *_state_0){
func (d *_state_0){
d0 := d.Send(0)
x := init_state_1(make(chan interface{}))
go plus_one(1)(x)
// FWD d x Start
d0x, x0 := x.Recv()
d1 := d0.Send(d0x)
x0.Recv()
d1.Send(nil)
return
// FWD d x End
}(c)
}}
func plus_one(n int) func (_x *_state_1) {
 return func (c *_state_1){
c0 := c.Send((n + 1))
c0.Send(nil)
}}
//Main compilation
func main () {
    c:= init_state_2(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_2){
d := init_state_0(make(chan interface{}))
go foo()(d)
x, d0 := d.Recv()
fmt.Printf("%v\n",x)
x1, d1 := d0.Recv()
fmt.Printf("%v\n",x1)
d1.Recv()
c.Send(nil)
}(c)
}
