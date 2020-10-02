package httpEngine

import (
	"github.com/RezaOptic/gin-project-structure/logic"
	"github.com/gin-gonic/gin"
	"gitlab.com/RezaOptic/go-utils/errorsHandler"

)

// Interface
type Interface interface {
	Get(c *gin.Context)
}

// 
type Struct struct {
	Self Interface
}

// New
func New() Interface {
	x := &Struct{}
	x.Self = x
	return x
}

func (s *Struct) Get(c *gin.Context) {
	err := logic.New(c).Get()
	if err != nil {
		errorsHandler.GinErrorResponseHandler(c, err)
		return
	}
}

