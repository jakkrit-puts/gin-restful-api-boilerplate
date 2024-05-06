package models

type Blog struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Topic  string `json:"topic" gorm:"type:text"`
	UserID uint   `json:"user_id"`
}
