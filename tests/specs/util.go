package specs

import "net/url"

func mustParseURL(s string) *url.URL {
	x, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return x
}
