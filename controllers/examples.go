package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jagadeeshtakkalaki/gin/db"
    "github.com/jagadeeshtakkalaki/gin/models"
)

func GetExamples(c *gin.Context) {
    var rows []models.Example
    if err := db.DB.Find(&rows).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, rows)
}

func GetExample(c *gin.Context) {
    id := c.Param("id")
    var row models.Example
    if err := db.DB.First(&row, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, row)
}

func CreateExample(c *gin.Context) {
    var input models.Example
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := db.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, input)
}

func UpdateExample(c *gin.Context) {
    id := c.Param("id")
    var row models.Example
    if err := db.DB.First(&row, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    var input models.Example
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    row.Name = input.Name
    if err := db.DB.Save(&row).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, row)
}

func DeleteExample(c *gin.Context) {
    id := c.Param("id")
    // Perform an unscoped (hard) delete so the row is removed from DB
    result := db.DB.Unscoped().Delete(&models.Example{}, id)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.Status(http.StatusNoContent)
}
