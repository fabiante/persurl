package driver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/fabiante/persurl/api/res"
	"github.com/fabiante/persurl/app"
	"github.com/fabiante/persurl/tests/dsl"
)

type HTTPDriver struct {
	BasePath string
	Client   *http.Client
}

func NewHTTPDriver(basePath string, transport http.RoundTripper) *HTTPDriver {
	client := &http.Client{
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// This prevents following any HTTP redirects.
			// We don't want to follow because we want to assert them
			return http.ErrUseLastResponse
		},
	}
	return &HTTPDriver{BasePath: basePath, Client: client}
}

func (driver *HTTPDriver) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	userKey := getUserKeyFromCtx(ctx)
	if userKey != "" {
		req.Header.Set("Persurl-Key", userKey)
	}

	return req, nil
}

func (driver *HTTPDriver) ResolvePURL(ctx context.Context, domain string, name string) (*url.URL, error) {
	req, err := driver.newRequest(ctx, http.MethodGet, fmt.Sprintf("%s/r/%s/%s", driver.BasePath, domain, name), nil)
	if err != nil {
		return nil, err
	}

	response, err := driver.Client.Do(req)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusFound:
		break
	case http.StatusNotFound:
		return nil, fmt.Errorf("%w: status %d returned", app.ErrNotFound, response.StatusCode)
	default:
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	loc, err := response.Location()
	if err != nil {
		return nil, fmt.Errorf("invalid location: %s", err)
	}

	return loc, nil
}

func (driver *HTTPDriver) SavePURL(ctx context.Context, purl *dsl.PURL) (string, error) {
	body := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(body).Encode(map[string]string{
		"target": purl.Target.String(),
	})
	if err != nil {
		return "", err
	}
	req, err := driver.newRequest(ctx, http.MethodPut, fmt.Sprintf("%s/a/domains/%s/purls/%s", driver.BasePath, purl.Domain, purl.Name), body)
	if err != nil {
		return "", err
	}

	response, err := driver.Client.Do(req)
	if err != nil {
		return "", err
	}

	switch response.StatusCode {
	case http.StatusOK:
		break
	case http.StatusBadRequest:
		return "", fmt.Errorf("%w: status %d returned", app.ErrBadRequest, response.StatusCode)
	default:
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var r res.SavePURLResponse
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		return "", fmt.Errorf("decoding response body failed: %w", err)
	}

	return r.Path, nil
}

func (driver *HTTPDriver) CreateDomain(ctx context.Context, name string) error {
	req, err := driver.newRequest(ctx, http.MethodPost, fmt.Sprintf("%s/a/domains/%s", driver.BasePath, name), nil)
	if err != nil {
		return err
	}

	response, err := driver.Client.Do(req)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("%w: status %d returned", app.ErrBadRequest, response.StatusCode)
	default:
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
}

const ctxKeyUserKey = "auth-user-key"

func getUserKeyFromCtx(ctx context.Context) string {
	value := ctx.Value(ctxKeyUserKey)
	if value == nil {
		return ""
	}
	return value.(string)
}

func CtxWithUserKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, ctxKeyUserKey, key)
}
