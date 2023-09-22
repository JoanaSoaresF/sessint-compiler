package main
import "fmt"
//Preamble generation
type _state_1 struct {
    c chan interface{}
}
func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c } }
func (x *_state_1) Send(v interface{}) { x.c <- v }
func (x *_state_1) Recv() interface{} { return <-x.c }

  type _state_0 struct {
    c  chan interface{}
    ls map[string]interface{}
  }
  func init_state_0(c chan interface{}) *_state_0 { m := make(map[string]interface{})
 	m["plus_one"] = init_state_1( c )
	m["plus_n"] = init_state_1( c )
   return &_state_0{ c, m } }
func (x *_state_0) Send(v string) { x.c <- v }
func (x *_state_0) Recv() string  { return (<-x.c).(string) }

  type _state_2 struct {
   c chan interface{}
   next *_state_1
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
func (x *_state_2) Send(v0 int, v1 int, v2 int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) };
 x.c <- []interface{}{v0, v1, v2}
return x.next }
func (x *_state_2) Recv() (int, int, int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) };ll := <- x.c
l := ll.([]interface{})
return l[0].(int), l[1].(int), l[2].(int), x.next }

type _state_5 struct {
   c chan interface{}
   next *_state_1
}

func init_state_5(c chan interface{}) *_state_5 { return &_state_5{c, nil} }
func (x *_state_5) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_5) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_4 struct {
   c chan interface{}
   next *_state_5
}

func init_state_4(c chan interface{}) *_state_4 { return &_state_4{c, nil} }
func (x *_state_4) Send(v int) *_state_5 {
   if x.next == nil { x.next = init_state_5(x.c) }; x.c <- v; return x.next }
func (x *_state_4) Recv() (int, *_state_5) {
   if x.next == nil { x.next = init_state_5(x.c) }; return (<-x.c).(int), x.next }

  type _state_3 struct {
   c chan interface{}
   next *_state_4
}

func init_state_3(c chan interface{}) *_state_3 { return &_state_3{c, nil} }
func (x *_state_3) Send(v int) *_state_4 {
   if x.next == nil { x.next = init_state_4(x.c) }; x.c <- v; return x.next }
func (x *_state_3) Recv() (int, *_state_4) {
   if x.next == nil { x.next = init_state_4(x.c) }; return (<-x.c).(int), x.next }

  //Declaration list compilation
func print_three(n int) func (_x *_state_0) {
 return func (c *_state_0){
p := init_state_3(make(chan interface{}))
go send_three(n)(p)
a1, p0 := p.Recv()
fmt.Printf("%v\n",(a1 + n))
label := c.Recv()
switch label {
case "plus_n" :
c0 := c.ls["plus_n"].(*_state_1)
a2, p1 := p0.Recv()
fmt.Printf("%v\n",a2)
a3, p2 := p1.Recv()
fmt.Printf("%v\n",a3)
p2.Recv()
c0.Send(nil)
case "plus_one" :
c0 := c.ls["plus_one"].(*_state_1)
b1, p1 := p0.Recv()
fmt.Printf("%v\n",(b1 + 1))
b2, p2 := p1.Recv()
fmt.Printf("%v\n",(b2 + 1))
p2.Recv()
c0.Send(nil)
}
}}
func print_three_optimized(n int) func (_x *_state_0) {
 return func (c *_state_0){
p := init_state_2(make(chan interface{}))
go send_three_optimized(n)(p)
a1, var_multi_0, var_multi_1, p0 := p.Recv()
fmt.Printf("%v\n",(a1 + n))
label := c.Recv()
switch label {
case "plus_n" :
c0 := c.ls["plus_n"].(*_state_1)
fmt.Printf("%v\n",var_multi_0)
fmt.Printf("%v\n",var_multi_1)
p0.Recv()
c0.Send(nil)
case "plus_one" :
c0 := c.ls["plus_one"].(*_state_1)
fmt.Printf("%v\n",(var_multi_0 + 1))
fmt.Printf("%v\n",(var_multi_1 + 1))
p0.Recv()
c0.Send(nil)
}
}}
func send_three_optimized(n int) func (_x *_state_2) {
 return func (c *_state_2){
c0 := c.Send(n, (n + 1), (n + 2))
c0.Send(nil)
}}
func send_three(n int) func (_x *_state_3) {
 return func (c *_state_3){
c0 := c.Send(n)
c1 := c0.Send((n + 1))
c2 := c1.Send((n + 2))
c2.Send(nil)
}}
//Main compilation
func main () {
    m:= init_state_1(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_1){
e := init_state_0(make(chan interface{}))
go print_three_optimized(1)(e)
e.Send("plus_n")
e0 := e.ls["plus_n"].(*_state_1)
e0.Recv()
m.Send(nil)
}(m)
}
