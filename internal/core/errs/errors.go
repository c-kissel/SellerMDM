package errs

import "errors"

// ErrNotFound is returned when the operation did not return any documents.
var ErrNotFound = errors.New("storage: no documents found in DB")

// ErrNameEmpty is returned when given document contains empty Name attribute.
var ErrNameEmpty = errors.New("storage: name attr of document is empty")

// ErrSKUEmpty is returned when given document contains empty Name attribute.
var ErrSKUEmpty = errors.New("storage: SKU attr of document is empty")

// ErrEmptyInput is returned when given input data is empty (typically an empty array).
var ErrEmptyInput = errors.New("input data is empty")
