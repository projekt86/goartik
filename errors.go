package goartik

import "errors"

var (
	// ErrNoInited : not inited error
	errNoInited = errors.New("You must first call Init function")
)
