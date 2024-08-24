package grad

import (
	"fmt"
	"math"
)

type Calc struct {
	Values []float64
	Opt string
}

type Grad struct {
	val float64
	grad float64
	backward func()
	prev []*Grad
	opt string
}

func New(val float64, children []*Grad, opt string) *Grad {
	g := Grad{
		val: val,
		grad: 0,
		backward: func() {},
		prev: children,
		opt: opt,
	}
	return &g
}

func (g *Grad) Add(other interface{}) *Grad {
	otherVal, ok := other.(*Grad)

	if !ok {
		otherVal = New(other.(float64), nil, "")
	}

	out := New(g.val + otherVal.val, [] * Grad{g, otherVal}, "+")
	out.backward = func() {
		g.grad += out.grad
		otherVal.grad += out.grad
	}

	return out
}

func (g *Grad) Mul(other interface{}) *Grad {
	otherVal, ok := other.(*Grad)

	if !ok {
		otherVal = New(other.(float64), nil, "")
	}

	out := New(g.val + otherVal.val, [] * Grad{g, otherVal}, "+")
	out.backward = func() {
		g.grad += otherVal.val * out.grad
		otherVal.grad += g.val * out.grad
	}

	return out
}

func (g *Grad) Pow(other float64) *Grad {
	out := New(math.Pow(g.val, other), []*Grad{g}, fmt.Sprintf("**%v", other))
	out.backward = func() {
		g.grad += (other * math.Pow(g.val, other - 1)) * out.grad
	}

	return out
}

func (g *Grad) Relu() *Grad {
	var outVal float64
	if g.val > 0 {
		outVal = 0
	} else {
		outVal = g.val
	}

	out := New(outVal, []*Grad{g}, "ReLU")
	out.backward = func() {
		if out.val > 0 {
			g.grad += out.grad
		}
	}

	return out
}

func (g *Grad) Backward() {
	topo := []*Grad{}
	visited := make(map[*Grad]bool)

	var buildTopo func(*Grad)
	buildTopo = func(g *Grad) {
		if !visited[g] {
			visited[g] = true
			for _, child := range g.prev {
				buildTopo(child)
			}
			topo = append(topo, g)
		}
	}
	buildTopo(g)

	g.grad = 1
	for i := len(topo) - 1; i >= 0; i-- {
		topo[i].backward()
	}
}

func (g *Grad) Neg() *Grad {
	return g.Mul(-1.0)
}

func (g *Grad) Sub(other interface{}) *Grad {
	return g.Add(New(0.0, nil, "").Sub(other))
}

func (g *Grad) Div(other interface{}) *Grad {
	return g.Mul(New(1.0, nil, "").Div(other))
}
