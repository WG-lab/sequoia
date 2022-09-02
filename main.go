package main

import (
	"fmt"

	"github.com/andy-wg/sequoia/config"
	"github.com/andy-wg/sequoia/logger"
	"github.com/andy-wg/sequoia/server"
)

func main() {

	err := config.InitConfig()
	if err != nil {
		fmt.Println("Init Config err, err = ", err)
		return
	}

	logger.InitLogger()

	err = server.Init()
	if err != nil {
		fmt.Println("Start Server err, err = ", err)
		return
	}
}
