package main

import (
	"github.com/kiyutink/sowhenthen/storage"
)

type Controller struct {
	storage storage.Storage
}

func newController(s storage.Storage) *Controller {
	return &Controller{s}
}
