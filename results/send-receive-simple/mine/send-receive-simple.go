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
   next *_state_0
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
func (x *_state_2) Send(v0 int, v1 int, v2 int, v3 int, v4 int, v5 int, v6 int, v7 int, v8 int, v9 int) *_state_0 {
   if x.next == nil { x.next = init_state_0(x.c) };
 x.c <- []interface{}{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9}
return x.next }
func (x *_state_2) Recv() (int, int, int, int, int, int, int, int, int, int, *_state_0) {
   if x.next == nil { x.next = init_state_0(x.c) };ll := <- x.c
l := ll.([]interface{})
return l[0].(int), l[1].(int), l[2].(int), l[3].(int), l[4].(int), l[5].(int), l[6].(int), l[7].(int), l[8].(int), l[9].(int), x.next }

type _state_1 struct {
   c chan interface{}
   next *_state_2
}

func init_state_1(c chan interface{}) *_state_1 { return &_state_1{c, nil} }
func (x *_state_1) Send(v0 int, v1 int, v2 int, v3 int, v4 int, v5 int, v6 int, v7 int, v8 int, v9 int) *_state_2 {
   if x.next == nil { x.next = init_state_2(x.c) };
 x.c <- []interface{}{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9}
return x.next }
func (x *_state_1) Recv() (int, int, int, int, int, int, int, int, int, int, *_state_2) {
   if x.next == nil { x.next = init_state_2(x.c) };ll := <- x.c
l := ll.([]interface{})
return l[0].(int), l[1].(int), l[2].(int), l[3].(int), l[4].(int), l[5].(int), l[6].(int), l[7].(int), l[8].(int), l[9].(int), x.next }

//Declaration list compilation
//Main compilation
func main () {
    c:= init_state_0(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_0){
d := init_state_1(make(chan interface{}))
go func (d *_state_1){
v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, d0 := d.Recv()
d1 := d0.Send(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)
d1.Send(nil)
}(d)
d0 := d.Send(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, d1 := d0.Recv()
fmt.Printf("%v\n",v1)
fmt.Printf("%v\n",v2)
fmt.Printf("%v\n",v3)
fmt.Printf("%v\n",v4)
fmt.Printf("%v\n",v5)
fmt.Printf("%v\n",v6)
fmt.Printf("%v\n",v7)
fmt.Printf("%v\n",v8)
fmt.Printf("%v\n",v9)
fmt.Printf("%v\n",v10)
fmt.Printf("%v\n",0)
d1.Recv()
c.Send(nil)
}(c)
}
