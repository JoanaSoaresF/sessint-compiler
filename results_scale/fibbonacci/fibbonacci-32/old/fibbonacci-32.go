package main
import ("fmt"
)
type _state_1 struct {
    c chan interface{}
    next *_state_0
  }
  
  func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c, nil } } 
  func (x *_state_1) Send(v int) *_state_0 { if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next}
  func (x *_state_1) Recv() (int, *_state_0) { if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int),x.next}

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

  func fib(a int) func (_x int) func (_x *_state_0) {
 return func (b int) func (_x *_state_0){
return func (c *_state_0){
label := c.Recv()
switch label {
case "next" :
c0 := c.ls["next"].(*_state_1)
c1 := c0.Send(b)
d := init_state_0(make(chan interface{}))
go fib(b)((a + b))(d)
// FWD c d Start
for {
dc1 := c1.Recv()
d.Send(dc1)
switch dc1 {
case "next":
d0 := d.ls["next"].(*_state_1)
c2 := c1.ls["next"].(*_state_1)
c2d0, c2_d0 := d0.Recv()
d = c2_d0
c1 = c2.Send(c2d0)
case "stop":
d0 := d.ls["stop"].(*_state_2)
c2 := c1.ls["stop"].(*_state_2)
d0.Recv()
c2.Send(nil)
return
}
}
// FWD c d End
case "stop" :
c0 := c.ls["stop"].(*_state_2)
c0.Send(nil)
}
}}
}
func main () {
c := init_state_2(make (chan interface{}))
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
q11.Send("next")
q12 := q11.ls["next"].(*_state_1)
x7, q13 := q12.Recv()
fmt.Printf("%v\n",x7)
q13.Send("next")
q14 := q13.ls["next"].(*_state_1)
x8, q15 := q14.Recv()
fmt.Printf("%v\n",x8)
q15.Send("next")
q16 := q15.ls["next"].(*_state_1)
x9, q17 := q16.Recv()
fmt.Printf("%v\n",x9)
q17.Send("next")
q18 := q17.ls["next"].(*_state_1)
x10, q19 := q18.Recv()
fmt.Printf("%v\n",x10)
q19.Send("next")
q20 := q19.ls["next"].(*_state_1)
x11, q21 := q20.Recv()
fmt.Printf("%v\n",x11)
q21.Send("next")
q22 := q21.ls["next"].(*_state_1)
x12, q23 := q22.Recv()
fmt.Printf("%v\n",x12)
q23.Send("next")
q24 := q23.ls["next"].(*_state_1)
x13, q25 := q24.Recv()
fmt.Printf("%v\n",x13)
q25.Send("next")
q26 := q25.ls["next"].(*_state_1)
x14, q27 := q26.Recv()
fmt.Printf("%v\n",x14)
q27.Send("next")
q28 := q27.ls["next"].(*_state_1)
x15, q29 := q28.Recv()
fmt.Printf("%v\n",x15)
q29.Send("next")
q30 := q29.ls["next"].(*_state_1)
x16, q31 := q30.Recv()
fmt.Printf("%v\n",x16)
q31.Send("next")
q32 := q31.ls["next"].(*_state_1)
x17, q33 := q32.Recv()
fmt.Printf("%v\n",x17)
q33.Send("next")
q34 := q33.ls["next"].(*_state_1)
x18, q35 := q34.Recv()
fmt.Printf("%v\n",x18)
q35.Send("next")
q36 := q35.ls["next"].(*_state_1)
x19, q37 := q36.Recv()
fmt.Printf("%v\n",x19)
q37.Send("next")
q38 := q37.ls["next"].(*_state_1)
x20, q39 := q38.Recv()
fmt.Printf("%v\n",x20)
q39.Send("next")
q40 := q39.ls["next"].(*_state_1)
x21, q41 := q40.Recv()
fmt.Printf("%v\n",x21)
q41.Send("next")
q42 := q41.ls["next"].(*_state_1)
x22, q43 := q42.Recv()
fmt.Printf("%v\n",x22)
q43.Send("next")
q44 := q43.ls["next"].(*_state_1)
x23, q45 := q44.Recv()
fmt.Printf("%v\n",x23)
q45.Send("next")
q46 := q45.ls["next"].(*_state_1)
x24, q47 := q46.Recv()
fmt.Printf("%v\n",x24)
q47.Send("next")
q48 := q47.ls["next"].(*_state_1)
x25, q49 := q48.Recv()
fmt.Printf("%v\n",x25)
q49.Send("next")
q50 := q49.ls["next"].(*_state_1)
x26, q51 := q50.Recv()
fmt.Printf("%v\n",x26)
q51.Send("next")
q52 := q51.ls["next"].(*_state_1)
x27, q53 := q52.Recv()
fmt.Printf("%v\n",x27)
q53.Send("next")
q54 := q53.ls["next"].(*_state_1)
x28, q55 := q54.Recv()
fmt.Printf("%v\n",x28)
q55.Send("next")
q56 := q55.ls["next"].(*_state_1)
x29, q57 := q56.Recv()
fmt.Printf("%v\n",x29)
q57.Send("next")
q58 := q57.ls["next"].(*_state_1)
x30, q59 := q58.Recv()
fmt.Printf("%v\n",x30)
q59.Send("next")
q60 := q59.ls["next"].(*_state_1)
x31, q61 := q60.Recv()
fmt.Printf("%v\n",x31)
q61.Send("next")
q62 := q61.ls["next"].(*_state_1)
x32, q63 := q62.Recv()
fmt.Printf("%v\n",x32)
q63.Send("stop")
q64 := q63.ls["stop"].(*_state_2)
q64.Recv()
c.Send(nil)
}(c)
}
