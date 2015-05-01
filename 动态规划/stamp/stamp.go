package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	stamp struct {
		val   int // 面值
		state int // 分配状态(state_x)
	}

	problem struct {
		stamps []*stamp // 邮票
		delta  int      // 差值(>0表示A的总分值较大;<0表示B的总分值较大)
	}
)

const (
	state_A = iota
	state_B
)

func main() {

	rand.Seed(time.Now().Unix())

	cnt := 10
	for idx := 0; idx < cnt; idx++ {
		this := newProblem()
		for this.findBest() {
		}
		this.output()
	}
}

func (this *problem) findBest() (changed bool) {

	minDelta := this.delta
	minIdxA, minIdxB := -1, -1

	exchange := func(idxA, valA, idxB, valB int) {
		delta := this.delta + 2*(valB-valA)
		// 交换后的面值差额更小
		if abs(delta) < abs(minDelta) {
			minDelta = delta
			minIdxA = idxA
			minIdxB = idxB
		}
	}

	for idxA, stA := range this.stamps {

		if stA.state == state_A {

			// -1,0 表示将邮票的分配状态改变,而不是交换两张邮票
			exchange(idxA, stA.val, -1, 0)

			for idxB, stB := range this.stamps {
				if stB.state == state_B {
					exchange(-1, 0, idxB, stB.val)
					// 尝试交换两张邮票
					exchange(idxA, stA.val, idxB, stB.val)
				}
			}
		}
	}

	// 执行交换
	if minIdxA >= 0 || minIdxB >= 0 {
		if minIdxA >= 0 {
			this.stamps[minIdxA].state = state_B
		}
		if minIdxB >= 0 {
			this.stamps[minIdxB].state = state_A
		}
		this.delta = minDelta
		changed = true
	}

	return
}

func (this *problem) output() {
	fmt.Println("\n分配方案: 差值=", abs(this.delta))
	var stra, strb string
	for _, st := range this.stamps {
		if st.state == state_A {
			stra += fmt.Sprintf("%4d", st.val)
		} else if st.state == state_B {
			strb += fmt.Sprintf("%4d", st.val)
		}
	}
	fmt.Println(" A:", stra)
	fmt.Println(" B:", strb)
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func newProblem() *problem {

	this := &problem{}
	cnt := rand.Intn(10) + 2
	this.stamps = make([]*stamp, cnt)

	fmt.Print("\n\n邮票面值: ")
	for idx := range this.stamps {
		st := &stamp{
			val:   rand.Intn(10) + 1,
			state: idx % 2, // 依次设定邮票的初始分配状态为A,B
		}
		// 设置初始分配差额: >0表示A的总分值较大;<0表示B的总分值较大
		if st.state == state_A {
			this.delta += st.val
		} else {
			this.delta -= st.val
		}
		this.stamps[idx] = st
		fmt.Printf("%4d", st.val)
	}

	return this
}
