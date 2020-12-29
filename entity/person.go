package entity

import "gorm.io/gorm"

// Person entity
type Person struct {
	gorm.Model
	ID        uint64 `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=101"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}
