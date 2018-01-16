// Package handler provides implementation for RPC handler with two endpoints.
package handler

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/matrosov-nikita/url-shortener/encoder"
	"github.com/matrosov-nikita/url-shortener/storage"
)

// URLHandler contains two endpoints to encode/decode URL.
type URLHandler struct {
	cacher  storage.Cacher
	storage storage.Storage
}

// New instance of RPC handler.
func New(cache storage.Cacher, storage storage.Storage) *URLHandler {
	return &URLHandler{cache, storage}
}

// URLRequest type with URL field.
type URLRequest struct {
	URL string `json:"url"`
}

// EncodedResponse type with ShortURL field.
type EncodedResponse struct {
	ShortURL string `json:"shortUrl"`
}

// DecodedResponse type with OriginURL field.
type DecodedResponse struct {
	OriginalURL string `json:"originUrl"`
}

// Encode method encodes given URL and return short version.
func (h *URLHandler) Encode(ctx context.Context, r *URLRequest, w *EncodedResponse) error {
	if !govalidator.IsURL(r.URL) {
		return errors.New("URL is not valid")
	}

	var count int64
	count, err := h.cacher.GetUniqueKey()
	if err != nil {
		if count, err = h.storage.Count(); err != nil {
			return err
		}
	}

	short := encoder.Encode(count)
	if err := h.storage.AddURL(short, r.URL); err != nil {
		return err
	}

	if err = h.cacher.IncrUniqueKey(); err != nil {
		return err
	}

	w.ShortURL = short
	return nil
}

// Decode method decodes short URL and returns original URL version.
func (h *URLHandler) Decode(ctx context.Context, r *URLRequest, w *DecodedResponse) error {
	origin, err := h.storage.GetURL(r.URL)
	w.OriginalURL = origin
	return err
}
