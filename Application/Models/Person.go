package Models

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	IdPerson  uint           `gorm:"primaryKey;autoIncrementIncrement;column:IDPERSON" json:"IdPerson"`
	Fname     string         `gorm:"column:FNAME" json:"FirstName"`
	Lname     string         `gorm:"column:LNAME" json:"LastName"`
	BirthDate *time.Time     `gorm:"column:BIRTHDATE;" json:"BirthDate"`
	CreatedAt time.Time      `gorm:"column:CREATEDAT;" json:"CreatedAt"`
	UpdatedAt time.Time      `gorm:"column:UPDATEAT;" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:DELETEDAT;" json:"DeletedAt"`
}
