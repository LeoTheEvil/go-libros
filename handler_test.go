package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/libros", ListarLibros)
	r.POST("/libros", CrearLibro)
	r.GET("/libros/:id", ObtenerLibro)
	r.PUT("/libros/:id", ActualizarLibro)
	r.DELETE("/libros/:id", EliminarLibro)
	return r
}

func resetData() {
	libros = []Libro{}
	nextID = 1
}

func TestCrearLibro(t *testing.T) {
	resetData()
	r := setupRouter()

	libro := Libro{ID: 1, Titulo: "Go Programming", Autor: "John Doe", Genero: "Programacion"}
	body, _ := json.Marshal(libro)
	req, _ := http.NewRequest("POST", "/libros", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var creado Libro
	err := json.Unmarshal(w.Body.Bytes(), &creado)
	assert.NoError(t, err)
	assert.Equal(t, 1, creado.ID)
	assert.Equal(t, "Go Programming", creado.Titulo)
	assert.Equal(t, "John Doe", creado.Autor)
	assert.Equal(t, "Programacion", creado.Genero)
}

func TestListarLibros(t *testing.T) {
	resetData()
	r := setupRouter()

	libros = append(libros, Libro{ID: 1, Titulo: "Test", Autor: "Autor", Genero: "Prueba"})

	req, _ := http.NewRequest("GET", "/libros", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var listado []Libro
	err := json.Unmarshal(w.Body.Bytes(), &listado)
	assert.NoError(t, err)
	assert.Len(t, listado, 1)
}

func TestObtenerLibro(t *testing.T) {
	resetData()
	r := setupRouter()

	libros = append(libros, Libro{ID: 1, Titulo: "Test", Autor: "Autor", Genero: "Prueba"})

	req, _ := http.NewRequest("GET", "/libros/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var libro Libro
	err := json.Unmarshal(w.Body.Bytes(), &libro)
	assert.NoError(t, err)
	assert.Equal(t, 1, libro.ID)
}

func TestActualizarLibro(t *testing.T) {
	resetData()
	r := setupRouter()

	libros = append(libros, Libro{ID: 1, Titulo: "Viejo", Autor: "Autor", Genero: "Genero"})

	updated := Libro{Titulo: "Nuevo", Autor: "Autor 2", Genero: "Generico"}
	body, _ := json.Marshal(updated)
	req, _ := http.NewRequest("PUT", "/libros/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var libroActualizado Libro
	err := json.Unmarshal(w.Body.Bytes(), &libroActualizado)
	assert.NoError(t, err)
	assert.Equal(t, "Nuevo", libroActualizado.Titulo)
	assert.Equal(t, "Autor 2", libroActualizado.Autor)
	assert.Equal(t, "Generico", libroActualizado.Genero)
}

func TestEliminarLibro(t *testing.T) {
	resetData()
	r := setupRouter()

	libros = append(libros, Libro{ID: 1, Titulo: "Eliminar", Autor: "Autor", Genero: "Genero"})

	req, _ := http.NewRequest("DELETE", "/libros/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Len(t, libros, 0)
}
