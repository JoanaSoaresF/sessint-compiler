package main
import ("fmt"
)
type _state_3 struct {
    c chan interface{}
  }
  func init_state_3(c chan interface{}) *_state_3 { return &_state_3{ c } } 
  func (x *_state_3) Send(v interface{}) { x.c <- v }
  func (x *_state_3) Recv() interface{} { return <-x.c }

  type _state_2 struct {
    c chan interface{}
    next *_state_3
  }
  
  func init_state_2(c chan interface{}) *_state_2 { return &_state_2{ c, nil } } 
  func (x *_state_2) Send(v int) *_state_3 { if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next}
  func (x *_state_2) Recv() (int, *_state_3) { if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int),x.next}

  type _state_1 struct {
    c chan interface{}
    next *_state_2
  }
  
  func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c, nil } } 
  func (x *_state_1) Send(v int) *_state_2 { if x.next == nil { x.next = init_state_2(x.c) }; x.c <- v; return x.next}
  func (x *_state_1) Recv() (int, *_state_2) { if x.next == nil { x.next = init_state_2(x.c) }; return (<-x.c).(int),x.next}

  type _state_0 struct {
    c chan interface{}
    next *_state_1
  }
  
  func init_state_0(c chan interface{}) *_state_0 { return &_state_0{ c, nil } } 
  func (x *_state_0) Send(v int) *_state_1 { if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next}
  func (x *_state_0) Recv() (int, *_state_1) { if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int),x.next}

  type _state_4 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_4(c chan interface{}) *_state_4 { m := make(map[string]interface{})
	m["plus_one"] = init_state_3( c )
	m["plus_n"] = init_state_3( c )
	return &_state_4{ c, m } }
  func (x *_state_4) Send(v string) { x.c <- v }
  func (x *_state_4) Recv() string  { return (<-x.c).(string) }

  func send_three(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send(n)
c1 := c0.Send((n + 1))
c2 := c1.Send((n + 2))
c2.Send(nil)
}}
func print_three(n int) func (_x *_state_4) {
 return func (c *_state_4){
p := init_state_0(make(chan interface{}))
go send_three(n)(p)
a1, p0 := p.Recv()
fmt.Printf("%v\n",(a1 + n))
label := c.Recv()
switch label {
case "plus_n" :
c0 := c.ls["plus_n"].(*_state_3)
a2, p1 := p0.Recv()
fmt.Printf("%v\n",a2)
a3, p2 := p1.Recv()
fmt.Printf("%v\n",a3)
p2.Recv()
c0.Send(nil)
case "plus_one" :
c0 := c.ls["plus_one"].(*_state_3)
b1, p1 := p0.Recv()
fmt.Printf("%v\n",(b1 + 1))
b2, p2 := p1.Recv()
fmt.Printf("%v\n",(b2 + 1))
p2.Recv()
c0.Send(nil)
}
}}
func main () {
m := init_state_3(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_3){
e := init_state_4(make(chan interface{}))
go print_three(1)(e)
e.Send("plus_n")
e0 := e.ls["plus_n"].(*_state_3)
e0.Recv()
m.Send(nil)
}(m)
}
