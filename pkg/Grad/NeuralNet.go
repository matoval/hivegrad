package grad

import (
  "fmt"
  "math/rand"
  "strings"
)

type Neuron struct {
  w []*Grad
  b *Grad
  nonlin bool
}

func NewNeuron(nin int, nonlin bool) *Neuron {
  w := make([]*Grad, nin)
  for i := range w {
    w[i] = New(rand.Float64()*2-1, nil, "")
  }

  return &Neuron{
    w: w,
    b: New(0, nil, ""),
    nonlin: nonlin,
  }
}

func (n *Neuron) Call(x []*Grad) *Grad {
  act := n.b
  for i, wi := range n.w {
    act  = act.Add(wi.Mul(x[i]))
  }

  if n.nonlin {
    return act.Relu()
  }
  return act
}

func (n *Neuron) Parameters() []*Grad  {
  return append(n.w, n.b)
}

func (n *Neuron) String() string  {
  if n.nonlin {
    return fmt.Sprintf("ReLUNeuron(%d)", len(n.w))
  }
  return fmt.Sprintf("LinearNeuron(%d)", len(n.w))
}

type Layer struct {
  neurons []*Neuron
}

func NewLayer(nin, nout int, nonlin bool) *Layer  {
  neurons := make([]*Neuron, nout)
  for i:= range neurons {
    neurons[i] = NewNeuron(nin, nonlin)
  }
  return &Layer{neurons: neurons}
}

func (l *Layer) Call(x []*Grad) []*Grad  {
  out := make([]*Grad, len(l.neurons))
  for i, n := range l.neurons {
    out[i] =n.Call(x)
  }

  return out
}

func (l *Layer) Parameters() []*Grad  {
  var params []*Grad
  for _, n := range l.neurons {
    params = append(params, n.Parameters()...)
  }
  return params 
}

func (l *Layer) String() string  {
  neurons := make([]string, len(l.neurons))
  for i, n := range l.neurons {
    neurons[i] = n.String()
  }

  return fmt.Sprintf("Layer of [%s]", strings.Join(neurons, ", "))
  
}

type MLP struct {
  layers []*Layer
}

func NewMLP(nin int, nouts []int) *MLP  {
  sizes := append([]int{nin}, nouts...)
  layers := make([]*Layer, len(nouts))
  for i:= range layers {
    layers[i] = NewLayer(sizes[i], sizes[i+1], i!= len(nouts)-1)
  }
  return &MLP{layers: layers}
}

func (m *MLP) Call(x []*Grad) []*Grad  {
  for _, layer := range m.layers {
    x = layer.Call(x)
  }
  return x 
}

func (m *MLP) Parameters() []*Grad  {
  var params []*Grad
  for _, layer := range m.layers {
    params = append(params, layer.Parameters()...)
  }
  return params 
}

func (m *MLP) String() string  {
  layers := make([]string, len(m.layers))
  for i, layer :=  range m.layers {
    layers[i] = layer.String()
  }
  return fmt.Sprintf("MLP of [%s]", strings.Join(layers, ", "))
}

