// Package handler provider implementation for RPC handler with two endpoints
package handler

import (
	"context"
	"errors"

	"github.com/url-shortener/encoder"
	"github.com/url-shortener/storage"
)

// UrlHandler contains two endpoints to encode/decode URL
type UrlHandler struct {
	cacher  storage.Cacher
	storage storage.Storage
}

// Request type that contains URL field
type Request struct {
	URL string `json:"url,omitempty"`
}

// EncodedResponse type that contains short version of URL
type EncodedResponse struct {
	ShortURL string `json:"shortUrl"`
}

// DecodedResponse type that contains original version of URL
type DecodedResponse struct {
	OriginURL string `json:"originUrl"`
}

// New instance of RPC handler
func New(cache storage.Cacher, storage storage.Storage) *UrlHandler {
	return &UrlHandler{cache, storage}
}

// Encode given URL and return short version
func (h *UrlHandler) Encode(ctx context.Context, r *Request, w *EncodedResponse) error {
	var count int64
	count, err := h.cacher.GetUniqueKey("key")
	if err != nil {
		if count, err = h.storage.Count(); err != nil {
			return err
		}
	}

	short := encoder.Encode(count)
	if err := h.storage.AddURL(short, r.URL); err != nil {
		return err
	}

	//TODO: Redis key set as env variable
	if err = h.cacher.IncrUniqueKey("key"); err != nil {
		return err
	}

	w.ShortURL = short
	return nil
}

// Decode returns original URL
func (h *UrlHandler) Decode(ctx context.Context, r *Request, w *DecodedResponse) error {
	origin, err := h.storage.GetURL(r.URL)
	if err != nil {
		return errors.New("404")
	}

	w.OriginURL = origin
	return nil
}
