package Models

import "time"

type Person struct {
	IdPerson  int       `form:"IdPerson" json:"IdPerson" xml:"IdPerson"  binding:"required"`
	Fname     string    `form:"FirstName" json:"FirstName" xml:"FirstName"`
	Lname     string    `form:"LastName" json:"LastName" xml:"LastName"`
	BirthDate time.Time `form:"BirthDate" json:"BirthDate" xml:"BirthDate"`
}
