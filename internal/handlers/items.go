package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/ghermosoj/go-api/internal/models"
)

var items = []models.Item{
	{ID: 1, Name: "Item 1"},
	{ID: 2, Name: "Item 2"},
}

func GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, item := range items {
        if item.ID == id {
            c.JSON(http.StatusOK, item)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func CreateItem(c *gin.Context) {
    var newItem models.Item
    if err := c.BindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newItem.ID = len(items) + 1
    items = append(items, newItem)
    c.JSON(http.StatusCreated, newItem)
}

func DeleteItem(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    for i, item := range items {
        if item.ID == id {
            items = append(items[:i], items[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Item eliminado"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Item no encontrado"})
}