package main
import ("fmt"
)
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
  
  func init_state_2(c chan interface{}) *_state_2 { return &_state_2{ c, nil } } 
  func (x *_state_2) Send(v int) *_state_0 { if x.next == nil { x.next = init_state_0(x.c) }; x.c <- v; return x.next}
  func (x *_state_2) Recv() (int, *_state_0) { if x.next == nil { x.next = init_state_0(x.c) }; return (<-x.c).(int),x.next}

  type _state_1 struct {
    c chan interface{}
    next *_state_2
  }
  
  func init_state_1(c chan interface{}) *_state_1 { return &_state_1{ c, nil } } 
  func (x *_state_1) Send(v int) *_state_2 { if x.next == nil { x.next = init_state_2(x.c) }; x.c <- v; return x.next}
  func (x *_state_1) Recv() (int, *_state_2) { if x.next == nil { x.next = init_state_2(x.c) }; return (<-x.c).(int),x.next}

  func main () {
c := init_state_0(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_0){
d := init_state_1(make(chan interface{}))
go func (d *_state_1){
v1, d0 := d.Recv()
d1 := d0.Send((v1 + 1))
d1.Send(nil)
}(d)
d0 := d.Send(1)
_, d1 := d0.Recv()
fmt.Printf("%v\n",0)
d1.Recv()
c.Send(nil)
}(c)
}
