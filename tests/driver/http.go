package driver

import (
	"github.com/fabiante/persurl/tests/dsl"
	"net/http"
	"net/url"
)

type HTTPDriver struct {
	BasePath string
	Client   *http.Client
}

func NewHTTPDriver(basePath string, client *http.Client) *HTTPDriver {
	return &HTTPDriver{BasePath: basePath, Client: client}
}

func (H *HTTPDriver) ResolvePURL(domain string, name string) (*url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (H *HTTPDriver) CreatePurl(purl *dsl.PURL) error {
	//TODO implement me
	panic("implement me")
}
