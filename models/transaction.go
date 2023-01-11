package models

import "time"

type Transaction struct {
	Id            int64     `gorm:"primaryKey" json:"id"`
	TransactionNo string    `gorm:"type:varchar(300)" json:"transaction_no"`
	ProductId     int64     `gorm:"type:int" json:"product_id"`
	Total         int64     `gorm:"type:int" json:"total"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
