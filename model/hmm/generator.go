package hmm

import (
	"fmt"
	"github.com/akualab/gjoa/model"
	"math/rand"
)

type Generator struct {
	hmm *HMM
	r   *rand.Rand
}

func NewGenerator(hmm *HMM, seed int64) (gen *Generator) {
	r := rand.New(rand.NewSource(seed))
	gen = &Generator{
		hmm: hmm,
		r:   r,
	}
	return gen
}

// Given n, the length of the seq, generates random sequence
// for a given hmm.
func (gen *Generator) Next(n int) ([][]float64, []int, error) {

	obs := make([][]float64, n)
	states := make([]int, n)
	r := gen.r
	logDist := gen.hmm.logInitProbs
	state0, err0 := model.RandIntFromLogDist(logDist, r)
	if err0 != nil {
		return nil, nil, fmt.Errorf("Error calling RandIntFromLogDist")
	}
	for i := 0; i < n; i++ {
		states[i] = state0
		g := gen.hmm.obsModels[state0]
		x, _, err0 := g.Random(r)
		obs[i] = x.([]float64)
		if err0 != nil {
			return nil, nil, fmt.Errorf("Error generating Random model")
		}
		dist := gen.hmm.logTransProbs[state0]
		state, err := model.RandIntFromLogDist(dist, r)
		if err != nil {
			return nil, nil, fmt.Errorf("Error calling RandIntFromLogDist")
		}
		state0 = state
	}
	return obs, states, nil
}
