package models

import "github.com/google/uuid"

type Processing struct {
	BaseModel
	VideoID     uuid.UUID `json:"videoId"`
	Status      string    `json:"status"`
	Steps       []ProcessingStep
	CurrentStep string `json:"currentStep"`
}

type ProcessingStep struct {
	BaseModel
	ProcessingID uuid.UUID `json:"processingId"`
	Status       string    `json:"status"`
	Step         string    `json:"step"`
}
