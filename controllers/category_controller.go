package controllers

import (
	"database/sql"
	"net/http"
	"sb-go-quiz/config"
	"sb-go-quiz/models"

	"github.com/gin-gonic/gin"
)

// GET all categories
func GetAllCategories(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		rows.Scan(&cat.ID, &cat.Name)
		categories = append(categories, cat)
	}

	c.JSON(http.StatusOK, categories)
}

// GET category by ID
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	err := config.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&cat.ID, &cat.Name)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

// POST create new category
func CreateCategory(c *gin.Context) {
	var cat models.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	_, err := config.DB.Exec("INSERT INTO categories (name, created_by) VALUES ($1, $2)", cat.Name, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambah kategori"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kategori berhasil ditambahkan"})
}

// DELETE category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus kategori"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}

// GET books by category ID
func GetBooksByCategory(c *gin.Context) {
	id := c.Param("id")

	rows, err := config.DB.Query("SELECT id, title, description FROM books WHERE category_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil buku"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Description)
		books = append(books, b)
	}

	c.JSON(http.StatusOK, books)
}
