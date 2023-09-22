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
func doMap() func (_x *_state_0, l *_state_0, r *_state_0) {
 return func (c *_state_0, l *_state_0, r *_state_0){
label := l.Recv()
switch label {
case "cons" :
l0 := l.ls["cons"].(*_state_2)
n, l1 := l0.Recv()
l_, l2 := l1.Recv()
l_0 := l_.(*_state_0)
l2.Recv()
r_ := init_state_0(make(chan interface{}))
go cons(((n * 2) + 1))(r_, r)
d := init_state_0(make(chan interface{}))
go doMap()(d, l_0, r_)
// FWD c d Start
for {
cd := d.Recv()
c.Send(cd)
switch cd {
case "null":
d0 := d.ls["null"].(*_state_1)
c0 := c.ls["null"].(*_state_1)
d0.Recv()
c0.Send(nil)
return
case "cons":
d0 := d.ls["cons"].(*_state_2)
c0 := c.ls["cons"].(*_state_2)
c0d0, d1 := d0.Recv()
c1 := c0.Send(c0d0)
c1d1, d2 := d1.Recv()
c2 := c1.Send(c1d1)
d2.Recv()
c2.Send(nil)
return
}
}
// FWD c d End
case "null" :
l0 := l.ls["null"].(*_state_1)
l0.Recv()
// FWD c r Start
for {
cr := r.Recv()
c.Send(cr)
switch cr {
case "null":
r0 := r.ls["null"].(*_state_1)
c0 := c.ls["null"].(*_state_1)
r0.Recv()
c0.Send(nil)
return
case "cons":
r0 := r.ls["cons"].(*_state_2)
c0 := c.ls["cons"].(*_state_2)
c0r0, r1 := r0.Recv()
c1 := c0.Send(c0r0)
c1r1, r2 := r1.Recv()
c2 := c1.Send(c1r1)
r2.Recv()
c2.Send(nil)
return
}
}
// FWD c r End
}
}}
func reduce(acc int) func (_x *_state_1, l *_state_0) {
 return func (c *_state_1, l *_state_0){
label := l.Recv()
switch label {
case "cons" :
l0 := l.ls["cons"].(*_state_2)
n, l1 := l0.Recv()
l_, l2 := l1.Recv()
l_0 := l_.(*_state_0)
l2.Recv()
d := init_state_1(make(chan interface{}))
go reduce((acc + n))(d, l_0)
// FWD c d Start
d.Recv()
c.Send(nil)
return
// FWD c d End
case "null" :
l0 := l.ls["null"].(*_state_1)
fmt.Printf("%v\n",acc)
l0.Recv()
c.Send(nil)
}
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
func main () {
c := init_state_1(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_1){
e0 := init_state_0(make(chan interface{}))
go null()(e0)
e1 := init_state_0(make(chan interface{}))
go cons(1)(e1, e0)
e2 := init_state_0(make(chan interface{}))
go cons(2)(e2, e1)
e3 := init_state_0(make(chan interface{}))
go cons(3)(e3, e2)
m := init_state_0(make(chan interface{}))
go null()(m)
doMapC := init_state_0(make(chan interface{}))
go doMap()(doMapC, e3, m)
result := init_state_1(make(chan interface{}))
go reduce(0)(result, doMapC)
result.Recv()
l1 := init_state_0(make(chan interface{}))
go null()(l1)
l2 := init_state_0(make(chan interface{}))
go cons(2)(l2, l1)
l3 := init_state_0(make(chan interface{}))
go cons(4)(l3, l2)
l4 := init_state_0(make(chan interface{}))
go cons(8)(l4, l3)
l5 := init_state_0(make(chan interface{}))
go cons(16)(l5, l4)
l6 := init_state_0(make(chan interface{}))
go cons(32)(l6, l5)
m1 := init_state_0(make(chan interface{}))
go null()(m1)
map1 := init_state_0(make(chan interface{}))
go doMap()(map1, l6, m1)
r := init_state_1(make(chan interface{}))
go reduce(0)(r, map1)
r.Recv()
c.Send(nil)
}(c)
}
