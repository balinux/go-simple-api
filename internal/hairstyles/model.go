package hairstyles

import "time"

type HairStyle struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HairStyleCategory struct {
	HairStyleID uint `gorm:"primaryKey" json:"hair_style_id"`
	CategoryID  uint `gorm:"primaryKey" json:"category_id"`
}
