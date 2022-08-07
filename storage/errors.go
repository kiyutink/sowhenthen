package storage

import "fmt"

type NotFoundError struct {
	Identifier string
	Err        error
}

func (nfe *NotFoundError) Error() string {
	return fmt.Sprintf("\"%v\" not found: %v", nfe.Identifier, nfe.Err)
}

func (nfe *NotFoundError) Unwrap() error {
	return nfe.Err
}

type InvalidRequestError struct {
	Message string
	Err     error
}

func (ir *InvalidRequestError) Error() string {
	return fmt.Sprintf("the storage request is invalid: %v", ir.Message)
}

func (ir *InvalidRequestError) Unwrap() error {
	return ir.Err
}
