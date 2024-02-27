package meta

import "errors"

var (
	ErrGoVersionNotSupported      = errors.New("go version not supported")
	ErrConfigFileTypeNotSupported = errors.New("config file type not supported")
	ErrMarshalConfig              = errors.New("marshal config")
)
