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

  func nats(n int) func (_x *_state_0) {
 return func (c *_state_0){
label := c.Recv()
switch label {
case "next" :
c0 := c.ls["next"].(*_state_1)
c1 := c0.Send(n)
d := init_state_0(make(chan interface{}))
go nats((n + 1))(d)
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
func sum() func (_x *_state_0, n *_state_0, m *_state_0) {
 return func (r_channel *_state_0, n *_state_0, m *_state_0){
label := r_channel.Recv()
switch label {
case "next" :
r_channel0 := r_channel.ls["next"].(*_state_1)
n.Send("next")
n0 := n.ls["next"].(*_state_1)
v1, n1 := n0.Recv()
m.Send("next")
m0 := m.ls["next"].(*_state_1)
v2, m1 := m0.Recv()
r_channel1 := r_channel0.Send((v1 + v2))
d_ := init_state_0(make(chan interface{}))
go sum()(d_, n1, m1)
// FWD r_channel d_ Start
for {
d_r_channel1 := r_channel1.Recv()
d_.Send(d_r_channel1)
switch d_r_channel1 {
case "next":
d_0 := d_.ls["next"].(*_state_1)
r_channel2 := r_channel1.ls["next"].(*_state_1)
r_channel2d_0, r_channel2_d_0 := d_0.Recv()
d_ = r_channel2_d_0
r_channel1 = r_channel2.Send(r_channel2d_0)
case "stop":
d_0 := d_.ls["stop"].(*_state_2)
r_channel2 := r_channel1.ls["stop"].(*_state_2)
d_0.Recv()
r_channel2.Send(nil)
return
}
}
// FWD r_channel d_ End
case "stop" :
r_channel0 := r_channel.ls["stop"].(*_state_2)
n.Send("stop")
n0 := n.ls["stop"].(*_state_2)
n0.Recv()
m.Send("stop")
m0 := m.ls["stop"].(*_state_2)
m0.Recv()
r_channel0.Send(nil)
}
}}
func main () {
c := init_state_2(make (chan interface{}))
go func () {
c.Recv()
}()
func (c *_state_2){
n := init_state_0(make(chan interface{}))
go nats(0)(n)
m := init_state_0(make(chan interface{}))
go nats(20)(m)
q := init_state_0(make(chan interface{}))
go sum()(q, n, m)
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
q11.Send("stop")
q12 := q11.ls["stop"].(*_state_2)
q12.Recv()
c.Send(nil)
}(c)
}
