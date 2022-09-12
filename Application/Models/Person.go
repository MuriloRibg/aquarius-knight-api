package Models

import "time"

type Person struct {
	IdPerson  int       `json:"IdPerson"`
	Fname     string    `json:"FirstName"`
	Lname     string    `json:"LastName"`
	BirthDate time.Time `json:"BirthDate"`
}
