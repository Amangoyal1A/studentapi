package models

type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Age      int    `gorm:"not null" json:"age"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}
