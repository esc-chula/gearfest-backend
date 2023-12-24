package domain

type User struct {
	UserID          string `gorm:"primaryKey"`
	UserName        string
	IsUserCompleted bool
	CocktailID      uint
	Checkins        []Checkin
}

type Checkin struct {
	CheckinID   uint `gorm:"primaryKey;autoincrement"`
	UserID     string
	LocationID string
}
