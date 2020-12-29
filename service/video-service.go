package service

import "Inexpediency/simple-gin-rest/entity"

// VideoService implementation
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// NewVideoService return new video service
func NewVideoService() VideoService {
	return &videoService{}
}

// Save saves video
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)

	return video
}

// FindAll finds all saved videos
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
