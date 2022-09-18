package Models

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	IdPerson  int64          `gorm:"primaryKey;autoIncrementIncrement" json:"IdPerson"`
	Fname     string         `gorm:"column:Fname" json:"FirstName"`
	Lname     string         `gorm:"column:Lname" json:"LastName"`
	BirthDate *time.Time     `gorm:"column:BirthDate;type:date" json:"BirthDate"`
	CreatedAt time.Time      `gorm:"column:CreatedAt;type:datetime" json:"CreatedAt"`
	UpdatedAt time.Time      `gorm:"column:UpdatedAt;type:datetime" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:DeletedAt;type:datetime" json:"DeletedAt"`
}
