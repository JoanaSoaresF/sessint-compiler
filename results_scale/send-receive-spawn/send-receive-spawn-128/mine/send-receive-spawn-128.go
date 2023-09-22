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
func (x *_state_0) Send(v0 int, v1 int, v2 int, v3 int, v4 int, v5 int, v6 int, v7 int, v8 int, v9 int, v10 int, v11 int, v12 int, v13 int, v14 int, v15 int, v16 int, v17 int, v18 int, v19 int, v20 int, v21 int, v22 int, v23 int, v24 int, v25 int, v26 int, v27 int, v28 int, v29 int, v30 int, v31 int, v32 int, v33 int, v34 int, v35 int, v36 int, v37 int, v38 int, v39 int, v40 int, v41 int, v42 int, v43 int, v44 int, v45 int, v46 int, v47 int, v48 int, v49 int, v50 int, v51 int, v52 int, v53 int, v54 int, v55 int, v56 int, v57 int, v58 int, v59 int, v60 int, v61 int, v62 int, v63 int, v64 int, v65 int, v66 int, v67 int, v68 int, v69 int, v70 int, v71 int, v72 int, v73 int, v74 int, v75 int, v76 int, v77 int, v78 int, v79 int, v80 int, v81 int, v82 int, v83 int, v84 int, v85 int, v86 int, v87 int, v88 int, v89 int, v90 int, v91 int, v92 int, v93 int, v94 int, v95 int, v96 int, v97 int, v98 int, v99 int, v100 int, v101 int, v102 int, v103 int, v104 int, v105 int, v106 int, v107 int, v108 int, v109 int, v110 int, v111 int, v112 int, v113 int, v114 int, v115 int, v116 int, v117 int, v118 int, v119 int, v120 int, v121 int, v122 int, v123 int, v124 int, v125 int, v126 int, v127 int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) };
 x.c <- []interface{}{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v90, v91, v92, v93, v94, v95, v96, v97, v98, v99, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v120, v121, v122, v123, v124, v125, v126, v127}
return x.next }
func (x *_state_0) Recv() (int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) };ll := <- x.c
l := ll.([]interface{})
return l[0].(int), l[1].(int), l[2].(int), l[3].(int), l[4].(int), l[5].(int), l[6].(int), l[7].(int), l[8].(int), l[9].(int), l[10].(int), l[11].(int), l[12].(int), l[13].(int), l[14].(int), l[15].(int), l[16].(int), l[17].(int), l[18].(int), l[19].(int), l[20].(int), l[21].(int), l[22].(int), l[23].(int), l[24].(int), l[25].(int), l[26].(int), l[27].(int), l[28].(int), l[29].(int), l[30].(int), l[31].(int), l[32].(int), l[33].(int), l[34].(int), l[35].(int), l[36].(int), l[37].(int), l[38].(int), l[39].(int), l[40].(int), l[41].(int), l[42].(int), l[43].(int), l[44].(int), l[45].(int), l[46].(int), l[47].(int), l[48].(int), l[49].(int), l[50].(int), l[51].(int), l[52].(int), l[53].(int), l[54].(int), l[55].(int), l[56].(int), l[57].(int), l[58].(int), l[59].(int), l[60].(int), l[61].(int), l[62].(int), l[63].(int), l[64].(int), l[65].(int), l[66].(int), l[67].(int), l[68].(int), l[69].(int), l[70].(int), l[71].(int), l[72].(int), l[73].(int), l[74].(int), l[75].(int), l[76].(int), l[77].(int), l[78].(int), l[79].(int), l[80].(int), l[81].(int), l[82].(int), l[83].(int), l[84].(int), l[85].(int), l[86].(int), l[87].(int), l[88].(int), l[89].(int), l[90].(int), l[91].(int), l[92].(int), l[93].(int), l[94].(int), l[95].(int), l[96].(int), l[97].(int), l[98].(int), l[99].(int), l[100].(int), l[101].(int), l[102].(int), l[103].(int), l[104].(int), l[105].(int), l[106].(int), l[107].(int), l[108].(int), l[109].(int), l[110].(int), l[111].(int), l[112].(int), l[113].(int), l[114].(int), l[115].(int), l[116].(int), l[117].(int), l[118].(int), l[119].(int), l[120].(int), l[121].(int), l[122].(int), l[123].(int), l[124].(int), l[125].(int), l[126].(int), l[127].(int), x.next }

type _state_129 struct {
   c chan interface{}
   next *_state_1
}

func init_state_129(c chan interface{}) *_state_129 { return &_state_129{c, nil} }
func (x *_state_129) Send(v int) *_state_1 {
   if x.next == nil { x.next = init_state_1(x.c) }; x.c <- v; return x.next }
func (x *_state_129) Recv() (int, *_state_1) {
   if x.next == nil { x.next = init_state_1(x.c) }; return (<-x.c).(int), x.next }

  type _state_128 struct {
   c chan interface{}
   next *_state_129
}

func init_state_128(c chan interface{}) *_state_128 { return &_state_128{c, nil} }
func (x *_state_128) Send(v int) *_state_129 {
   if x.next == nil { x.next = init_state_129(x.c) }; x.c <- v; return x.next }
func (x *_state_128) Recv() (int, *_state_129) {
   if x.next == nil { x.next = init_state_129(x.c) }; return (<-x.c).(int), x.next }

  type _state_127 struct {
   c chan interface{}
   next *_state_128
}

func init_state_127(c chan interface{}) *_state_127 { return &_state_127{c, nil} }
func (x *_state_127) Send(v int) *_state_128 {
   if x.next == nil { x.next = init_state_128(x.c) }; x.c <- v; return x.next }
func (x *_state_127) Recv() (int, *_state_128) {
   if x.next == nil { x.next = init_state_128(x.c) }; return (<-x.c).(int), x.next }

  type _state_126 struct {
   c chan interface{}
   next *_state_127
}

func init_state_126(c chan interface{}) *_state_126 { return &_state_126{c, nil} }
func (x *_state_126) Send(v int) *_state_127 {
   if x.next == nil { x.next = init_state_127(x.c) }; x.c <- v; return x.next }
func (x *_state_126) Recv() (int, *_state_127) {
   if x.next == nil { x.next = init_state_127(x.c) }; return (<-x.c).(int), x.next }

  type _state_125 struct {
   c chan interface{}
   next *_state_126
}

func init_state_125(c chan interface{}) *_state_125 { return &_state_125{c, nil} }
func (x *_state_125) Send(v int) *_state_126 {
   if x.next == nil { x.next = init_state_126(x.c) }; x.c <- v; return x.next }
func (x *_state_125) Recv() (int, *_state_126) {
   if x.next == nil { x.next = init_state_126(x.c) }; return (<-x.c).(int), x.next }

  type _state_124 struct {
   c chan interface{}
   next *_state_125
}

func init_state_124(c chan interface{}) *_state_124 { return &_state_124{c, nil} }
func (x *_state_124) Send(v int) *_state_125 {
   if x.next == nil { x.next = init_state_125(x.c) }; x.c <- v; return x.next }
func (x *_state_124) Recv() (int, *_state_125) {
   if x.next == nil { x.next = init_state_125(x.c) }; return (<-x.c).(int), x.next }

  type _state_123 struct {
   c chan interface{}
   next *_state_124
}

func init_state_123(c chan interface{}) *_state_123 { return &_state_123{c, nil} }
func (x *_state_123) Send(v int) *_state_124 {
   if x.next == nil { x.next = init_state_124(x.c) }; x.c <- v; return x.next }
func (x *_state_123) Recv() (int, *_state_124) {
   if x.next == nil { x.next = init_state_124(x.c) }; return (<-x.c).(int), x.next }

  type _state_122 struct {
   c chan interface{}
   next *_state_123
}

func init_state_122(c chan interface{}) *_state_122 { return &_state_122{c, nil} }
func (x *_state_122) Send(v int) *_state_123 {
   if x.next == nil { x.next = init_state_123(x.c) }; x.c <- v; return x.next }
func (x *_state_122) Recv() (int, *_state_123) {
   if x.next == nil { x.next = init_state_123(x.c) }; return (<-x.c).(int), x.next }

  type _state_121 struct {
   c chan interface{}
   next *_state_122
}

func init_state_121(c chan interface{}) *_state_121 { return &_state_121{c, nil} }
func (x *_state_121) Send(v int) *_state_122 {
   if x.next == nil { x.next = init_state_122(x.c) }; x.c <- v; return x.next }
func (x *_state_121) Recv() (int, *_state_122) {
   if x.next == nil { x.next = init_state_122(x.c) }; return (<-x.c).(int), x.next }

  type _state_120 struct {
   c chan interface{}
   next *_state_121
}

func init_state_120(c chan interface{}) *_state_120 { return &_state_120{c, nil} }
func (x *_state_120) Send(v int) *_state_121 {
   if x.next == nil { x.next = init_state_121(x.c) }; x.c <- v; return x.next }
func (x *_state_120) Recv() (int, *_state_121) {
   if x.next == nil { x.next = init_state_121(x.c) }; return (<-x.c).(int), x.next }

  type _state_119 struct {
   c chan interface{}
   next *_state_120
}

func init_state_119(c chan interface{}) *_state_119 { return &_state_119{c, nil} }
func (x *_state_119) Send(v int) *_state_120 {
   if x.next == nil { x.next = init_state_120(x.c) }; x.c <- v; return x.next }
func (x *_state_119) Recv() (int, *_state_120) {
   if x.next == nil { x.next = init_state_120(x.c) }; return (<-x.c).(int), x.next }

  type _state_118 struct {
   c chan interface{}
   next *_state_119
}

func init_state_118(c chan interface{}) *_state_118 { return &_state_118{c, nil} }
func (x *_state_118) Send(v int) *_state_119 {
   if x.next == nil { x.next = init_state_119(x.c) }; x.c <- v; return x.next }
func (x *_state_118) Recv() (int, *_state_119) {
   if x.next == nil { x.next = init_state_119(x.c) }; return (<-x.c).(int), x.next }

  type _state_117 struct {
   c chan interface{}
   next *_state_118
}

func init_state_117(c chan interface{}) *_state_117 { return &_state_117{c, nil} }
func (x *_state_117) Send(v int) *_state_118 {
   if x.next == nil { x.next = init_state_118(x.c) }; x.c <- v; return x.next }
func (x *_state_117) Recv() (int, *_state_118) {
   if x.next == nil { x.next = init_state_118(x.c) }; return (<-x.c).(int), x.next }

  type _state_116 struct {
   c chan interface{}
   next *_state_117
}

func init_state_116(c chan interface{}) *_state_116 { return &_state_116{c, nil} }
func (x *_state_116) Send(v int) *_state_117 {
   if x.next == nil { x.next = init_state_117(x.c) }; x.c <- v; return x.next }
func (x *_state_116) Recv() (int, *_state_117) {
   if x.next == nil { x.next = init_state_117(x.c) }; return (<-x.c).(int), x.next }

  type _state_115 struct {
   c chan interface{}
   next *_state_116
}

func init_state_115(c chan interface{}) *_state_115 { return &_state_115{c, nil} }
func (x *_state_115) Send(v int) *_state_116 {
   if x.next == nil { x.next = init_state_116(x.c) }; x.c <- v; return x.next }
func (x *_state_115) Recv() (int, *_state_116) {
   if x.next == nil { x.next = init_state_116(x.c) }; return (<-x.c).(int), x.next }

  type _state_114 struct {
   c chan interface{}
   next *_state_115
}

func init_state_114(c chan interface{}) *_state_114 { return &_state_114{c, nil} }
func (x *_state_114) Send(v int) *_state_115 {
   if x.next == nil { x.next = init_state_115(x.c) }; x.c <- v; return x.next }
func (x *_state_114) Recv() (int, *_state_115) {
   if x.next == nil { x.next = init_state_115(x.c) }; return (<-x.c).(int), x.next }

  type _state_113 struct {
   c chan interface{}
   next *_state_114
}

func init_state_113(c chan interface{}) *_state_113 { return &_state_113{c, nil} }
func (x *_state_113) Send(v int) *_state_114 {
   if x.next == nil { x.next = init_state_114(x.c) }; x.c <- v; return x.next }
func (x *_state_113) Recv() (int, *_state_114) {
   if x.next == nil { x.next = init_state_114(x.c) }; return (<-x.c).(int), x.next }

  type _state_112 struct {
   c chan interface{}
   next *_state_113
}

func init_state_112(c chan interface{}) *_state_112 { return &_state_112{c, nil} }
func (x *_state_112) Send(v int) *_state_113 {
   if x.next == nil { x.next = init_state_113(x.c) }; x.c <- v; return x.next }
func (x *_state_112) Recv() (int, *_state_113) {
   if x.next == nil { x.next = init_state_113(x.c) }; return (<-x.c).(int), x.next }

  type _state_111 struct {
   c chan interface{}
   next *_state_112
}

func init_state_111(c chan interface{}) *_state_111 { return &_state_111{c, nil} }
func (x *_state_111) Send(v int) *_state_112 {
   if x.next == nil { x.next = init_state_112(x.c) }; x.c <- v; return x.next }
func (x *_state_111) Recv() (int, *_state_112) {
   if x.next == nil { x.next = init_state_112(x.c) }; return (<-x.c).(int), x.next }

  type _state_110 struct {
   c chan interface{}
   next *_state_111
}

func init_state_110(c chan interface{}) *_state_110 { return &_state_110{c, nil} }
func (x *_state_110) Send(v int) *_state_111 {
   if x.next == nil { x.next = init_state_111(x.c) }; x.c <- v; return x.next }
func (x *_state_110) Recv() (int, *_state_111) {
   if x.next == nil { x.next = init_state_111(x.c) }; return (<-x.c).(int), x.next }

  type _state_109 struct {
   c chan interface{}
   next *_state_110
}

func init_state_109(c chan interface{}) *_state_109 { return &_state_109{c, nil} }
func (x *_state_109) Send(v int) *_state_110 {
   if x.next == nil { x.next = init_state_110(x.c) }; x.c <- v; return x.next }
func (x *_state_109) Recv() (int, *_state_110) {
   if x.next == nil { x.next = init_state_110(x.c) }; return (<-x.c).(int), x.next }

  type _state_108 struct {
   c chan interface{}
   next *_state_109
}

func init_state_108(c chan interface{}) *_state_108 { return &_state_108{c, nil} }
func (x *_state_108) Send(v int) *_state_109 {
   if x.next == nil { x.next = init_state_109(x.c) }; x.c <- v; return x.next }
func (x *_state_108) Recv() (int, *_state_109) {
   if x.next == nil { x.next = init_state_109(x.c) }; return (<-x.c).(int), x.next }

  type _state_107 struct {
   c chan interface{}
   next *_state_108
}

func init_state_107(c chan interface{}) *_state_107 { return &_state_107{c, nil} }
func (x *_state_107) Send(v int) *_state_108 {
   if x.next == nil { x.next = init_state_108(x.c) }; x.c <- v; return x.next }
func (x *_state_107) Recv() (int, *_state_108) {
   if x.next == nil { x.next = init_state_108(x.c) }; return (<-x.c).(int), x.next }

  type _state_106 struct {
   c chan interface{}
   next *_state_107
}

func init_state_106(c chan interface{}) *_state_106 { return &_state_106{c, nil} }
func (x *_state_106) Send(v int) *_state_107 {
   if x.next == nil { x.next = init_state_107(x.c) }; x.c <- v; return x.next }
func (x *_state_106) Recv() (int, *_state_107) {
   if x.next == nil { x.next = init_state_107(x.c) }; return (<-x.c).(int), x.next }

  type _state_105 struct {
   c chan interface{}
   next *_state_106
}

func init_state_105(c chan interface{}) *_state_105 { return &_state_105{c, nil} }
func (x *_state_105) Send(v int) *_state_106 {
   if x.next == nil { x.next = init_state_106(x.c) }; x.c <- v; return x.next }
func (x *_state_105) Recv() (int, *_state_106) {
   if x.next == nil { x.next = init_state_106(x.c) }; return (<-x.c).(int), x.next }

  type _state_104 struct {
   c chan interface{}
   next *_state_105
}

func init_state_104(c chan interface{}) *_state_104 { return &_state_104{c, nil} }
func (x *_state_104) Send(v int) *_state_105 {
   if x.next == nil { x.next = init_state_105(x.c) }; x.c <- v; return x.next }
func (x *_state_104) Recv() (int, *_state_105) {
   if x.next == nil { x.next = init_state_105(x.c) }; return (<-x.c).(int), x.next }

  type _state_103 struct {
   c chan interface{}
   next *_state_104
}

func init_state_103(c chan interface{}) *_state_103 { return &_state_103{c, nil} }
func (x *_state_103) Send(v int) *_state_104 {
   if x.next == nil { x.next = init_state_104(x.c) }; x.c <- v; return x.next }
func (x *_state_103) Recv() (int, *_state_104) {
   if x.next == nil { x.next = init_state_104(x.c) }; return (<-x.c).(int), x.next }

  type _state_102 struct {
   c chan interface{}
   next *_state_103
}

func init_state_102(c chan interface{}) *_state_102 { return &_state_102{c, nil} }
func (x *_state_102) Send(v int) *_state_103 {
   if x.next == nil { x.next = init_state_103(x.c) }; x.c <- v; return x.next }
func (x *_state_102) Recv() (int, *_state_103) {
   if x.next == nil { x.next = init_state_103(x.c) }; return (<-x.c).(int), x.next }

  type _state_101 struct {
   c chan interface{}
   next *_state_102
}

func init_state_101(c chan interface{}) *_state_101 { return &_state_101{c, nil} }
func (x *_state_101) Send(v int) *_state_102 {
   if x.next == nil { x.next = init_state_102(x.c) }; x.c <- v; return x.next }
func (x *_state_101) Recv() (int, *_state_102) {
   if x.next == nil { x.next = init_state_102(x.c) }; return (<-x.c).(int), x.next }

  type _state_100 struct {
   c chan interface{}
   next *_state_101
}

func init_state_100(c chan interface{}) *_state_100 { return &_state_100{c, nil} }
func (x *_state_100) Send(v int) *_state_101 {
   if x.next == nil { x.next = init_state_101(x.c) }; x.c <- v; return x.next }
func (x *_state_100) Recv() (int, *_state_101) {
   if x.next == nil { x.next = init_state_101(x.c) }; return (<-x.c).(int), x.next }

  type _state_99 struct {
   c chan interface{}
   next *_state_100
}

func init_state_99(c chan interface{}) *_state_99 { return &_state_99{c, nil} }
func (x *_state_99) Send(v int) *_state_100 {
   if x.next == nil { x.next = init_state_100(x.c) }; x.c <- v; return x.next }
func (x *_state_99) Recv() (int, *_state_100) {
   if x.next == nil { x.next = init_state_100(x.c) }; return (<-x.c).(int), x.next }

  type _state_98 struct {
   c chan interface{}
   next *_state_99
}

func init_state_98(c chan interface{}) *_state_98 { return &_state_98{c, nil} }
func (x *_state_98) Send(v int) *_state_99 {
   if x.next == nil { x.next = init_state_99(x.c) }; x.c <- v; return x.next }
func (x *_state_98) Recv() (int, *_state_99) {
   if x.next == nil { x.next = init_state_99(x.c) }; return (<-x.c).(int), x.next }

  type _state_97 struct {
   c chan interface{}
   next *_state_98
}

func init_state_97(c chan interface{}) *_state_97 { return &_state_97{c, nil} }
func (x *_state_97) Send(v int) *_state_98 {
   if x.next == nil { x.next = init_state_98(x.c) }; x.c <- v; return x.next }
func (x *_state_97) Recv() (int, *_state_98) {
   if x.next == nil { x.next = init_state_98(x.c) }; return (<-x.c).(int), x.next }

  type _state_96 struct {
   c chan interface{}
   next *_state_97
}

func init_state_96(c chan interface{}) *_state_96 { return &_state_96{c, nil} }
func (x *_state_96) Send(v int) *_state_97 {
   if x.next == nil { x.next = init_state_97(x.c) }; x.c <- v; return x.next }
func (x *_state_96) Recv() (int, *_state_97) {
   if x.next == nil { x.next = init_state_97(x.c) }; return (<-x.c).(int), x.next }

  type _state_95 struct {
   c chan interface{}
   next *_state_96
}

func init_state_95(c chan interface{}) *_state_95 { return &_state_95{c, nil} }
func (x *_state_95) Send(v int) *_state_96 {
   if x.next == nil { x.next = init_state_96(x.c) }; x.c <- v; return x.next }
func (x *_state_95) Recv() (int, *_state_96) {
   if x.next == nil { x.next = init_state_96(x.c) }; return (<-x.c).(int), x.next }

  type _state_94 struct {
   c chan interface{}
   next *_state_95
}

func init_state_94(c chan interface{}) *_state_94 { return &_state_94{c, nil} }
func (x *_state_94) Send(v int) *_state_95 {
   if x.next == nil { x.next = init_state_95(x.c) }; x.c <- v; return x.next }
func (x *_state_94) Recv() (int, *_state_95) {
   if x.next == nil { x.next = init_state_95(x.c) }; return (<-x.c).(int), x.next }

  type _state_93 struct {
   c chan interface{}
   next *_state_94
}

func init_state_93(c chan interface{}) *_state_93 { return &_state_93{c, nil} }
func (x *_state_93) Send(v int) *_state_94 {
   if x.next == nil { x.next = init_state_94(x.c) }; x.c <- v; return x.next }
func (x *_state_93) Recv() (int, *_state_94) {
   if x.next == nil { x.next = init_state_94(x.c) }; return (<-x.c).(int), x.next }

  type _state_92 struct {
   c chan interface{}
   next *_state_93
}

func init_state_92(c chan interface{}) *_state_92 { return &_state_92{c, nil} }
func (x *_state_92) Send(v int) *_state_93 {
   if x.next == nil { x.next = init_state_93(x.c) }; x.c <- v; return x.next }
func (x *_state_92) Recv() (int, *_state_93) {
   if x.next == nil { x.next = init_state_93(x.c) }; return (<-x.c).(int), x.next }

  type _state_91 struct {
   c chan interface{}
   next *_state_92
}

func init_state_91(c chan interface{}) *_state_91 { return &_state_91{c, nil} }
func (x *_state_91) Send(v int) *_state_92 {
   if x.next == nil { x.next = init_state_92(x.c) }; x.c <- v; return x.next }
func (x *_state_91) Recv() (int, *_state_92) {
   if x.next == nil { x.next = init_state_92(x.c) }; return (<-x.c).(int), x.next }

  type _state_90 struct {
   c chan interface{}
   next *_state_91
}

func init_state_90(c chan interface{}) *_state_90 { return &_state_90{c, nil} }
func (x *_state_90) Send(v int) *_state_91 {
   if x.next == nil { x.next = init_state_91(x.c) }; x.c <- v; return x.next }
func (x *_state_90) Recv() (int, *_state_91) {
   if x.next == nil { x.next = init_state_91(x.c) }; return (<-x.c).(int), x.next }

  type _state_89 struct {
   c chan interface{}
   next *_state_90
}

func init_state_89(c chan interface{}) *_state_89 { return &_state_89{c, nil} }
func (x *_state_89) Send(v int) *_state_90 {
   if x.next == nil { x.next = init_state_90(x.c) }; x.c <- v; return x.next }
func (x *_state_89) Recv() (int, *_state_90) {
   if x.next == nil { x.next = init_state_90(x.c) }; return (<-x.c).(int), x.next }

  type _state_88 struct {
   c chan interface{}
   next *_state_89
}

func init_state_88(c chan interface{}) *_state_88 { return &_state_88{c, nil} }
func (x *_state_88) Send(v int) *_state_89 {
   if x.next == nil { x.next = init_state_89(x.c) }; x.c <- v; return x.next }
func (x *_state_88) Recv() (int, *_state_89) {
   if x.next == nil { x.next = init_state_89(x.c) }; return (<-x.c).(int), x.next }

  type _state_87 struct {
   c chan interface{}
   next *_state_88
}

func init_state_87(c chan interface{}) *_state_87 { return &_state_87{c, nil} }
func (x *_state_87) Send(v int) *_state_88 {
   if x.next == nil { x.next = init_state_88(x.c) }; x.c <- v; return x.next }
func (x *_state_87) Recv() (int, *_state_88) {
   if x.next == nil { x.next = init_state_88(x.c) }; return (<-x.c).(int), x.next }

  type _state_86 struct {
   c chan interface{}
   next *_state_87
}

func init_state_86(c chan interface{}) *_state_86 { return &_state_86{c, nil} }
func (x *_state_86) Send(v int) *_state_87 {
   if x.next == nil { x.next = init_state_87(x.c) }; x.c <- v; return x.next }
func (x *_state_86) Recv() (int, *_state_87) {
   if x.next == nil { x.next = init_state_87(x.c) }; return (<-x.c).(int), x.next }

  type _state_85 struct {
   c chan interface{}
   next *_state_86
}

func init_state_85(c chan interface{}) *_state_85 { return &_state_85{c, nil} }
func (x *_state_85) Send(v int) *_state_86 {
   if x.next == nil { x.next = init_state_86(x.c) }; x.c <- v; return x.next }
func (x *_state_85) Recv() (int, *_state_86) {
   if x.next == nil { x.next = init_state_86(x.c) }; return (<-x.c).(int), x.next }

  type _state_84 struct {
   c chan interface{}
   next *_state_85
}

func init_state_84(c chan interface{}) *_state_84 { return &_state_84{c, nil} }
func (x *_state_84) Send(v int) *_state_85 {
   if x.next == nil { x.next = init_state_85(x.c) }; x.c <- v; return x.next }
func (x *_state_84) Recv() (int, *_state_85) {
   if x.next == nil { x.next = init_state_85(x.c) }; return (<-x.c).(int), x.next }

  type _state_83 struct {
   c chan interface{}
   next *_state_84
}

func init_state_83(c chan interface{}) *_state_83 { return &_state_83{c, nil} }
func (x *_state_83) Send(v int) *_state_84 {
   if x.next == nil { x.next = init_state_84(x.c) }; x.c <- v; return x.next }
func (x *_state_83) Recv() (int, *_state_84) {
   if x.next == nil { x.next = init_state_84(x.c) }; return (<-x.c).(int), x.next }

  type _state_82 struct {
   c chan interface{}
   next *_state_83
}

func init_state_82(c chan interface{}) *_state_82 { return &_state_82{c, nil} }
func (x *_state_82) Send(v int) *_state_83 {
   if x.next == nil { x.next = init_state_83(x.c) }; x.c <- v; return x.next }
func (x *_state_82) Recv() (int, *_state_83) {
   if x.next == nil { x.next = init_state_83(x.c) }; return (<-x.c).(int), x.next }

  type _state_81 struct {
   c chan interface{}
   next *_state_82
}

func init_state_81(c chan interface{}) *_state_81 { return &_state_81{c, nil} }
func (x *_state_81) Send(v int) *_state_82 {
   if x.next == nil { x.next = init_state_82(x.c) }; x.c <- v; return x.next }
func (x *_state_81) Recv() (int, *_state_82) {
   if x.next == nil { x.next = init_state_82(x.c) }; return (<-x.c).(int), x.next }

  type _state_80 struct {
   c chan interface{}
   next *_state_81
}

func init_state_80(c chan interface{}) *_state_80 { return &_state_80{c, nil} }
func (x *_state_80) Send(v int) *_state_81 {
   if x.next == nil { x.next = init_state_81(x.c) }; x.c <- v; return x.next }
func (x *_state_80) Recv() (int, *_state_81) {
   if x.next == nil { x.next = init_state_81(x.c) }; return (<-x.c).(int), x.next }

  type _state_79 struct {
   c chan interface{}
   next *_state_80
}

func init_state_79(c chan interface{}) *_state_79 { return &_state_79{c, nil} }
func (x *_state_79) Send(v int) *_state_80 {
   if x.next == nil { x.next = init_state_80(x.c) }; x.c <- v; return x.next }
func (x *_state_79) Recv() (int, *_state_80) {
   if x.next == nil { x.next = init_state_80(x.c) }; return (<-x.c).(int), x.next }

  type _state_78 struct {
   c chan interface{}
   next *_state_79
}

func init_state_78(c chan interface{}) *_state_78 { return &_state_78{c, nil} }
func (x *_state_78) Send(v int) *_state_79 {
   if x.next == nil { x.next = init_state_79(x.c) }; x.c <- v; return x.next }
func (x *_state_78) Recv() (int, *_state_79) {
   if x.next == nil { x.next = init_state_79(x.c) }; return (<-x.c).(int), x.next }

  type _state_77 struct {
   c chan interface{}
   next *_state_78
}

func init_state_77(c chan interface{}) *_state_77 { return &_state_77{c, nil} }
func (x *_state_77) Send(v int) *_state_78 {
   if x.next == nil { x.next = init_state_78(x.c) }; x.c <- v; return x.next }
func (x *_state_77) Recv() (int, *_state_78) {
   if x.next == nil { x.next = init_state_78(x.c) }; return (<-x.c).(int), x.next }

  type _state_76 struct {
   c chan interface{}
   next *_state_77
}

func init_state_76(c chan interface{}) *_state_76 { return &_state_76{c, nil} }
func (x *_state_76) Send(v int) *_state_77 {
   if x.next == nil { x.next = init_state_77(x.c) }; x.c <- v; return x.next }
func (x *_state_76) Recv() (int, *_state_77) {
   if x.next == nil { x.next = init_state_77(x.c) }; return (<-x.c).(int), x.next }

  type _state_75 struct {
   c chan interface{}
   next *_state_76
}

func init_state_75(c chan interface{}) *_state_75 { return &_state_75{c, nil} }
func (x *_state_75) Send(v int) *_state_76 {
   if x.next == nil { x.next = init_state_76(x.c) }; x.c <- v; return x.next }
func (x *_state_75) Recv() (int, *_state_76) {
   if x.next == nil { x.next = init_state_76(x.c) }; return (<-x.c).(int), x.next }

  type _state_74 struct {
   c chan interface{}
   next *_state_75
}

func init_state_74(c chan interface{}) *_state_74 { return &_state_74{c, nil} }
func (x *_state_74) Send(v int) *_state_75 {
   if x.next == nil { x.next = init_state_75(x.c) }; x.c <- v; return x.next }
func (x *_state_74) Recv() (int, *_state_75) {
   if x.next == nil { x.next = init_state_75(x.c) }; return (<-x.c).(int), x.next }

  type _state_73 struct {
   c chan interface{}
   next *_state_74
}

func init_state_73(c chan interface{}) *_state_73 { return &_state_73{c, nil} }
func (x *_state_73) Send(v int) *_state_74 {
   if x.next == nil { x.next = init_state_74(x.c) }; x.c <- v; return x.next }
func (x *_state_73) Recv() (int, *_state_74) {
   if x.next == nil { x.next = init_state_74(x.c) }; return (<-x.c).(int), x.next }

  type _state_72 struct {
   c chan interface{}
   next *_state_73
}

func init_state_72(c chan interface{}) *_state_72 { return &_state_72{c, nil} }
func (x *_state_72) Send(v int) *_state_73 {
   if x.next == nil { x.next = init_state_73(x.c) }; x.c <- v; return x.next }
func (x *_state_72) Recv() (int, *_state_73) {
   if x.next == nil { x.next = init_state_73(x.c) }; return (<-x.c).(int), x.next }

  type _state_71 struct {
   c chan interface{}
   next *_state_72
}

func init_state_71(c chan interface{}) *_state_71 { return &_state_71{c, nil} }
func (x *_state_71) Send(v int) *_state_72 {
   if x.next == nil { x.next = init_state_72(x.c) }; x.c <- v; return x.next }
func (x *_state_71) Recv() (int, *_state_72) {
   if x.next == nil { x.next = init_state_72(x.c) }; return (<-x.c).(int), x.next }

  type _state_70 struct {
   c chan interface{}
   next *_state_71
}

func init_state_70(c chan interface{}) *_state_70 { return &_state_70{c, nil} }
func (x *_state_70) Send(v int) *_state_71 {
   if x.next == nil { x.next = init_state_71(x.c) }; x.c <- v; return x.next }
func (x *_state_70) Recv() (int, *_state_71) {
   if x.next == nil { x.next = init_state_71(x.c) }; return (<-x.c).(int), x.next }

  type _state_69 struct {
   c chan interface{}
   next *_state_70
}

func init_state_69(c chan interface{}) *_state_69 { return &_state_69{c, nil} }
func (x *_state_69) Send(v int) *_state_70 {
   if x.next == nil { x.next = init_state_70(x.c) }; x.c <- v; return x.next }
func (x *_state_69) Recv() (int, *_state_70) {
   if x.next == nil { x.next = init_state_70(x.c) }; return (<-x.c).(int), x.next }

  type _state_68 struct {
   c chan interface{}
   next *_state_69
}

func init_state_68(c chan interface{}) *_state_68 { return &_state_68{c, nil} }
func (x *_state_68) Send(v int) *_state_69 {
   if x.next == nil { x.next = init_state_69(x.c) }; x.c <- v; return x.next }
func (x *_state_68) Recv() (int, *_state_69) {
   if x.next == nil { x.next = init_state_69(x.c) }; return (<-x.c).(int), x.next }

  type _state_67 struct {
   c chan interface{}
   next *_state_68
}

func init_state_67(c chan interface{}) *_state_67 { return &_state_67{c, nil} }
func (x *_state_67) Send(v int) *_state_68 {
   if x.next == nil { x.next = init_state_68(x.c) }; x.c <- v; return x.next }
func (x *_state_67) Recv() (int, *_state_68) {
   if x.next == nil { x.next = init_state_68(x.c) }; return (<-x.c).(int), x.next }

  type _state_66 struct {
   c chan interface{}
   next *_state_67
}

func init_state_66(c chan interface{}) *_state_66 { return &_state_66{c, nil} }
func (x *_state_66) Send(v int) *_state_67 {
   if x.next == nil { x.next = init_state_67(x.c) }; x.c <- v; return x.next }
func (x *_state_66) Recv() (int, *_state_67) {
   if x.next == nil { x.next = init_state_67(x.c) }; return (<-x.c).(int), x.next }

  type _state_65 struct {
   c chan interface{}
   next *_state_66
}

func init_state_65(c chan interface{}) *_state_65 { return &_state_65{c, nil} }
func (x *_state_65) Send(v int) *_state_66 {
   if x.next == nil { x.next = init_state_66(x.c) }; x.c <- v; return x.next }
func (x *_state_65) Recv() (int, *_state_66) {
   if x.next == nil { x.next = init_state_66(x.c) }; return (<-x.c).(int), x.next }

  type _state_64 struct {
   c chan interface{}
   next *_state_65
}

func init_state_64(c chan interface{}) *_state_64 { return &_state_64{c, nil} }
func (x *_state_64) Send(v int) *_state_65 {
   if x.next == nil { x.next = init_state_65(x.c) }; x.c <- v; return x.next }
func (x *_state_64) Recv() (int, *_state_65) {
   if x.next == nil { x.next = init_state_65(x.c) }; return (<-x.c).(int), x.next }

  type _state_63 struct {
   c chan interface{}
   next *_state_64
}

func init_state_63(c chan interface{}) *_state_63 { return &_state_63{c, nil} }
func (x *_state_63) Send(v int) *_state_64 {
   if x.next == nil { x.next = init_state_64(x.c) }; x.c <- v; return x.next }
func (x *_state_63) Recv() (int, *_state_64) {
   if x.next == nil { x.next = init_state_64(x.c) }; return (<-x.c).(int), x.next }

  type _state_62 struct {
   c chan interface{}
   next *_state_63
}

func init_state_62(c chan interface{}) *_state_62 { return &_state_62{c, nil} }
func (x *_state_62) Send(v int) *_state_63 {
   if x.next == nil { x.next = init_state_63(x.c) }; x.c <- v; return x.next }
func (x *_state_62) Recv() (int, *_state_63) {
   if x.next == nil { x.next = init_state_63(x.c) }; return (<-x.c).(int), x.next }

  type _state_61 struct {
   c chan interface{}
   next *_state_62
}

func init_state_61(c chan interface{}) *_state_61 { return &_state_61{c, nil} }
func (x *_state_61) Send(v int) *_state_62 {
   if x.next == nil { x.next = init_state_62(x.c) }; x.c <- v; return x.next }
func (x *_state_61) Recv() (int, *_state_62) {
   if x.next == nil { x.next = init_state_62(x.c) }; return (<-x.c).(int), x.next }

  type _state_60 struct {
   c chan interface{}
   next *_state_61
}

func init_state_60(c chan interface{}) *_state_60 { return &_state_60{c, nil} }
func (x *_state_60) Send(v int) *_state_61 {
   if x.next == nil { x.next = init_state_61(x.c) }; x.c <- v; return x.next }
func (x *_state_60) Recv() (int, *_state_61) {
   if x.next == nil { x.next = init_state_61(x.c) }; return (<-x.c).(int), x.next }

  type _state_59 struct {
   c chan interface{}
   next *_state_60
}

func init_state_59(c chan interface{}) *_state_59 { return &_state_59{c, nil} }
func (x *_state_59) Send(v int) *_state_60 {
   if x.next == nil { x.next = init_state_60(x.c) }; x.c <- v; return x.next }
func (x *_state_59) Recv() (int, *_state_60) {
   if x.next == nil { x.next = init_state_60(x.c) }; return (<-x.c).(int), x.next }

  type _state_58 struct {
   c chan interface{}
   next *_state_59
}

func init_state_58(c chan interface{}) *_state_58 { return &_state_58{c, nil} }
func (x *_state_58) Send(v int) *_state_59 {
   if x.next == nil { x.next = init_state_59(x.c) }; x.c <- v; return x.next }
func (x *_state_58) Recv() (int, *_state_59) {
   if x.next == nil { x.next = init_state_59(x.c) }; return (<-x.c).(int), x.next }

  type _state_57 struct {
   c chan interface{}
   next *_state_58
}

func init_state_57(c chan interface{}) *_state_57 { return &_state_57{c, nil} }
func (x *_state_57) Send(v int) *_state_58 {
   if x.next == nil { x.next = init_state_58(x.c) }; x.c <- v; return x.next }
func (x *_state_57) Recv() (int, *_state_58) {
   if x.next == nil { x.next = init_state_58(x.c) }; return (<-x.c).(int), x.next }

  type _state_56 struct {
   c chan interface{}
   next *_state_57
}

func init_state_56(c chan interface{}) *_state_56 { return &_state_56{c, nil} }
func (x *_state_56) Send(v int) *_state_57 {
   if x.next == nil { x.next = init_state_57(x.c) }; x.c <- v; return x.next }
func (x *_state_56) Recv() (int, *_state_57) {
   if x.next == nil { x.next = init_state_57(x.c) }; return (<-x.c).(int), x.next }

  type _state_55 struct {
   c chan interface{}
   next *_state_56
}

func init_state_55(c chan interface{}) *_state_55 { return &_state_55{c, nil} }
func (x *_state_55) Send(v int) *_state_56 {
   if x.next == nil { x.next = init_state_56(x.c) }; x.c <- v; return x.next }
func (x *_state_55) Recv() (int, *_state_56) {
   if x.next == nil { x.next = init_state_56(x.c) }; return (<-x.c).(int), x.next }

  type _state_54 struct {
   c chan interface{}
   next *_state_55
}

func init_state_54(c chan interface{}) *_state_54 { return &_state_54{c, nil} }
func (x *_state_54) Send(v int) *_state_55 {
   if x.next == nil { x.next = init_state_55(x.c) }; x.c <- v; return x.next }
func (x *_state_54) Recv() (int, *_state_55) {
   if x.next == nil { x.next = init_state_55(x.c) }; return (<-x.c).(int), x.next }

  type _state_53 struct {
   c chan interface{}
   next *_state_54
}

func init_state_53(c chan interface{}) *_state_53 { return &_state_53{c, nil} }
func (x *_state_53) Send(v int) *_state_54 {
   if x.next == nil { x.next = init_state_54(x.c) }; x.c <- v; return x.next }
func (x *_state_53) Recv() (int, *_state_54) {
   if x.next == nil { x.next = init_state_54(x.c) }; return (<-x.c).(int), x.next }

  type _state_52 struct {
   c chan interface{}
   next *_state_53
}

func init_state_52(c chan interface{}) *_state_52 { return &_state_52{c, nil} }
func (x *_state_52) Send(v int) *_state_53 {
   if x.next == nil { x.next = init_state_53(x.c) }; x.c <- v; return x.next }
func (x *_state_52) Recv() (int, *_state_53) {
   if x.next == nil { x.next = init_state_53(x.c) }; return (<-x.c).(int), x.next }

  type _state_51 struct {
   c chan interface{}
   next *_state_52
}

func init_state_51(c chan interface{}) *_state_51 { return &_state_51{c, nil} }
func (x *_state_51) Send(v int) *_state_52 {
   if x.next == nil { x.next = init_state_52(x.c) }; x.c <- v; return x.next }
func (x *_state_51) Recv() (int, *_state_52) {
   if x.next == nil { x.next = init_state_52(x.c) }; return (<-x.c).(int), x.next }

  type _state_50 struct {
   c chan interface{}
   next *_state_51
}

func init_state_50(c chan interface{}) *_state_50 { return &_state_50{c, nil} }
func (x *_state_50) Send(v int) *_state_51 {
   if x.next == nil { x.next = init_state_51(x.c) }; x.c <- v; return x.next }
func (x *_state_50) Recv() (int, *_state_51) {
   if x.next == nil { x.next = init_state_51(x.c) }; return (<-x.c).(int), x.next }

  type _state_49 struct {
   c chan interface{}
   next *_state_50
}

func init_state_49(c chan interface{}) *_state_49 { return &_state_49{c, nil} }
func (x *_state_49) Send(v int) *_state_50 {
   if x.next == nil { x.next = init_state_50(x.c) }; x.c <- v; return x.next }
func (x *_state_49) Recv() (int, *_state_50) {
   if x.next == nil { x.next = init_state_50(x.c) }; return (<-x.c).(int), x.next }

  type _state_48 struct {
   c chan interface{}
   next *_state_49
}

func init_state_48(c chan interface{}) *_state_48 { return &_state_48{c, nil} }
func (x *_state_48) Send(v int) *_state_49 {
   if x.next == nil { x.next = init_state_49(x.c) }; x.c <- v; return x.next }
func (x *_state_48) Recv() (int, *_state_49) {
   if x.next == nil { x.next = init_state_49(x.c) }; return (<-x.c).(int), x.next }

  type _state_47 struct {
   c chan interface{}
   next *_state_48
}

func init_state_47(c chan interface{}) *_state_47 { return &_state_47{c, nil} }
func (x *_state_47) Send(v int) *_state_48 {
   if x.next == nil { x.next = init_state_48(x.c) }; x.c <- v; return x.next }
func (x *_state_47) Recv() (int, *_state_48) {
   if x.next == nil { x.next = init_state_48(x.c) }; return (<-x.c).(int), x.next }

  type _state_46 struct {
   c chan interface{}
   next *_state_47
}

func init_state_46(c chan interface{}) *_state_46 { return &_state_46{c, nil} }
func (x *_state_46) Send(v int) *_state_47 {
   if x.next == nil { x.next = init_state_47(x.c) }; x.c <- v; return x.next }
func (x *_state_46) Recv() (int, *_state_47) {
   if x.next == nil { x.next = init_state_47(x.c) }; return (<-x.c).(int), x.next }

  type _state_45 struct {
   c chan interface{}
   next *_state_46
}

func init_state_45(c chan interface{}) *_state_45 { return &_state_45{c, nil} }
func (x *_state_45) Send(v int) *_state_46 {
   if x.next == nil { x.next = init_state_46(x.c) }; x.c <- v; return x.next }
func (x *_state_45) Recv() (int, *_state_46) {
   if x.next == nil { x.next = init_state_46(x.c) }; return (<-x.c).(int), x.next }

  type _state_44 struct {
   c chan interface{}
   next *_state_45
}

func init_state_44(c chan interface{}) *_state_44 { return &_state_44{c, nil} }
func (x *_state_44) Send(v int) *_state_45 {
   if x.next == nil { x.next = init_state_45(x.c) }; x.c <- v; return x.next }
func (x *_state_44) Recv() (int, *_state_45) {
   if x.next == nil { x.next = init_state_45(x.c) }; return (<-x.c).(int), x.next }

  type _state_43 struct {
   c chan interface{}
   next *_state_44
}

func init_state_43(c chan interface{}) *_state_43 { return &_state_43{c, nil} }
func (x *_state_43) Send(v int) *_state_44 {
   if x.next == nil { x.next = init_state_44(x.c) }; x.c <- v; return x.next }
func (x *_state_43) Recv() (int, *_state_44) {
   if x.next == nil { x.next = init_state_44(x.c) }; return (<-x.c).(int), x.next }

  type _state_42 struct {
   c chan interface{}
   next *_state_43
}

func init_state_42(c chan interface{}) *_state_42 { return &_state_42{c, nil} }
func (x *_state_42) Send(v int) *_state_43 {
   if x.next == nil { x.next = init_state_43(x.c) }; x.c <- v; return x.next }
func (x *_state_42) Recv() (int, *_state_43) {
   if x.next == nil { x.next = init_state_43(x.c) }; return (<-x.c).(int), x.next }

  type _state_41 struct {
   c chan interface{}
   next *_state_42
}

func init_state_41(c chan interface{}) *_state_41 { return &_state_41{c, nil} }
func (x *_state_41) Send(v int) *_state_42 {
   if x.next == nil { x.next = init_state_42(x.c) }; x.c <- v; return x.next }
func (x *_state_41) Recv() (int, *_state_42) {
   if x.next == nil { x.next = init_state_42(x.c) }; return (<-x.c).(int), x.next }

  type _state_40 struct {
   c chan interface{}
   next *_state_41
}

func init_state_40(c chan interface{}) *_state_40 { return &_state_40{c, nil} }
func (x *_state_40) Send(v int) *_state_41 {
   if x.next == nil { x.next = init_state_41(x.c) }; x.c <- v; return x.next }
func (x *_state_40) Recv() (int, *_state_41) {
   if x.next == nil { x.next = init_state_41(x.c) }; return (<-x.c).(int), x.next }

  type _state_39 struct {
   c chan interface{}
   next *_state_40
}

func init_state_39(c chan interface{}) *_state_39 { return &_state_39{c, nil} }
func (x *_state_39) Send(v int) *_state_40 {
   if x.next == nil { x.next = init_state_40(x.c) }; x.c <- v; return x.next }
func (x *_state_39) Recv() (int, *_state_40) {
   if x.next == nil { x.next = init_state_40(x.c) }; return (<-x.c).(int), x.next }

  type _state_38 struct {
   c chan interface{}
   next *_state_39
}

func init_state_38(c chan interface{}) *_state_38 { return &_state_38{c, nil} }
func (x *_state_38) Send(v int) *_state_39 {
   if x.next == nil { x.next = init_state_39(x.c) }; x.c <- v; return x.next }
func (x *_state_38) Recv() (int, *_state_39) {
   if x.next == nil { x.next = init_state_39(x.c) }; return (<-x.c).(int), x.next }

  type _state_37 struct {
   c chan interface{}
   next *_state_38
}

func init_state_37(c chan interface{}) *_state_37 { return &_state_37{c, nil} }
func (x *_state_37) Send(v int) *_state_38 {
   if x.next == nil { x.next = init_state_38(x.c) }; x.c <- v; return x.next }
func (x *_state_37) Recv() (int, *_state_38) {
   if x.next == nil { x.next = init_state_38(x.c) }; return (<-x.c).(int), x.next }

  type _state_36 struct {
   c chan interface{}
   next *_state_37
}

func init_state_36(c chan interface{}) *_state_36 { return &_state_36{c, nil} }
func (x *_state_36) Send(v int) *_state_37 {
   if x.next == nil { x.next = init_state_37(x.c) }; x.c <- v; return x.next }
func (x *_state_36) Recv() (int, *_state_37) {
   if x.next == nil { x.next = init_state_37(x.c) }; return (<-x.c).(int), x.next }

  type _state_35 struct {
   c chan interface{}
   next *_state_36
}

func init_state_35(c chan interface{}) *_state_35 { return &_state_35{c, nil} }
func (x *_state_35) Send(v int) *_state_36 {
   if x.next == nil { x.next = init_state_36(x.c) }; x.c <- v; return x.next }
func (x *_state_35) Recv() (int, *_state_36) {
   if x.next == nil { x.next = init_state_36(x.c) }; return (<-x.c).(int), x.next }

  type _state_34 struct {
   c chan interface{}
   next *_state_35
}

func init_state_34(c chan interface{}) *_state_34 { return &_state_34{c, nil} }
func (x *_state_34) Send(v int) *_state_35 {
   if x.next == nil { x.next = init_state_35(x.c) }; x.c <- v; return x.next }
func (x *_state_34) Recv() (int, *_state_35) {
   if x.next == nil { x.next = init_state_35(x.c) }; return (<-x.c).(int), x.next }

  type _state_33 struct {
   c chan interface{}
   next *_state_34
}

func init_state_33(c chan interface{}) *_state_33 { return &_state_33{c, nil} }
func (x *_state_33) Send(v int) *_state_34 {
   if x.next == nil { x.next = init_state_34(x.c) }; x.c <- v; return x.next }
func (x *_state_33) Recv() (int, *_state_34) {
   if x.next == nil { x.next = init_state_34(x.c) }; return (<-x.c).(int), x.next }

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
c0 := c.Send((n + 1), (n + 2), (n + 3), (n + 4), (n + 5), (n + 6), (n + 7), (n + 8), (n + 9), (n + 10), (n + 11), (n + 12), (n + 13), (n + 14), (n + 15), (n + 16), (n + 17), (n + 18), (n + 19), (n + 20), (n + 21), (n + 22), (n + 23), (n + 24), (n + 25), (n + 26), (n + 27), (n + 28), (n + 29), (n + 30), (n + 31), (n + 32), (n + 33), (n + 34), (n + 35), (n + 36), (n + 37), (n + 38), (n + 39), (n + 40), (n + 41), (n + 42), (n + 43), (n + 44), (n + 45), (n + 46), (n + 47), (n + 48), (n + 49), (n + 50), (n + 51), (n + 52), (n + 53), (n + 54), (n + 55), (n + 56), (n + 57), (n + 58), (n + 59), (n + 60), (n + 61), (n + 62), (n + 63), (n + 64), (n + 65), (n + 66), (n + 67), (n + 68), (n + 69), (n + 70), (n + 71), (n + 72), (n + 73), (n + 74), (n + 75), (n + 76), (n + 77), (n + 78), (n + 79), (n + 80), (n + 81), (n + 82), (n + 83), (n + 84), (n + 85), (n + 86), (n + 87), (n + 88), (n + 89), (n + 90), (n + 91), (n + 92), (n + 93), (n + 94), (n + 95), (n + 96), (n + 97), (n + 98), (n + 99), (n + 100), (n + 101), (n + 102), (n + 103), (n + 104), (n + 105), (n + 106), (n + 107), (n + 108), (n + 109), (n + 110), (n + 111), (n + 112), (n + 113), (n + 114), (n + 115), (n + 116), (n + 117), (n + 118), (n + 119), (n + 120), (n + 121), (n + 122), (n + 123), (n + 124), (n + 125), (n + 126), (n + 127), (n + 128))
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
c64 := c63.Send((n + 65))
c65 := c64.Send((n + 66))
c66 := c65.Send((n + 67))
c67 := c66.Send((n + 68))
c68 := c67.Send((n + 69))
c69 := c68.Send((n + 70))
c70 := c69.Send((n + 71))
c71 := c70.Send((n + 72))
c72 := c71.Send((n + 73))
c73 := c72.Send((n + 74))
c74 := c73.Send((n + 75))
c75 := c74.Send((n + 76))
c76 := c75.Send((n + 77))
c77 := c76.Send((n + 78))
c78 := c77.Send((n + 79))
c79 := c78.Send((n + 80))
c80 := c79.Send((n + 81))
c81 := c80.Send((n + 82))
c82 := c81.Send((n + 83))
c83 := c82.Send((n + 84))
c84 := c83.Send((n + 85))
c85 := c84.Send((n + 86))
c86 := c85.Send((n + 87))
c87 := c86.Send((n + 88))
c88 := c87.Send((n + 89))
c89 := c88.Send((n + 90))
c90 := c89.Send((n + 91))
c91 := c90.Send((n + 92))
c92 := c91.Send((n + 93))
c93 := c92.Send((n + 94))
c94 := c93.Send((n + 95))
c95 := c94.Send((n + 96))
c96 := c95.Send((n + 97))
c97 := c96.Send((n + 98))
c98 := c97.Send((n + 99))
c99 := c98.Send((n + 100))
c100 := c99.Send((n + 101))
c101 := c100.Send((n + 102))
c102 := c101.Send((n + 103))
c103 := c102.Send((n + 104))
c104 := c103.Send((n + 105))
c105 := c104.Send((n + 106))
c106 := c105.Send((n + 107))
c107 := c106.Send((n + 108))
c108 := c107.Send((n + 109))
c109 := c108.Send((n + 110))
c110 := c109.Send((n + 111))
c111 := c110.Send((n + 112))
c112 := c111.Send((n + 113))
c113 := c112.Send((n + 114))
c114 := c113.Send((n + 115))
c115 := c114.Send((n + 116))
c116 := c115.Send((n + 117))
c117 := c116.Send((n + 118))
c118 := c117.Send((n + 119))
c119 := c118.Send((n + 120))
c120 := c119.Send((n + 121))
c121 := c120.Send((n + 122))
c122 := c121.Send((n + 123))
c123 := c122.Send((n + 124))
c124 := c123.Send((n + 125))
c125 := c124.Send((n + 126))
c126 := c125.Send((n + 127))
c127 := c126.Send((n + 128))
c127.Send(nil)
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
a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28, a29, a30, a31, a32, a33, a34, a35, a36, a37, a38, a39, a40, a41, a42, a43, a44, a45, a46, a47, a48, a49, a50, a51, a52, a53, a54, a55, a56, a57, a58, a59, a60, a61, a62, a63, a64, a65, a66, a67, a68, a69, a70, a71, a72, a73, a74, a75, a76, a77, a78, a79, a80, a81, a82, a83, a84, a85, a86, a87, a88, a89, a90, a91, a92, a93, a94, a95, a96, a97, a98, a99, a100, a101, a102, a103, a104, a105, a106, a107, a108, a109, a110, a111, a112, a113, a114, a115, a116, a117, a118, a119, a120, a121, a122, a123, a124, a125, a126, a127, a128, e0 := e.Recv()
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
fmt.Printf("%v\n",a33)
fmt.Printf("%v\n",a34)
fmt.Printf("%v\n",a35)
fmt.Printf("%v\n",a36)
fmt.Printf("%v\n",a37)
fmt.Printf("%v\n",a38)
fmt.Printf("%v\n",a39)
fmt.Printf("%v\n",a40)
fmt.Printf("%v\n",a41)
fmt.Printf("%v\n",a42)
fmt.Printf("%v\n",a43)
fmt.Printf("%v\n",a44)
fmt.Printf("%v\n",a45)
fmt.Printf("%v\n",a46)
fmt.Printf("%v\n",a47)
fmt.Printf("%v\n",a48)
fmt.Printf("%v\n",a49)
fmt.Printf("%v\n",a50)
fmt.Printf("%v\n",a51)
fmt.Printf("%v\n",a52)
fmt.Printf("%v\n",a53)
fmt.Printf("%v\n",a54)
fmt.Printf("%v\n",a55)
fmt.Printf("%v\n",a56)
fmt.Printf("%v\n",a57)
fmt.Printf("%v\n",a58)
fmt.Printf("%v\n",a59)
fmt.Printf("%v\n",a60)
fmt.Printf("%v\n",a61)
fmt.Printf("%v\n",a62)
fmt.Printf("%v\n",a63)
fmt.Printf("%v\n",a64)
fmt.Printf("%v\n",a65)
fmt.Printf("%v\n",a66)
fmt.Printf("%v\n",a67)
fmt.Printf("%v\n",a68)
fmt.Printf("%v\n",a69)
fmt.Printf("%v\n",a70)
fmt.Printf("%v\n",a71)
fmt.Printf("%v\n",a72)
fmt.Printf("%v\n",a73)
fmt.Printf("%v\n",a74)
fmt.Printf("%v\n",a75)
fmt.Printf("%v\n",a76)
fmt.Printf("%v\n",a77)
fmt.Printf("%v\n",a78)
fmt.Printf("%v\n",a79)
fmt.Printf("%v\n",a80)
fmt.Printf("%v\n",a81)
fmt.Printf("%v\n",a82)
fmt.Printf("%v\n",a83)
fmt.Printf("%v\n",a84)
fmt.Printf("%v\n",a85)
fmt.Printf("%v\n",a86)
fmt.Printf("%v\n",a87)
fmt.Printf("%v\n",a88)
fmt.Printf("%v\n",a89)
fmt.Printf("%v\n",a90)
fmt.Printf("%v\n",a91)
fmt.Printf("%v\n",a92)
fmt.Printf("%v\n",a93)
fmt.Printf("%v\n",a94)
fmt.Printf("%v\n",a95)
fmt.Printf("%v\n",a96)
fmt.Printf("%v\n",a97)
fmt.Printf("%v\n",a98)
fmt.Printf("%v\n",a99)
fmt.Printf("%v\n",a100)
fmt.Printf("%v\n",a101)
fmt.Printf("%v\n",a102)
fmt.Printf("%v\n",a103)
fmt.Printf("%v\n",a104)
fmt.Printf("%v\n",a105)
fmt.Printf("%v\n",a106)
fmt.Printf("%v\n",a107)
fmt.Printf("%v\n",a108)
fmt.Printf("%v\n",a109)
fmt.Printf("%v\n",a110)
fmt.Printf("%v\n",a111)
fmt.Printf("%v\n",a112)
fmt.Printf("%v\n",a113)
fmt.Printf("%v\n",a114)
fmt.Printf("%v\n",a115)
fmt.Printf("%v\n",a116)
fmt.Printf("%v\n",a117)
fmt.Printf("%v\n",a118)
fmt.Printf("%v\n",a119)
fmt.Printf("%v\n",a120)
fmt.Printf("%v\n",a121)
fmt.Printf("%v\n",a122)
fmt.Printf("%v\n",a123)
fmt.Printf("%v\n",a124)
fmt.Printf("%v\n",a125)
fmt.Printf("%v\n",a126)
fmt.Printf("%v\n",a127)
fmt.Printf("%v\n",a128)
e0.Recv()
m.Send(nil)
}(m)
}
