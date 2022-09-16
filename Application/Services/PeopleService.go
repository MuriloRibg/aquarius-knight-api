package Services

import (
	"aquarius-knight-api/DataBase"
	"aquarius-knight-api/Models"
	"github.com/gin-gonic/gin"
)

func InsertPerson(c *gin.Context) (error, Models.Person) {
	var Person Models.Person
	if err := c.ShouldBindJSON(&Person); err != nil {
		return err, Person
	}
	DataBase.DB.Create(&Person)
	return nil, Person
}

func DeletePerson(c *gin.Context) (error, string) {
	id := c.Params.ByName("id")
	var Person Models.Person
	resp := DataBase.DB.Delete(&Person, id)
	return resp.Error, id
}

func EditPerson(c *gin.Context) (error, Models.Person) {
	id := c.Params.ByName("id")
	var Person Models.Person
	if err := c.ShouldBindJSON(&Person); err != nil {
		return err, Person
	}
	DataBase.DB.Model(&Person).Where("id = ?", id).Updates(Person)
	return nil, Person
}
