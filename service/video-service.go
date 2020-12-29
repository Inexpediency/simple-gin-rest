package service

import (
	"Inexpediency/simple-gin-rest/entity"
	"Inexpediency/simple-gin-rest/repository"
)

// VideoService implementation
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Delete(entity.Video)
	Update(entity.Video)
}

type videoService struct {
	videoRepository repository.VideoRepository
}

// NewVideoService return new video service
func NewVideoService() VideoService {
	return &videoService{
		videoRepository: repository.NewVideoRepository(),
	}
}

// Save saves video
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)

	return video
}

// FindAll finds all saved videos
func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}

// Update updates video in db
func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

// Delete deletes video from db
func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}
