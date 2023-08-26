package models

import "time"

type Order struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time
	ProductReferer int     `json:"product_id"`
	Product        Product `gorm:"foreignKey:ProductReferer"`
	UserReferer    int     `json:"user_id"`
	User           User    `gorm:"foreignKey:UserReferer"`
}
