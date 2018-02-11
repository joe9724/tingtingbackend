// Code generated by go-swagger; DO NOT EDIT.

package version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// VersionListURL generates an URL for the version list operation
type VersionListURL struct {
	PageIndex *int64
	PageSize  *int64
	Queryid   *int64
	Type      *int64
	Userid    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *VersionListURL) WithBasePath(bp string) *VersionListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *VersionListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *VersionListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/version/list"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/TingtingBackend/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var pageIndex string
	if o.PageIndex != nil {
		pageIndex = swag.FormatInt64(*o.PageIndex)
	}
	if pageIndex != "" {
		qs.Set("pageIndex", pageIndex)
	}

	var pageSize string
	if o.PageSize != nil {
		pageSize = swag.FormatInt64(*o.PageSize)
	}
	if pageSize != "" {
		qs.Set("pageSize", pageSize)
	}

	var queryid string
	if o.Queryid != nil {
		queryid = swag.FormatInt64(*o.Queryid)
	}
	if queryid != "" {
		qs.Set("queryid", queryid)
	}

	var typeVar string
	if o.Type != nil {
		typeVar = swag.FormatInt64(*o.Type)
	}
	if typeVar != "" {
		qs.Set("type", typeVar)
	}

	var userid string
	if o.Userid != nil {
		userid = *o.Userid
	}
	if userid != "" {
		qs.Set("userid", userid)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *VersionListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *VersionListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *VersionListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on VersionListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on VersionListURL")
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
func (o *VersionListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
