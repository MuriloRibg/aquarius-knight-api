package Models

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	IdPerson  int            `gorm:"primaryKey;autoIncrementIncrement" json:"IdPerson"`
	Fname     string         `gorm:"column:Fname" json:"FirstName"`
	Lname     string         `gorm:"column:Lname" json:"LastName"`
	BirthDate *time.Time     `gorm:"column:BirthDate;type:date" json:"-"`
	CreatedAt time.Time      `gorm:"column:CreatedAt;type:datetime" json:"-"`
	UpdatedAt time.Time      `gorm:"column:UpdatedAt;type:datetime" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:DeletedAt;type:datetime" json:"-"`
}
