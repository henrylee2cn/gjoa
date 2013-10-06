package model

import (
	"encoding/json"
	"github.com/golang/glog"
	"io"
	"math/rand"
)

// Returns a random observations generated by the model.
// obs is the random observation.
// sequence applies to Markov models.
type Generator interface {
	Random(r *rand.Rand) (obs interface{}, sequence []int, e error)
}

// A read-only model.
type Modeler interface {
	Prob(obs interface{}) float64
	LogProb(obs interface{}) float64
	Name() string
	//String() string
	NumElements() int
	Trainable() bool
	Write(w io.Writer) error
	Values() interface{}
	Generator
}

// A trainable model.
type Trainer interface {
	Modeler

	Update(a []float64, w float64) error
	Estimate() error
	Clear() error
	SetName(name string)
	NumSamples() float64
}

// A trainable sequence model.
type SequenceTrainer interface {
	Modeler

	Update(seq [][]float64, w float64) error
	Estimate() error
	Clear() error
	SetName(name string)
	NumSamples() float64
}

// Implements basic functionality for models. Model implementations can embed
// this type to us the methods.
type Base struct{}

func (base *Base) WriteModel(w io.Writer, m Modeler) error {

	b, err := json.Marshal(m.Values())
	if err != nil {
		glog.Fatal(err)
	}
	_, e := w.Write(b)
	return e
}
