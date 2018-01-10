package handler

import (
	"context"
	"url-shortener/encoder"
	"url-shortener/storage"
)

// URLHandler contains two endpoints to encode/decode URL
type URLHandler struct {
	cacher  storage.Cacher
	storage storage.Storage
}

// Request type that contains URL field only
type Request struct {
	URL string `json:"url,omitempty"`
}

// EncodedResponse type that contains short version of URL
type EncodedResponse struct {
	ShortURL string `json:"shortUrl,omitempty"`
}

// DecodedResponse type that contains short version of URL
type DecodedResponse struct {
	OriginURL string `json:"originUrl,omitempty"`
}

// New instance of handler
func New(cache storage.Cacher, storage storage.Storage) *URLHandler {
	return &URLHandler{cache, storage}
}

// Encode endpoint to encode given URL and returns shortened one
func (h *URLHandler) Encode(ctx context.Context, r *Request, w *EncodedResponse) error {
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

	if err = h.cacher.SetUniqueKey("key"); err != nil {
		return err
	}

	w.ShortURL = short
	return nil
}

// Decode returns original URL
func (h *URLHandler) Decode(ctx context.Context, r *Request, w *DecodedResponse) error {
	origin, err := h.storage.GetURL(r.URL)
	if err != nil {
		return err
	}

	w.OriginURL = origin
	return nil
}
