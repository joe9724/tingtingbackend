// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// OrderListURL generates an URL for the order list operation
type OrderListURL struct {
	AlbumName *string
	EndTime   *int64
	MemberID  *int64
	OrderNo   *string
	PageIndex *int64
	PageSize  *int64
	PayType   *int64
	StartTime *int64
	Userid    *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *OrderListURL) WithBasePath(bp string) *OrderListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *OrderListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *OrderListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/order/list"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/TingtingBackend/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var albumName string
	if o.AlbumName != nil {
		albumName = *o.AlbumName
	}
	if albumName != "" {
		qs.Set("albumName", albumName)
	}

	var endTime string
	if o.EndTime != nil {
		endTime = swag.FormatInt64(*o.EndTime)
	}
	if endTime != "" {
		qs.Set("endTime", endTime)
	}

	var memberID string
	if o.MemberID != nil {
		memberID = swag.FormatInt64(*o.MemberID)
	}
	if memberID != "" {
		qs.Set("memberId", memberID)
	}

	var orderNo string
	if o.OrderNo != nil {
		orderNo = *o.OrderNo
	}
	if orderNo != "" {
		qs.Set("orderNo", orderNo)
	}

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

	var payType string
	if o.PayType != nil {
		payType = swag.FormatInt64(*o.PayType)
	}
	if payType != "" {
		qs.Set("payType", payType)
	}

	var startTime string
	if o.StartTime != nil {
		startTime = swag.FormatInt64(*o.StartTime)
	}
	if startTime != "" {
		qs.Set("startTime", startTime)
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
func (o *OrderListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *OrderListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *OrderListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on OrderListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on OrderListURL")
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
func (o *OrderListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
