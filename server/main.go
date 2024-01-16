package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/controllers"
	"github.com/dkrest1/task-manager/routes"
	"github.com/joho/godotenv"
)

func main() {

	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env  file found!")
	}

	// init DB
	configs.InitDB()

	// Controllers
	userController := controllers.NewUserController()
	taskController := controllers.NewTaskController()

	// croutes
	appRoutes := routes.Routes(userController, taskController)

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in env")
	}

	fmt.Printf("Server is live and running on port %v 🚀🚀🚀", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), appRoutes)

	if err != nil {
		log.Fatal(err)
	}

}

