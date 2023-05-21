package server

import (
	"github.com/NetfluxESIR/backend/internal/models"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func fromVideoAPIModel(video gen.Video) (models.Video, error) {
	labels := datatypes.JSONMap{}
	if video.Labels != nil {
		err := labels.Scan(video.Labels)
		if err != nil {
			return models.Video{}, err
		}
	}
	metadata := datatypes.JSONMap{}
	if video.Metadata != nil {
		err := metadata.Scan(video.Metadata)
		if err != nil {
			return models.Video{}, err
		}
	}
	if video.CaptionUrl == nil {
		video.CaptionUrl = new(string)
	}
	if video.Id == nil {
		video.Id = &uuid.Nil
	}
	if video.Language == nil {
		video.Language = new(string)
	}
	if video.VideoUrl == nil {
		video.VideoUrl = new(string)
	}
	return models.Video{
		CaptionUrl:  *video.CaptionUrl,
		Description: video.Description,
		BaseModel: models.BaseModel{
			ID: *video.Id,
		},
		Labels:   labels,
		Language: *video.Language,
		Metadata: metadata,
		Title:    video.Title,
		VideoUrl: *video.VideoUrl,
	}, nil
}

func toVideoAPIModel(video models.Video) gen.Video {
	labels := map[string]string{}
	for k, v := range video.Labels {
		labels[k] = v.(string)
	}
	metadata := map[string]string{}
	for k, v := range video.Metadata {
		metadata[k] = v.(string)
	}
	return gen.Video{
		CaptionUrl:   &video.CaptionUrl,
		Description:  video.Description,
		Id:           &video.ID,
		Labels:       &labels,
		Language:     &video.Language,
		Metadata:     &metadata,
		ProcessingId: &video.Processing.ID,
		Title:        video.Title,
		VideoUrl:     &video.VideoUrl,
	}
}

func toVideosAPIModel(videos []models.Video) []gen.Video {
	videosAPI := []gen.Video{}
	for _, video := range videos {
		videosAPI = append(videosAPI, toVideoAPIModel(video))
	}
	return videosAPI
}

func toProcessingAPIModel(processing models.Processing) gen.Processing {
	return gen.Processing{
		CurrentStep: (*gen.ProcessingCurrentStep)(&processing.CurrentStep),
		Id:          &processing.ID,
		Status:      (*gen.ProcessingStatus)(&processing.Status),
		Steps:       toProcessingStepsAPIModel(processing.Steps),
		VideoId:     &processing.VideoID,
	}
}

func toProcessingStepsAPIModel(steps []models.ProcessingStep) *[]gen.ProcessingStep {
	stepsAPI := []gen.ProcessingStep{}
	for _, step := range steps {
		stepsAPI = append(stepsAPI, toProcessingStepAPIModel(step))
	}
	return &stepsAPI
}

func toProcessingStepAPIModel(step models.ProcessingStep) gen.ProcessingStep {
	return gen.ProcessingStep{
		Log:    &step.Log,
		Status: (*gen.ProcessingStepStatus)(&step.Status),
		Step:   (*gen.ProcessingStepStep)(&step.Step),
	}
}

func fromProcessingStepAPIModel(step gen.ProcessingStep) models.ProcessingStep {
	if step.Log == nil {
		step.Log = new(string)
	}
	if step.Status == nil {
		step.Status = new(gen.ProcessingStepStatus)
	}
	if step.Step == nil {
		step.Step = new(gen.ProcessingStepStep)
	}
	return models.ProcessingStep{
		Log:    *step.Log,
		Status: string(*step.Status),
		Step:   string(*step.Step),
	}
}
