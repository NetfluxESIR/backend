package server

import (
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetProcessing(c *gin.Context, videoId string) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateProcessingStatus(c *gin.Context, videoId openapi_types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateProcessingStep(c *gin.Context, videoId openapi_types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) LoginUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) RegisterUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetVideos(c *gin.Context, params gen.GetVideosParams) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreateVideo(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteVideo(c *gin.Context, videoId openapi_types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetVideo(c *gin.Context, videoId openapi_types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateVideo(c *gin.Context, videoId openapi_types.UUID) {
	//TODO implement me
	panic("implement me")
}
