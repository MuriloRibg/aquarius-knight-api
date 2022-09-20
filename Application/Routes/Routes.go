package Routes

import (
	"aquarius-knight-api/Middleware"
	"aquarius-knight-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/swaggo/files"
	_ "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger"
	"os"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

//swag init --pd --parseInternal --parseDepth 1

type Route struct {
	peopleRoute PeopleRoute
}

func (route *Route) HandleRequests() {
	r := gin.Default()

	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))
	request := appinsights.NewRequestTelemetry("GET", "https://myapp.azurewebsites.net/", 1, "Success")
	client.Track(request)

	//Informações do swagger
	docs.SwaggerInfo.Title = "Api do Murilove"
	docs.SwaggerInfo.Description = "Api REST com Gin e Gorm"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "aquarius-knight-api.azurewebsites.net/v1"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	r.Use(Middleware.CORSMiddleware())
	r.GET("/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.peopleRoute.Route(r)

	err := r.Run(getPort())
	if err != nil {
		panic(err.Error())
	}
}

func getPort() string {
	p := os.Getenv("HTTP_PLATFORM_PORT")
	if p != "" {
		return ":" + p
	}
	return ":80"
}
