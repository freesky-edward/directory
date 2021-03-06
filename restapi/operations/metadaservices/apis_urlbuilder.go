// Code generated by go-swagger; DO NOT EDIT.

package metadaservices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ApisURL generates an URL for the apis operation
type ApisURL struct {
	Keywords *string
	Length   *int64
	Start    *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ApisURL) WithBasePath(bp string) *ApisURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ApisURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ApisURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/apis"

	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var keywords string
	if o.Keywords != nil {
		keywords = *o.Keywords
	}
	if keywords != "" {
		qs.Set("keywords", keywords)
	}

	var length string
	if o.Length != nil {
		length = swag.FormatInt64(*o.Length)
	}
	if length != "" {
		qs.Set("length", length)
	}

	var start string
	if o.Start != nil {
		start = swag.FormatInt64(*o.Start)
	}
	if start != "" {
		qs.Set("start", start)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ApisURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ApisURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ApisURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ApisURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ApisURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ApisURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
