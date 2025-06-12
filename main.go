package main

import (
	_ "go-libros/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Libros
// @version 1.0
// @description API para gestionar libros (ABM)

// @hosst localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/libros", ListarLibros)
	r.POST("/libros", CrearLibro)
	r.GET("/libros/:id", ObtenerLibro)
	r.PUT("/libros/:id", ActualizarLibro)
	r.DELETE("/libros/:id", EliminarLibro)

	r.Run(":8080")
}
