package main
import "fmt"
//Preamble generation
type _state_1 struct {
   c chan interface{}
   next *_state_0
}

func init_state_1(c chan interface{}) *_state_1 { return &_state_1{c, nil} }
func (x *_state_1) Send(v int) *_state_0 {
   if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next }
func (x *_state_1) Recv() (int, *_state_0) {
   if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int), x.next }

  type _state_2 struct {
    c chan interface{}
}
func init_state_2(c chan interface{}) *_state_2 { return &_state_2{ c } }
func (x *_state_2) Send(v interface{}) { x.c <- v }
func (x *_state_2) Recv() interface{} { return <-x.c }

  type _state_0 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_0(c chan interface{}) *_state_0 { m := make(map[string]interface{})
 	m["stop"] = init_state_2( c )
	m["next"] = init_state_1( c )
   return &_state_0{ c, m } }
func (x *_state_0) Send(v string) { x.c <- v }
func (x *_state_0) Recv() string  { return (<-x.c).(string) }

  //Declaration list compilation
func fib(a int) func (_x int) func (_x *_state_0) {
 return func (b int) func (_x *_state_0){
return func (c *_state_0){
label := c.Recv()
switch label {
case "next" :
c0 := c.ls["next"].(*_state_1)
c1 := c0.Send(b)
fib(b)((a + b))(c1)
case "stop" :
c0 := c.ls["stop"].(*_state_2)
c0.Send(nil)
}
}}
}
//Main compilation
func main () {
    c:= init_state_2(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_2){
q := init_state_0(make(chan interface{}))
go fib(0)(1)(q)
q.Send("next")
q0 := q.ls["next"].(*_state_1)
x1, q1 := q0.Recv()
fmt.Printf("%v\n",x1)
q1.Send("next")
q2 := q1.ls["next"].(*_state_1)
x2, q3 := q2.Recv()
fmt.Printf("%v\n",x2)
q3.Send("stop")
q4 := q3.ls["stop"].(*_state_2)
q4.Recv()
c.Send(nil)
}(c)
}
