package entity

// Person entity
type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1, lte=101"`
	Email    string `json:"email" binding:"required,email"`
}
