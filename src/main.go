package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.Run(os.Getenv(":8080"))
}
