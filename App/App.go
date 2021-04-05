package App

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {

	fmt.Println("Looking for MapsURLs..")

	MapsURLs()

	PORT := ":8080"

	router.Run(PORT)

	fmt.Println("Running on port: ", PORT[1:])

}
