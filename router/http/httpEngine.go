package httpEngine

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run(Port string) {
	engine := gin.Default()

	engine.Use(gin.Recovery())

	fmt.Println(engine.Run(fmt.Sprintf(":%s", Port)))
}
