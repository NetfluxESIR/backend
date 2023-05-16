package models

import "github.com/google/uuid"

type Video struct {
	BaseModel
	UserID      uuid.UUID         `json:"userId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	CaptionUrl  string            `json:"captionUrl"`
	VideoUrl    string            `json:"videoUrl"`
	Language    string            `json:"language"`
	Labels      map[string]string `json:"labels"`
	Metadata    map[string]string `json:"metadata"`
	Processing  Processing        `json:"processing"`
}
