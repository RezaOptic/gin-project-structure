package logic

import (
	"context"
	"errors"
	"github.com/RezaOptic/gin-project-structure/constant"
)

// Interface
type Interface interface {
	Get() error
}

// 
type Logic struct {
	Context context.Context
	Self    Interface
}

// New
func New(ctx context.Context) Interface {
	x := &Logic{Context: ctx}
	x.Self = x
	return x
}

func (o Logic) Get() error {
	return errors.New(constant.NewError)
}

