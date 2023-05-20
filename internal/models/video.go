package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Video struct {
	BaseModel
	UserID      uuid.UUID         `json:"userId"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	CaptionUrl  string            `json:"captionUrl"`
	VideoUrl    string            `json:"videoUrl"`
	Language    string            `json:"language"`
	Labels      datatypes.JSONMap `json:"labels"`
	Metadata    datatypes.JSONMap `json:"metadata"`
	Processing  Processing        `json:"processing"`
}

func (v *Video) Merge(video Video) {
	if video.Title != "" {
		v.Title = video.Title
	}
	if video.Description != "" {
		v.Description = video.Description
	}
	if video.CaptionUrl != "" {
		v.CaptionUrl = video.CaptionUrl
	}
	if video.VideoUrl != "" {
		v.VideoUrl = video.VideoUrl
	}
	if video.Language != "" {
		v.Language = video.Language
	}
	if video.Labels != nil {
		v.Labels = video.Labels
	}
	if video.Metadata != nil {
		v.Metadata = video.Metadata
	}
}
