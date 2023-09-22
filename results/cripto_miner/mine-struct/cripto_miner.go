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
   c chan interface{}
   next *_state_1
}

func init_state_0(c chan interface{}) *_state_0 { return &_state_0{c, nil} }
func (x *_state_0) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_0) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_3 struct {
   c chan interface{}
   next *_state_0
}

func init_state_3(c chan interface{}) *_state_3 { return &_state_3{c, nil} }
func (x *_state_3) Send(v int) *_state_0 {
   if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next }
func (x *_state_3) Recv() (int, *_state_0) {
   if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int), x.next }

  type _state_2 struct {
   c chan interface{}
   next *_state_3
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
func (x *_state_2) Send(v int) *_state_3 {
   if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next }
func (x *_state_2) Recv() (int, *_state_3) {
   if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int), x.next }

  type _state_4 struct {
   c chan interface{}
   next *_state_0
}

func init_state_4(c chan interface{}) *_state_4 { return &_state_4{c, nil} }
type _multisend_type__state_4 struct {
v0 int
v1 int
}
func (x *_state_4) Send(v0 int, v1 int) *_state_0 {
   if x.next == nil { x.next = init_state_0(x.c) };
 x.c <- _multisend_type__state_4{v0, v1}
return x.next }
func (x *_state_4) Recv() (int, int, *_state_0) {
   if x.next == nil { x.next = init_state_0(x.c) };ll := <- x.c
l := ll.(_multisend_type__state_4)
return l.v0, l.v1, x.next }

//Declaration list compilation
func master_optimized(x int) func (_x *_state_0) {
 return func (c *_state_0){
w1 := init_state_4(make(chan interface{}))
go worker_optimized()(w1)
w2 := init_state_4(make(chan interface{}))
go worker_optimized()(w2)
w3 := init_state_4(make(chan interface{}))
go worker_optimized()(w3)
w4 := init_state_4(make(chan interface{}))
go worker_optimized()(w4)
w10 := w1.Send(0, (x / 4))
w20 := w2.Send((x / 4), (x / 2))
w30 := w3.Send((x / 2), ((3 * x) / 4))
w40 := w4.Send(((3 * x) / 4), x)
res1, w11 := w10.Recv()
w11.Recv()
res2, w21 := w20.Recv()
w21.Recv()
res3, w31 := w30.Recv()
w31.Recv()
res4, w41 := w40.Recv()
w41.Recv()
c0 := c.Send(solve(res1)(res2)(res3)(res4))
c0.Send(nil)
}}
func master(x int) func (_x *_state_0) {
 return func (c *_state_0){
w1 := init_state_2(make(chan interface{}))
go worker()(w1)
w2 := init_state_2(make(chan interface{}))
go worker()(w2)
w3 := init_state_2(make(chan interface{}))
go worker()(w3)
w4 := init_state_2(make(chan interface{}))
go worker()(w4)
w10 := w1.Send(0)
w11 := w10.Send((x / 4))
w20 := w2.Send((x / 4))
w21 := w20.Send((x / 2))
w30 := w3.Send((x / 2))
w31 := w30.Send(((3 * x) / 4))
w40 := w4.Send(((3 * x) / 4))
w41 := w40.Send(x)
res1, w12 := w11.Recv()
w12.Recv()
res2, w22 := w21.Recv()
w22.Recv()
res3, w32 := w31.Recv()
w32.Recv()
res4, w42 := w41.Recv()
w42.Recv()
c0 := c.Send(solve(res1)(res2)(res3)(res4))
c0.Send(nil)
}}
func worker()func (_x *_state_2) {
 return func (w *_state_2){
low, w0 := w.Recv()
high, w1 := w0.Recv()
w2 := w1.Send(partial_solve(low)(high))
w2.Send(nil)
}}
func worker_optimized()func (_x *_state_4) {
 return func (w *_state_4){
low, high, w0 := w.Recv()
w1 := w0.Send(partial_solve(low)(high))
w1.Send(nil)
}}
func solve(a int) func (_x int) func (_x int) func (_x int) int {
 return func (b int) func (_x int) func (_x int) int{
return func (c int) func (_x int) int{
return func (d int) int{
return (((a + b) + c) + d)}
}
}
}
func partial_solve(a int) func (_x int) int {
 return func (b int) int{
return (a * b)}
}
//Main compilation
func main () {
    c:= init_state_1(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_1){
d := init_state_0(make(chan interface{}))
go master_optimized(16)(d)
res, d0 := d.Recv()
fmt.Printf("%v\n",res)
d0.Recv()
c.Send(nil)
}(c)
}
