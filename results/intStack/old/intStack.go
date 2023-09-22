package main
import ("fmt"
)
type _state_1 struct {
    c chan interface{}
  }
  func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c } } 
  func (x *_state_1) Send(v interface{}) { x.c <- v }
  func (x *_state_1) Recv() interface{} { return <-x.c }

  type _state_3 struct {
    c chan interface{}
    next *_state_1
  }
  
  func init_state_3(c chan interface{}) *_state_3 { return &_state_3{ c, nil } } 
  func (x *_state_3) Send(v interface{}) *_state_1 { if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next}
  func (x *_state_3) Recv() (interface{}, *_state_1) { if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(interface{}),x.next}

  type _state_2 struct {
    c chan interface{}
    next *_state_3
  }
  
  func init_state_2(c chan interface{}) *_state_2 { return &_state_2{ c, nil } } 
  func (x *_state_2) Send(v int) *_state_3 { if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next}
  func (x *_state_2) Recv() (int, *_state_3) { if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int),x.next}

  type _state_0 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_0(c chan interface{}) *_state_0 { m := make(map[string]interface{})
	m["cons"] = init_state_2( c )
	m["null"] = init_state_1( c )
	return &_state_0{ c, m } }
  func (x *_state_0) Send(v string) { x.c <- v }
  func (x *_state_0) Recv() string  { return (<-x.c).(string) }

  type _state_5 struct {
    c chan interface{}
    next *_state_4
  }
  
  func init_state_5(c chan interface{}) *_state_5 { return &_state_5{ c, nil } } 
  func (x *_state_5) Send(v int) *_state_4 { if x.next == nil { x.next = init_state_4(x.c) }; x.c <- v; return x.next}
  func (x *_state_5) Recv() (int, *_state_4) { if x.next == nil { x.next = init_state_4(x.c) }; return (<-x.c).(int),x.next}

  type _state_7 struct {
    c chan interface{}
    next *_state_4
  }
  
  func init_state_7(c chan interface{}) *_state_7 { return &_state_7{ c, nil } } 
  func (x *_state_7) Send(v interface{}) *_state_4 { if x.next == nil { x.next = init_state_4(x.c) }; x.c <- v; return x.next}
  func (x *_state_7) Recv() (interface{}, *_state_4) { if x.next == nil { x.next = init_state_4(x.c) }; return (<-x.c).(interface{}),x.next}

  type _state_8 struct {
    c chan interface{}
    next *_state_4
  }
  
  func init_state_8(c chan interface{}) *_state_8 { return &_state_8{ c, nil } } 
  func (x *_state_8) Send(v int) *_state_4 { if x.next == nil { x.next = init_state_4(x.c) }; x.c <- v; return x.next}
  func (x *_state_8) Recv() (int, *_state_4) { if x.next == nil { x.next = init_state_4(x.c) }; return (<-x.c).(int),x.next}

  type _state_6 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_6(c chan interface{}) *_state_6 { m := make(map[string]interface{})
	m["some"] = init_state_8( c )
	m["none"] = init_state_7( c )
	return &_state_6{ c, m } }
  func (x *_state_6) Send(v string) { x.c <- v }
  func (x *_state_6) Recv() string  { return (<-x.c).(string) }

  type _state_4 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_4(c chan interface{}) *_state_4 { m := make(map[string]interface{})
	m["dealloc"] = init_state_1( c )
	m["pop"] = init_state_6( c )
	m["push"] = init_state_5( c )
	return &_state_4{ c, m } }
  func (x *_state_4) Send(v string) { x.c <- v }
  func (x *_state_4) Recv() string  { return (<-x.c).(string) }

  func null() func (_x *_state_0) {
 return func (c *_state_0){
c.Send("null")
c0 := c.ls["null"].(*_state_1)
c0.Send(nil)
}}
func cons(v int) func (_x *_state_0, l *_state_0) {
 return func (c *_state_0, l *_state_0){
c.Send("cons")
c0 := c.ls["cons"].(*_state_2)
c1 := c0.Send(v)
fmt.Printf("%v\n",v)
l_ := init_state_0(make(chan interface{}))
go func (l_ *_state_0, l *_state_0){
// FWD l_ l Start
for {
l_l := l.Recv()
l_.Send(l_l)
switch l_l {
case "null":
l0 := l.ls["null"].(*_state_1)
l_0 := l_.ls["null"].(*_state_1)
l0.Recv()
l_0.Send(nil)
return
case "cons":
l0 := l.ls["cons"].(*_state_2)
l_0 := l_.ls["cons"].(*_state_2)
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
func dealloc() func (_x *_state_1, l *_state_0) {
 return func (c *_state_1, l *_state_0){
label := l.Recv()
switch label {
case "cons" :
l0 := l.ls["cons"].(*_state_2)
_, l1 := l0.Recv()
l_, l2 := l1.Recv()
l_0 := l_.(*_state_0)
l2.Recv()
d := init_state_1(make(chan interface{}))
go dealloc()(d, l_0)
// FWD c d Start
d.Recv()
c.Send(nil)
return
// FWD c d End
case "null" :
l0 := l.ls["null"].(*_state_1)
l0.Recv()
c.Send(nil)
}
}}
func stack() func (_x *_state_4, l *_state_0) {
 return func (c *_state_4, l *_state_0){
label := c.Recv()
switch label {
case "push" :
c0 := c.ls["push"].(*_state_5)
v, c1 := c0.Recv()
l_ := init_state_0(make(chan interface{}))
go cons(v)(l_, l)
d := init_state_4(make(chan interface{}))
go stack()(d, l_)
// FWD c d Start
for {
dc1 := c1.Recv()
d.Send(dc1)
switch dc1 {
case "push":
d0 := d.ls["push"].(*_state_5)
c2 := c1.ls["push"].(*_state_5)
d0c2, c2_d0 := c2.Recv()
c1 = c2_d0
d = d0.Send(d0c2)
case "pop":
d0 := d.ls["pop"].(*_state_6)
c2 := c1.ls["pop"].(*_state_6)
c2d0 := d0.Recv()
c2.Send(c2d0)
switch c2d0 {
case "none":
d1 := d0.ls["none"].(*_state_7)
c3 := c2.ls["none"].(*_state_7)
c3d1, c3_d1 := d1.Recv()
d = c3_d1
c1 = c3.Send(c3d1)
case "some":
d1 := d0.ls["some"].(*_state_8)
c3 := c2.ls["some"].(*_state_8)
c3d1, c3_d1 := d1.Recv()
d = c3_d1
c1 = c3.Send(c3d1)
}
case "dealloc":
d0 := d.ls["dealloc"].(*_state_1)
c2 := c1.ls["dealloc"].(*_state_1)
d0.Recv()
c2.Send(nil)
return
}
}
// FWD c d End
case "pop" :
c0 := c.ls["pop"].(*_state_6)
label := l.Recv()
switch label {
case "cons" :
l0 := l.ls["cons"].(*_state_2)
v, l1 := l0.Recv()
l_, l2 := l1.Recv()
l_0 := l_.(*_state_0)
l2.Recv()
c0.Send("some")
c1 := c0.ls["some"].(*_state_8)
c2 := c1.Send(v)
d := init_state_4(make(chan interface{}))
go stack()(d, l_0)
// FWD c d Start
for {
dc2 := c2.Recv()
d.Send(dc2)
switch dc2 {
case "push":
d0 := d.ls["push"].(*_state_5)
c3 := c2.ls["push"].(*_state_5)
d0c3, c3_d0 := c3.Recv()
c2 = c3_d0
d = d0.Send(d0c3)
case "pop":
d0 := d.ls["pop"].(*_state_6)
c3 := c2.ls["pop"].(*_state_6)
c3d0 := d0.Recv()
c3.Send(c3d0)
switch c3d0 {
case "none":
d1 := d0.ls["none"].(*_state_7)
c4 := c3.ls["none"].(*_state_7)
c4d1, c4_d1 := d1.Recv()
d = c4_d1
c2 = c4.Send(c4d1)
case "some":
d1 := d0.ls["some"].(*_state_8)
c4 := c3.ls["some"].(*_state_8)
c4d1, c4_d1 := d1.Recv()
d = c4_d1
c2 = c4.Send(c4d1)
}
case "dealloc":
d0 := d.ls["dealloc"].(*_state_1)
c3 := c2.ls["dealloc"].(*_state_1)
d0.Recv()
c3.Send(nil)
return
}
}
// FWD c d End
case "null" :
l0 := l.ls["null"].(*_state_1)
l0.Recv()
c0.Send("none")
c1 := c0.ls["none"].(*_state_7)
c2 := c1.Send(struct{}{})
l_ := init_state_0(make(chan interface{}))
go null()(l_)
d := init_state_4(make(chan interface{}))
go stack()(d, l_)
// FWD c d Start
for {
dc2 := c2.Recv()
d.Send(dc2)
switch dc2 {
case "push":
d0 := d.ls["push"].(*_state_5)
c3 := c2.ls["push"].(*_state_5)
d0c3, c3_d0 := c3.Recv()
c2 = c3_d0
d = d0.Send(d0c3)
case "pop":
d0 := d.ls["pop"].(*_state_6)
c3 := c2.ls["pop"].(*_state_6)
c3d0 := d0.Recv()
c3.Send(c3d0)
switch c3d0 {
case "none":
d1 := d0.ls["none"].(*_state_7)
c4 := c3.ls["none"].(*_state_7)
c4d1, c4_d1 := d1.Recv()
d = c4_d1
c2 = c4.Send(c4d1)
case "some":
d1 := d0.ls["some"].(*_state_8)
c4 := c3.ls["some"].(*_state_8)
c4d1, c4_d1 := d1.Recv()
d = c4_d1
c2 = c4.Send(c4d1)
}
case "dealloc":
d0 := d.ls["dealloc"].(*_state_1)
c3 := c2.ls["dealloc"].(*_state_1)
d0.Recv()
c3.Send(nil)
return
}
}
// FWD c d End
}
case "dealloc" :
c0 := c.ls["dealloc"].(*_state_1)
d := init_state_1(make(chan interface{}))
go dealloc()(d, l)
// FWD c d Start
d.Recv()
c0.Send(nil)
return
// FWD c d End
}
}}
func main () {
c := init_state_1(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_1){
l := init_state_0(make(chan interface{}))
go null()(l)
s := init_state_4(make(chan interface{}))
go stack()(s, l)
s.Send("push")
s0 := s.ls["push"].(*_state_5)
s1 := s0.Send(1)
s1.Send("push")
s2 := s1.ls["push"].(*_state_5)
s3 := s2.Send(2)
s3.Send("pop")
s4 := s3.ls["pop"].(*_state_6)
label := s4.Recv()
switch label {
case "some" :
s5 := s4.ls["some"].(*_state_8)
_, s6 := s5.Recv()
s6.Send("dealloc")
s7 := s6.ls["dealloc"].(*_state_1)
s7.Recv()
c.Send(nil)
case "none" :
s5 := s4.ls["none"].(*_state_7)
_, s6 := s5.Recv()
s6.Send("dealloc")
s7 := s6.ls["dealloc"].(*_state_1)
s7.Recv()
c.Send(nil)
}
}(c)
}
