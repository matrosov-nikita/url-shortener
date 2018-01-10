package handler

import (
	"context"
	"url-shortener/shortener"
	"url-shortener/storage"
)

// URLHandler contains two endpoints to encode/decode URL
type URLHandler struct {
	storage storage.Storage
}

// Request type that contains URL field only
type Request struct {
	URL string
}

// Response type that contains shortened version of URL
type Response struct {
	ShortedURL string
}

// New instance of handler
func New(storage storage.Storage) *URLHandler {
	return &URLHandler{storage}
}

// Encode endpoint to encode given URL and returns shortened one
func (h *URLHandler) Encode(ctx context.Context, r *Request, w *Response) error {
	key, err := h.storage.GetUniqueKey()
	if err != nil {
		return err
	}
	w.ShortedURL = shortener.Encode(r.URL, key)
	return nil
}

// Decode endpoint to encode given URL and returns shortened one
func (h *URLHandler) Decode(ctx context.Context, r *Request, w *Response) error {
	return nil
}
