package domain

type User struct {
	UserID          string `gorm:"primaryKey"`
	UserName        string
	IsUserCompleted bool
	CocktailID      uint
	Checkins        []Checkin	
}

type Checkin struct {
	CheckinID  uint `gorm:"primaryKey;autoincrement"`
	UserID     string
	LocationID uint
}

type CreateCheckinDTO struct {
	UserID     string `json:"user_id" binding:"required"`
	LocationID uint   `json:"location_id" binding:"required"`
}

type CreateUserCompletedDTO struct {
	IsUserCompleted  bool `json:"is_user_completed"`
}
type CreateUserNameDTO struct {
	UserName  string `json:"user_name"`
}