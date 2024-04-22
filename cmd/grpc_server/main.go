package main

import (
	"context"
	"log"

	"github.com/kirillmc/trainings-auth/internal/app"
)

//var configPath string

//func init() {
//	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
//	flag.Parse()
//}

func main() {

	ctx := context.Background()

	a, err := app.NewApp(ctx, "prod.env")
	//a, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("failed to init app1: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
