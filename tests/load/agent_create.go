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

// CreateAgent creates PURLs
type CreateAgent struct {
	Id int

	Domain         string
	CreateInterval time.Duration

	API dsl.AdminAPI
}

func NewCreateAgent(id int, domain string, createInterval time.Duration, API dsl.AdminAPI) *CreateAgent {
	return &CreateAgent{Id: id, Domain: domain, CreateInterval: createInterval, API: API}
}

func (a *CreateAgent) Run(t *testing.T, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	dsl.GivenExistingDomain(t, a.API, a.Domain)

	target, _ := url.Parse("https://google.com")
	i := 0

	for {
		name := fmt.Sprintf("purl-%d", i)

		err := a.API.SavePURL(dsl.NewPURL(a.Domain, name, target))
		require.NoError(t, err)

		i += 1
		select {
		case <-time.After(a.CreateInterval):
			break
		case <-done:
			return
		}
	}
}
