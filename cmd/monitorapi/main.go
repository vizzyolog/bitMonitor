package main

import (
	"context"
	"fmt"

	"app/controller"
	"app/interfaces/repo"
	"app/interfaces/repo/mongodb"
	"app/interfaces/router"
	"app/usecases/monitorUseCase"
)

func main() {
	db, closeDB, err := mongodb.NewMongo("mongodb://root:example@mongo:27017/", true)
	if err != nil {
		fmt.Println("error connect to mongodb ", err)
		panic(err)
	}
	defer closeDB()

	ctx := context.Background()
	repo := repo.NewMonitorStore(db)
	// Switch to redis
	// repo := repository.NewRedisRepository()
	monitorUseCase := monitorUseCase.NewMonitorUserCase(repo)
	monitorController := controller.NewMonitorController(ctx, monitorUseCase)
	httpRouter := router.NewRouter()

	httpRouter.GetDevice("/device", monitorController.FindAll)
	httpRouter.GetDevice("/device/:id", monitorController.FindById)
	httpRouter.AddNewDevice("/device", monitorController.AddMonitor)
	httpRouter.Start(":8080")
}
