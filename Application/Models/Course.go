package Models

import "time"

type Course struct {
	IdCourse      int           `json:"IdCourse"`
	NameCour      string        `json:"NameCourse"`
	DscCour       string        `json:"DescriptionCourse"`
	DurationCour  time.Duration `json:"DurationCourse"`
	SituationCour string        `json:"SituationCourse"`
	IdPerson      Person        `json:"Person"`
	IdInsti       Institution   `json:"Institution"`
}
