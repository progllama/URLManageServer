package server

import (
	"fmt"
	"strconv"
	route "url_manager/routes"

	"github.com/gin-gonic/gin"
)

var (
	port int = 8000
)

func SetPort(p int) {
	port = p
}

func getPort() int {
	return port
}

func Run() {
	engine := getGinEngine()
	registerRoutes(engine)
	runEngine(engine)
}

func registerRoutes(e *gin.Engine) {
	route.RegisterRoutes(e)
}

func runEngine(e *gin.Engine) {
	e.Run(
		getFormattedPort(),
	)
}

func getGinEngine() *gin.Engine {
	return gin.Default()
}

func getFormattedPort() string {
	p := getPort()
	sp := strconv.Itoa(p)
	return fmt.Sprintf(":%s", sp)
}
