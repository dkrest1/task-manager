package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/controllers"
	"github.com/dkrest1/task-manager/routes"
	_ "github.com/dkrest1/task-manager/docs"
	_ "github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	// Init DB
	configs.InitDB()

	// Controllers
	userController := controllers.NewUserController()
	taskController := controllers.NewTaskController()
	authController := controllers.NewAuthController()

	// croutes
	appRoutes := routes.Routes(userController, taskController, authController)

	// enable cors
	handler := cors.Default().Handler(appRoutes)

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in env")
	}

	fmt.Printf("Server is live and running on port %v 🚀🚀🚀", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler)

	if err != nil {
		log.Fatal(err)
	}

}

