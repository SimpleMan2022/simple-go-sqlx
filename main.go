package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-gin-sqlx/controllers"
	"go-gin-sqlx/repository"
	"go-gin-sqlx/usecase"
	"log"
	"time"
)

func main() {
	// membuka koneksi database
	username := "root"
	password := ""
	db_name := "go-sqlx-api"
	host := "localhost"
	port := 3306

	// "root:@(localhost:3306)/go-sqlx-api"
	connectionStr := fmt.Sprintf("%s:%s@(%s:%d)/%s", username, password, host, port, db_name)
	db, err := sqlx.Connect("mysql", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(30)

	//konfigurasi gin
	router := gin.New()
	// implement clean architecture
	pegawaiRepository := repository.NewPegawaiRepository(db)
	pegawaiUsecase := usecase.NewPegawaiUsecase(pegawaiRepository)
	controller := controllers.NewPegawaiController(pegawaiUsecase)

	usersRepository := repository.NewUsersRepository(db)
	usersUsecase := usecase.NewUsersUsecase(usersRepository)
	usersController := controllers.NewUsersController(usersUsecase)

	router.POST("/login", usersController.Login)
	v1 := router.Group("/v1")
	v1.GET("/pegawai", controller.GetAllPegawai)
	v1.GET("/pegawai/:id", controller.FindPegawaiByid)
	v1.POST("pegawai", controller.CreatePegawai)
	v1.PUT("/pegawai/:id", controller.UpdatePegawai)
	v1.DELETE("/pegawai/:id", controller.DeletePegawai)
	router.Run(":7000")
}
