package bitree

import(
	"fmt"
)

const (
	LH = +1
	EH = 0
	RH = -1
)

type BBiTree struct {
	Data int
	BF int8
	Left *BBiTree
	Right *BBiTree
}

func (t *BBiTree)LRotate() *BBiTree {
	rc := t.Right
	t.Right = rc.Left
	rc.Left = t
	t = rc
	return t
}

func (t *BBiTree)RRotate() *BBiTree {
	lc := t.Left
	t.Left = lc.Right
	lc.Right = t
	t = lc
	return t
}

func (t *BBiTree)LBalance() *BBiTree {
	lc := t.Left
	switch lc.BF {
	case LH:
		lc.BF = EH
		t.BF = EH
		t = t.RRotate()
	case RH:
		rg := lc.Right
		switch rg.BF {
		case LH:
			lc.BF = EH
			t.BF = RH
		case RH:
			lc.BF = LH
			t.BF = EH
		}
		rg.BF = EH
		t.Left = lc.LRotate()
		t = t.RRotate()
	}
	return t
}

func (t *BBiTree)RBalance() *BBiTree{
	rc := t.Right
	switch rc.BF {
	case LH:
		lg := rc.Left
		switch lg.BF {
		case LH:
			rc.BF = RH
			t.BF = EH
		case RH:
			rc.BF = EH
			t.BF = LH
		}
		lg.BF = EH
		t.Right = rc.RRotate()
		t = t.LRotate()
	case RH:
		rc.BF = EH
		t.BF = EH
		t = t.LRotate()
	}
	return t
}

func (t *BBiTree)Insert(data int) (*BBiTree, bool) {
	taller := false
	if t==nil {
		t = new(BBiTree)
		t.Data = data
		t.BF = EH
		taller = true
	} else if data<t.Data {
		t.Left, taller = t.Left.Insert(data)
		if taller {
			switch t.BF {
			case LH:
				t = t.LBalance()
				taller = false
			case EH:
				t.BF = LH
				taller = true
			case RH:
				t.BF = EH
				taller = false
			}
		}
	} else if data>t.Data {
		t.Right, taller = t.Right.Insert(data)
		if taller {
			switch t.BF {
			case LH:
				t.BF = EH
				taller = false
			case EH:
				t.BF = RH
				taller = true
			case RH:
				t = t.RBalance()
				taller = false
			}
		}
	}
	return t, taller
}

func (t *BBiTree)Delete() {

}

func (t *BBiTree)PreTraverse() {
	/*
	if t==nil { return }
	stack := make([]*BBiTree, 0)
	stack = append(stack, t)
	for len(stack)>0 {
		t = stack[len(stack)-1]
		for t!=nil {
			stack = append(stack, t.Left)
			t = t.Left
		}
		stack = stack[:len(stack)-1]
		if len(stack)>0 {
			t = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%d ", t.Data)
			stack = append(stack, t.Right)
		}
	}
	fmt.Printf("\n")
	*/
	if t==nil { return }
	stack := make([]*BBiTree, 0)
	for t!=nil||len(stack)>0 {
		if t!=nil {
			stack = append(stack, t)
			t = t.Left
		} else {
			t = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%d ", t.Data)
			t = t.Right
		}
	}
	fmt.Printf("\n")
}