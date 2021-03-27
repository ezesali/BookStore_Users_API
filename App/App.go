package App

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {

	MapsURLs()
	router.Run(":8080")

}
