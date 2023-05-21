package server

import (
	"github.com/NetfluxESIR/backend/internal/models"
	"github.com/NetfluxESIR/backend/pkg/api/gen"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *Server) GetProcessing(c *gin.Context, videoId string) {
	processing, err := s.db.GetProcessing(c, videoId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, toProcessingAPIModel(processing))
}

func (s *Server) UpdateProcessingStatus(c *gin.Context, videoId openapi_types.UUID) {
	var processingStatus gen.ProcessingStatus
	if err := c.ShouldBind(&processingStatus); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	processing, err := s.db.GetProcessing(c, videoId.String())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	processing.Status = string(processingStatus)
	if err := s.db.UpdateProcessing(c, processing); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nil)
}

func (s *Server) UpdateProcessingStep(c *gin.Context, videoId openapi_types.UUID) {
	var processingStep gen.ProcessingStep
	if err := c.ShouldBindJSON(&processingStep); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	processing, err := s.db.GetProcessing(c, videoId.String())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if processing.Steps == nil {
		processing.Steps = []models.ProcessingStep{}
	}
	processing.Steps = append(processing.Steps, fromProcessingStepAPIModel(processingStep))
	if err := s.db.UpdateProcessing(c, processing); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nil)
}

func (s *Server) LoginUser(c *gin.Context) {
	var user gen.Account
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	getUser, err := s.db.GetUser(c, string(user.Email))
	if err != nil {
		return
	}
	if getUser.Email == "" {
		c.JSON(400, gin.H{"error": "Account not found"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(getUser.HashedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := generateToken(getUser.ID.String())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = s.db.UpdateToken(c, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token.Token})
}

func (s *Server) RegisterUser(c *gin.Context) {
	var user gen.Account
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := hashAndSalt([]byte(user.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userDbModel := models.Account{
		Email:          string(user.Email),
		HashedPassword: hashedPassword,
		Role:           string(*user.Role),
	}
	err = s.db.RegisterUser(c, &userDbModel)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := generateToken(userDbModel.ID.String())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = s.db.UpdateToken(c, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"token": token.Token})
}

func (s *Server) GetVideos(c *gin.Context, params gen.GetVideosParams) {
	if params.Limit == nil {
		limit := 10
		params.Limit = &limit
	}
	if params.Offset == nil {
		offset := 0
		params.Offset = &offset
	}
	list, err := s.db.GetVideoList(c, c.GetString("userId"), *params.Limit, *params.Offset)
	if err != nil {
		return
	}
	c.JSON(200, toVideosAPIModel(list))
}

func (s *Server) CreateVideo(c *gin.Context) {
	var video gen.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userUUID, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	s.logger.Info("Creating video with title: " + video.Title)
	videoDbModel, err := fromVideoAPIModel(video)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	s.logger.Info("Video model: " + videoDbModel.Title)
	videoDbModel.UserID = userUUID
	videoDbModel.Processing = models.Processing{
		Status:      string(gen.ProcessingStatusPENDING),
		Steps:       []models.ProcessingStep{},
		CurrentStep: string(gen.ProcessingStepStepNONE),
	}
	err = s.db.CreateVideo(c, &videoDbModel)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, toVideoAPIModel(videoDbModel))
}

func (s *Server) DeleteVideo(c *gin.Context, videoId openapi_types.UUID) {
	userUUID, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = s.db.DeleteVideo(c, userUUID.String(), videoId.String())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}

func (s *Server) GetVideo(c *gin.Context, videoId openapi_types.UUID) {
	userUUID, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	video, err := s.db.GetVideo(c, userUUID.String(), videoId.String())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, toVideoAPIModel(video))
}

func (s *Server) UpdateVideo(c *gin.Context, videoId openapi_types.UUID) {
	var video gen.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userUUID, err := uuid.Parse(c.GetString("userId"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	videoInDb, err := s.db.GetVideo(c, userUUID.String(), videoId.String())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	videoReceived, err := fromVideoAPIModel(video)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	videoInDb.Merge(videoReceived)
	err = s.db.UpdateVideo(c, &videoInDb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, toVideoAPIModel(videoInDb))
}

func (s *Server) GetPresignedUrl(c *gin.Context, key string) {
	svc := s3.New(s.s3)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.cfg.S3Bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"url": urlStr})
}
