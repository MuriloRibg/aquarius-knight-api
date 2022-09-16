package Controllers

import (
	"aquarius-knight-api/DataBase"
	"aquarius-knight-api/Models"
	"aquarius-knight-api/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
)

// List godoc
// @Summary      Listar todos os alunos. | List all people.
// @Description  Rota para buscar todos as pessoas. | Route to fetch all people.
// @Tags         Person
// @Accept       json
// @Produce      json
// @Success      200  {object}  Models.Person
// @Router       /people [get]
func List(c *gin.Context) {
	var People []Models.Person
	DataBase.DB.Find(&People)
	c.JSON(http.StatusOK, People)
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
func Search(c *gin.Context) {
	id := c.Params.ByName("id")

	var Person Models.Person
	DataBase.DB.Find(&Person, id)
	if Person.IdPerson > 0 {
		c.JSON(http.StatusOK, Person)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"Mensagem de erro": fmt.Sprintf("Pessoa com id = %s, não encontrado!", id),
		"Error message":    fmt.Sprintf("Person with id = %s, not found!", id),
	})
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
func Insert(c *gin.Context) {
	err, Person := Services.InsertPerson(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Person)
	return
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
func Delete(c *gin.Context) {
	err, id := Services.DeletePerson(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Mesagem": fmt.Sprintf("Pessoa com o id = %s, apagada com sucesso!\nPerson with the id = %s, successfully deleted!", id),
		"Message": fmt.Sprintf("Person with the id = %s, successfully deleted!", id),
	})
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
func Edit(c *gin.Context) {
	err, Person := Services.EditPerson(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Person)
}
