package utils

import (
	"net/url"
	"strings"
)

type URLBuilder struct {
	scheme string
	host   string
	path   []string
	query  url.Values
}

func NewURLBuilder(sceme, host string) *URLBuilder {
	return &URLBuilder{
		scheme: sceme,
		host:   host,
	}
}

func (u *URLBuilder) AddPath(path string) *URLBuilder {
	u.path = append(u.path, path)
	return u
}

func (u *URLBuilder) AddQuery(key, value string) *URLBuilder {
	if u.query == nil {
		u.query = url.Values{}
	}
	u.query.Add(key, value)
	return u
}

func (u *URLBuilder) Build() string {
	path := strings.Join(u.path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	link := &url.URL{
		Scheme:   u.scheme,
		Host:     u.host,
		Path:     path,
		RawQuery: u.query.Encode(),
	}

	return link.String()
}
