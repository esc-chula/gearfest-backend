package domain

type User struct {
	UserID string `gorm:"primaryKey"`
	UserName string
	IsUserCompleted bool
	Checkins []Checkin `gorm:"foreignKey:UserID;references:UserID"`
}

type Checkin struct {
	Checkin uint `gorm:"primaryKey"`
	UserID string
	LocationID string
	User User `gorm:"foreignKey:UserID;references:UserID"`
}