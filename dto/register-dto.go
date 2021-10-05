package dto

//used when a client posts to /register
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password"  form:"password, omitempty" validate:"min:6"`
}
