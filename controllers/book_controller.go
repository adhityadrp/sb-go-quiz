package controllers

import (
	"database/sql"
	"net/http"
	"sb-go-quiz/config"
	"sb-go-quiz/models"

	"github.com/gin-gonic/gin"
)

// GET all books
func GetAllBooks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, description, release_year, price, total_page, thickness FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Description, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness)
		books = append(books, b)
	}

	c.JSON(http.StatusOK, books)
}

// GET book by ID
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Book
	err := config.DB.QueryRow("SELECT id, title, description, release_year, price, total_page, thickness FROM books WHERE id = $1", id).
		Scan(&b.ID, &b.Title, &b.Description, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, b)
}

// POST create book
func CreateBook(c *gin.Context) {
	var b models.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// Validasi tahun rilis
	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year harus antara 1980 - 2024"})
		return
	}

	// Tentukan thickness
	if b.TotalPage > 100 {
		b.Thickness = "tebal"
	} else {
		b.Thickness = "tipis"
	}

	_, err := config.DB.Exec(`
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 'admin')
	`, b.Title, b.Description, b.ImageURL, b.ReleaseYear, b.Price, b.TotalPage, b.Thickness, b.CategoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan buku"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Buku berhasil ditambahkan"})
}

// DELETE book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus buku"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}
