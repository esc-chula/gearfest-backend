package controller

type CreateUserDTO struct {
	Id   string `json:"user_id"`
	Name string `json:"user_name"`
}

type CreateCheckinDTO struct {
	UserID string `json:"user_id" binding:"required"`
	LocationID string `json:"location_id" binding:"required"`
}
