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

// New return new video service
func New() VideoService {
	return &videoService
}

// Save saves video
func (service *VideoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)

	return entity
}

// FindAll finds all saved videos
func (service *VideoService) FindAll() []entity.Video {
	return service.videos
}
