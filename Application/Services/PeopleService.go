package Services

import (
	"aquarius-knight-api/DataBase"
	"aquarius-knight-api/Models"
	"github.com/gin-gonic/gin"
)

type PeopleService struct {
	Person Models.Person
	People []Models.Person
}

func (ps *PeopleService) ListPeople() ([]Models.Person, error) {
	resp := DataBase.DB.Find(&ps.People)
	return ps.People, resp.Error
}

func (ps *PeopleService) SearchPeople(c *gin.Context) (Models.Person, error, string) {
	id := c.Params.ByName("id")
	err := DataBase.DB.Find(&ps.Person, id)
	return ps.Person, err.Error, id
}

func (ps *PeopleService) InsertPerson(c *gin.Context) (error, Models.Person) {
	err := c.ShouldBindJSON(&ps.Person)
	DataBase.DB.Create(&ps.Person)
	return err, ps.Person
}

func (ps *PeopleService) DeletePerson(c *gin.Context) (error, string) {
	id := c.Params.ByName("id")
	resp := DataBase.DB.Delete(&ps.Person, id)
	return resp.Error, id
}

func (ps *PeopleService) EditPerson(c *gin.Context) (error, Models.Person) {
	id := c.Params.ByName("id")
	err := c.ShouldBindJSON(&ps.Person)
	DataBase.DB.Model(&ps.Person).Where("id = ?", id).Updates(ps.Person)
	return err, ps.Person
}
