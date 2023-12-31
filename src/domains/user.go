package domains

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

// create user
type CreateUser struct {
	UserID   string `json:"user_id" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
}

type CreateCheckinDTO struct {
	UserID     string `json:"user_id" binding:"required"`
	LocationID uint   `json:"location_id" binding:"required"`
}

type CreateUserCompletedDTO struct {
	CocktailID uint `json:"cocktail_id" binding:"required"`
}
type CreateUserNameDTO struct {
	UserName string `json:"user_name" binding:"required"`
}
