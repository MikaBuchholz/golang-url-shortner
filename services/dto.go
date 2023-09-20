package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UploadUrlDTO struct {
	Payload string `json:"payload"`
}

type UploadUrlResponseDTO struct {
	ShortenedUrl string `json:"id"`
	ViewCount    int32  `json:"view_count"`
}

type GetUrlDTO struct {
	ShortenedUrl string
}

type GetUrlResponseDTO struct {
	TrueUrl   string `json:"true_url"`
	ViewCount int32  `json:"view_count"`
}

func (a *GetUrlDTO) ReadAndValidate(r *http.Request) error {
	shrt_url_id := chi.URLParam(r, "id")

	if shrt_url_id == "" {
		return fmt.Errorf("no checksum provided")
	}

	a.ShortenedUrl = shrt_url_id

	return a.Validate()
}

func (data *GetUrlDTO) Validate() error {
	if data.ShortenedUrl == "" {
		return fmt.Errorf("payload cannot be empty")
	}
	return nil
}

func (a *UploadUrlDTO) ReadAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(a)

	if err != nil {
		return fmt.Errorf("body is empty")
	}

	return a.Validate()
}

func (data *UploadUrlDTO) Validate() error {
	if data.Payload == "" {
		return fmt.Errorf("payload cannot be empty")
	}
	return nil
}
