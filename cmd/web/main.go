package main

import (
	"_core/auth"
	lib "_lib"
	web "_web"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	client := lib.InitDB(logger)
	defer client.Close()
	defer logger.Sync()

	as := auth.NewAuthService(client, "my secret", logger)

	router := web.NewRouter(as, logger)

	router.Engine.Logger.Fatal(router.Engine.Start(":1337"))
}
