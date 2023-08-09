package constants

import "errors"

var InvalidProductIdError = errors.New("invalid product id")
var ProductIsAlreadyExistsError = errors.New("product is already exists")
var ProductDoesNotExistsInCartError = errors.New("product does not exists in cart")
var InvalidIndexError = errors.New("invalid index")
