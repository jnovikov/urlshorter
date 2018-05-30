package url

import (
	"net/http"
	"path"
	"net/url"
	"fmt"
)

func ParsePath(r *http.Request) string {
	return path.Base(r.URL.Path)
}

func Absolute(host, port, relative string) string {
	r := new(url.URL)
	r.Scheme = "http"
	if port == "80" {
		r.Host = host
	} else {
		r.Host = fmt.Sprintf("%s:%s", host, port)
	}
	r.Path = relative
	return r.String()
}
