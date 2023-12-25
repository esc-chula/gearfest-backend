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
	LocationID string
}

type CreateCheckinDTO struct {
	UserID     string `json:"user_id" binding:"required"`
	LocationID string `json:"location_id" binding:"required"`
}
