package Models

type Contact struct {
	NumPhone int    `json:"NumberPhoto"`
	Email    string `json:"Email"`
	IdPerson Person `json:"Person"`
}
