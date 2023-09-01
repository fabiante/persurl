package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fabiante/persurl/tests/dsl"
)

type HTTPDriver struct {
	BasePath string
	Client   *http.Client
}

func NewHTTPDriver(basePath string, client *http.Client) *HTTPDriver {
	return &HTTPDriver{BasePath: basePath, Client: client}
}

func (driver *HTTPDriver) ResolvePURL(domain string, name string) (*url.URL, error) {
	res, err := driver.Client.Get(driver.purlPath(domain, name))
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusFound:
		break
	case http.StatusNotFound:
		return nil, fmt.Errorf("%w: status %d returned", dsl.ErrNotFound, res.StatusCode)
	default:
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	loc, err := res.Location()
	if err != nil {
		return nil, fmt.Errorf("invalid location: %s", err)
	}

	return loc, nil
}

func (driver *HTTPDriver) purlPath(domain string, name string) string {
	return fmt.Sprintf("%s/r/%s/%s", driver.BasePath, domain, name)
}

func (driver *HTTPDriver) adminPath(domain string, name string) string {
	return fmt.Sprintf("%s/a/%s/%s", driver.BasePath, domain, name)
}

func (driver *HTTPDriver) CreatePurl(purl *dsl.PURL) error {
	body := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(body).Encode(map[string]string{
		"target": purl.Target.String(),
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, driver.adminPath(purl.Domain, purl.Name), body)
	if err != nil {
		return err
	}

	res, err := driver.Client.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
}
