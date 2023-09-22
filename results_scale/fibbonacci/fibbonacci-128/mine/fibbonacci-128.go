package main

import "fmt"

// Preamble generation
type _state_1 struct {
	c    chan interface{}
	next *_state_0
}

func init_state_1(c chan interface{}) *_state_1 { return &_state_1{c, nil} }
func (x *_state_1) Send(v int) *_state_0 {
	if x.next == nil {
		x.next = init_state_0(x.c)
	}
	x.c <- v
	return x.next
}
func (x *_state_1) Recv() (int, *_state_0) {
	if x.next == nil {
		x.next = init_state_0(x.c)
	}
	return (<-x.c).(int), x.next
}

type _state_2 struct {
	c chan interface{}
}

func init_state_2(c chan interface{}) *_state_2 { return &_state_2{c} }
func (x *_state_2) Send(v interface{})          { x.c <- v }
func (x *_state_2) Recv() interface{}           { return <-x.c }

type _state_0 struct {
	c  chan interface{}
	ls map[string]interface{}
}

func init_state_0(c chan interface{}) *_state_0 {
	m := make(map[string]interface{})
	m["stop"] = init_state_2(c)
	m["next"] = init_state_1(c)
	return &_state_0{c, m}
}
func (x *_state_0) Send(v string) { x.c <- v }
func (x *_state_0) Recv() string  { return (<-x.c).(string) }

//Declaration list compilation
func fib(a int) func(_x int) func(_x *_state_0) {
	return func(b int) func(_x *_state_0) {
		return func(c *_state_0) {
			label := c.Recv()
			switch label {
			case "next":
				c0 := c.ls["next"].(*_state_1)
				c1 := c0.Send(b)
				fib(b)((a + b))(c1)
			case "stop":
				c0 := c.ls["stop"].(*_state_2)
				c0.Send(nil)
			}
		}
	}
}

// Main compilation
func main() {
	c := init_state_2(make(chan interface{}))
	go func() {
		c.Recv()
	}()
	func(c *_state_2) {
		q := init_state_0(make(chan interface{}))
		go fib(0)(1)(q)
		q.Send("next")
		q0 := q.ls["next"].(*_state_1)
		x1, q1 := q0.Recv()
		fmt.Printf("%v\n", x1)
		q1.Send("next")
		q2 := q1.ls["next"].(*_state_1)
		x2, q3 := q2.Recv()
		fmt.Printf("%v\n", x2)
		q3.Send("next")
		q4 := q3.ls["next"].(*_state_1)
		x3, q5 := q4.Recv()
		fmt.Printf("%v\n", x3)
		q5.Send("next")
		q6 := q5.ls["next"].(*_state_1)
		x4, q7 := q6.Recv()
		fmt.Printf("%v\n", x4)
		q7.Send("next")
		q8 := q7.ls["next"].(*_state_1)
		x5, q9 := q8.Recv()
		fmt.Printf("%v\n", x5)
		q9.Send("next")
		q10 := q9.ls["next"].(*_state_1)
		x6, q11 := q10.Recv()
		fmt.Printf("%v\n", x6)
		q11.Send("next")
		q12 := q11.ls["next"].(*_state_1)
		x7, q13 := q12.Recv()
		fmt.Printf("%v\n", x7)
		q13.Send("next")
		q14 := q13.ls["next"].(*_state_1)
		x8, q15 := q14.Recv()
		fmt.Printf("%v\n", x8)
		q15.Send("next")
		q16 := q15.ls["next"].(*_state_1)
		x9, q17 := q16.Recv()
		fmt.Printf("%v\n", x9)
		q17.Send("next")
		q18 := q17.ls["next"].(*_state_1)
		x10, q19 := q18.Recv()
		fmt.Printf("%v\n", x10)
		q19.Send("next")
		q20 := q19.ls["next"].(*_state_1)
		x11, q21 := q20.Recv()
		fmt.Printf("%v\n", x11)
		q21.Send("next")
		q22 := q21.ls["next"].(*_state_1)
		x12, q23 := q22.Recv()
		fmt.Printf("%v\n", x12)
		q23.Send("next")
		q24 := q23.ls["next"].(*_state_1)
		x13, q25 := q24.Recv()
		fmt.Printf("%v\n", x13)
		q25.Send("next")
		q26 := q25.ls["next"].(*_state_1)
		x14, q27 := q26.Recv()
		fmt.Printf("%v\n", x14)
		q27.Send("next")
		q28 := q27.ls["next"].(*_state_1)
		x15, q29 := q28.Recv()
		fmt.Printf("%v\n", x15)
		q29.Send("next")
		q30 := q29.ls["next"].(*_state_1)
		x16, q31 := q30.Recv()
		fmt.Printf("%v\n", x16)
		q31.Send("next")
		q32 := q31.ls["next"].(*_state_1)
		x17, q33 := q32.Recv()
		fmt.Printf("%v\n", x17)
		q33.Send("next")
		q34 := q33.ls["next"].(*_state_1)
		x18, q35 := q34.Recv()
		fmt.Printf("%v\n", x18)
		q35.Send("next")
		q36 := q35.ls["next"].(*_state_1)
		x19, q37 := q36.Recv()
		fmt.Printf("%v\n", x19)
		q37.Send("next")
		q38 := q37.ls["next"].(*_state_1)
		x20, q39 := q38.Recv()
		fmt.Printf("%v\n", x20)
		q39.Send("next")
		q40 := q39.ls["next"].(*_state_1)
		x21, q41 := q40.Recv()
		fmt.Printf("%v\n", x21)
		q41.Send("next")
		q42 := q41.ls["next"].(*_state_1)
		x22, q43 := q42.Recv()
		fmt.Printf("%v\n", x22)
		q43.Send("next")
		q44 := q43.ls["next"].(*_state_1)
		x23, q45 := q44.Recv()
		fmt.Printf("%v\n", x23)
		q45.Send("next")
		q46 := q45.ls["next"].(*_state_1)
		x24, q47 := q46.Recv()
		fmt.Printf("%v\n", x24)
		q47.Send("next")
		q48 := q47.ls["next"].(*_state_1)
		x25, q49 := q48.Recv()
		fmt.Printf("%v\n", x25)
		q49.Send("next")
		q50 := q49.ls["next"].(*_state_1)
		x26, q51 := q50.Recv()
		fmt.Printf("%v\n", x26)
		q51.Send("next")
		q52 := q51.ls["next"].(*_state_1)
		x27, q53 := q52.Recv()
		fmt.Printf("%v\n", x27)
		q53.Send("next")
		q54 := q53.ls["next"].(*_state_1)
		x28, q55 := q54.Recv()
		fmt.Printf("%v\n", x28)
		q55.Send("next")
		q56 := q55.ls["next"].(*_state_1)
		x29, q57 := q56.Recv()
		fmt.Printf("%v\n", x29)
		q57.Send("next")
		q58 := q57.ls["next"].(*_state_1)
		x30, q59 := q58.Recv()
		fmt.Printf("%v\n", x30)
		q59.Send("next")
		q60 := q59.ls["next"].(*_state_1)
		x31, q61 := q60.Recv()
		fmt.Printf("%v\n", x31)
		q61.Send("next")
		q62 := q61.ls["next"].(*_state_1)
		x32, q63 := q62.Recv()
		fmt.Printf("%v\n", x32)
		q63.Send("next")
		q64 := q63.ls["next"].(*_state_1)
		x33, q65 := q64.Recv()
		fmt.Printf("%v\n", x33)
		q65.Send("next")
		q66 := q65.ls["next"].(*_state_1)
		x34, q67 := q66.Recv()
		fmt.Printf("%v\n", x34)
		q67.Send("next")
		q68 := q67.ls["next"].(*_state_1)
		x35, q69 := q68.Recv()
		fmt.Printf("%v\n", x35)
		q69.Send("next")
		q70 := q69.ls["next"].(*_state_1)
		x36, q71 := q70.Recv()
		fmt.Printf("%v\n", x36)
		q71.Send("next")
		q72 := q71.ls["next"].(*_state_1)
		x37, q73 := q72.Recv()
		fmt.Printf("%v\n", x37)
		q73.Send("next")
		q74 := q73.ls["next"].(*_state_1)
		x38, q75 := q74.Recv()
		fmt.Printf("%v\n", x38)
		q75.Send("next")
		q76 := q75.ls["next"].(*_state_1)
		x39, q77 := q76.Recv()
		fmt.Printf("%v\n", x39)
		q77.Send("next")
		q78 := q77.ls["next"].(*_state_1)
		x40, q79 := q78.Recv()
		fmt.Printf("%v\n", x40)
		q79.Send("next")
		q80 := q79.ls["next"].(*_state_1)
		x41, q81 := q80.Recv()
		fmt.Printf("%v\n", x41)
		q81.Send("next")
		q82 := q81.ls["next"].(*_state_1)
		x42, q83 := q82.Recv()
		fmt.Printf("%v\n", x42)
		q83.Send("next")
		q84 := q83.ls["next"].(*_state_1)
		x43, q85 := q84.Recv()
		fmt.Printf("%v\n", x43)
		q85.Send("next")
		q86 := q85.ls["next"].(*_state_1)
		x44, q87 := q86.Recv()
		fmt.Printf("%v\n", x44)
		q87.Send("next")
		q88 := q87.ls["next"].(*_state_1)
		x45, q89 := q88.Recv()
		fmt.Printf("%v\n", x45)
		q89.Send("next")
		q90 := q89.ls["next"].(*_state_1)
		x46, q91 := q90.Recv()
		fmt.Printf("%v\n", x46)
		q91.Send("next")
		q92 := q91.ls["next"].(*_state_1)
		x47, q93 := q92.Recv()
		fmt.Printf("%v\n", x47)
		q93.Send("next")
		q94 := q93.ls["next"].(*_state_1)
		x48, q95 := q94.Recv()
		fmt.Printf("%v\n", x48)
		q95.Send("next")
		q96 := q95.ls["next"].(*_state_1)
		x49, q97 := q96.Recv()
		fmt.Printf("%v\n", x49)
		q97.Send("next")
		q98 := q97.ls["next"].(*_state_1)
		x50, q99 := q98.Recv()
		fmt.Printf("%v\n", x50)
		q99.Send("next")
		q100 := q99.ls["next"].(*_state_1)
		x51, q101 := q100.Recv()
		fmt.Printf("%v\n", x51)
		q101.Send("next")
		q102 := q101.ls["next"].(*_state_1)
		x52, q103 := q102.Recv()
		fmt.Printf("%v\n", x52)
		q103.Send("next")
		q104 := q103.ls["next"].(*_state_1)
		x53, q105 := q104.Recv()
		fmt.Printf("%v\n", x53)
		q105.Send("next")
		q106 := q105.ls["next"].(*_state_1)
		x54, q107 := q106.Recv()
		fmt.Printf("%v\n", x54)
		q107.Send("next")
		q108 := q107.ls["next"].(*_state_1)
		x55, q109 := q108.Recv()
		fmt.Printf("%v\n", x55)
		q109.Send("next")
		q110 := q109.ls["next"].(*_state_1)
		x56, q111 := q110.Recv()
		fmt.Printf("%v\n", x56)
		q111.Send("next")
		q112 := q111.ls["next"].(*_state_1)
		x57, q113 := q112.Recv()
		fmt.Printf("%v\n", x57)
		q113.Send("next")
		q114 := q113.ls["next"].(*_state_1)
		x58, q115 := q114.Recv()
		fmt.Printf("%v\n", x58)
		q115.Send("next")
		q116 := q115.ls["next"].(*_state_1)
		x59, q117 := q116.Recv()
		fmt.Printf("%v\n", x59)
		q117.Send("next")
		q118 := q117.ls["next"].(*_state_1)
		x60, q119 := q118.Recv()
		fmt.Printf("%v\n", x60)
		q119.Send("next")
		q120 := q119.ls["next"].(*_state_1)
		x61, q121 := q120.Recv()
		fmt.Printf("%v\n", x61)
		q121.Send("next")
		q122 := q121.ls["next"].(*_state_1)
		x62, q123 := q122.Recv()
		fmt.Printf("%v\n", x62)
		q123.Send("next")
		q124 := q123.ls["next"].(*_state_1)
		x63, q125 := q124.Recv()
		fmt.Printf("%v\n", x63)
		q125.Send("next")
		q126 := q125.ls["next"].(*_state_1)
		x64, q127 := q126.Recv()
		fmt.Printf("%v\n", x64)
		q127.Send("next")
		q128 := q127.ls["next"].(*_state_1)
		x65, q129 := q128.Recv()
		fmt.Printf("%v\n", x65)
		q129.Send("next")
		q130 := q129.ls["next"].(*_state_1)
		x66, q131 := q130.Recv()
		fmt.Printf("%v\n", x66)
		q131.Send("next")
		q132 := q131.ls["next"].(*_state_1)
		x67, q133 := q132.Recv()
		fmt.Printf("%v\n", x67)
		q133.Send("next")
		q134 := q133.ls["next"].(*_state_1)
		x68, q135 := q134.Recv()
		fmt.Printf("%v\n", x68)
		q135.Send("next")
		q136 := q135.ls["next"].(*_state_1)
		x69, q137 := q136.Recv()
		fmt.Printf("%v\n", x69)
		q137.Send("next")
		q138 := q137.ls["next"].(*_state_1)
		x70, q139 := q138.Recv()
		fmt.Printf("%v\n", x70)
		q139.Send("next")
		q140 := q139.ls["next"].(*_state_1)
		x71, q141 := q140.Recv()
		fmt.Printf("%v\n", x71)
		q141.Send("next")
		q142 := q141.ls["next"].(*_state_1)
		x72, q143 := q142.Recv()
		fmt.Printf("%v\n", x72)
		q143.Send("next")
		q144 := q143.ls["next"].(*_state_1)
		x73, q145 := q144.Recv()
		fmt.Printf("%v\n", x73)
		q145.Send("next")
		q146 := q145.ls["next"].(*_state_1)
		x74, q147 := q146.Recv()
		fmt.Printf("%v\n", x74)
		q147.Send("next")
		q148 := q147.ls["next"].(*_state_1)
		x75, q149 := q148.Recv()
		fmt.Printf("%v\n", x75)
		q149.Send("next")
		q150 := q149.ls["next"].(*_state_1)
		x76, q151 := q150.Recv()
		fmt.Printf("%v\n", x76)
		q151.Send("next")
		q152 := q151.ls["next"].(*_state_1)
		x77, q153 := q152.Recv()
		fmt.Printf("%v\n", x77)
		q153.Send("next")
		q154 := q153.ls["next"].(*_state_1)
		x78, q155 := q154.Recv()
		fmt.Printf("%v\n", x78)
		q155.Send("next")
		q156 := q155.ls["next"].(*_state_1)
		x79, q157 := q156.Recv()
		fmt.Printf("%v\n", x79)
		q157.Send("next")
		q158 := q157.ls["next"].(*_state_1)
		x80, q159 := q158.Recv()
		fmt.Printf("%v\n", x80)
		q159.Send("next")
		q160 := q159.ls["next"].(*_state_1)
		x81, q161 := q160.Recv()
		fmt.Printf("%v\n", x81)
		q161.Send("next")
		q162 := q161.ls["next"].(*_state_1)
		x82, q163 := q162.Recv()
		fmt.Printf("%v\n", x82)
		q163.Send("next")
		q164 := q163.ls["next"].(*_state_1)
		x83, q165 := q164.Recv()
		fmt.Printf("%v\n", x83)
		q165.Send("next")
		q166 := q165.ls["next"].(*_state_1)
		x84, q167 := q166.Recv()
		fmt.Printf("%v\n", x84)
		q167.Send("next")
		q168 := q167.ls["next"].(*_state_1)
		x85, q169 := q168.Recv()
		fmt.Printf("%v\n", x85)
		q169.Send("next")
		q170 := q169.ls["next"].(*_state_1)
		x86, q171 := q170.Recv()
		fmt.Printf("%v\n", x86)
		q171.Send("next")
		q172 := q171.ls["next"].(*_state_1)
		x87, q173 := q172.Recv()
		fmt.Printf("%v\n", x87)
		q173.Send("next")
		q174 := q173.ls["next"].(*_state_1)
		x88, q175 := q174.Recv()
		fmt.Printf("%v\n", x88)
		q175.Send("next")
		q176 := q175.ls["next"].(*_state_1)
		x89, q177 := q176.Recv()
		fmt.Printf("%v\n", x89)
		q177.Send("next")
		q178 := q177.ls["next"].(*_state_1)
		x90, q179 := q178.Recv()
		fmt.Printf("%v\n", x90)
		q179.Send("next")
		q180 := q179.ls["next"].(*_state_1)
		x91, q181 := q180.Recv()
		fmt.Printf("%v\n", x91)
		q181.Send("next")
		q182 := q181.ls["next"].(*_state_1)
		x92, q183 := q182.Recv()
		fmt.Printf("%v\n", x92)
		q183.Send("next")
		q184 := q183.ls["next"].(*_state_1)
		x93, q185 := q184.Recv()
		fmt.Printf("%v\n", x93)
		q185.Send("next")
		q186 := q185.ls["next"].(*_state_1)
		x94, q187 := q186.Recv()
		fmt.Printf("%v\n", x94)
		q187.Send("next")
		q188 := q187.ls["next"].(*_state_1)
		x95, q189 := q188.Recv()
		fmt.Printf("%v\n", x95)
		q189.Send("next")
		q190 := q189.ls["next"].(*_state_1)
		x96, q191 := q190.Recv()
		fmt.Printf("%v\n", x96)
		q191.Send("next")
		q192 := q191.ls["next"].(*_state_1)
		x97, q193 := q192.Recv()
		fmt.Printf("%v\n", x97)
		q193.Send("next")
		q194 := q193.ls["next"].(*_state_1)
		x98, q195 := q194.Recv()
		fmt.Printf("%v\n", x98)
		q195.Send("next")
		q196 := q195.ls["next"].(*_state_1)
		x99, q197 := q196.Recv()
		fmt.Printf("%v\n", x99)
		q197.Send("next")
		q198 := q197.ls["next"].(*_state_1)
		x100, q199 := q198.Recv()
		fmt.Printf("%v\n", x100)
		q199.Send("next")
		q200 := q199.ls["next"].(*_state_1)
		x101, q201 := q200.Recv()
		fmt.Printf("%v\n", x101)
		q201.Send("next")
		q202 := q201.ls["next"].(*_state_1)
		x102, q203 := q202.Recv()
		fmt.Printf("%v\n", x102)
		q203.Send("next")
		q204 := q203.ls["next"].(*_state_1)
		x103, q205 := q204.Recv()
		fmt.Printf("%v\n", x103)
		q205.Send("next")
		q206 := q205.ls["next"].(*_state_1)
		x104, q207 := q206.Recv()
		fmt.Printf("%v\n", x104)
		q207.Send("next")
		q208 := q207.ls["next"].(*_state_1)
		x105, q209 := q208.Recv()
		fmt.Printf("%v\n", x105)
		q209.Send("next")
		q210 := q209.ls["next"].(*_state_1)
		x106, q211 := q210.Recv()
		fmt.Printf("%v\n", x106)
		q211.Send("next")
		q212 := q211.ls["next"].(*_state_1)
		x107, q213 := q212.Recv()
		fmt.Printf("%v\n", x107)
		q213.Send("next")
		q214 := q213.ls["next"].(*_state_1)
		x108, q215 := q214.Recv()
		fmt.Printf("%v\n", x108)
		q215.Send("next")
		q216 := q215.ls["next"].(*_state_1)
		x109, q217 := q216.Recv()
		fmt.Printf("%v\n", x109)
		q217.Send("next")
		q218 := q217.ls["next"].(*_state_1)
		x110, q219 := q218.Recv()
		fmt.Printf("%v\n", x110)
		q219.Send("next")
		q220 := q219.ls["next"].(*_state_1)
		x111, q221 := q220.Recv()
		fmt.Printf("%v\n", x111)
		q221.Send("next")
		q222 := q221.ls["next"].(*_state_1)
		x112, q223 := q222.Recv()
		fmt.Printf("%v\n", x112)
		q223.Send("next")
		q224 := q223.ls["next"].(*_state_1)
		x113, q225 := q224.Recv()
		fmt.Printf("%v\n", x113)
		q225.Send("next")
		q226 := q225.ls["next"].(*_state_1)
		x114, q227 := q226.Recv()
		fmt.Printf("%v\n", x114)
		q227.Send("next")
		q228 := q227.ls["next"].(*_state_1)
		x115, q229 := q228.Recv()
		fmt.Printf("%v\n", x115)
		q229.Send("next")
		q230 := q229.ls["next"].(*_state_1)
		x116, q231 := q230.Recv()
		fmt.Printf("%v\n", x116)
		q231.Send("next")
		q232 := q231.ls["next"].(*_state_1)
		x117, q233 := q232.Recv()
		fmt.Printf("%v\n", x117)
		q233.Send("next")
		q234 := q233.ls["next"].(*_state_1)
		x118, q235 := q234.Recv()
		fmt.Printf("%v\n", x118)
		q235.Send("next")
		q236 := q235.ls["next"].(*_state_1)
		x119, q237 := q236.Recv()
		fmt.Printf("%v\n", x119)
		q237.Send("next")
		q238 := q237.ls["next"].(*_state_1)
		x120, q239 := q238.Recv()
		fmt.Printf("%v\n", x120)
		q239.Send("next")
		q240 := q239.ls["next"].(*_state_1)
		x121, q241 := q240.Recv()
		fmt.Printf("%v\n", x121)
		q241.Send("next")
		q242 := q241.ls["next"].(*_state_1)
		x122, q243 := q242.Recv()
		fmt.Printf("%v\n", x122)
		q243.Send("next")
		q244 := q243.ls["next"].(*_state_1)
		x123, q245 := q244.Recv()
		fmt.Printf("%v\n", x123)
		q245.Send("next")
		q246 := q245.ls["next"].(*_state_1)
		x124, q247 := q246.Recv()
		fmt.Printf("%v\n", x124)
		q247.Send("next")
		q248 := q247.ls["next"].(*_state_1)
		x125, q249 := q248.Recv()
		fmt.Printf("%v\n", x125)
		q249.Send("next")
		q250 := q249.ls["next"].(*_state_1)
		x126, q251 := q250.Recv()
		fmt.Printf("%v\n", x126)
		q251.Send("next")
		q252 := q251.ls["next"].(*_state_1)
		x127, q253 := q252.Recv()
		fmt.Printf("%v\n", x127)
		q253.Send("next")
		q254 := q253.ls["next"].(*_state_1)
		x128, q255 := q254.Recv()
		fmt.Printf("%v\n", x128)
		q255.Send("stop")
		q256 := q255.ls["stop"].(*_state_2)
		q256.Recv()
		c.Send(nil)
	}(c)
}