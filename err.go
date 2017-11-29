package main

import "errors"

var (
	errMissingName   = errors.New("missing alias name")
	errMissingValue  = errors.New("missing alias value")
	errNoSuchAlias   = errors.New("no such alias")
	errNoChanges     = errors.New("no changes")
	errInvalidAction = errors.New("invalid action. type -h or --help to see usage")
)
