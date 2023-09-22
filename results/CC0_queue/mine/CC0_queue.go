package main
import "fmt"
//Preamble generation
type _state_0 struct {
    c chan interface{}
}
func init_state_0(c chan interface{}) *_state_0 { return &_state_0{ c } }
func (x *_state_0) Send(v interface{}) { x.c <- v }
func (x *_state_0) Recv() interface{} { return <-x.c }

  type _state_2 struct {
   c chan interface{}
   next *_state_1
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
func (x *_state_2) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_2) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_4 struct {
   c chan interface{}
   next *_state_1
}

func init_state_4(c chan interface{}) *_state_4 { return &_state_4{c, nil} }
func (x *_state_4) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_4) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_3 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_3(c chan interface{}) *_state_3 { m := make(map[string]interface{})
 	m["none"] = init_state_0( c )
	m["some"] = init_state_4( c )
   return &_state_3{ c, m } }
func (x *_state_3) Send(v string) { x.c <- v }
func (x *_state_3) Recv() string  { return (<-x.c).(string) }

  type _state_1 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_1(c chan interface{}) *_state_1 { m := make(map[string]interface{})
 	m["deq"] = init_state_3( c )
	m["enq"] = init_state_2( c )
   return &_state_1{ c, m } }
func (x *_state_1) Send(v string) { x.c <- v }
func (x *_state_1) Recv() string  { return (<-x.c).(string) }

  //Declaration list compilation
func dealloc()func (_x *_state_0, q *_state_1) {
 return func (c *_state_0, q *_state_1){
for {
 q.Send("deq")
q0 := q.ls["deq"].(*_state_3)
label := q0.Recv()
switch label {
case "none":
q1 := q0.ls["none"].(*_state_0)
q1.Recv()
c.Send(nil)
break
case "some":
q1 := q0.ls["some"].(*_state_4)
_, q2 := q1.Recv()
//Update arguments
//Update channels
c = c
q = q2
 }
}
}
}
func empty()func (_x *_state_1) {
 return func (q *_state_1){
label := q.Recv()
switch label {
case "enq" :
q0 := q.ls["enq"].(*_state_2)
y, q1 := q0.Recv()
fmt.Printf("%v\n",y)
e := init_state_1(make(chan interface{}))
go empty()(e)
elem(y)(q1, e)
case "deq" :
q0 := q.ls["deq"].(*_state_3)
q0.Send("none")
q1 := q0.ls["none"].(*_state_0)
q1.Send(nil)
}
}}
func elem(x int) func (_x *_state_1, r *_state_1) {
 return func (q *_state_1, r *_state_1){
for {
 label := q.Recv()
switch label {
case "enq" :
q0 := q.ls["enq"].(*_state_2)
y, q1 := q0.Recv()
r.Send("enq")
r0 := r.ls["enq"].(*_state_2)
r1 := r0.Send(y)
//Update arguments
x = x
//Update channels
q = q1
r = r1
 case "deq" :
q0 := q.ls["deq"].(*_state_3)
q0.Send("some")
q1 := q0.ls["some"].(*_state_4)
q2 := q1.Send(x)
// FWD q r Start
for {
rq2 := q2.Recv()
r.Send(rq2)
switch rq2 {
case "enq":
r0 := r.ls["enq"].(*_state_2)
q3 := q2.ls["enq"].(*_state_2)
r0q3, q3_r0 := q3.Recv()
q2 = q3_r0
r = r0.Send(r0q3)
case "deq":
r0 := r.ls["deq"].(*_state_3)
q3 := q2.ls["deq"].(*_state_3)
q3r0 := r0.Recv()
q3.Send(q3r0)
switch q3r0 {
case "some":
r1 := r0.ls["some"].(*_state_4)
q4 := q3.ls["some"].(*_state_4)
q4r1, q4_r1 := r1.Recv()
r = q4_r1
q2 = q4.Send(q4r1)
case "none":
r1 := r0.ls["none"].(*_state_0)
q4 := q3.ls["none"].(*_state_0)
r1.Recv()
q4.Send(nil)
return
}
}
}
// FWD q r End
}
}
}
}
//Main compilation
func main () {
    c:= init_state_0(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_0){
q := init_state_1(make(chan interface{}))
go empty()(q)
q.Send("enq")
q0 := q.ls["enq"].(*_state_2)
q1 := q0.Send(1)
q1.Send("enq")
q2 := q1.ls["enq"].(*_state_2)
q3 := q2.Send(2)
q3.Send("enq")
q4 := q3.ls["enq"].(*_state_2)
q5 := q4.Send(3)
d := init_state_0(make(chan interface{}))
go dealloc()(d, q5)
d.Recv()
c.Send(nil)
}(c)
}
