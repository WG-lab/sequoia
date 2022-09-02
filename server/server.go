package server

import (
	"errors"

	"github.com/andy-wg/sequoia/config"
	"github.com/gin-gonic/gin"
)

func Init() error {
	gin.SetMode(config.Config.Server.GinMode)
	r := NewRouter()
	if r == nil {
		return errors.New("Create New Router Error")
	}

	err := r.Run(config.Config.Server.Port)
	if err != nil {
		return err
	}
}
