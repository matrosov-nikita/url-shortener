package handler

import (
	"context"
	"fmt"
	"url-shortener/storage"
)

type Handler struct {
	storage storage.Storage
}

type Request struct {
	Payload string
}

// New instance of handler
func New(storage storage.Storage) *Handler {
	return &Handler{storage}
}

func (h *Handler) Hello(ctx context.Context, r *Request, w *Request) error {
	fmt.Println("here")
	return nil
}
