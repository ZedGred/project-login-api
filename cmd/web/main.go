package main

import (
	"fmt"
	"project/internal/config"
)

func main(){
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	log := config.NewLog(viperConfig)
	db := config.NewMongo(viperConfig,log)
	fmt.Println("anjing")
	validate := config.NewValidator(viperConfig)
	jwt := config.NewJwt(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		Config : viperConfig,
		DB: db,
		Validate: validate,
		Jwt: jwt,
		App: app,
		Log:log,
	})

	webPort := viperConfig.GetString("web.port")
	err := app.Listen(webPort)
	if err != nil{
		log.Fatalf("Failed Start Server : %v",err)
	}
}


