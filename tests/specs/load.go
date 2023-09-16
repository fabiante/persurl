package specs

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/fabiante/persurl/tests/dsl"
	"github.com/fabiante/persurl/tests/load"
)

// TestLoad creates various stages of load on the given dsl.API.
//
// It does this by starting one or more agents which interact with the dsl.API.
//
// The agents are started in parallel and run for a given duration. This is done iteratively
// with an increasing amount of agents and thus load.
func TestLoad(t *testing.T, api dsl.API) {
	t.Run("load", func(t *testing.T) {
		tests := []struct {
			CreateAgents   int
			CreateInterval time.Duration
			Duration       time.Duration
		}{
			{
				CreateAgents:   1,
				CreateInterval: 50 * time.Millisecond,
				Duration:       3 * time.Second,
			},
			{
				CreateAgents:   2,
				CreateInterval: 50 * time.Millisecond,
				Duration:       3 * time.Second,
			},
			{
				CreateAgents:   5,
				CreateInterval: 50 * time.Millisecond,
				Duration:       3 * time.Second,
			},
			{
				CreateAgents:   15,
				CreateInterval: 50 * time.Millisecond,
				Duration:       3 * time.Second,
			},
			{
				CreateAgents:   25,
				CreateInterval: 50 * time.Millisecond,
				Duration:       3 * time.Second,
			},
			{
				CreateAgents:   50,
				CreateInterval: 50 * time.Millisecond,
				Duration:       3 * time.Second,
			},
		}

		for i, test := range tests {
			t.Run(fmt.Sprintf("tests[%d] create:%d", i, test.CreateAgents), func(t *testing.T) {
				ctx, cancel := context.WithCancel(context.Background())
				done := ctx.Done()
				wg := &sync.WaitGroup{}

				for j := 0; j < test.CreateAgents; j++ {
					agent := load.NewCreateAgent(j, fmt.Sprintf("agent-%d-%d", i, j), test.CreateInterval, api)
					wg.Add(1)
					go agent.Run(t, done, wg)
				}

				time.Sleep(test.Duration)
				cancel()
				wg.Wait()
			})
		}
	})
}
