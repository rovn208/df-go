package util

import "errors"

var InputTypeMissingError = errors.New("at least one of input type is set")
var MoreThanOneInputTypeError = errors.New("more than one input type is set")
