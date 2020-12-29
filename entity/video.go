package entity

// Video entity
type Video struct {
	Title       string `json:"title" binding:"min=2,max=100"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
