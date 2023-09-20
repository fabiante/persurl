package load

import (
	"fmt"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/fabiante/persurl/tests/dsl"
	"github.com/stretchr/testify/require"
)

// CreateAgent creates PURLs in a given domain with a given interval.
//
// Use this agent to simulate a user continuously generating PURLs.
type CreateAgent struct {
	Id int

	Domain         string
	CreateInterval time.Duration

	API dsl.AdminAPI
}

func NewCreateAgent(id int, domain string, createInterval time.Duration, API dsl.AdminAPI) *CreateAgent {
	return &CreateAgent{Id: id, Domain: domain, CreateInterval: createInterval, API: API}
}

// Run starts this agent's work loop.
//
// Run ensures that the agent's Domain exists before starting to create PURLs.
//
// The agent runs until a message is sent to the done channel. It will then decrement the given wait group.
//
// You should use this method in a dedicated goroutine as this is a blocking function.
func (a *CreateAgent) Run(t *testing.T, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	dsl.GivenExistingDomain(t, a.API, a.Domain)

	target, _ := url.Parse("https://google.com")
	i := 0

	for {
		name := fmt.Sprintf("purl-%d", i)

		path, err := a.API.SavePURL(dsl.NewPURL(a.Domain, name, target))
		require.NoError(t, err)
		require.NotEmpty(t, path)

		i += 1
		select {
		case <-time.After(a.CreateInterval):
			break
		case <-done:
			return
		}
	}
}
