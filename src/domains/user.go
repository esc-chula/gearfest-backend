package domains

type User struct {
	UserID          string    `json:"user_id" gorm:"primaryKey"`
	UserName        string    `json:"user_name"`
	IsUserCompleted bool      `json:"is_user_completed"`
	CocktailID      uint      `json:"cocktail_id"`
	Checkins        []Checkin `json:"checkins"`
}

type Checkin struct {
	CheckinID  uint   `json:"checkin_id" gorm:"primaryKey;autoincrement"`
	UserID     string `json:"user_id"`
	LocationID uint   `json:"location_id"`
}

type CreateCheckinDTO struct {
	LocationID *uint `json:"location_id" binding:"required"`
}

type CreateUserCompletedDTO struct {
	CocktailID uint `json:"cocktail_id" binding:"required"`
}

type CreateUserNameDTO struct {
	UserName string `json:"user_name" binding:"required"`
}
