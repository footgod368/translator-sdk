package utils

import (
	"net/url"
)

func AppendURLParams(urlStr string, params map[string]string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}

	query := u.Query()
	for key, value := range params {
		if !query.Has(key) {
			query.Add(key, value)
		}
	}
	u.RawQuery = query.Encode()

	return u.String()
}
