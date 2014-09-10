package anomalize

import (
	"github.com/bmizerany/assert"
	"github.com/drewlanenga/govector"
	"math/rand"
	"testing"
)

// Generates a random walk given the number of steps desired, a starting point,
// and the desired standard deviation.
func randomWalk(nsteps int, start float64, sd float64) (govector.Vector, error) {
	walk := make(govector.Vector, nsteps)
	walk[0] = float64(start)

	i := 1
	for i < nsteps {
		step := rand.NormFloat64() * sd
		walk[i] = cap(walk[i-1]+step, 0.0, 1.0)

		i++
	}
	return walk, nil
}

func TestAnomalizer(t *testing.T) {
	conf := &AnomalizerConf{
		Sensitivity:   0.5,
		UpperBound:    5,
		LowerBound:    0,
		ActiveSize:    1,
		ReferenceSize: 4,
		Methods:       []string{"diff", "fence", "rank", "magnitude"},
	}

	// initialize with empty data now
	data := []float64{0.1, 2.05, 1.5, 2.5, 2.6, 2.55}

	anomalizer, err := NewAnomalizer(conf, data)
	assert.Equal(t, nil, err, "Error initializing new anomalizer")

	prob := anomalizer.Push(8.0)
	assert.Tf(t, prob > 0.5, "Anomalizer returned a probability that was too small")
}
