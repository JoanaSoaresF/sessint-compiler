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
func filter()func (_x *_state_0, n_channel *_state_0) {
 return func (r_channel *_state_0, n_channel *_state_0){
for {
 label := r_channel.Recv()
switch label {
case "next" :
r_channel0 := r_channel.ls["next"].(*_state_1)
n_channel.Send("next")
n_channel0 := n_channel.ls["next"].(*_state_1)
v, n_channel1 := n_channel0.Recv()
if is_odd(v) {
r_channel1 := r_channel0.Send(v)
//Update arguments
//Update channels
r_channel = r_channel1
n_channel = n_channel1
 } else {
r_channel1 := r_channel0.Send((-1))
//Update arguments
//Update channels
r_channel = r_channel1
n_channel = n_channel1
 }
case "stop" :
r_channel0 := r_channel.ls["stop"].(*_state_2)
n_channel.Send("stop")
n_channel0 := n_channel.ls["stop"].(*_state_2)
n_channel0.Recv()
r_channel0.Send(nil)
break
}
}
}
}
func is_odd(i int) bool {
 return (!((i - (2 * (i / 2))) == 0))}
func nats(n int) func (_x *_state_0) {
 return func (c *_state_0){
for {
 label := c.Recv()
switch label {
case "next" :
c0 := c.ls["next"].(*_state_1)
c1 := c0.Send(n)
//Update arguments
n = (n + 1)
//Update channels
c = c1
case "stop" :
c0 := c.ls["stop"].(*_state_2)
c0.Send(nil)
break
}
}
}
}
//Main compilation
func main () {
    c:= init_state_2(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_2){
d := init_state_0(make(chan interface{}))
go nats(0)(d)
q := init_state_0(make(chan interface{}))
go filter()(q, d)
q.Send("next")
q0 := q.ls["next"].(*_state_1)
x1, q1 := q0.Recv()
fmt.Printf("%v\n",x1)
q1.Send("next")
q2 := q1.ls["next"].(*_state_1)
x2, q3 := q2.Recv()
fmt.Printf("%v\n",x2)
q3.Send("next")
q4 := q3.ls["next"].(*_state_1)
x3, q5 := q4.Recv()
fmt.Printf("%v\n",x3)
q5.Send("next")
q6 := q5.ls["next"].(*_state_1)
x4, q7 := q6.Recv()
fmt.Printf("%v\n",x4)
q7.Send("next")
q8 := q7.ls["next"].(*_state_1)
x5, q9 := q8.Recv()
fmt.Printf("%v\n",x5)
q9.Send("next")
q10 := q9.ls["next"].(*_state_1)
x6, q11 := q10.Recv()
fmt.Printf("%v\n",x6)
q11.Send("stop")
q12 := q11.ls["stop"].(*_state_2)
q12.Recv()
c.Send(nil)
}(c)
}
