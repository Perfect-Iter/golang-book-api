package dto

//used when user posts from /login
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password"  form:"password, omitempty" validate:"min:6"`
}
