package main
import "fmt"
//Preamble generation
type _state_0 struct {
   c chan interface{}
   next *_state_0
}

func init_state_0(c chan interface{}) *_state_0 { return &_state_0{c, nil} }
func (x *_state_0) Send(v int) *_state_0 {
   if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x }
func (x *_state_0) Recv() (int, *_state_0) {
   if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int), x }

  type _state_1 struct {
    c chan interface{}
}
func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c } }
func (x *_state_1) Send(v interface{}) { x.c <- v }
func (x *_state_1) Recv() interface{} { return <-x.c }

  //Declaration list compilation
func nats(n int) func (_x *_state_0) {
 return func (c *_state_0){
for {
 c0 := c.Send(n)
fmt.Printf("%v\n",n)
//Update arguments
n = (n + 1)
//Update channels
c = c0
}
}
}
//Main compilation
func main () {
    c:= init_state_1(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_1){
c.Send(nil)
}(c)
}
