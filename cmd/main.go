package main

import (
	"job_portal/internal/repository"
	"job_portal/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main(){
	if err :=godotenv.Load(); err!=nil{
		log.Fatal("Error loading .env file")
	}

	db, err :=repository.InitDb() 
	if err!=nil {
		panic(err)
	}

	defer db.Close()
	r :=gin.Default()
	routes.InitRoutes(r,db);


	port:= os.Getenv("SERVER_PORT")
	if port == ""{
		r.Run(":8080")
	}
	r.Run(":"+port)
}