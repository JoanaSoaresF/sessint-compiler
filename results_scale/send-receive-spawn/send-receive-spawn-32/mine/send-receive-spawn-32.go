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
func (x *_state_0) Send(v0 int, v1 int, v2 int, v3 int, v4 int, v5 int, v6 int, v7 int, v8 int, v9 int, v10 int, v11 int, v12 int, v13 int, v14 int, v15 int, v16 int, v17 int, v18 int, v19 int, v20 int, v21 int, v22 int, v23 int, v24 int, v25 int, v26 int, v27 int, v28 int, v29 int, v30 int, v31 int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) };
 x.c <- []interface{}{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31}
return x.next }
func (x *_state_0) Recv() (int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) };ll := <- x.c
l := ll.([]interface{})
return l[0].(int), l[1].(int), l[2].(int), l[3].(int), l[4].(int), l[5].(int), l[6].(int), l[7].(int), l[8].(int), l[9].(int), l[10].(int), l[11].(int), l[12].(int), l[13].(int), l[14].(int), l[15].(int), l[16].(int), l[17].(int), l[18].(int), l[19].(int), l[20].(int), l[21].(int), l[22].(int), l[23].(int), l[24].(int), l[25].(int), l[26].(int), l[27].(int), l[28].(int), l[29].(int), l[30].(int), l[31].(int), x.next }

type _state_33 struct {
   c chan interface{}
   next *_state_1
}

func init_state_33(c chan interface{}) *_state_33 { return &_state_33{c, nil} }
func (x *_state_33) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_33) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_32 struct {
   c chan interface{}
   next *_state_33
}

func init_state_32(c chan interface{}) *_state_32 { return &_state_32{c, nil} }
func (x *_state_32) Send(v int) *_state_33 {
   if x.next == nil { x.next = init_state_33(x.c) }; x.c <- v; return x.next }
func (x *_state_32) Recv() (int, *_state_33) {
   if x.next == nil { x.next = init_state_33(x.c) }; return (<-x.c).(int), x.next }

  type _state_31 struct {
   c chan interface{}
   next *_state_32
}

func init_state_31(c chan interface{}) *_state_31 { return &_state_31{c, nil} }
func (x *_state_31) Send(v int) *_state_32 {
   if x.next == nil { x.next = init_state_32(x.c) }; x.c <- v; return x.next }
func (x *_state_31) Recv() (int, *_state_32) {
   if x.next == nil { x.next = init_state_32(x.c) }; return (<-x.c).(int), x.next }

  type _state_30 struct {
   c chan interface{}
   next *_state_31
}

func init_state_30(c chan interface{}) *_state_30 { return &_state_30{c, nil} }
func (x *_state_30) Send(v int) *_state_31 {
   if x.next == nil { x.next = init_state_31(x.c) }; x.c <- v; return x.next }
func (x *_state_30) Recv() (int, *_state_31) {
   if x.next == nil { x.next = init_state_31(x.c) }; return (<-x.c).(int), x.next }

  type _state_29 struct {
   c chan interface{}
   next *_state_30
}

func init_state_29(c chan interface{}) *_state_29 { return &_state_29{c, nil} }
func (x *_state_29) Send(v int) *_state_30 {
   if x.next == nil { x.next = init_state_30(x.c) }; x.c <- v; return x.next }
func (x *_state_29) Recv() (int, *_state_30) {
   if x.next == nil { x.next = init_state_30(x.c) }; return (<-x.c).(int), x.next }

  type _state_28 struct {
   c chan interface{}
   next *_state_29
}

func init_state_28(c chan interface{}) *_state_28 { return &_state_28{c, nil} }
func (x *_state_28) Send(v int) *_state_29 {
   if x.next == nil { x.next = init_state_29(x.c) }; x.c <- v; return x.next }
func (x *_state_28) Recv() (int, *_state_29) {
   if x.next == nil { x.next = init_state_29(x.c) }; return (<-x.c).(int), x.next }

  type _state_27 struct {
   c chan interface{}
   next *_state_28
}

func init_state_27(c chan interface{}) *_state_27 { return &_state_27{c, nil} }
func (x *_state_27) Send(v int) *_state_28 {
   if x.next == nil { x.next = init_state_28(x.c) }; x.c <- v; return x.next }
func (x *_state_27) Recv() (int, *_state_28) {
   if x.next == nil { x.next = init_state_28(x.c) }; return (<-x.c).(int), x.next }

  type _state_26 struct {
   c chan interface{}
   next *_state_27
}

func init_state_26(c chan interface{}) *_state_26 { return &_state_26{c, nil} }
func (x *_state_26) Send(v int) *_state_27 {
   if x.next == nil { x.next = init_state_27(x.c) }; x.c <- v; return x.next }
func (x *_state_26) Recv() (int, *_state_27) {
   if x.next == nil { x.next = init_state_27(x.c) }; return (<-x.c).(int), x.next }

  type _state_25 struct {
   c chan interface{}
   next *_state_26
}

func init_state_25(c chan interface{}) *_state_25 { return &_state_25{c, nil} }
func (x *_state_25) Send(v int) *_state_26 {
   if x.next == nil { x.next = init_state_26(x.c) }; x.c <- v; return x.next }
func (x *_state_25) Recv() (int, *_state_26) {
   if x.next == nil { x.next = init_state_26(x.c) }; return (<-x.c).(int), x.next }

  type _state_24 struct {
   c chan interface{}
   next *_state_25
}

func init_state_24(c chan interface{}) *_state_24 { return &_state_24{c, nil} }
func (x *_state_24) Send(v int) *_state_25 {
   if x.next == nil { x.next = init_state_25(x.c) }; x.c <- v; return x.next }
func (x *_state_24) Recv() (int, *_state_25) {
   if x.next == nil { x.next = init_state_25(x.c) }; return (<-x.c).(int), x.next }

  type _state_23 struct {
   c chan interface{}
   next *_state_24
}

func init_state_23(c chan interface{}) *_state_23 { return &_state_23{c, nil} }
func (x *_state_23) Send(v int) *_state_24 {
   if x.next == nil { x.next = init_state_24(x.c) }; x.c <- v; return x.next }
func (x *_state_23) Recv() (int, *_state_24) {
   if x.next == nil { x.next = init_state_24(x.c) }; return (<-x.c).(int), x.next }

  type _state_22 struct {
   c chan interface{}
   next *_state_23
}

func init_state_22(c chan interface{}) *_state_22 { return &_state_22{c, nil} }
func (x *_state_22) Send(v int) *_state_23 {
   if x.next == nil { x.next = init_state_23(x.c) }; x.c <- v; return x.next }
func (x *_state_22) Recv() (int, *_state_23) {
   if x.next == nil { x.next = init_state_23(x.c) }; return (<-x.c).(int), x.next }

  type _state_21 struct {
   c chan interface{}
   next *_state_22
}

func init_state_21(c chan interface{}) *_state_21 { return &_state_21{c, nil} }
func (x *_state_21) Send(v int) *_state_22 {
   if x.next == nil { x.next = init_state_22(x.c) }; x.c <- v; return x.next }
func (x *_state_21) Recv() (int, *_state_22) {
   if x.next == nil { x.next = init_state_22(x.c) }; return (<-x.c).(int), x.next }

  type _state_20 struct {
   c chan interface{}
   next *_state_21
}

func init_state_20(c chan interface{}) *_state_20 { return &_state_20{c, nil} }
func (x *_state_20) Send(v int) *_state_21 {
   if x.next == nil { x.next = init_state_21(x.c) }; x.c <- v; return x.next }
func (x *_state_20) Recv() (int, *_state_21) {
   if x.next == nil { x.next = init_state_21(x.c) }; return (<-x.c).(int), x.next }

  type _state_19 struct {
   c chan interface{}
   next *_state_20
}

func init_state_19(c chan interface{}) *_state_19 { return &_state_19{c, nil} }
func (x *_state_19) Send(v int) *_state_20 {
   if x.next == nil { x.next = init_state_20(x.c) }; x.c <- v; return x.next }
func (x *_state_19) Recv() (int, *_state_20) {
   if x.next == nil { x.next = init_state_20(x.c) }; return (<-x.c).(int), x.next }

  type _state_18 struct {
   c chan interface{}
   next *_state_19
}

func init_state_18(c chan interface{}) *_state_18 { return &_state_18{c, nil} }
func (x *_state_18) Send(v int) *_state_19 {
   if x.next == nil { x.next = init_state_19(x.c) }; x.c <- v; return x.next }
func (x *_state_18) Recv() (int, *_state_19) {
   if x.next == nil { x.next = init_state_19(x.c) }; return (<-x.c).(int), x.next }

  type _state_17 struct {
   c chan interface{}
   next *_state_18
}

func init_state_17(c chan interface{}) *_state_17 { return &_state_17{c, nil} }
func (x *_state_17) Send(v int) *_state_18 {
   if x.next == nil { x.next = init_state_18(x.c) }; x.c <- v; return x.next }
func (x *_state_17) Recv() (int, *_state_18) {
   if x.next == nil { x.next = init_state_18(x.c) }; return (<-x.c).(int), x.next }

  type _state_16 struct {
   c chan interface{}
   next *_state_17
}

func init_state_16(c chan interface{}) *_state_16 { return &_state_16{c, nil} }
func (x *_state_16) Send(v int) *_state_17 {
   if x.next == nil { x.next = init_state_17(x.c) }; x.c <- v; return x.next }
func (x *_state_16) Recv() (int, *_state_17) {
   if x.next == nil { x.next = init_state_17(x.c) }; return (<-x.c).(int), x.next }

  type _state_15 struct {
   c chan interface{}
   next *_state_16
}

func init_state_15(c chan interface{}) *_state_15 { return &_state_15{c, nil} }
func (x *_state_15) Send(v int) *_state_16 {
   if x.next == nil { x.next = init_state_16(x.c) }; x.c <- v; return x.next }
func (x *_state_15) Recv() (int, *_state_16) {
   if x.next == nil { x.next = init_state_16(x.c) }; return (<-x.c).(int), x.next }

  type _state_14 struct {
   c chan interface{}
   next *_state_15
}

func init_state_14(c chan interface{}) *_state_14 { return &_state_14{c, nil} }
func (x *_state_14) Send(v int) *_state_15 {
   if x.next == nil { x.next = init_state_15(x.c) }; x.c <- v; return x.next }
func (x *_state_14) Recv() (int, *_state_15) {
   if x.next == nil { x.next = init_state_15(x.c) }; return (<-x.c).(int), x.next }

  type _state_13 struct {
   c chan interface{}
   next *_state_14
}

func init_state_13(c chan interface{}) *_state_13 { return &_state_13{c, nil} }
func (x *_state_13) Send(v int) *_state_14 {
   if x.next == nil { x.next = init_state_14(x.c) }; x.c <- v; return x.next }
func (x *_state_13) Recv() (int, *_state_14) {
   if x.next == nil { x.next = init_state_14(x.c) }; return (<-x.c).(int), x.next }

  type _state_12 struct {
   c chan interface{}
   next *_state_13
}

func init_state_12(c chan interface{}) *_state_12 { return &_state_12{c, nil} }
func (x *_state_12) Send(v int) *_state_13 {
   if x.next == nil { x.next = init_state_13(x.c) }; x.c <- v; return x.next }
func (x *_state_12) Recv() (int, *_state_13) {
   if x.next == nil { x.next = init_state_13(x.c) }; return (<-x.c).(int), x.next }

  type _state_11 struct {
   c chan interface{}
   next *_state_12
}

func init_state_11(c chan interface{}) *_state_11 { return &_state_11{c, nil} }
func (x *_state_11) Send(v int) *_state_12 {
   if x.next == nil { x.next = init_state_12(x.c) }; x.c <- v; return x.next }
func (x *_state_11) Recv() (int, *_state_12) {
   if x.next == nil { x.next = init_state_12(x.c) }; return (<-x.c).(int), x.next }

  type _state_10 struct {
   c chan interface{}
   next *_state_11
}

func init_state_10(c chan interface{}) *_state_10 { return &_state_10{c, nil} }
func (x *_state_10) Send(v int) *_state_11 {
   if x.next == nil { x.next = init_state_11(x.c) }; x.c <- v; return x.next }
func (x *_state_10) Recv() (int, *_state_11) {
   if x.next == nil { x.next = init_state_11(x.c) }; return (<-x.c).(int), x.next }

  type _state_9 struct {
   c chan interface{}
   next *_state_10
}

func init_state_9(c chan interface{}) *_state_9 { return &_state_9{c, nil} }
func (x *_state_9) Send(v int) *_state_10 {
   if x.next == nil { x.next = init_state_10(x.c) }; x.c <- v; return x.next }
func (x *_state_9) Recv() (int, *_state_10) {
   if x.next == nil { x.next = init_state_10(x.c) }; return (<-x.c).(int), x.next }

  type _state_8 struct {
   c chan interface{}
   next *_state_9
}

func init_state_8(c chan interface{}) *_state_8 { return &_state_8{c, nil} }
func (x *_state_8) Send(v int) *_state_9 {
   if x.next == nil { x.next = init_state_9(x.c) }; x.c <- v; return x.next }
func (x *_state_8) Recv() (int, *_state_9) {
   if x.next == nil { x.next = init_state_9(x.c) }; return (<-x.c).(int), x.next }

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
   c chan interface{}
   next *_state_7
}

func init_state_6(c chan interface{}) *_state_6 { return &_state_6{c, nil} }
func (x *_state_6) Send(v int) *_state_7 {
   if x.next == nil { x.next = init_state_7(x.c) }; x.c <- v; return x.next }
func (x *_state_6) Recv() (int, *_state_7) {
   if x.next == nil { x.next = init_state_7(x.c) }; return (<-x.c).(int), x.next }

  type _state_5 struct {
   c chan interface{}
   next *_state_6
}

func init_state_5(c chan interface{}) *_state_5 { return &_state_5{c, nil} }
func (x *_state_5) Send(v int) *_state_6 {
   if x.next == nil { x.next = init_state_6(x.c) }; x.c <- v; return x.next }
func (x *_state_5) Recv() (int, *_state_6) {
   if x.next == nil { x.next = init_state_6(x.c) }; return (<-x.c).(int), x.next }

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

  type _state_2 struct {
   c chan interface{}
   next *_state_3
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c, nil} }
func (x *_state_2) Send(v int) *_state_3 {
   if x.next == nil { x.next = init_state_3(x.c) }; x.c <- v; return x.next }
func (x *_state_2) Recv() (int, *_state_3) {
   if x.next == nil { x.next = init_state_3(x.c) }; return (<-x.c).(int), x.next }

  //Declaration list compilation
func send_ints_optimized(n int) func (_x *_state_0) {
 return func (c *_state_0){
c0 := c.Send((n + 1), (n + 2), (n + 3), (n + 4), (n + 5), (n + 6), (n + 7), (n + 8), (n + 9), (n + 10), (n + 11), (n + 12), (n + 13), (n + 14), (n + 15), (n + 16), (n + 17), (n + 18), (n + 19), (n + 20), (n + 21), (n + 22), (n + 23), (n + 24), (n + 25), (n + 26), (n + 27), (n + 28), (n + 29), (n + 30), (n + 31), (n + 32))
c0.Send(nil)
}}
func send_ints(n int) func (_x *_state_2) {
 return func (c *_state_2){
c0 := c.Send((n + 1))
c1 := c0.Send((n + 2))
c2 := c1.Send((n + 3))
c3 := c2.Send((n + 4))
c4 := c3.Send((n + 5))
c5 := c4.Send((n + 6))
c6 := c5.Send((n + 7))
c7 := c6.Send((n + 8))
c8 := c7.Send((n + 9))
c9 := c8.Send((n + 10))
c10 := c9.Send((n + 11))
c11 := c10.Send((n + 12))
c12 := c11.Send((n + 13))
c13 := c12.Send((n + 14))
c14 := c13.Send((n + 15))
c15 := c14.Send((n + 16))
c16 := c15.Send((n + 17))
c17 := c16.Send((n + 18))
c18 := c17.Send((n + 19))
c19 := c18.Send((n + 20))
c20 := c19.Send((n + 21))
c21 := c20.Send((n + 22))
c22 := c21.Send((n + 23))
c23 := c22.Send((n + 24))
c24 := c23.Send((n + 25))
c25 := c24.Send((n + 26))
c26 := c25.Send((n + 27))
c27 := c26.Send((n + 28))
c28 := c27.Send((n + 29))
c29 := c28.Send((n + 30))
c30 := c29.Send((n + 31))
c31 := c30.Send((n + 32))
c31.Send(nil)
}}
//Main compilation
func main () {
    m:= init_state_1(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_1){
e := init_state_0(make(chan interface{}))
go send_ints_optimized(1)(e)
a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31, a32, e0 := e.Recv()
fmt.Printf("%v\n",a1)
fmt.Printf("%v\n",a2)
fmt.Printf("%v\n",a3)
fmt.Printf("%v\n",a4)
fmt.Printf("%v\n",a5)
fmt.Printf("%v\n",a6)
fmt.Printf("%v\n",a7)
fmt.Printf("%v\n",a8)
fmt.Printf("%v\n",a9)
fmt.Printf("%v\n",a10)
fmt.Printf("%v\n",a11)
fmt.Printf("%v\n",a12)
fmt.Printf("%v\n",a13)
fmt.Printf("%v\n",a14)
fmt.Printf("%v\n",a15)
fmt.Printf("%v\n",a16)
fmt.Printf("%v\n",a17)
fmt.Printf("%v\n",a18)
fmt.Printf("%v\n",a19)
fmt.Printf("%v\n",a20)
fmt.Printf("%v\n",a21)
fmt.Printf("%v\n",a22)
fmt.Printf("%v\n",a23)
fmt.Printf("%v\n",a24)
fmt.Printf("%v\n",a25)
fmt.Printf("%v\n",a26)
fmt.Printf("%v\n",a27)
fmt.Printf("%v\n",a28)
fmt.Printf("%v\n",a29)
fmt.Printf("%v\n",a30)
fmt.Printf("%v\n",a31)
fmt.Printf("%v\n",a32)
e0.Recv()
m.Send(nil)
}(m)
}
