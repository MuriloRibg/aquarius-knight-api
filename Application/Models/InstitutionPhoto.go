package Models

type InstitutionPhoto struct {
	IdPhoto   int    `json:"IdPhoto"`
	IdInsti   int    `json:"IdInstitution"`
	NamePhoto string `json:"NamePhoto"`
	DscPhoto  string `json:"DescriptionPhoto"`
	Image     byte   `json:"Image"`
}
