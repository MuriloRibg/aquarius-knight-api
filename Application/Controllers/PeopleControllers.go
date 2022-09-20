package Controllers

import (
	"aquarius-knight-api/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
)

type PeopleController struct {
	peopleService Services.PeopleService
}

// List godoc
// @Summary      Listar todos os alunos. | List all people.
// @Description  Rota para buscar todos as pessoas. | Route to fetch all people.
// @Tags         Person
// @Accept       json
// @Produce      json
// @Success      200  {object}  Models.Person
// @Router       /people [get]
func (pc *PeopleController) List(c *gin.Context) {
	if people, err := pc.peopleService.ListPeople(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
	} else {
		c.JSONP(http.StatusOK, people)
	}

}

// Search godoc
// @Summary      Recuperar uma pessoa. | Retrieve person by id.
// @Description  Get pessoa por ID. | Search person for id
// @Tags         Person
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Person ID"
// @Success      200  {array}   Models.Person
// @Failure      404  {object}  httputil.HTTPError
// @Router       /person/{id} [get]
func (pc *PeopleController) Search(c *gin.Context) {
	if person, err, id := pc.peopleService.SearchPeople(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Mesage:": err.Error()})
	} else if person.IdPerson > 0 {
		c.JSON(http.StatusOK, person)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Mensagem de erro": fmt.Sprintf("Pessoa com id = %s, não encontrado!", id),
			"Error message":    fmt.Sprintf("Person with id = %s, not found!", id),
		})
	}
}

// Insert godoc
// @Summary      Inserir uma nova pessoa. | Insert new person.
// @Description  Post pessoa/person
// @Tags         Person
// @Accept       json
// @Produce      json
// @Param        person	body	Models.Person  true  "Modelo do pessoa. | Models of person."
// @Success      200  {array}   Models.Person
// @Failure      400  {object}  httputil.HTTPError
// @Router       /person [post]
func (pc *PeopleController) Insert(c *gin.Context) {
	if err, Person := pc.peopleService.InsertPerson(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error message": err.Error()})
	} else {
		c.JSON(http.StatusOK, Person)
	}
}

// Delete godoc
// @Summary      Deletar uma pessoa. | Delete a person.
// @Description  Delete pessoa por Id. | Delete person by Id.
// @Tags         Person
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Person ID"
// @Success      200
// @Failure      400  {object}  httputil.HTTPError
// @Router       /person/{id} [delete]
func (pc *PeopleController) Delete(c *gin.Context) {
	if err, id := pc.peopleService.DeletePerson(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Mesagem": fmt.Sprintf("Pessoa com o id = %s, apagada com sucesso!\nPerson with the id = %s, successfully deleted!", id, id),
			"Message": fmt.Sprintf("Person with the id = %s, successfully deleted!", id),
		})
	}
}

// Edit godoc
// @Summary      Editar uma pessoa. | Edit a person.
// @Description  Put person por Id. | Edit person by Id.
// @Tags         Person
// @Accept       json
// @Produce      json
// @Param        person	body	Models.Person	true	"Modelo de pessoa. | Models of person."
// @Success      200  {array}   Models.Person
// @Failure      400  {object}  httputil.HTTPError
// @Router       /person/{id} [put]
func (pc *PeopleController) Edit(c *gin.Context) {
	if err, Person := pc.peopleService.EditPerson(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		c.JSON(http.StatusOK, Person)
	}
}
