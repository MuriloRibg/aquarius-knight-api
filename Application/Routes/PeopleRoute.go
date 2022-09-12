package Routes

import (
	"aquarius-knight-api/Controllers"
	"github.com/gin-gonic/gin"
)

func PeopleRoute(r *gin.Engine) {
	r.GET("/v1/people", Controllers.List)
	r.GET("/v1/person/:id", Controllers.Search)
	r.POST("/v1/person", Controllers.Insert)
	r.DELETE("/v1/person/:id", Controllers.Delete)
	r.PUT("/v1/person/:id", Controllers.Edit)
}
