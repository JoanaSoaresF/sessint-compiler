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

  //Declaration list compilation
func plus_two(n int) func (_x *_state_0) {
 return func (c *_state_0){
plus_one((n + 1))(c)
}}
func plus_one(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send((n + 1))
c0.Send(nil)
}}
//Main compilation
func main () {
    m:= init_state_1(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_1){
d1d := init_state_0(make(chan interface{}))
go plus_two(1)(d1d)
a1, d1d0 := d1d.Recv()
fmt.Printf("%v\n",a1)
d1d0.Recv()
d2d := init_state_0(make(chan interface{}))
go plus_two(2)(d2d)
a2, d2d0 := d2d.Recv()
fmt.Printf("%v\n",a2)
d2d0.Recv()
d3d := init_state_0(make(chan interface{}))
go plus_two(3)(d3d)
a3, d3d0 := d3d.Recv()
fmt.Printf("%v\n",a3)
d3d0.Recv()
d4d := init_state_0(make(chan interface{}))
go plus_two(4)(d4d)
a4, d4d0 := d4d.Recv()
fmt.Printf("%v\n",a4)
d4d0.Recv()
d5d := init_state_0(make(chan interface{}))
go plus_two(5)(d5d)
a5, d5d0 := d5d.Recv()
fmt.Printf("%v\n",a5)
d5d0.Recv()
d6d := init_state_0(make(chan interface{}))
go plus_two(6)(d6d)
a6, d6d0 := d6d.Recv()
fmt.Printf("%v\n",a6)
d6d0.Recv()
d7d := init_state_0(make(chan interface{}))
go plus_two(7)(d7d)
a7, d7d0 := d7d.Recv()
fmt.Printf("%v\n",a7)
d7d0.Recv()
d8d := init_state_0(make(chan interface{}))
go plus_two(8)(d8d)
a8, d8d0 := d8d.Recv()
fmt.Printf("%v\n",a8)
d8d0.Recv()
d9d := init_state_0(make(chan interface{}))
go plus_two(9)(d9d)
a9, d9d0 := d9d.Recv()
fmt.Printf("%v\n",a9)
d9d0.Recv()
d10d := init_state_0(make(chan interface{}))
go plus_two(10)(d10d)
a10, d10d0 := d10d.Recv()
fmt.Printf("%v\n",a10)
d10d0.Recv()
d11d := init_state_0(make(chan interface{}))
go plus_two(11)(d11d)
a11, d11d0 := d11d.Recv()
fmt.Printf("%v\n",a11)
d11d0.Recv()
d12d := init_state_0(make(chan interface{}))
go plus_two(12)(d12d)
a12, d12d0 := d12d.Recv()
fmt.Printf("%v\n",a12)
d12d0.Recv()
d13d := init_state_0(make(chan interface{}))
go plus_two(13)(d13d)
a13, d13d0 := d13d.Recv()
fmt.Printf("%v\n",a13)
d13d0.Recv()
d14d := init_state_0(make(chan interface{}))
go plus_two(14)(d14d)
a14, d14d0 := d14d.Recv()
fmt.Printf("%v\n",a14)
d14d0.Recv()
d15d := init_state_0(make(chan interface{}))
go plus_two(15)(d15d)
a15, d15d0 := d15d.Recv()
fmt.Printf("%v\n",a15)
d15d0.Recv()
d16d := init_state_0(make(chan interface{}))
go plus_two(16)(d16d)
a16, d16d0 := d16d.Recv()
fmt.Printf("%v\n",a16)
d16d0.Recv()
d17d := init_state_0(make(chan interface{}))
go plus_two(17)(d17d)
a17, d17d0 := d17d.Recv()
fmt.Printf("%v\n",a17)
d17d0.Recv()
d18d := init_state_0(make(chan interface{}))
go plus_two(18)(d18d)
a18, d18d0 := d18d.Recv()
fmt.Printf("%v\n",a18)
d18d0.Recv()
d19d := init_state_0(make(chan interface{}))
go plus_two(19)(d19d)
a19, d19d0 := d19d.Recv()
fmt.Printf("%v\n",a19)
d19d0.Recv()
d20d := init_state_0(make(chan interface{}))
go plus_two(20)(d20d)
a20, d20d0 := d20d.Recv()
fmt.Printf("%v\n",a20)
d20d0.Recv()
d21d := init_state_0(make(chan interface{}))
go plus_two(21)(d21d)
a21, d21d0 := d21d.Recv()
fmt.Printf("%v\n",a21)
d21d0.Recv()
d22d := init_state_0(make(chan interface{}))
go plus_two(22)(d22d)
a22, d22d0 := d22d.Recv()
fmt.Printf("%v\n",a22)
d22d0.Recv()
d23d := init_state_0(make(chan interface{}))
go plus_two(23)(d23d)
a23, d23d0 := d23d.Recv()
fmt.Printf("%v\n",a23)
d23d0.Recv()
d24d := init_state_0(make(chan interface{}))
go plus_two(24)(d24d)
a24, d24d0 := d24d.Recv()
fmt.Printf("%v\n",a24)
d24d0.Recv()
d25d := init_state_0(make(chan interface{}))
go plus_two(25)(d25d)
a25, d25d0 := d25d.Recv()
fmt.Printf("%v\n",a25)
d25d0.Recv()
d26d := init_state_0(make(chan interface{}))
go plus_two(26)(d26d)
a26, d26d0 := d26d.Recv()
fmt.Printf("%v\n",a26)
d26d0.Recv()
d27d := init_state_0(make(chan interface{}))
go plus_two(27)(d27d)
a27, d27d0 := d27d.Recv()
fmt.Printf("%v\n",a27)
d27d0.Recv()
d28d := init_state_0(make(chan interface{}))
go plus_two(28)(d28d)
a28, d28d0 := d28d.Recv()
fmt.Printf("%v\n",a28)
d28d0.Recv()
d29d := init_state_0(make(chan interface{}))
go plus_two(29)(d29d)
a29, d29d0 := d29d.Recv()
fmt.Printf("%v\n",a29)
d29d0.Recv()
d30d := init_state_0(make(chan interface{}))
go plus_two(30)(d30d)
a30, d30d0 := d30d.Recv()
fmt.Printf("%v\n",a30)
d30d0.Recv()
d31d := init_state_0(make(chan interface{}))
go plus_two(31)(d31d)
a31, d31d0 := d31d.Recv()
fmt.Printf("%v\n",a31)
d31d0.Recv()
d32d := init_state_0(make(chan interface{}))
go plus_two(32)(d32d)
a32, d32d0 := d32d.Recv()
fmt.Printf("%v\n",a32)
d32d0.Recv()
d33d := init_state_0(make(chan interface{}))
go plus_two(33)(d33d)
a33, d33d0 := d33d.Recv()
fmt.Printf("%v\n",a33)
d33d0.Recv()
d34d := init_state_0(make(chan interface{}))
go plus_two(34)(d34d)
a34, d34d0 := d34d.Recv()
fmt.Printf("%v\n",a34)
d34d0.Recv()
d35d := init_state_0(make(chan interface{}))
go plus_two(35)(d35d)
a35, d35d0 := d35d.Recv()
fmt.Printf("%v\n",a35)
d35d0.Recv()
d36d := init_state_0(make(chan interface{}))
go plus_two(36)(d36d)
a36, d36d0 := d36d.Recv()
fmt.Printf("%v\n",a36)
d36d0.Recv()
d37d := init_state_0(make(chan interface{}))
go plus_two(37)(d37d)
a37, d37d0 := d37d.Recv()
fmt.Printf("%v\n",a37)
d37d0.Recv()
d38d := init_state_0(make(chan interface{}))
go plus_two(38)(d38d)
a38, d38d0 := d38d.Recv()
fmt.Printf("%v\n",a38)
d38d0.Recv()
d39d := init_state_0(make(chan interface{}))
go plus_two(39)(d39d)
a39, d39d0 := d39d.Recv()
fmt.Printf("%v\n",a39)
d39d0.Recv()
d40d := init_state_0(make(chan interface{}))
go plus_two(40)(d40d)
a40, d40d0 := d40d.Recv()
fmt.Printf("%v\n",a40)
d40d0.Recv()
d41d := init_state_0(make(chan interface{}))
go plus_two(41)(d41d)
a41, d41d0 := d41d.Recv()
fmt.Printf("%v\n",a41)
d41d0.Recv()
d42d := init_state_0(make(chan interface{}))
go plus_two(42)(d42d)
a42, d42d0 := d42d.Recv()
fmt.Printf("%v\n",a42)
d42d0.Recv()
d43d := init_state_0(make(chan interface{}))
go plus_two(43)(d43d)
a43, d43d0 := d43d.Recv()
fmt.Printf("%v\n",a43)
d43d0.Recv()
d44d := init_state_0(make(chan interface{}))
go plus_two(44)(d44d)
a44, d44d0 := d44d.Recv()
fmt.Printf("%v\n",a44)
d44d0.Recv()
d45d := init_state_0(make(chan interface{}))
go plus_two(45)(d45d)
a45, d45d0 := d45d.Recv()
fmt.Printf("%v\n",a45)
d45d0.Recv()
d46d := init_state_0(make(chan interface{}))
go plus_two(46)(d46d)
a46, d46d0 := d46d.Recv()
fmt.Printf("%v\n",a46)
d46d0.Recv()
d47d := init_state_0(make(chan interface{}))
go plus_two(47)(d47d)
a47, d47d0 := d47d.Recv()
fmt.Printf("%v\n",a47)
d47d0.Recv()
d48d := init_state_0(make(chan interface{}))
go plus_two(48)(d48d)
a48, d48d0 := d48d.Recv()
fmt.Printf("%v\n",a48)
d48d0.Recv()
d49d := init_state_0(make(chan interface{}))
go plus_two(49)(d49d)
a49, d49d0 := d49d.Recv()
fmt.Printf("%v\n",a49)
d49d0.Recv()
d50d := init_state_0(make(chan interface{}))
go plus_two(50)(d50d)
a50, d50d0 := d50d.Recv()
fmt.Printf("%v\n",a50)
d50d0.Recv()
d51d := init_state_0(make(chan interface{}))
go plus_two(51)(d51d)
a51, d51d0 := d51d.Recv()
fmt.Printf("%v\n",a51)
d51d0.Recv()
d52d := init_state_0(make(chan interface{}))
go plus_two(52)(d52d)
a52, d52d0 := d52d.Recv()
fmt.Printf("%v\n",a52)
d52d0.Recv()
d53d := init_state_0(make(chan interface{}))
go plus_two(53)(d53d)
a53, d53d0 := d53d.Recv()
fmt.Printf("%v\n",a53)
d53d0.Recv()
d54d := init_state_0(make(chan interface{}))
go plus_two(54)(d54d)
a54, d54d0 := d54d.Recv()
fmt.Printf("%v\n",a54)
d54d0.Recv()
d55d := init_state_0(make(chan interface{}))
go plus_two(55)(d55d)
a55, d55d0 := d55d.Recv()
fmt.Printf("%v\n",a55)
d55d0.Recv()
d56d := init_state_0(make(chan interface{}))
go plus_two(56)(d56d)
a56, d56d0 := d56d.Recv()
fmt.Printf("%v\n",a56)
d56d0.Recv()
d57d := init_state_0(make(chan interface{}))
go plus_two(57)(d57d)
a57, d57d0 := d57d.Recv()
fmt.Printf("%v\n",a57)
d57d0.Recv()
d58d := init_state_0(make(chan interface{}))
go plus_two(58)(d58d)
a58, d58d0 := d58d.Recv()
fmt.Printf("%v\n",a58)
d58d0.Recv()
d59d := init_state_0(make(chan interface{}))
go plus_two(59)(d59d)
a59, d59d0 := d59d.Recv()
fmt.Printf("%v\n",a59)
d59d0.Recv()
d60d := init_state_0(make(chan interface{}))
go plus_two(60)(d60d)
a60, d60d0 := d60d.Recv()
fmt.Printf("%v\n",a60)
d60d0.Recv()
d61d := init_state_0(make(chan interface{}))
go plus_two(61)(d61d)
a61, d61d0 := d61d.Recv()
fmt.Printf("%v\n",a61)
d61d0.Recv()
d62d := init_state_0(make(chan interface{}))
go plus_two(62)(d62d)
a62, d62d0 := d62d.Recv()
fmt.Printf("%v\n",a62)
d62d0.Recv()
d63d := init_state_0(make(chan interface{}))
go plus_two(63)(d63d)
a63, d63d0 := d63d.Recv()
fmt.Printf("%v\n",a63)
d63d0.Recv()
d64d := init_state_0(make(chan interface{}))
go plus_two(64)(d64d)
a64, d64d0 := d64d.Recv()
fmt.Printf("%v\n",a64)
d64d0.Recv()
m.Send(nil)
}(m)
}
