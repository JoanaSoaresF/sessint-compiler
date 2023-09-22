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
 	m["halt"] = init_state_2( c )
	m["val"] = init_state_1( c )
	m["inc"] = &_state_0{ c, m }
   return &_state_0{ c, m } }
func (x *_state_0) Send(v string) { x.c <- v }
func (x *_state_0) Recv() string  { return (<-x.c).(string) }

  //Declaration list compilation
func epsilon()func (_x *_state_0) {
 return func (c *_state_0){
for {
 label := c.Recv()
switch label {
case "inc" :
c0 := c.ls["inc"].(*_state_0)
e := init_state_0(make(chan interface{}))
go epsilon()(e)
bit(1)(c0, e)
case "val" :
c0 := c.ls["val"].(*_state_1)
c1 := c0.Send(0)
//Update arguments
//Update channels
c = c1
case "halt" :
c0 := c.ls["halt"].(*_state_2)
c0.Send(nil)
break
}
}
}
}
func bit(b int) func (_x *_state_0, d *_state_0) {
 return func (c *_state_0, d *_state_0){
for {
 label := c.Recv()
switch label {
case "inc" :
c0 := c.ls["inc"].(*_state_0)
if (b == 0) {
//Update arguments
b = 1
//Update channels
c = c0
d = d
 } else {
d.Send("inc")
d0 := d.ls["inc"].(*_state_0)
//Update arguments
b = 0
//Update channels
c = c0
d = d0
 }
case "val" :
c0 := c.ls["val"].(*_state_1)
d.Send("val")
d0 := d.ls["val"].(*_state_1)
n, d1 := d0.Recv()
c1 := c0.Send(((2 * n) + b))
//Update arguments
b = b
//Update channels
c = c1
d = d1
 case "halt" :
c0 := c.ls["halt"].(*_state_2)
d.Send("halt")
d0 := d.ls["halt"].(*_state_2)
d0.Recv()
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
q := init_state_0(make(chan interface{}))
go epsilon()(q)
q.Send("inc")
q0 := q.ls["inc"].(*_state_0)
q0.Send("val")
q1 := q0.ls["val"].(*_state_1)
x1, q2 := q1.Recv()
fmt.Printf("%v\n",x1)
q2.Send("inc")
q3 := q2.ls["inc"].(*_state_0)
q3.Send("val")
q4 := q3.ls["val"].(*_state_1)
x2, q5 := q4.Recv()
fmt.Printf("%v\n",x2)
q5.Send("inc")
q6 := q5.ls["inc"].(*_state_0)
q6.Send("val")
q7 := q6.ls["val"].(*_state_1)
x3, q8 := q7.Recv()
fmt.Printf("%v\n",x3)
q8.Send("inc")
q9 := q8.ls["inc"].(*_state_0)
q9.Send("val")
q10 := q9.ls["val"].(*_state_1)
x4, q11 := q10.Recv()
fmt.Printf("%v\n",x4)
q11.Send("inc")
q12 := q11.ls["inc"].(*_state_0)
q12.Send("val")
q13 := q12.ls["val"].(*_state_1)
x5, q14 := q13.Recv()
fmt.Printf("%v\n",x5)
q14.Send("inc")
q15 := q14.ls["inc"].(*_state_0)
q15.Send("inc")
q16 := q15.ls["inc"].(*_state_0)
q16.Send("inc")
q17 := q16.ls["inc"].(*_state_0)
q17.Send("inc")
q18 := q17.ls["inc"].(*_state_0)
q18.Send("inc")
q19 := q18.ls["inc"].(*_state_0)
q19.Send("inc")
q20 := q19.ls["inc"].(*_state_0)
q20.Send("inc")
q21 := q20.ls["inc"].(*_state_0)
q21.Send("inc")
q22 := q21.ls["inc"].(*_state_0)
q22.Send("val")
q23 := q22.ls["val"].(*_state_1)
x6, q24 := q23.Recv()
fmt.Printf("%v\n",x6)
q24.Send("halt")
q25 := q24.ls["halt"].(*_state_2)
q25.Recv()
c.Send(nil)
}(c)
}
