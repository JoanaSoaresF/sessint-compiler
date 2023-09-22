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

  type _state_3 struct {
   c chan interface{}
   next *_state_0
}

func init_state_3(c chan interface{}) *_state_3 { return &_state_3{c, nil} }
func (x *_state_3) Send(v interface{}) *_state_0 {
   if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next }
func (x *_state_3) Recv() (interface{}, *_state_0) {
   if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(interface{}), x.next }

  type _state_4 struct {
   c chan interface{}
   next *_state_0
}

func init_state_4(c chan interface{}) *_state_4 { return &_state_4{c, nil} }
func (x *_state_4) Send(v int) *_state_0 {
   if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next }
func (x *_state_4) Recv() (int, *_state_0) {
   if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int), x.next }

  type _state_2 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_2(c chan interface{}) *_state_2 { m := make(map[string]interface{})
 	m["some"] = init_state_4( c )
	m["none"] = init_state_3( c )
   return &_state_2{ c, m } }
func (x *_state_2) Send(v string) { x.c <- v }
func (x *_state_2) Recv() string  { return (<-x.c).(string) }

  type _state_5 struct {
    c chan interface{}
}
func init_state_5(c chan interface{}) *_state_5 { return &_state_5{ c } }
func (x *_state_5) Send(v interface{}) { x.c <- v }
func (x *_state_5) Recv() interface{} { return <-x.c }

  type _state_0 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_0(c chan interface{}) *_state_0 { m := make(map[string]interface{})
 	m["dealloc"] = init_state_5( c )
	m["pop"] = init_state_2( c )
	m["push"] = init_state_1( c )
   return &_state_0{ c, m } }
func (x *_state_0) Send(v string) { x.c <- v }
func (x *_state_0) Recv() string  { return (<-x.c).(string) }

  type _state_8 struct {
   c chan interface{}
   next *_state_5
}

func init_state_8(c chan interface{}) *_state_8 { return &_state_8{c, nil} }
func (x *_state_8) Send(v interface{}) *_state_5 {
   if x.next == nil { x.next = init_state_5(x.c) }; x.c <- v; return x.next }
func (x *_state_8) Recv() (interface{}, *_state_5) {
   if x.next == nil { x.next = init_state_5(x.c) }; return (<-x.c).(interface{}), x.next }

  type _state_7 struct {
   c chan interface{}
   next *_state_8
}

func init_state_7(c chan interface{}) *_state_7 { return &_state_7{c, nil} }
func (x *_state_7) Send(v int) *_state_8 {
   if x.next == nil { x.next = init_state_8(x.c) }; x.c <- v; return x.next }
func (x *_state_7) Recv() (int, *_state_8) {
   if x.next == nil { x.next = init_state_8(x.c) }; return (<-x.c).(int), x.next }

  type _state_6 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_6(c chan interface{}) *_state_6 { m := make(map[string]interface{})
 	m["cons"] = init_state_7( c )
	m["null"] = init_state_5( c )
   return &_state_6{ c, m } }
func (x *_state_6) Send(v string) { x.c <- v }
func (x *_state_6) Recv() string  { return (<-x.c).(string) }

  //Declaration list compilation
func stack()func (_x *_state_0, l *_state_6) {
 return func (c *_state_0, l *_state_6){
for {
 label := c.Recv()
switch label {
case "push" :
c0 := c.ls["push"].(*_state_1)
v, c1 := c0.Recv()
l_ := init_state_6(make(chan interface{}))
go cons(v)(l_, l)
//Update arguments
//Update channels
c = c1
l = l_
 case "pop" :
c0 := c.ls["pop"].(*_state_2)
label := l.Recv()
switch label {
case "cons" :
l0 := l.ls["cons"].(*_state_7)
v, l1 := l0.Recv()
l_, l2 := l1.Recv()
l_0 := l_.(*_state_6)
l2.Recv()
c0.Send("some")
c1 := c0.ls["some"].(*_state_4)
c2 := c1.Send(v)
//Update arguments
//Update channels
c = c2
l = l_0
 case "null" :
l0 := l.ls["null"].(*_state_5)
l0.Recv()
c0.Send("none")
c1 := c0.ls["none"].(*_state_3)
c2 := c1.Send(struct{}{})
l_ := init_state_6(make(chan interface{}))
go null()(l_)
//Update arguments
//Update channels
c = c2
l = l_
 }
case "dealloc" :
c0 := c.ls["dealloc"].(*_state_5)
dealloc()(c0, l)
}
}
}
}
func dealloc()func (_x *_state_5, l *_state_6) {
 return func (c *_state_5, l *_state_6){
for {
 label := l.Recv()
switch label {
case "cons" :
l0 := l.ls["cons"].(*_state_7)
_, l1 := l0.Recv()
l_, l2 := l1.Recv()
l_0 := l_.(*_state_6)
l2.Recv()
//Update arguments
//Update channels
c = c
l = l_0
 case "null" :
l0 := l.ls["null"].(*_state_5)
l0.Recv()
c.Send(nil)
break
}
}
}
}
func cons(v int) func (_x *_state_6, l *_state_6) {
 return func (c *_state_6, l *_state_6){
c.Send("cons")
c0 := c.ls["cons"].(*_state_7)
c1 := c0.Send(v)
fmt.Printf("%v\n",v)
l_ := init_state_6(make(chan interface{}))
go func (l_ *_state_6, l *_state_6){
// FWD l_ l Start
for {
l_l := l.Recv()
l_.Send(l_l)
switch l_l {
case "null":
l0 := l.ls["null"].(*_state_5)
l_0 := l_.ls["null"].(*_state_5)
l0.Recv()
l_0.Send(nil)
return
case "cons":
l0 := l.ls["cons"].(*_state_7)
l_0 := l_.ls["cons"].(*_state_7)
l_0l0, l1 := l0.Recv()
l_1 := l_0.Send(l_0l0)
l_1l1, l2 := l1.Recv()
l_2 := l_1.Send(l_1l1)
l2.Recv()
l_2.Send(nil)
return
}
}
// FWD l_ l End
}(l_, l)
c2 := c1.Send(l_)
c2.Send(nil)
}}
func null()func (_x *_state_6) {
 return func (c *_state_6){
c.Send("null")
c0 := c.ls["null"].(*_state_5)
c0.Send(nil)
}}
//Main compilation
func main () {
    c:= init_state_5(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_5){
l := init_state_6(make(chan interface{}))
go null()(l)
s := init_state_0(make(chan interface{}))
go stack()(s, l)
s.Send("push")
s0 := s.ls["push"].(*_state_1)
s1 := s0.Send(1)
s1.Send("push")
s2 := s1.ls["push"].(*_state_1)
s3 := s2.Send(2)
s3.Send("pop")
s4 := s3.ls["pop"].(*_state_2)
label := s4.Recv()
switch label {
case "some":
s5 := s4.ls["some"].(*_state_4)
_, s6 := s5.Recv()
s6.Send("dealloc")
s7 := s6.ls["dealloc"].(*_state_5)
s7.Recv()
c.Send(nil)
case "none":
s5 := s4.ls["none"].(*_state_3)
_, s6 := s5.Recv()
s6.Send("dealloc")
s7 := s6.ls["dealloc"].(*_state_5)
s7.Recv()
c.Send(nil)
}
}(c)
}
