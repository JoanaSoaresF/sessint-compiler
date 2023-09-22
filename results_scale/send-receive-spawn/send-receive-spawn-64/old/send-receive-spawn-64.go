package main
import ("fmt"
)
type _state_64 struct {
    c chan interface{}
  }
  func init_state_64(c chan interface{}) *_state_64 { return &_state_64{ c } } 
  func (x *_state_64) Send(v interface{}) { x.c <- v }
  func (x *_state_64) Recv() interface{} { return <-x.c }

  type _state_63 struct {
    c chan interface{}
    next *_state_64
  }
  
  func init_state_63(c chan interface{}) *_state_63 { return &_state_63{ c, nil } } 
  func (x *_state_63) Send(v int) *_state_64 { if x.next == nil { x.next = init_state_64(x.c) }; x.c <- v; return x.next}
  func (x *_state_63) Recv() (int, *_state_64) { if x.next == nil { x.next = init_state_64(x.c) }; return (<-x.c).(int),x.next}

  type _state_62 struct {
    c chan interface{}
    next *_state_63
  }
  
  func init_state_62(c chan interface{}) *_state_62 { return &_state_62{ c, nil } } 
  func (x *_state_62) Send(v int) *_state_63 { if x.next == nil { x.next = init_state_63(x.c) }; x.c <- v; return x.next}
  func (x *_state_62) Recv() (int, *_state_63) { if x.next == nil { x.next = init_state_63(x.c) }; return (<-x.c).(int),x.next}

  type _state_61 struct {
    c chan interface{}
    next *_state_62
  }
  
  func init_state_61(c chan interface{}) *_state_61 { return &_state_61{ c, nil } } 
  func (x *_state_61) Send(v int) *_state_62 { if x.next == nil { x.next = init_state_62(x.c) }; x.c <- v; return x.next}
  func (x *_state_61) Recv() (int, *_state_62) { if x.next == nil { x.next = init_state_62(x.c) }; return (<-x.c).(int),x.next}

  type _state_60 struct {
    c chan interface{}
    next *_state_61
  }
  
  func init_state_60(c chan interface{}) *_state_60 { return &_state_60{ c, nil } } 
  func (x *_state_60) Send(v int) *_state_61 { if x.next == nil { x.next = init_state_61(x.c) }; x.c <- v; return x.next}
  func (x *_state_60) Recv() (int, *_state_61) { if x.next == nil { x.next = init_state_61(x.c) }; return (<-x.c).(int),x.next}

  type _state_59 struct {
    c chan interface{}
    next *_state_60
  }
  
  func init_state_59(c chan interface{}) *_state_59 { return &_state_59{ c, nil } } 
  func (x *_state_59) Send(v int) *_state_60 { if x.next == nil { x.next = init_state_60(x.c) }; x.c <- v; return x.next}
  func (x *_state_59) Recv() (int, *_state_60) { if x.next == nil { x.next = init_state_60(x.c) }; return (<-x.c).(int),x.next}

  type _state_58 struct {
    c chan interface{}
    next *_state_59
  }
  
  func init_state_58(c chan interface{}) *_state_58 { return &_state_58{ c, nil } } 
  func (x *_state_58) Send(v int) *_state_59 { if x.next == nil { x.next = init_state_59(x.c) }; x.c <- v; return x.next}
  func (x *_state_58) Recv() (int, *_state_59) { if x.next == nil { x.next = init_state_59(x.c) }; return (<-x.c).(int),x.next}

  type _state_57 struct {
    c chan interface{}
    next *_state_58
  }
  
  func init_state_57(c chan interface{}) *_state_57 { return &_state_57{ c, nil } } 
  func (x *_state_57) Send(v int) *_state_58 { if x.next == nil { x.next = init_state_58(x.c) }; x.c <- v; return x.next}
  func (x *_state_57) Recv() (int, *_state_58) { if x.next == nil { x.next = init_state_58(x.c) }; return (<-x.c).(int),x.next}

  type _state_56 struct {
    c chan interface{}
    next *_state_57
  }
  
  func init_state_56(c chan interface{}) *_state_56 { return &_state_56{ c, nil } } 
  func (x *_state_56) Send(v int) *_state_57 { if x.next == nil { x.next = init_state_57(x.c) }; x.c <- v; return x.next}
  func (x *_state_56) Recv() (int, *_state_57) { if x.next == nil { x.next = init_state_57(x.c) }; return (<-x.c).(int),x.next}

  type _state_55 struct {
    c chan interface{}
    next *_state_56
  }
  
  func init_state_55(c chan interface{}) *_state_55 { return &_state_55{ c, nil } } 
  func (x *_state_55) Send(v int) *_state_56 { if x.next == nil { x.next = init_state_56(x.c) }; x.c <- v; return x.next}
  func (x *_state_55) Recv() (int, *_state_56) { if x.next == nil { x.next = init_state_56(x.c) }; return (<-x.c).(int),x.next}

  type _state_54 struct {
    c chan interface{}
    next *_state_55
  }
  
  func init_state_54(c chan interface{}) *_state_54 { return &_state_54{ c, nil } } 
  func (x *_state_54) Send(v int) *_state_55 { if x.next == nil { x.next = init_state_55(x.c) }; x.c <- v; return x.next}
  func (x *_state_54) Recv() (int, *_state_55) { if x.next == nil { x.next = init_state_55(x.c) }; return (<-x.c).(int),x.next}

  type _state_53 struct {
    c chan interface{}
    next *_state_54
  }
  
  func init_state_53(c chan interface{}) *_state_53 { return &_state_53{ c, nil } } 
  func (x *_state_53) Send(v int) *_state_54 { if x.next == nil { x.next = init_state_54(x.c) }; x.c <- v; return x.next}
  func (x *_state_53) Recv() (int, *_state_54) { if x.next == nil { x.next = init_state_54(x.c) }; return (<-x.c).(int),x.next}

  type _state_52 struct {
    c chan interface{}
    next *_state_53
  }
  
  func init_state_52(c chan interface{}) *_state_52 { return &_state_52{ c, nil } } 
  func (x *_state_52) Send(v int) *_state_53 { if x.next == nil { x.next = init_state_53(x.c) }; x.c <- v; return x.next}
  func (x *_state_52) Recv() (int, *_state_53) { if x.next == nil { x.next = init_state_53(x.c) }; return (<-x.c).(int),x.next}

  type _state_51 struct {
    c chan interface{}
    next *_state_52
  }
  
  func init_state_51(c chan interface{}) *_state_51 { return &_state_51{ c, nil } } 
  func (x *_state_51) Send(v int) *_state_52 { if x.next == nil { x.next = init_state_52(x.c) }; x.c <- v; return x.next}
  func (x *_state_51) Recv() (int, *_state_52) { if x.next == nil { x.next = init_state_52(x.c) }; return (<-x.c).(int),x.next}

  type _state_50 struct {
    c chan interface{}
    next *_state_51
  }
  
  func init_state_50(c chan interface{}) *_state_50 { return &_state_50{ c, nil } } 
  func (x *_state_50) Send(v int) *_state_51 { if x.next == nil { x.next = init_state_51(x.c) }; x.c <- v; return x.next}
  func (x *_state_50) Recv() (int, *_state_51) { if x.next == nil { x.next = init_state_51(x.c) }; return (<-x.c).(int),x.next}

  type _state_49 struct {
    c chan interface{}
    next *_state_50
  }
  
  func init_state_49(c chan interface{}) *_state_49 { return &_state_49{ c, nil } } 
  func (x *_state_49) Send(v int) *_state_50 { if x.next == nil { x.next = init_state_50(x.c) }; x.c <- v; return x.next}
  func (x *_state_49) Recv() (int, *_state_50) { if x.next == nil { x.next = init_state_50(x.c) }; return (<-x.c).(int),x.next}

  type _state_48 struct {
    c chan interface{}
    next *_state_49
  }
  
  func init_state_48(c chan interface{}) *_state_48 { return &_state_48{ c, nil } } 
  func (x *_state_48) Send(v int) *_state_49 { if x.next == nil { x.next = init_state_49(x.c) }; x.c <- v; return x.next}
  func (x *_state_48) Recv() (int, *_state_49) { if x.next == nil { x.next = init_state_49(x.c) }; return (<-x.c).(int),x.next}

  type _state_47 struct {
    c chan interface{}
    next *_state_48
  }
  
  func init_state_47(c chan interface{}) *_state_47 { return &_state_47{ c, nil } } 
  func (x *_state_47) Send(v int) *_state_48 { if x.next == nil { x.next = init_state_48(x.c) }; x.c <- v; return x.next}
  func (x *_state_47) Recv() (int, *_state_48) { if x.next == nil { x.next = init_state_48(x.c) }; return (<-x.c).(int),x.next}

  type _state_46 struct {
    c chan interface{}
    next *_state_47
  }
  
  func init_state_46(c chan interface{}) *_state_46 { return &_state_46{ c, nil } } 
  func (x *_state_46) Send(v int) *_state_47 { if x.next == nil { x.next = init_state_47(x.c) }; x.c <- v; return x.next}
  func (x *_state_46) Recv() (int, *_state_47) { if x.next == nil { x.next = init_state_47(x.c) }; return (<-x.c).(int),x.next}

  type _state_45 struct {
    c chan interface{}
    next *_state_46
  }
  
  func init_state_45(c chan interface{}) *_state_45 { return &_state_45{ c, nil } } 
  func (x *_state_45) Send(v int) *_state_46 { if x.next == nil { x.next = init_state_46(x.c) }; x.c <- v; return x.next}
  func (x *_state_45) Recv() (int, *_state_46) { if x.next == nil { x.next = init_state_46(x.c) }; return (<-x.c).(int),x.next}

  type _state_44 struct {
    c chan interface{}
    next *_state_45
  }
  
  func init_state_44(c chan interface{}) *_state_44 { return &_state_44{ c, nil } } 
  func (x *_state_44) Send(v int) *_state_45 { if x.next == nil { x.next = init_state_45(x.c) }; x.c <- v; return x.next}
  func (x *_state_44) Recv() (int, *_state_45) { if x.next == nil { x.next = init_state_45(x.c) }; return (<-x.c).(int),x.next}

  type _state_43 struct {
    c chan interface{}
    next *_state_44
  }
  
  func init_state_43(c chan interface{}) *_state_43 { return &_state_43{ c, nil } } 
  func (x *_state_43) Send(v int) *_state_44 { if x.next == nil { x.next = init_state_44(x.c) }; x.c <- v; return x.next}
  func (x *_state_43) Recv() (int, *_state_44) { if x.next == nil { x.next = init_state_44(x.c) }; return (<-x.c).(int),x.next}

  type _state_42 struct {
    c chan interface{}
    next *_state_43
  }
  
  func init_state_42(c chan interface{}) *_state_42 { return &_state_42{ c, nil } } 
  func (x *_state_42) Send(v int) *_state_43 { if x.next == nil { x.next = init_state_43(x.c) }; x.c <- v; return x.next}
  func (x *_state_42) Recv() (int, *_state_43) { if x.next == nil { x.next = init_state_43(x.c) }; return (<-x.c).(int),x.next}

  type _state_41 struct {
    c chan interface{}
    next *_state_42
  }
  
  func init_state_41(c chan interface{}) *_state_41 { return &_state_41{ c, nil } } 
  func (x *_state_41) Send(v int) *_state_42 { if x.next == nil { x.next = init_state_42(x.c) }; x.c <- v; return x.next}
  func (x *_state_41) Recv() (int, *_state_42) { if x.next == nil { x.next = init_state_42(x.c) }; return (<-x.c).(int),x.next}

  type _state_40 struct {
    c chan interface{}
    next *_state_41
  }
  
  func init_state_40(c chan interface{}) *_state_40 { return &_state_40{ c, nil } } 
  func (x *_state_40) Send(v int) *_state_41 { if x.next == nil { x.next = init_state_41(x.c) }; x.c <- v; return x.next}
  func (x *_state_40) Recv() (int, *_state_41) { if x.next == nil { x.next = init_state_41(x.c) }; return (<-x.c).(int),x.next}

  type _state_39 struct {
    c chan interface{}
    next *_state_40
  }
  
  func init_state_39(c chan interface{}) *_state_39 { return &_state_39{ c, nil } } 
  func (x *_state_39) Send(v int) *_state_40 { if x.next == nil { x.next = init_state_40(x.c) }; x.c <- v; return x.next}
  func (x *_state_39) Recv() (int, *_state_40) { if x.next == nil { x.next = init_state_40(x.c) }; return (<-x.c).(int),x.next}

  type _state_38 struct {
    c chan interface{}
    next *_state_39
  }
  
  func init_state_38(c chan interface{}) *_state_38 { return &_state_38{ c, nil } } 
  func (x *_state_38) Send(v int) *_state_39 { if x.next == nil { x.next = init_state_39(x.c) }; x.c <- v; return x.next}
  func (x *_state_38) Recv() (int, *_state_39) { if x.next == nil { x.next = init_state_39(x.c) }; return (<-x.c).(int),x.next}

  type _state_37 struct {
    c chan interface{}
    next *_state_38
  }
  
  func init_state_37(c chan interface{}) *_state_37 { return &_state_37{ c, nil } } 
  func (x *_state_37) Send(v int) *_state_38 { if x.next == nil { x.next = init_state_38(x.c) }; x.c <- v; return x.next}
  func (x *_state_37) Recv() (int, *_state_38) { if x.next == nil { x.next = init_state_38(x.c) }; return (<-x.c).(int),x.next}

  type _state_36 struct {
    c chan interface{}
    next *_state_37
  }
  
  func init_state_36(c chan interface{}) *_state_36 { return &_state_36{ c, nil } } 
  func (x *_state_36) Send(v int) *_state_37 { if x.next == nil { x.next = init_state_37(x.c) }; x.c <- v; return x.next}
  func (x *_state_36) Recv() (int, *_state_37) { if x.next == nil { x.next = init_state_37(x.c) }; return (<-x.c).(int),x.next}

  type _state_35 struct {
    c chan interface{}
    next *_state_36
  }
  
  func init_state_35(c chan interface{}) *_state_35 { return &_state_35{ c, nil } } 
  func (x *_state_35) Send(v int) *_state_36 { if x.next == nil { x.next = init_state_36(x.c) }; x.c <- v; return x.next}
  func (x *_state_35) Recv() (int, *_state_36) { if x.next == nil { x.next = init_state_36(x.c) }; return (<-x.c).(int),x.next}

  type _state_34 struct {
    c chan interface{}
    next *_state_35
  }
  
  func init_state_34(c chan interface{}) *_state_34 { return &_state_34{ c, nil } } 
  func (x *_state_34) Send(v int) *_state_35 { if x.next == nil { x.next = init_state_35(x.c) }; x.c <- v; return x.next}
  func (x *_state_34) Recv() (int, *_state_35) { if x.next == nil { x.next = init_state_35(x.c) }; return (<-x.c).(int),x.next}

  type _state_33 struct {
    c chan interface{}
    next *_state_34
  }
  
  func init_state_33(c chan interface{}) *_state_33 { return &_state_33{ c, nil } } 
  func (x *_state_33) Send(v int) *_state_34 { if x.next == nil { x.next = init_state_34(x.c) }; x.c <- v; return x.next}
  func (x *_state_33) Recv() (int, *_state_34) { if x.next == nil { x.next = init_state_34(x.c) }; return (<-x.c).(int),x.next}

  type _state_32 struct {
    c chan interface{}
    next *_state_33
  }
  
  func init_state_32(c chan interface{}) *_state_32 { return &_state_32{ c, nil } } 
  func (x *_state_32) Send(v int) *_state_33 { if x.next == nil { x.next = init_state_33(x.c) }; x.c <- v; return x.next}
  func (x *_state_32) Recv() (int, *_state_33) { if x.next == nil { x.next = init_state_33(x.c) }; return (<-x.c).(int),x.next}

  type _state_31 struct {
    c chan interface{}
    next *_state_32
  }
  
  func init_state_31(c chan interface{}) *_state_31 { return &_state_31{ c, nil } } 
  func (x *_state_31) Send(v int) *_state_32 { if x.next == nil { x.next = init_state_32(x.c) }; x.c <- v; return x.next}
  func (x *_state_31) Recv() (int, *_state_32) { if x.next == nil { x.next = init_state_32(x.c) }; return (<-x.c).(int),x.next}

  type _state_30 struct {
    c chan interface{}
    next *_state_31
  }
  
  func init_state_30(c chan interface{}) *_state_30 { return &_state_30{ c, nil } } 
  func (x *_state_30) Send(v int) *_state_31 { if x.next == nil { x.next = init_state_31(x.c) }; x.c <- v; return x.next}
  func (x *_state_30) Recv() (int, *_state_31) { if x.next == nil { x.next = init_state_31(x.c) }; return (<-x.c).(int),x.next}

  type _state_29 struct {
    c chan interface{}
    next *_state_30
  }
  
  func init_state_29(c chan interface{}) *_state_29 { return &_state_29{ c, nil } } 
  func (x *_state_29) Send(v int) *_state_30 { if x.next == nil { x.next = init_state_30(x.c) }; x.c <- v; return x.next}
  func (x *_state_29) Recv() (int, *_state_30) { if x.next == nil { x.next = init_state_30(x.c) }; return (<-x.c).(int),x.next}

  type _state_28 struct {
    c chan interface{}
    next *_state_29
  }
  
  func init_state_28(c chan interface{}) *_state_28 { return &_state_28{ c, nil } } 
  func (x *_state_28) Send(v int) *_state_29 { if x.next == nil { x.next = init_state_29(x.c) }; x.c <- v; return x.next}
  func (x *_state_28) Recv() (int, *_state_29) { if x.next == nil { x.next = init_state_29(x.c) }; return (<-x.c).(int),x.next}

  type _state_27 struct {
    c chan interface{}
    next *_state_28
  }
  
  func init_state_27(c chan interface{}) *_state_27 { return &_state_27{ c, nil } } 
  func (x *_state_27) Send(v int) *_state_28 { if x.next == nil { x.next = init_state_28(x.c) }; x.c <- v; return x.next}
  func (x *_state_27) Recv() (int, *_state_28) { if x.next == nil { x.next = init_state_28(x.c) }; return (<-x.c).(int),x.next}

  type _state_26 struct {
    c chan interface{}
    next *_state_27
  }
  
  func init_state_26(c chan interface{}) *_state_26 { return &_state_26{ c, nil } } 
  func (x *_state_26) Send(v int) *_state_27 { if x.next == nil { x.next = init_state_27(x.c) }; x.c <- v; return x.next}
  func (x *_state_26) Recv() (int, *_state_27) { if x.next == nil { x.next = init_state_27(x.c) }; return (<-x.c).(int),x.next}

  type _state_25 struct {
    c chan interface{}
    next *_state_26
  }
  
  func init_state_25(c chan interface{}) *_state_25 { return &_state_25{ c, nil } } 
  func (x *_state_25) Send(v int) *_state_26 { if x.next == nil { x.next = init_state_26(x.c) }; x.c <- v; return x.next}
  func (x *_state_25) Recv() (int, *_state_26) { if x.next == nil { x.next = init_state_26(x.c) }; return (<-x.c).(int),x.next}

  type _state_24 struct {
    c chan interface{}
    next *_state_25
  }
  
  func init_state_24(c chan interface{}) *_state_24 { return &_state_24{ c, nil } } 
  func (x *_state_24) Send(v int) *_state_25 { if x.next == nil { x.next = init_state_25(x.c) }; x.c <- v; return x.next}
  func (x *_state_24) Recv() (int, *_state_25) { if x.next == nil { x.next = init_state_25(x.c) }; return (<-x.c).(int),x.next}

  type _state_23 struct {
    c chan interface{}
    next *_state_24
  }
  
  func init_state_23(c chan interface{}) *_state_23 { return &_state_23{ c, nil } } 
  func (x *_state_23) Send(v int) *_state_24 { if x.next == nil { x.next = init_state_24(x.c) }; x.c <- v; return x.next}
  func (x *_state_23) Recv() (int, *_state_24) { if x.next == nil { x.next = init_state_24(x.c) }; return (<-x.c).(int),x.next}

  type _state_22 struct {
    c chan interface{}
    next *_state_23
  }
  
  func init_state_22(c chan interface{}) *_state_22 { return &_state_22{ c, nil } } 
  func (x *_state_22) Send(v int) *_state_23 { if x.next == nil { x.next = init_state_23(x.c) }; x.c <- v; return x.next}
  func (x *_state_22) Recv() (int, *_state_23) { if x.next == nil { x.next = init_state_23(x.c) }; return (<-x.c).(int),x.next}

  type _state_21 struct {
    c chan interface{}
    next *_state_22
  }
  
  func init_state_21(c chan interface{}) *_state_21 { return &_state_21{ c, nil } } 
  func (x *_state_21) Send(v int) *_state_22 { if x.next == nil { x.next = init_state_22(x.c) }; x.c <- v; return x.next}
  func (x *_state_21) Recv() (int, *_state_22) { if x.next == nil { x.next = init_state_22(x.c) }; return (<-x.c).(int),x.next}

  type _state_20 struct {
    c chan interface{}
    next *_state_21
  }
  
  func init_state_20(c chan interface{}) *_state_20 { return &_state_20{ c, nil } } 
  func (x *_state_20) Send(v int) *_state_21 { if x.next == nil { x.next = init_state_21(x.c) }; x.c <- v; return x.next}
  func (x *_state_20) Recv() (int, *_state_21) { if x.next == nil { x.next = init_state_21(x.c) }; return (<-x.c).(int),x.next}

  type _state_19 struct {
    c chan interface{}
    next *_state_20
  }
  
  func init_state_19(c chan interface{}) *_state_19 { return &_state_19{ c, nil } } 
  func (x *_state_19) Send(v int) *_state_20 { if x.next == nil { x.next = init_state_20(x.c) }; x.c <- v; return x.next}
  func (x *_state_19) Recv() (int, *_state_20) { if x.next == nil { x.next = init_state_20(x.c) }; return (<-x.c).(int),x.next}

  type _state_18 struct {
    c chan interface{}
    next *_state_19
  }
  
  func init_state_18(c chan interface{}) *_state_18 { return &_state_18{ c, nil } } 
  func (x *_state_18) Send(v int) *_state_19 { if x.next == nil { x.next = init_state_19(x.c) }; x.c <- v; return x.next}
  func (x *_state_18) Recv() (int, *_state_19) { if x.next == nil { x.next = init_state_19(x.c) }; return (<-x.c).(int),x.next}

  type _state_17 struct {
    c chan interface{}
    next *_state_18
  }
  
  func init_state_17(c chan interface{}) *_state_17 { return &_state_17{ c, nil } } 
  func (x *_state_17) Send(v int) *_state_18 { if x.next == nil { x.next = init_state_18(x.c) }; x.c <- v; return x.next}
  func (x *_state_17) Recv() (int, *_state_18) { if x.next == nil { x.next = init_state_18(x.c) }; return (<-x.c).(int),x.next}

  type _state_16 struct {
    c chan interface{}
    next *_state_17
  }
  
  func init_state_16(c chan interface{}) *_state_16 { return &_state_16{ c, nil } } 
  func (x *_state_16) Send(v int) *_state_17 { if x.next == nil { x.next = init_state_17(x.c) }; x.c <- v; return x.next}
  func (x *_state_16) Recv() (int, *_state_17) { if x.next == nil { x.next = init_state_17(x.c) }; return (<-x.c).(int),x.next}

  type _state_15 struct {
    c chan interface{}
    next *_state_16
  }
  
  func init_state_15(c chan interface{}) *_state_15 { return &_state_15{ c, nil } } 
  func (x *_state_15) Send(v int) *_state_16 { if x.next == nil { x.next = init_state_16(x.c) }; x.c <- v; return x.next}
  func (x *_state_15) Recv() (int, *_state_16) { if x.next == nil { x.next = init_state_16(x.c) }; return (<-x.c).(int),x.next}

  type _state_14 struct {
    c chan interface{}
    next *_state_15
  }
  
  func init_state_14(c chan interface{}) *_state_14 { return &_state_14{ c, nil } } 
  func (x *_state_14) Send(v int) *_state_15 { if x.next == nil { x.next = init_state_15(x.c) }; x.c <- v; return x.next}
  func (x *_state_14) Recv() (int, *_state_15) { if x.next == nil { x.next = init_state_15(x.c) }; return (<-x.c).(int),x.next}

  type _state_13 struct {
    c chan interface{}
    next *_state_14
  }
  
  func init_state_13(c chan interface{}) *_state_13 { return &_state_13{ c, nil } } 
  func (x *_state_13) Send(v int) *_state_14 { if x.next == nil { x.next = init_state_14(x.c) }; x.c <- v; return x.next}
  func (x *_state_13) Recv() (int, *_state_14) { if x.next == nil { x.next = init_state_14(x.c) }; return (<-x.c).(int),x.next}

  type _state_12 struct {
    c chan interface{}
    next *_state_13
  }
  
  func init_state_12(c chan interface{}) *_state_12 { return &_state_12{ c, nil } } 
  func (x *_state_12) Send(v int) *_state_13 { if x.next == nil { x.next = init_state_13(x.c) }; x.c <- v; return x.next}
  func (x *_state_12) Recv() (int, *_state_13) { if x.next == nil { x.next = init_state_13(x.c) }; return (<-x.c).(int),x.next}

  type _state_11 struct {
    c chan interface{}
    next *_state_12
  }
  
  func init_state_11(c chan interface{}) *_state_11 { return &_state_11{ c, nil } } 
  func (x *_state_11) Send(v int) *_state_12 { if x.next == nil { x.next = init_state_12(x.c) }; x.c <- v; return x.next}
  func (x *_state_11) Recv() (int, *_state_12) { if x.next == nil { x.next = init_state_12(x.c) }; return (<-x.c).(int),x.next}

  type _state_10 struct {
    c chan interface{}
    next *_state_11
  }
  
  func init_state_10(c chan interface{}) *_state_10 { return &_state_10{ c, nil } } 
  func (x *_state_10) Send(v int) *_state_11 { if x.next == nil { x.next = init_state_11(x.c) }; x.c <- v; return x.next}
  func (x *_state_10) Recv() (int, *_state_11) { if x.next == nil { x.next = init_state_11(x.c) }; return (<-x.c).(int),x.next}

  type _state_9 struct {
    c chan interface{}
    next *_state_10
  }
  
  func init_state_9(c chan interface{}) *_state_9 { return &_state_9{ c, nil } } 
  func (x *_state_9) Send(v int) *_state_10 { if x.next == nil { x.next = init_state_10(x.c) }; x.c <- v; return x.next}
  func (x *_state_9) Recv() (int, *_state_10) { if x.next == nil { x.next = init_state_10(x.c) }; return (<-x.c).(int),x.next}

  type _state_8 struct {
    c chan interface{}
    next *_state_9
  }
  
  func init_state_8(c chan interface{}) *_state_8 { return &_state_8{ c, nil } } 
  func (x *_state_8) Send(v int) *_state_9 { if x.next == nil { x.next = init_state_9(x.c) }; x.c <- v; return x.next}
  func (x *_state_8) Recv() (int, *_state_9) { if x.next == nil { x.next = init_state_9(x.c) }; return (<-x.c).(int),x.next}

  type _state_7 struct {
    c chan interface{}
    next *_state_8
  }
  
  func init_state_7(c chan interface{}) *_state_7 { return &_state_7{ c, nil } } 
  func (x *_state_7) Send(v int) *_state_8 { if x.next == nil { x.next = init_state_8(x.c) }; x.c <- v; return x.next}
  func (x *_state_7) Recv() (int, *_state_8) { if x.next == nil { x.next = init_state_8(x.c) }; return (<-x.c).(int),x.next}

  type _state_6 struct {
    c chan interface{}
    next *_state_7
  }
  
  func init_state_6(c chan interface{}) *_state_6 { return &_state_6{ c, nil } } 
  func (x *_state_6) Send(v int) *_state_7 { if x.next == nil { x.next = init_state_7(x.c) }; x.c <- v; return x.next}
  func (x *_state_6) Recv() (int, *_state_7) { if x.next == nil { x.next = init_state_7(x.c) }; return (<-x.c).(int),x.next}

  type _state_5 struct {
    c chan interface{}
    next *_state_6
  }
  
  func init_state_5(c chan interface{}) *_state_5 { return &_state_5{ c, nil } } 
  func (x *_state_5) Send(v int) *_state_6 { if x.next == nil { x.next = init_state_6(x.c) }; x.c <- v; return x.next}
  func (x *_state_5) Recv() (int, *_state_6) { if x.next == nil { x.next = init_state_6(x.c) }; return (<-x.c).(int),x.next}

  type _state_4 struct {
    c chan interface{}
    next *_state_5
  }
  
  func init_state_4(c chan interface{}) *_state_4 { return &_state_4{ c, nil } } 
  func (x *_state_4) Send(v int) *_state_5 { if x.next == nil { x.next = init_state_5(x.c) }; x.c <- v; return x.next}
  func (x *_state_4) Recv() (int, *_state_5) { if x.next == nil { x.next = init_state_5(x.c) }; return (<-x.c).(int),x.next}

  type _state_3 struct {
    c chan interface{}
    next *_state_4
  }
  
  func init_state_3(c chan interface{}) *_state_3 { return &_state_3{ c, nil } } 
  func (x *_state_3) Send(v int) *_state_4 { if x.next == nil { x.next = init_state_4(x.c) }; x.c <- v; return x.next}
  func (x *_state_3) Recv() (int, *_state_4) { if x.next == nil { x.next = init_state_4(x.c) }; return (<-x.c).(int),x.next}

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

  func send_ints(n int) func (_x *_state_0) {
 return func (c *_state_0){
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
c32 := c31.Send((n + 33))
c33 := c32.Send((n + 34))
c34 := c33.Send((n + 35))
c35 := c34.Send((n + 36))
c36 := c35.Send((n + 37))
c37 := c36.Send((n + 38))
c38 := c37.Send((n + 39))
c39 := c38.Send((n + 40))
c40 := c39.Send((n + 41))
c41 := c40.Send((n + 42))
c42 := c41.Send((n + 43))
c43 := c42.Send((n + 44))
c44 := c43.Send((n + 45))
c45 := c44.Send((n + 46))
c46 := c45.Send((n + 47))
c47 := c46.Send((n + 48))
c48 := c47.Send((n + 49))
c49 := c48.Send((n + 50))
c50 := c49.Send((n + 51))
c51 := c50.Send((n + 52))
c52 := c51.Send((n + 53))
c53 := c52.Send((n + 54))
c54 := c53.Send((n + 55))
c55 := c54.Send((n + 56))
c56 := c55.Send((n + 57))
c57 := c56.Send((n + 58))
c58 := c57.Send((n + 59))
c59 := c58.Send((n + 60))
c60 := c59.Send((n + 61))
c61 := c60.Send((n + 62))
c62 := c61.Send((n + 63))
c63 := c62.Send((n + 64))
c63.Send(nil)
}}
func main () {
m := init_state_64(make (chan interface{}))
go func () {
m.Recv()
}()
func (m *_state_64){
e := init_state_0(make(chan interface{}))
go send_ints(1)(e)
a1, e0 := e.Recv()
fmt.Printf("%v\n",a1)
a2, e1 := e0.Recv()
fmt.Printf("%v\n",a2)
a3, e2 := e1.Recv()
fmt.Printf("%v\n",a3)
a4, e3 := e2.Recv()
fmt.Printf("%v\n",a4)
a5, e4 := e3.Recv()
fmt.Printf("%v\n",a5)
a6, e5 := e4.Recv()
fmt.Printf("%v\n",a6)
a7, e6 := e5.Recv()
fmt.Printf("%v\n",a7)
a8, e7 := e6.Recv()
fmt.Printf("%v\n",a8)
a9, e8 := e7.Recv()
fmt.Printf("%v\n",a9)
a10, e9 := e8.Recv()
fmt.Printf("%v\n",a10)
a11, e10 := e9.Recv()
fmt.Printf("%v\n",a11)
a12, e11 := e10.Recv()
fmt.Printf("%v\n",a12)
a13, e12 := e11.Recv()
fmt.Printf("%v\n",a13)
a14, e13 := e12.Recv()
fmt.Printf("%v\n",a14)
a15, e14 := e13.Recv()
fmt.Printf("%v\n",a15)
a16, e15 := e14.Recv()
fmt.Printf("%v\n",a16)
a17, e16 := e15.Recv()
fmt.Printf("%v\n",a17)
a18, e17 := e16.Recv()
fmt.Printf("%v\n",a18)
a19, e18 := e17.Recv()
fmt.Printf("%v\n",a19)
a20, e19 := e18.Recv()
fmt.Printf("%v\n",a20)
a21, e20 := e19.Recv()
fmt.Printf("%v\n",a21)
a22, e21 := e20.Recv()
fmt.Printf("%v\n",a22)
a23, e22 := e21.Recv()
fmt.Printf("%v\n",a23)
a24, e23 := e22.Recv()
fmt.Printf("%v\n",a24)
a25, e24 := e23.Recv()
fmt.Printf("%v\n",a25)
a26, e25 := e24.Recv()
fmt.Printf("%v\n",a26)
a27, e26 := e25.Recv()
fmt.Printf("%v\n",a27)
a28, e27 := e26.Recv()
fmt.Printf("%v\n",a28)
a29, e28 := e27.Recv()
fmt.Printf("%v\n",a29)
a30, e29 := e28.Recv()
fmt.Printf("%v\n",a30)
a31, e30 := e29.Recv()
fmt.Printf("%v\n",a31)
a32, e31 := e30.Recv()
fmt.Printf("%v\n",a32)
a33, e32 := e31.Recv()
fmt.Printf("%v\n",a33)
a34, e33 := e32.Recv()
fmt.Printf("%v\n",a34)
a35, e34 := e33.Recv()
fmt.Printf("%v\n",a35)
a36, e35 := e34.Recv()
fmt.Printf("%v\n",a36)
a37, e36 := e35.Recv()
fmt.Printf("%v\n",a37)
a38, e37 := e36.Recv()
fmt.Printf("%v\n",a38)
a39, e38 := e37.Recv()
fmt.Printf("%v\n",a39)
a40, e39 := e38.Recv()
fmt.Printf("%v\n",a40)
a41, e40 := e39.Recv()
fmt.Printf("%v\n",a41)
a42, e41 := e40.Recv()
fmt.Printf("%v\n",a42)
a43, e42 := e41.Recv()
fmt.Printf("%v\n",a43)
a44, e43 := e42.Recv()
fmt.Printf("%v\n",a44)
a45, e44 := e43.Recv()
fmt.Printf("%v\n",a45)
a46, e45 := e44.Recv()
fmt.Printf("%v\n",a46)
a47, e46 := e45.Recv()
fmt.Printf("%v\n",a47)
a48, e47 := e46.Recv()
fmt.Printf("%v\n",a48)
a49, e48 := e47.Recv()
fmt.Printf("%v\n",a49)
a50, e49 := e48.Recv()
fmt.Printf("%v\n",a50)
a51, e50 := e49.Recv()
fmt.Printf("%v\n",a51)
a52, e51 := e50.Recv()
fmt.Printf("%v\n",a52)
a53, e52 := e51.Recv()
fmt.Printf("%v\n",a53)
a54, e53 := e52.Recv()
fmt.Printf("%v\n",a54)
a55, e54 := e53.Recv()
fmt.Printf("%v\n",a55)
a56, e55 := e54.Recv()
fmt.Printf("%v\n",a56)
a57, e56 := e55.Recv()
fmt.Printf("%v\n",a57)
a58, e57 := e56.Recv()
fmt.Printf("%v\n",a58)
a59, e58 := e57.Recv()
fmt.Printf("%v\n",a59)
a60, e59 := e58.Recv()
fmt.Printf("%v\n",a60)
a61, e60 := e59.Recv()
fmt.Printf("%v\n",a61)
a62, e61 := e60.Recv()
fmt.Printf("%v\n",a62)
a63, e62 := e61.Recv()
fmt.Printf("%v\n",a63)
a64, e63 := e62.Recv()
fmt.Printf("%v\n",a64)
e63.Recv()
m.Send(nil)
}(m)
}
