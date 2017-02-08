package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// DeleteMachineURL generates an URL for the delete machine operation
type DeleteMachineURL struct {
	UUID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteMachineURL) WithBasePath(bp string) *DeleteMachineURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteMachineURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteMachineURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/machines/{uuid}"

	uuid := o.UUID
	if uuid != "" {
		_path = strings.Replace(_path, "{uuid}", uuid, -1)
	} else {
		return nil, errors.New("UUID is required on DeleteMachineURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v3"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteMachineURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteMachineURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteMachineURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteMachineURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteMachineURL")
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
func (o *DeleteMachineURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}