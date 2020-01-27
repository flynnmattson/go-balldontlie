package balldontlie

import (
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

func addOptions(path string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return path, nil
	}

	pathURL, err := url.Parse(path)
	if err != nil {
		return path, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return path, err
	}

	pathURL.RawQuery = qs.Encode()
	return pathURL.String(), nil
}
