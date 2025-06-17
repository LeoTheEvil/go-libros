package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var libros = []Libro{}
var nextID = 1

// ListarLibros godoc
// @Summry Lista todos los libros
// @Produce json
// @Success 200 {array} Libro
// @Router /libros [get]
func ListarLibros(c *gin.Context) {
	c.JSON(http.StatusOK, libros)
}

// CrearLibro godoc
// @Summary Crea un nuevo libro
// @Accept json
// @Produce json
// @Param libro body Libro true "Libro a crear"
// @Success 201 {object} Libro
// @Router /libros [post]
func CrearLibro(c *gin.Context) {
	var libro Libro
	if err := c.ShouldBindJSON(&libro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	libro.ID = nextID
	nextID++
	libros = append(libros, libro)
	c.JSON(http.StatusCreated, libro)
}

// ObtenerLibro godoc
// @Summary Obtiene un libro por ID
// @Produce json
// @Param id path int true "ID del libro"
// @Success 200 {object} Libro
// @Failure 404 {string} string "Libro no encontrado"
// @Router /libros/{id} [get]
func ObtenerLibro(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, I := range libros {
		if I.ID == id {
			c.JSON(http.StatusOK, I)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
}

// ActualizarLibro godoc
// @Summary Actualiza un libro
// @Accept json
// @Produce json
// @Param id path int trrue "ID del libro"
// @Param libro body Libro true "Libro actualizado"
// @Success 200 {object} Libro
// @Router /libros/{id} [put]
func ActualizarLibro(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var actualizado Libro
	if err := c.ShouldBindJSON(&actualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, I := range libros {
		if I.ID == id {
			actualizado.ID = id
			libros[i] = actualizado
			c.JSON(http.StatusOK, actualizado)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
}

// EliminarLibro godoc
// @Summary Elimina un libro
// @Param id path int true "ID del libro"
// @Success 204
// @Router /libros/{id} [delete]
func EliminarLibro(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, I := range libros {
		if I.ID == id {
			libros = append(libros[:i], libros[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
}
