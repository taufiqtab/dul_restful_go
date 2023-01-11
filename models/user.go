package models

type User struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	FullName  string `gorm:"type:varchar(255)" json:"full_name"`
	Birthdate string `gorm:"type:date" json:"birthdate"`
}
