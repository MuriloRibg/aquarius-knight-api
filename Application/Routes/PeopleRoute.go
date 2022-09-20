package Routes

import (
	"aquarius-knight-api/Controllers"
	"github.com/gin-gonic/gin"
)

type PeopleRoute struct {
	peopleController Controllers.PeopleController
}

func (pr *PeopleRoute) Route(r *gin.Engine) {
	r.GET("/v1/people", pr.peopleController.List)
	r.GET("/v1/person/:id", pr.peopleController.Search)
	r.POST("/v1/person", pr.peopleController.Insert)
	r.DELETE("/v1/person/:id", pr.peopleController.Delete)
	r.PUT("/v1/person/:id", pr.peopleController.Edit)
}
